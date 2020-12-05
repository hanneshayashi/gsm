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

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// eventsInsertBatchCmd represents the batchInsert command
var eventsInsertBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch inserts events using a CSV file as input.",
	Long:  "https://developers.google.com/calendar/v3/reference/events/insert",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, eventFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *calendar.Event, cap)
		final := []*calendar.Event{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						event, err := mapToEvent(m)
						if err != nil {
							log.Printf("Error building event object: %v", err)
						}
						result, err := gsmcalendar.InsertEvent(m["calendarId"].GetString(), m["sendUpdates"].GetString(), m["fields"].GetString(), event, m["conferenceDataVersion"].GetInt64(), m["maxAttendees"].GetInt64(), m["supportsAttachments"].GetBool())
						if err != nil {
							log.Println(err)
						} else {
							results <- result
						}
					}
					wg.Done()
				}()
			}
			wg.Wait()
			close(results)
		}()
		for res := range results {
			final = append(final, res)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(eventsInsertCmd, eventsInsertBatchCmd, eventFlags, eventFlagsALL, batchFlags)
}
