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

	"github.com/flowchartsman/retry"
	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// orgUnitsInsertBatchCmd represents the batch command
var orgUnitsInsertBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch inserts organizational units using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/orgunits/insert",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		l := len(csv)
		results := make(chan *admin.OrgUnit, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []*admin.OrgUnit{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, orgUnitFlags, line, "insert")
				maps <- m
			}
			close(maps)
			wg1.Done()
		}()
		wg2.Add(1)
		retrier := retry.NewRetrier(10, 250*time.Millisecond, 60*time.Second)
		for i := 0; i < gsmhelpers.MaxThreads(l); i++ {
			wg2.Add(1)
			go func() {
				for m := range maps {
					var err error
					o, err := mapToOrgUnit(m)
					if err != nil {
						log.Printf("Error building org unit object: %v\n", err)
						continue
					}
					operation := func() error {
						result, err := gsmadmin.InsertOrgUnit(m["customerId"].GetString(), m["fields"].GetString(), o)
						if err != nil {
							log.Println(err)
							gerr := err.(*googleapi.Error)
							if gerr.Code == 403 {
								return err
							}
							return nil
						}
						results <- result
						return nil
					}
					err = retrier.Run(operation)
					if err != nil {
						log.Fatal(err)
					}
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
	orgUnitsInsertCmd.AddCommand(orgUnitsInsertBatchCmd)
	flags := orgUnitsInsertBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(orgUnitFlags, flags, "insert")
	markFlagsRequired(orgUnitsInsertBatchCmd, orgUnitFlags, "insert")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(orgUnitsInsertBatchCmd, batchFlags, "")
}
