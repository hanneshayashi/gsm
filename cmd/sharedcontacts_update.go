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

// sharedContactsUpdateCmd represents the update command
var sharedContactsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a shared contact",
	Long: `Only supplied values will be updated, but multi-value fields must be supplied with ALL values.
Example: sharedContacts update --phoneNumber "phoneNumber=+12348;primary=false;label=Mobile" --url https://www.google.com/m8/feeds/contacts/example.org/base/a1034b28e4f62f3`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		s, _, err := gsmadmin.GetSharedContact(flags["url"].GetString())
		if err != nil {
			log.Fatalf("Error getting shared contact: %v", err)
		}
		s, err = mapToSharedContact(flags, s)
		if err != nil {
			log.Fatalf("Error building shared contact object: %v", err)
		}
		result, _, err := gsmadmin.UpdateSharedContact(flags["url"].GetString(), s)
		if err != nil {
			log.Fatalf("Error creating shared contact: %v", err)
		}
		if flags["json"].GetBool() {
			fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
		} else {
			fmt.Println(gsmhelpers.PrettyPrint(result, "xml"))
		}
	},
}

func init() {
	gsmhelpers.InitCommand(sharedContactsCmd, sharedContactsUpdateCmd, sharedContactFlags)
}
