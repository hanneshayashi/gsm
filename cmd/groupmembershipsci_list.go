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

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// groupMembershipsCiListCmd represents the list command
var groupMembershipsCiListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists the Memberships within a Group.",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/groups.memberships/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent, er := getGroupCiName(flags["parent"].GetString(), flags["email"].GetString())
		if er != nil {
			log.Fatalf("Error determining group name: %v", er)
		}
		gsmhelpers.StreamOrCollect(gsmci.ListMembers(parent, flags["fields"].GetString(), flags["view"].GetString()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(groupMembershipsCiCmd, groupMembershipsCiListCmd, groupMembershipCiFlags)
}
