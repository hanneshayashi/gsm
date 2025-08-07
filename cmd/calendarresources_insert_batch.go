/*
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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// calendarResourcesInsertBatchCmd represents the batch command
var calendarResourcesInsertBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch inserts calendar resources using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/admin/directory/reference/rest/v1/resources.calendars/insert",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, calendarResourceFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *admin.CalendarResource, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						c, err := mapToCalendarResource(m)
						if err != nil {
							log.Printf("Error building calendarResource object: %v\n", err)
							continue
						}
						result, err := gsmadmin.InsertCalendarResource(m["customer"].GetString(), m["fields"].GetString(), c)
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
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for r := range results {
				err := enc.Encode(r)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*admin.CalendarResource{}
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
	gsmhelpers.InitBatchCommand(calendarResourcesInsertCmd, calendarResourcesInsertBatchCmd, calendarResourceFlags, calendarResourceFlagsALL, batchFlags)
}
