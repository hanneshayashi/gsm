/*
Copyright © 2020-2022 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// groupMembershipsCiGetMembershipGraphCmd represents the getMembershipGraph command
var groupMembershipsCiGetMembershipGraphCmd = &cobra.Command{
	Use:               "getMembershipGraph",
	Short:             "Get a membership graph of just a member or both a member and a group.",
	Long:              `Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/groups.memberships/getMembershipGraph`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent, err := getGroupCiName(flags["parent"].GetString(), flags["email"].GetString())
		if err != nil {
			log.Fatalf("Error determining group name: %v", err)
		}
		result, err := gsmci.GetMembershipGraph(parent, flags["query"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting membership graph: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupMembershipsCiCmd, groupMembershipsCiGetMembershipGraphCmd, groupMembershipCiFlags)
}
