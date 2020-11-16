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

// eventsInsertBatchCmd represents the batchInsert command
var eventsInsertBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch inserts events using a CSV file as input.",
	Long:  "https://developers.google.com/calendar/v3/reference/events/insert",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		err = gsmhelpers.CheckBatchFlags(flags, eventFlags, int64(len(csv[0])))
		if err != nil {
			log.Fatalf("Error with batch flag index: %v\n", err)
		}
		l := len(csv)
		results := make(chan *calendar.Event, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []*calendar.Event{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, eventFlags, line, "insert")
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
					event, err := mapToEvent(m)
					if err != nil {
						log.Printf("Error building event object: %v", err)
					}
					errKey := fmt.Sprintf("%s - %s:", m["calendarId"].GetString(), event.Id)
					operation := func() error {
						result, err := gsmcalendar.InsertEvent(m["calendarId"].GetString(), m["sendUpdates"].GetString(), m["fields"].GetString(), event, m["conferenceDataVersion"].GetInt64(), m["maxAttendees"].GetInt64(), m["supportsAttachments"].GetBool())
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
	eventsInsertCmd.AddCommand(eventsInsertBatchCmd)
	flags := eventsInsertBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(eventFlags, flags, "insert")
	markFlagsRequired(eventsInsertBatchCmd, eventFlags, "insert")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(eventsInsertBatchCmd, batchFlags, "")
}
