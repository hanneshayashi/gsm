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

// groupMembershipsCiSearchTransitiveMembershipsCmd represents the searchTransitiveMemberships command
var groupMembershipsCiSearchTransitiveMembershipsCmd = &cobra.Command{
	Use:   "searchTransitiveMemberships",
	Short: "Search transitive memberships of a group.",
	Long:  `https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups.memberships/searchTransitiveMemberships`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent, err := getGroupCiName(flags["parent"].GetString(), flags["email"].GetString())
		if err != nil {
			log.Fatalf("%v", err)
		}
		result, err := gsmci.SearchTransitiveMemberships(parent, flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error searching transitive groups: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	groupMembershipsCiCmd.AddCommand(groupMembershipsCiSearchTransitiveMembershipsCmd)
	gsmhelpers.AddFlags(groupMembershipCiFlags, groupMembershipsCiSearchTransitiveMembershipsCmd.Flags(), groupMembershipsCiSearchTransitiveMembershipsCmd.Use)
	markFlagsRequired(groupMembershipsCiSearchTransitiveMembershipsCmd, groupMembershipCiFlags, "")
}
