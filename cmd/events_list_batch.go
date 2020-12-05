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
	"github.com/spf13/viper"
	"google.golang.org/api/calendar/v3"
)

// eventsListBatchCmd represents the batch command
var eventsListBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch lists events using a CSV file as input.",
	Long:  "https://developers.google.com/calendar/v3/reference/events/list",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, eventFlags, viper.GetInt("threads"))
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			CalendarID string            `json:"calendarId,omitempty"`
			Events     []*calendar.Event `json:"events,omitempty"`
		}
		results := make(chan resultStruct, cap)
		final := []resultStruct{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmcalendar.ListEvents(m["calendarId"].GetString(), m["iCalUID"].GetString(), m["orderBy"].GetString(), m["q"].GetString(), m["timeZone"].GetString(), m["timeMax"].GetString(), m["timeMin"].GetString(), m["updatedMin"].GetString(), m["fields"].GetString(), m["privateExtendedProperty"].GetStringSlice(), m["sharedExtendedProperty"].GetStringSlice(), m["maxAttendees"].GetInt64(), m["showDeleted"].GetBool(), m["showHiddenInvitations"].GetBool(), m["singleEvents"].GetBool())
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{CalendarID: m["calendarId"].GetString(), Events: result}
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
	gsmhelpers.InitBatchCommand(eventsListCmd, eventsListBatchCmd, eventFlags, eventFlagsALL, batchFlags)
}
