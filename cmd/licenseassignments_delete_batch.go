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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlicensing"

	"github.com/spf13/cobra"
)

// licenseAssignmentsDeleteBatchCmd represents the batch command
var licenseAssignmentsDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch deletes user license asignments using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/admin-sdk/licensing/reference/rest/v1/licenseAssignments/delete",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, licenseAssignmentFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			ProductID string `json:"productId,omitempty"`
			SkuID     string `json:"skuId,omitempty"`
			UserID    string `json:"userId,omitempty"`
			Result    bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						productID := m["productId"].GetString()
						skuID := m["skuId"].GetString()
						userID := m["userId"].GetString()
						result, err := gsmlicensing.DeleteLicenseAssignment(productID, skuID, userID)
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{ProductID: productID, SkuID: skuID, UserID: userID, Result: result}
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
	gsmhelpers.InitBatchCommand(licenseAssignmentsDeleteCmd, licenseAssignmentsDeleteBatchCmd, licenseAssignmentFlags, licenseAssignmentFlagsALL, batchFlags)
}
