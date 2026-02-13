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

// domainAliasesDeleteBatchCmd represents the batch command
var domainAliasesDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch deletes domain aliases of the customer using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/admin/directory/reference/rest/v1/domainAliases/delete",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, domainAliasFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			Customer        string `json:"customer,omitempty"`
			DomainAliasName string `json:"domainAliasName,omitempty"`
			Result          bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for range cap {
				wg.Go(func() {
					for m := range maps {
						customer := m["customer"].GetString()
						domainAliasName := m["domainAliasName"].GetString()
						result, err := gsmadmin.DeleteDomainAlias(customer, domainAliasName)
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{Customer: customer, DomainAliasName: domainAliasName, Result: result}
					}
				})
			}
			wg.Wait()
			close(results)
		}()
		gsmhelpers.StreamOrCollect(gsmhelpers.ChanToIter(results, nil), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(domainAliasesDeleteCmd, domainAliasesDeleteBatchCmd, domainAliasFlags, domainAliasFlagsALL, batchFlags)
}
