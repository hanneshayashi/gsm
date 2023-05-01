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

	"github.com/hanneshayashi/gsm/gsmcibeta"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	cibeta "google.golang.org/api/cloudidentity/v1beta1"

	"github.com/spf13/cobra"
)

// orgUnitsMembershipsListBatchCmd represents the batch command
var orgUnitsMembershipsListBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch list Shared Drives in organizational units using a CSV file as input.",
	Long:  "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1beta1/orgUnits.memberships/list",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, orgUnitsMembershipFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			Parent      string `json:"parent,omitempty"`
			Filter      string `json:"filter,omitempty"`
			Customer    string `json:"customer,omitempty"`
			Memberships []*cibeta.OrgMembership
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						parent := gsmhelpers.EnsurePrefix(m["parent"].GetString(), "orgUnits/")
						filter := m["filter"].GetString()
						customer := m["customer"].GetString()
						result, err := gsmcibeta.ListOrgUnitMemberships(parent, customer, filter, m["fields"].GetString(), gsmhelpers.MaxThreads(0))
						r := resultStruct{Parent: parent, Filter: filter, Customer: customer}
						for i := range result {
							r.Memberships = append(r.Memberships, i)
						}
						e := <-err
						if e != nil {
							log.Println(e)
						} else {
							results <- r
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
	gsmhelpers.InitBatchCommand(orgUnitsMembershipsListCmd, orgUnitsMembershipsListBatchCmd, orgUnitsMembershipFlags, orgUnitsMembershipFlagsALL, batchFlags)
}
