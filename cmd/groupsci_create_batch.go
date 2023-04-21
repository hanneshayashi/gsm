/*
Copyright © 2020-2023 Hannes Hayashi

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
	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/googleapi"

	"github.com/spf13/cobra"
)

// groupsCiCreateBatchCmd represents the batch command
var groupsCiCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch creates groups using a CSV file as input.",
	Long:  "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/groups/create",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, groupCiFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *googleapi.RawMessage, cap)
		customerID, err := gsmadmin.GetOwnCustomerID()
		if err != nil {
			log.Printf("Error determining customer ID: %v\n", err)
		}
		parent := "customers/" + customerID
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						g, err := mapToGroupCi(m)
						if g.Parent == "" {
							g.Parent = parent
						}
						if err != nil {
							log.Printf("Error building group object: %v\n", err)
							continue
						}
						if len(g.Labels) == 0 {
							g.Labels = map[string]string{
								"cloudidentity.googleapis.com/groups.discussion_forum": "",
							}
						}
						result, err := gsmci.CreateGroup(g, m["initialGroupConfig"].GetString(), m["fields"].GetString())
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
				err := enc.Encode(r)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*googleapi.RawMessage{}
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
	gsmhelpers.InitBatchCommand(groupsCiCreateCmd, groupsCiCreateBatchCmd, groupCiFlags, groupCiFlagsALL, batchFlags)
}
