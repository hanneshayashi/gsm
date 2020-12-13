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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// sharedContactsListCmd represents the list command
var sharedContactsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all shared contacts in your domain",
	Long:  `Example: gsm sharedContacts list --domain "example.org"`,
	Annotations: map[string]string{
		"crescendoFlags": "--json",
	},
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.ListSharedContacts(flags["domain"].GetString())
		if err != nil {
			log.Fatalf("Error listing shared contacts: %v", err)
		}
		if flags["json"].GetBool() {
			gsmhelpers.StreamOutput(result, "json", compressOutput)
		} else {
			gsmhelpers.StreamOutput(result, "xml", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(sharedContactsCmd, sharedContactsListCmd, sharedContactFlags)
}
