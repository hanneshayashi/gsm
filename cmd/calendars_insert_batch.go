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

	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/calendar/v3"

	"github.com/spf13/cobra"
)

// calendarsInsertBatchCmd represents the batch command
var calendarsInsertBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch inserts secondary calendars using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/calendar/api/v3/reference/calendars/insert",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, calendarFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *calendar.Calendar, cap)
		go func() {
			for range cap {
				wg.Go(func() {
					for m := range maps {
						c, err := mapToCalendar(m)
						if err != nil {
							log.Printf("Error building calendar object: %v\n", err)
							continue
						}
						result, err := gsmcalendar.InsertCalendar(c, m["fields"].GetString())
						if err != nil {
							log.Println(err)
						} else {
							results <- result
						}
					}
				})
			}
			wg.Wait()
			close(results)
		}()
		gsmhelpers.StreamOrCollect(gsmhelpers.ChanToIter(results, nil), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(calendarsInsertCmd, calendarsInsertBatchCmd, calendarFlags, calendarFlagsALL, batchFlags)
}
