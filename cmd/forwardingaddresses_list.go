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
	"gsm/gsmgmail"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// forwardingAddressesListCmd represents the list command
var forwardingAddressesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the forwarding addresses for the specified account.",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.settings.forwardingAddresses/list",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmail.ListForwardingAddresses(flags["userId"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error listing forwarding address for user %s: %v", flags["userId"].GetString(), err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(forwardingAddressesCmd, forwardingAddressesListCmd, forwardingAddressFlags)
}
