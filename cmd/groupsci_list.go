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
	"gsm/gsmci"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// groupsCiListCmd represents the list command
var groupsCiListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the Groups under a customer or namespace.",
	Long:  "https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/list",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent := flags["parent"].GetString()
		if parent == "" {
			customerID, err := gsmadmin.GetOwnCustomerID()
			if err != nil {
				log.Fatalf("Error determining customer ID: %v", err)
			}
			parent = "customers/" + customerID
		}
		result, err := gsmci.ListGroups(parent, flags["view"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error updating group %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	gsmhelpers.InitCommand(groupsCiCmd, groupsCiListCmd, groupCiFlags)
}
