/*
Copyright © 2020-2024 Hannes Hayashi

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

	"github.com/spf13/cobra"
)

// orgUnitsMembershipsMoveCmd represents the move command
var orgUnitsMembershipsMoveCmd = &cobra.Command{
	Use: "move",
	Short: `Move an OrgMembership to a new OrgUnit.
NOTE: This is an atomic copy-and-delete.
The resource will have a new copy under the destination OrgUnit and be deleted from the source OrgUnit.
The resource can only be searched under the destination OrgUnit afterwards.`,
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1beta1/orgUnits.memberships/move",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		orgMembershipMoveRequest, err := mapToOrgMembershipMoveRequest(flags)
		if err != nil {
			log.Fatalf("Error building org unit membership move request object: %v", err)
		}
		var name string
		if flags["name"].IsSet() {
			name = flags["name"].GetString()
		} else if flags["driveId"].IsSet() {
			name = "orgUnits/-/memberships/shared_drive;" + flags["driveId"].GetString()
		}
		result, err := gsmcibeta.MoveOrgUnitMemberships(name, flags["fields"].GetString(), orgMembershipMoveRequest)
		if err != nil {
			log.Fatalf("Error moving org unit membership: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(orgUnitsMembershipsCmd, orgUnitsMembershipsMoveCmd, orgUnitsMembershipFlags)
}
