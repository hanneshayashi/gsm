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
)

// calendarListsDeleteBatchCmd represents the batch command
var calendarListsDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch returns calendars from the user's calendar list using a CSV file as input.",
	Long:  "https://developers.google.com/calendar/v3/reference/calendarList/delete",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		l := len(csv)
		type resultStruct struct {
			CalendarID string `json:"calendarId,omitempty"`
			Result     bool   `json:"result"`
		}
		results := make(chan resultStruct, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []resultStruct{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, calendarListFlags, line, "delete")
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
					operation := func() error {
						result, err := gsmcalendar.DeleteCalendarListEntry(m["calendarId"].GetString())
						if err != nil {
							retryable := gsmhelpers.ErrorIsRetryable(err)
							if retryable {
								log.Println("Retrying after", err)
								return err
							}
							log.Println("Giving up after", err)
							return nil
						}
						results <- resultStruct{CalendarID: m["calendarId"].GetString(), Result: result}
						return nil
					}
					err = retrier.Run(operation)
					if err != nil {
						log.Println("Max retry reached. Giving up after", err)
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
	calendarListsDeleteCmd.AddCommand(calendarListsDeleteBatchCmd)
	flags := calendarListsDeleteBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(calendarListFlags, flags, "delete")
	markFlagsRequired(calendarListsDeleteBatchCmd, calendarListFlags, "delete")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(calendarListsDeleteBatchCmd, batchFlags, "")
}
