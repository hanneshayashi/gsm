/*
Package cmd contains the commands available to the end user
Copyright © 2020 Hannes Hayashi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"gsm/gsmcalendar"
	"gsm/gsmhelpers"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// calendarListsGetBatchCmd represents the batch command
var calendarListsGetBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch returns calendars from the user's calendar list using a CSV file as input.",
	Long:  "https://developers.google.com/calendar/v3/reference/calendarList/get",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		err = gsmhelpers.CheckBatchFlags(flags, calendarListFlags, int64(len(csv[0])))
		if err != nil {
			log.Fatalf("Error with batch flag index: %v\n", err)
		}
		l := len(csv)
		results := make(chan *calendar.CalendarListEntry, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []*calendar.CalendarListEntry{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, calendarListFlags, line, "get")
				maps <- m
			}
			close(maps)
			wg1.Done()
		}()
		wg2.Add(1)
		retrier := gsmhelpers.NewStandardRetrier()
		for i := 0; i < gsmhelpers.MaxThreads(l); i++ {
			wg2.Add(1)
			go func() {
				for m := range maps {
					var err error
					errKey := fmt.Sprintf("%s:", m["calendarId"].GetString())
					operation := func() error {
						result, err := gsmcalendar.GetCalendarListEntry(m["calendarId"].GetString(), m["fields"].GetString())
						if err != nil {
							retryable := gsmhelpers.ErrorIsRetryable(err)
							if retryable {
								log.Println(errKey, "Retrying after", err)
								return err
							}
							log.Println(errKey, "Giving up after", err)
							return nil
						}
						results <- result
						return nil
					}
					err = retrier.Run(operation)
					if err != nil {
						log.Println(errKey, "Max retries reached. Giving up after", err)
					}
					time.Sleep(200 * time.Millisecond)
				}
				wg2.Done()
			}()
		}
		wg3.Add(1)
		go func() {
			for res := range results {
				final = append(final, res)
			}
			wg3.Done()
		}()
		wg2.Done()
		wg1.Wait()
		wg2.Wait()
		close(results)
		wg3.Wait()
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json"))
	},
}

func init() {
	calendarListsGetCmd.AddCommand(calendarListsGetBatchCmd)
	flags := calendarListsGetBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(calendarListFlags, flags, "get")
	markFlagsRequired(calendarListsGetBatchCmd, calendarListFlags, "get")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(calendarListsGetBatchCmd, batchFlags, "")
}
