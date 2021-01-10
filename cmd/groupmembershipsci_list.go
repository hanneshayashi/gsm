/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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
	ci "google.golang.org/api/cloudidentity/v1beta1"

	"github.com/spf13/cobra"
)

// groupMembershipsCiListCmd represents the list command
var groupMembershipsCiListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists members of a (dynamic) group",
	Long:              "https://cloud.google.com/identity/docs/how-to/retrieve-list-dynamic-groups",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent, er := getGroupCiName(flags["parent"].GetString(), flags["email"].GetString())
		if er != nil {
			log.Fatalf("Error determining group name: %v", er)
		}
		result, err := gsmci.ListMembers(parent, flags["fields"].GetString(), flags["view"].GetString(), gsmhelpers.MaxThreads(0))
		if err != nil {
			log.Fatalf("Error listing members: %v", err)
		}
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*ci.Membership{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing members: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupMembershipsCiCmd, groupMembershipsCiListCmd, groupMembershipCiFlags)
}
