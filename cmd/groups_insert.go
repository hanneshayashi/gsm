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

// groupsInsertCmd represents the insert command
var groupsInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Creates a group.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/groups/insert",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		g, err := mapToGroup(flags)
		if err != nil {
			log.Fatalf("Error building group object: %v", err)
		}
		result, err := gsmadmin.InsertGroup(g, flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error inserting group %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	groupsCmd.AddCommand(groupsInsertCmd)
	gsmhelpers.AddFlags(groupFlags, groupsInsertCmd.Flags(), groupsInsertCmd.Use)
	markFlagsRequired(groupsInsertCmd, groupFlags, "")
}
