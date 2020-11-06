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
	"fmt"
	"gsm/gsmci"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// groupMembershipsCiModifyMembershipRolesCmd represents the modifyMembershipRoles command
var groupMembershipsCiModifyMembershipRolesCmd = &cobra.Command{
	Use:   "modifyMembershipRoles",
	Short: "Modifies the MembershipRoles of a Membership.",
	Long:  `https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups.memberships/modifyMembershipRoles`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		m, err := mapToModifyMembershipRolesRequestCi(flags)
		if err != nil {
			log.Fatalf("Error building group membership object: %v", err)
		}
		result, err := gsmci.ModifyMembershipRoles(flags["name"].GetString(), flags["fields"].GetString(), m)
		if err != nil {
			log.Fatalf("Error creating membership %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	groupMembershipsCiCmd.AddCommand(groupMembershipsCiModifyMembershipRolesCmd)
	gsmhelpers.AddFlags(groupMembershipCiFlags, groupMembershipsCiModifyMembershipRolesCmd.Flags(), groupMembershipsCiModifyMembershipRolesCmd.Use)
	markFlagsRequired(groupMembershipsCiModifyMembershipRolesCmd, groupMembershipCiFlags, "")
}
