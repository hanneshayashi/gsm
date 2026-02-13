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
	"github.com/hanneshayashi/gsm/gsmcibeta"
	"github.com/hanneshayashi/gsm/gsmhelpers"

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
		gsmhelpers.StreamOrCollect(gsmcibeta.ListOrgUnitMemberships(gsmhelpers.EnsurePrefix(flags["parent"].GetString(), "orgUnits/"), flags["customer"].GetString(), flags["filter"].GetString(), flags["fields"].GetString()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(orgUnitsMembershipsCmd, orgUnitsMembershipsListCmd, orgUnitsMembershipFlags)
}
