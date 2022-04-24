/*
Copyright © 2020-2022 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmgroupssettings"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/groupssettings/v1"
)

// groupSettingsPatchBatchCmd represents the patch command
var groupSettingsPatchBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch patches groups' settings using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/admin-sdk/groups-settings/v1/reference/groups/patch",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, groupSettingFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *groupssettings.Groups, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						g, err := mapToGroupSettings(m)
						if err != nil {
							log.Printf("Error building group settings object: %v", err)
							continue
						}
						result, err := gsmgroupssettings.PatchGroupSettings(m["groupUniqueId"].GetString(), m["fields"].GetString(), g)
						if err != nil {
							log.Println(err)
						} else {
							if m["ignoreDeprecated"].GetBool() {
								result = ignoreDeprecatedGroupSettings(result)
							}
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
			final := []*groupssettings.Groups{}
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
	gsmhelpers.InitBatchCommand(groupSettingsPatchCmd, groupSettingsPatchBatchCmd, groupSettingFlags, groupSettingFlagsALL, batchFlags)
}
