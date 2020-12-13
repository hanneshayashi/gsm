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
	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// groupMembershipsCiGetCmd represents the get command
var groupMembershipsCiGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a Membership.",
	Long:  `https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups.memberships/get`,	
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		name, err := getGroupCiName(flags["name"].GetString(), flags["email"].GetString())
		if err != nil {
			log.Fatalf("%v", err)
		}
		result, err := gsmci.GetMembership(name, flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error deleting membership %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(groupMembershipsCiCmd, groupMembershipsCiGetCmd, groupMembershipCiFlags)
}
