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

// calendarSettingsGetBatchCmd represents the batch command
var calendarSettingsGetBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch gets calendar settings using a CSV file as input",
	Long:  "Implements the API documented at https://developers.google.com/workspace/calendar/api/v3/reference/settings/get",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, calendarSettingFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *calendar.Setting, cap)
		go func() {
			for range cap {
				wg.Go(func() {
					for m := range maps {
						result, err := gsmcalendar.GetSetting(m["setting"].GetString(), m["fields"].GetString())
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
	gsmhelpers.InitBatchCommand(calendarSettingsGetCmd, calendarSettingsGetBatchCmd, calendarSettingFlags, calendarSettingFlagsALL, batchFlags)
}
