/*
Copyright © 2020-2021 Hannes Hayashi

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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// eventsListBatchCmd represents the batch command
var eventsListBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch lists events using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/calendar/v3/reference/events/list",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, eventFlags)
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
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						calendarID := m["calendarId"].GetString()
						result, err := gsmcalendar.ListEvents(calendarID, m["iCalUID"].GetString(), m["orderBy"].GetString(), m["q"].GetString(), m["timeZone"].GetString(), m["timeMax"].GetString(), m["timeMin"].GetString(), m["updatedMin"].GetString(), m["fields"].GetString(), m["privateExtendedProperty"].GetStringSlice(), m["sharedExtendedProperty"].GetStringSlice(), m["maxAttendees"].GetInt64(), m["showDeleted"].GetBool(), m["showHiddenInvitations"].GetBool(), m["singleEvents"].GetBool(), cap)
						r := resultStruct{CalendarID: calendarID}
						for i := range result {
							r.Events = append(r.Events, i)
						}
						e := <-err
						if e != nil {
							log.Println(e)
						} else {
							results <- r
						}
					}
					wg.Done()
				}()
			}
			wg.Wait()
			close(results)
		}()
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for r := range results {
				err := enc.Encode(r)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []resultStruct{}
			for res := range results {
				final = append(final, res)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	gsmhelpers.InitBatchCommand(eventsListCmd, eventsListBatchCmd, eventFlags, eventFlagsALL, batchFlags)
}
