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
	"gsm/gsmci"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// groupsCiDeleteCmd represents the delete command
var groupsCiDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a Group.",
	Long:  "https://cloud.google.com/identity/docs/how-to/delete-dynamic-groups#python",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		name, err := getGroupCiName(flags["name"].GetString(), flags["email"].GetString())
		if err != nil {
			log.Fatalf("Error resolving group name: %v", err)
		}
		result, err := gsmci.DeleteGroup(name)
		if err != nil {
			log.Fatalf("Error deleting group: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(groupsCiCmd, groupsCiDeleteCmd, groupCiFlags)
}
