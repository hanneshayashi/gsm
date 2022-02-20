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
	"fmt"
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmcibeta"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// groupsCiUpdateSecuritySettingsBatchCmd represents the batch command
var groupsCiUpdateSecuritySettingsBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch retrieves groups' security settings (member restrictions) using a CSV file as input.",
	Long:  "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/updateSecuritySettings",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, groupCiFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan map[string]interface{}, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						name, err := getGroupCiName(m["name"].GetString(), m["email"].GetString())
						if err != nil {
							log.Printf("Error determining group name: %v\n", err)
							continue
						}
						securitySettings, err := mapToSecuritySettings(m)
						if err != nil {
							log.Fatalf("Error building security settings object: %v", err)
						}
						result, err := gsmcibeta.UpdateSecuritySettings(fmt.Sprintf("%s/securitySettings", name), m["updateMask"].GetString(), m["fields"].GetString(), securitySettings)
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
			final := []map[string]interface{}{}
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
	gsmhelpers.InitBatchCommand(groupsCiUpdateSecuritySettingsCmd, groupsCiUpdateSecuritySettingsBatchCmd, groupCiFlags, groupCiFlagsALL, batchFlags)
}