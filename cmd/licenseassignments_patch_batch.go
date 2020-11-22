/*
Package cmd contains the commands available to the end user
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
	"fmt"
	"gsm/gsmhelpers"
	"gsm/gsmlicensing"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/licensing/v1"
)

// licenseAssignmentsPatchBatchCmd represents the batch command
var licenseAssignmentsPatchBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Patch patches users' license asignments using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/licensing/v1/reference/licenseAssignments/patch",
	Run: func(cmd *cobra.Command, args []string) {
		retrier := gsmhelpers.NewStandardRetrier()
		var wg sync.WaitGroup
		maps, err := gsmhelpers.GetBatchMaps(cmd, licenseAssignmentFlags, viper.GetInt("threads"))
		cap := cap(maps)
		if err != nil {
			log.Fatalln(err)
		}
		type resultStruct struct {
			ProductID string `json:"productId,omitempty"`
			SkuID     string `json:"skuId,omitempty"`
			UserID    string `json:"userId,omitempty"`
			Result    bool   `json:"result"`
		}
		results := make(chan *licensing.LicenseAssignment, cap)
		final := []*licensing.LicenseAssignment{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						var err error
						licenseAssignmentPatch, err := mapToLicenseAssignment(m)
						if err != nil {
							log.Printf("Error building licenseAssignmentPatch object: %v\n", err)
							continue
						}
						errKey := fmt.Sprintf("%s - %s - %s:", m["productId"].GetString(), m["skuId"].GetString(), m["userId"].GetString())
						operation := func() error {
							result, err := gsmlicensing.PatchLicenseAssignment(m["productId"].GetString(), m["skuId"].GetString(), m["userId"].GetString(), m["fields"].GetString(), licenseAssignmentPatch)
							if err != nil {
								retryable := gsmhelpers.ErrorIsRetryable(err)
								if retryable {
									log.Println(errKey, "Retrying after", err)
									return err
								}
								log.Println(errKey, "Giving up after", err)
								return nil
							}
							results <- result
							return nil
						}
						err = retrier.Run(operation)
						if err != nil {
							log.Println(errKey, "Max retries reached. Giving up after", err)
						}
						time.Sleep(200 * time.Millisecond)
					}
					wg.Done()
				}()
			}
			wg.Wait()
			close(results)
		}()
		for res := range results {
			final = append(final, res)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json"))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(licenseAssignmentsPatchCmd, licenseAssignmentsPatchBatchCmd, licenseAssignmentFlags, licenseAssignmentFlagsALL, batchFlags)
}
