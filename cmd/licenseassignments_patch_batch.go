/*
Package cmd contains the commands available to the end user
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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlicensing"

	"github.com/spf13/cobra"
	"google.golang.org/api/licensing/v1"
)

// licenseAssignmentsPatchBatchCmd represents the batch command
var licenseAssignmentsPatchBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Patch patches users' license asignments using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/licensing/v1/reference/licenseAssignments/patch",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, licenseAssignmentFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *licensing.LicenseAssignment, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						licenseAssignmentPatch, err := mapToLicenseAssignment(m)
						if err != nil {
							log.Printf("Error building licenseAssignmentPatch object: %v\n", err)
							continue
						}
						result, err := gsmlicensing.PatchLicenseAssignment(m["productId"].GetString(), m["skuId"].GetString(), m["userId"].GetString(), m["fields"].GetString(), licenseAssignmentPatch)
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
				enc.Encode(r)
			}
		} else {
			final := []*licensing.LicenseAssignment{}
			for res := range results {
				final = append(final, res)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitBatchCommand(licenseAssignmentsPatchCmd, licenseAssignmentsPatchBatchCmd, licenseAssignmentFlags, licenseAssignmentFlagsALL, batchFlags)
}
