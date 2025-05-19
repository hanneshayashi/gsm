/*
Copyright © 2020-2025 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmcibeta"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	cibeta "google.golang.org/api/cloudidentity/v1beta1"

	"github.com/spf13/cobra"
)

// orgUnitsMembershipsListCmd represents the list command
var orgUnitsMembershipsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List OrgMembership resources in an OrgUnit treated as 'parent'.",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1beta1/orgUnits.memberships/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmcibeta.ListOrgUnitMemberships(gsmhelpers.EnsurePrefix(flags["parent"].GetString(), "orgUnits/"), flags["customer"].GetString(), flags["filter"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*cibeta.OrgMembership{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing org unit memberships: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(orgUnitsMembershipsCmd, orgUnitsMembershipsListCmd, orgUnitsMembershipFlags)
}
