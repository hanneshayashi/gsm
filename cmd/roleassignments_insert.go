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
	"fmt"
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// roleAssignmentsInsertCmd represents the insert command
var roleAssignmentsInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "A brief description of your command",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/roleAssignments/insert",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToRoleAssignment(flags)
		if err != nil {
			log.Fatalf("Error building role assignment object: %v", err)
		}
		result, err := gsmadmin.InsertRoleAssignment(flags["customer"].GetString(), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error deleting role assignment %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	roleAssignmentsCmd.AddCommand(roleAssignmentsInsertCmd)
	gsmhelpers.AddFlags(roleAssignmentFlags, roleAssignmentsInsertCmd.Flags(), roleAssignmentsInsertCmd.Use)
	markFlagsRequired(roleAssignmentsInsertCmd, roleAssignmentFlags, "")
}
