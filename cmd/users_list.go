/*
Copyright © 2020-2023 Hannes Hayashi

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

// usersListCmd represents the list command
var usersListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Retrieves a paginated list of either deleted users or all users in a domain.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.ListUsers(flags["showDeleted"].GetBool(), flags["query"].GetString(), flags["domain"].GetString(), flags["customer"].GetString(), flags["fields"].GetString(), flags["projection"].GetString(), flags["orderBy"].GetString(), flags["sortOrder"].GetString(), flags["viewType"].GetString(), flags["customFieldMask"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*admin.User{}
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
			log.Fatalf("Error listing users: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(usersCmd, usersListCmd, userFlags)
}
