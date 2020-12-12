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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// schemasInsertBatchCmd represents the batch command
var schemasInsertBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch inserts schemas using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/schemas/insert",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, schemaFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *admin.Schema, cap)
		final := []*admin.Schema{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						s, err := mapToSchema(m)
						if err != nil {
							log.Printf("Error building schema object: %v\n", err)
							continue
						}
						result, err := gsmadmin.InsertSchema(m["customerId"].GetString(), m["fields"].GetString(), s)
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
		for res := range results {
			final = append(final, res)
		}
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(schemasInsertCmd, schemasInsertBatchCmd, schemaFlags, schemaFlagsALL, batchFlags)
}