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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// forwardingAddressesListCmd represents the list command
var forwardingAddressesListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists the forwarding addresses for the specified account.",
	Long:              "Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users.settings.forwardingAddresses/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmail.ListForwardingAddresses(flags["userId"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error listing forwarding address for user %s: %v", flags["userId"].GetString(), err)
		}
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err = enc.Encode(result[i])
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			err = gsmhelpers.Output(result, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	gsmhelpers.InitCommand(forwardingAddressesCmd, forwardingAddressesListCmd, forwardingAddressFlags)
}
