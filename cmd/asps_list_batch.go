/*
Copyright © 2020-2024 Hannes Hayashi

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

// aspsListBatchCmd represents the batch command
var aspsListBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch lists ASPs issued by a user using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/asps/list",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, aspFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			UserKey string       `json:"userKey,omitempty"`
			Asps    []*admin.Asp `json:"asps,omitempty"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						userKey := m["userKey"].GetString()
						result, err := gsmadmin.ListAsps(userKey, m["fields"].GetString())
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{UserKey: userKey, Asps: result}
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
	gsmhelpers.InitBatchCommand(aspsListCmd, aspsListBatchCmd, aspFlags, aspFlagsALL, batchFlags)
}
