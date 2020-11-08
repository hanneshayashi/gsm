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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// resourcesBuildingsPatchBatchCmd represents the batch command
var resourcesBuildingsPatchBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch retrieves buildings using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/resources/buildings/patch",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		l := len(csv)
		results := make(chan *admin.Building, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []*admin.Building{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, resourcesBuildingFlags, line, "patch")
				maps <- m
			}
			close(maps)
			wg1.Done()
		}()
		wg2.Add(1)
		retrier := gsmhelpers.NewStandardRetrier()
		for i := 0; i < gsmhelpers.MaxThreads(l); i++ {
			wg2.Add(1)
			go func() {
				for m := range maps {
					var err error
					b, err := mapToBuilding(m)
					if err != nil {
						log.Printf("Error building resourcesBuilding object: %v\n", err)
						continue
					}
					errKey := fmt.Sprintf("%s - %s:", m["customer"].GetString(), m["buildingId"].GetString())
					operation := func() error {
						result, err := gsmadmin.PatchResourcesBuilding(m["customer"].GetString(), m["buildingId"].GetString(), m["coordinatesSource"].GetString(), m["fields"].GetString(), b)
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
				wg2.Done()
			}()
		}
		wg3.Add(1)
		go func() {
			for res := range results {
				final = append(final, res)
			}
			wg3.Done()
		}()
		wg2.Done()
		wg1.Wait()
		wg2.Wait()
		close(results)
		wg3.Wait()
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json"))
	},
}

func init() {
	resourcesBuildingsPatchCmd.AddCommand(resourcesBuildingsPatchBatchCmd)
	flags := resourcesBuildingsPatchBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(resourcesBuildingFlags, flags, "patch")
	markFlagsRequired(resourcesBuildingsPatchBatchCmd, resourcesBuildingFlags, "patch")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(resourcesBuildingsPatchBatchCmd, batchFlags, "")
}
