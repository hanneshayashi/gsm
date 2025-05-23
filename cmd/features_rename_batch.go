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
)

// featuresRenameBatchCmd represents the batch command
var featuresRenameBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch renames feature resources using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/admin/directory/reference/rest/v1/resources.features/rename",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, featureFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			Customer   string `json:"customer,omitempty"`
			Featurekey string `json:"featurekey,omitempty"`
			Result     bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						f, err := mapToFeatureRename(m)
						if err != nil {
							log.Printf("Error building feature rename object: %v\n", err)
							continue
						}
						customer := m["customer"].GetString()
						result, err := gsmadmin.RenameFeature(customer, m["oldName"].GetString(), f)
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{Featurekey: m["featurekey"].GetString(), Customer: customer, Result: result}
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
	gsmhelpers.InitBatchCommand(featuresRenameCmd, featuresRenameBatchCmd, featureFlags, featureFlagsALL, batchFlags)
}
