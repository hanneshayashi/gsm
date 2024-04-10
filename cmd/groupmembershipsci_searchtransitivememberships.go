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

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	ci "google.golang.org/api/cloudidentity/v1"

	"github.com/spf13/cobra"
)

// groupMembershipsCiSearchTransitiveMembershipsCmd represents the searchTransitiveMemberships command
var groupMembershipsCiSearchTransitiveMembershipsCmd = &cobra.Command{
	Use:               "searchTransitiveMemberships",
	Short:             "Search transitive memberships of a group.",
	Long:              `Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/groups.memberships/searchTransitiveMemberships`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent, er := getGroupCiName(flags["parent"].GetString(), flags["email"].GetString())
		if er != nil {
			log.Fatalf("Error determining group name: %v", er)
		}
		result, err := gsmci.SearchTransitiveMemberships(parent, flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*ci.MemberRelation{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error searching for transitive members: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupMembershipsCiCmd, groupMembershipsCiSearchTransitiveMembershipsCmd, groupMembershipCiFlags)
}
