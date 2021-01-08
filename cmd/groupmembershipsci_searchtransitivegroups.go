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
	"log"

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	ci "google.golang.org/api/cloudidentity/v1beta1"

	"github.com/spf13/cobra"
)

// groupMembershipsCiSearchTransitiveGroupsCmd represents the searchTransitiveGroups command
var groupMembershipsCiSearchTransitiveGroupsCmd = &cobra.Command{
	Use:               "searchTransitiveGroups",
	Short:             "Search transitive groups of a member.",
	Long:              `https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups.memberships/searchTransitiveGroups`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent, er := getGroupCiName(flags["parent"].GetString(), flags["email"].GetString())
		if er != nil {
			log.Fatalf("Error determining group name: %v", er)
		}
		result, err := gsmci.SearchTransitiveGroups(parent, flags["query"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*ci.GroupRelation{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error searching for transitive groups: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupMembershipsCiCmd, groupMembershipsCiSearchTransitiveGroupsCmd, groupMembershipCiFlags)
}
