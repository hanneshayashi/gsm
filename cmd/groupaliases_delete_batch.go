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
)

// groupAliasesDeleteBatchCmd represents the batch command
var groupAliasesDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch deletes group aliases using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/groups/aliases/delete",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, groupAliasFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			GroupKey string `json:"groupKey,omitempty"`
			Alias    string `json:"alias,omitempty"`
			Result   bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		final := []resultStruct{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmadmin.DeleteGroupAlias(m["groupKey"].GetString(), m["alias"].GetString())
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{Alias: m["alias"].GetString(), GroupKey: m["groupKey"].GetString(), Result: result}
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
	gsmhelpers.InitBatchCommand(groupAliasesDeleteCmd, groupAliasesDeleteBatchCmd, groupAliasFlags, groupAliasFlagsALL, batchFlags)
}
