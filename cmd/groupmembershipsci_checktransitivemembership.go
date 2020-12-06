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
	"gsm/gsmci"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// groupMembershipsCiCheckTransitiveMembershipCmd represents the checkTransitiveMembership command
var groupMembershipsCiCheckTransitiveMembershipCmd = &cobra.Command{
	Use:   "checkTransitiveMembership",
	Short: "Check a potential member for membership in a group.",
	Long: `A member has membership to a group as long as there is a single viewable transitive membership between the group and the member.
The actor must have view permissions to at least one transitive membership between the member and group.

https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups.memberships/checkTransitiveMembership`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent, err := getGroupCiName(flags["parent"].GetString(), flags["email"].GetString())
		if err != nil {
			log.Fatalf("%v", err)
		}
		result, err := gsmci.CheckTransitiveMembership(parent, flags["query"].GetString())
		if err != nil {
			log.Fatalf("Error checking transitive group membership %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(groupMembershipsCiCmd, groupMembershipsCiCheckTransitiveMembershipCmd, groupMembershipCiFlags)
}
