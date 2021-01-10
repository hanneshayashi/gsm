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
	"strconv"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	admin "google.golang.org/api/admin/directory/v1"

	"github.com/spf13/cobra"
)

// roleAssignmentsListCmd represents the list command
var roleAssignmentsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Retrieves a paginated list of all roleAssignments.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/roleAssignments/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		roleID := flags["roleId"].GetInt64()
		var roleIDString string
		if roleID != 0 {
			roleIDString = strconv.FormatInt(roleID, 10)
		}
		result, err := gsmadmin.ListRoleAssignments(flags["customer"].GetString(), roleIDString, flags["userKey"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*admin.RoleAssignment{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing role assignments: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(roleAssignmentsCmd, roleAssignmentsListCmd, roleAssignmentFlags)
}
