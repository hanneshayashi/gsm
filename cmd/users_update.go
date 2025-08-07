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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// usersUpdateCmd represents the update command
var usersUpdateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Updates a user using patch semantics.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/admin/directory/reference/rest/v1/users/update",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		u, err := mapToUser(flags)
		if err != nil {
			log.Fatalf("Error building user object: %v", err)
		}
		result, err := gsmadmin.UpdateUser(flags["userKey"].GetString(), flags["fields"].GetString(), u)
		if err != nil {
			log.Fatalf("Error updating user: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(usersCmd, usersUpdateCmd, userFlags)
}
