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
						var iterErr error
						r := resultStruct{Parent: parent, Filter: filter, Customer: customer}
						for i, err := range gsmcibeta.ListOrgUnitMemberships(parent, customer, filter, m["fields"].GetString()) {
							if err != nil {
								iterErr = err
								break
							}
							r.Memberships = append(r.Memberships, i)
						}
						if iterErr != nil {
							log.Println(iterErr)
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
		gsmhelpers.StreamOrCollect(gsmhelpers.ChanToIter(results, nil), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(orgUnitsMembershipsListCmd, orgUnitsMembershipsListBatchCmd, orgUnitsMembershipFlags, orgUnitsMembershipFlagsALL, batchFlags)
}
