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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	admin "google.golang.org/api/admin/directory/v1"

	"github.com/spf13/cobra"
)

// membersListCmd represents the list command
var membersListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Retrieves a paginated list of all members in a group.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/members/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.ListMembers(flags["groupKey"].GetString(), flags["roles"].GetString(), flags["fields"].GetString(), flags["includeDerivedMembership"].GetBool(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*admin.Member{}
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
			log.Fatalf("Error listing members: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(membersCmd, membersListCmd, memberFlags)
}
