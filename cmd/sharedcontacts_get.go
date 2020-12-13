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

// sharedContactsGetCmd represents the get command
var sharedContactsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a Domain Shared Contact via its URL / ID",
	Long:  "",
	Annotations: map[string]string{
		"crescendoFlags": "--json",
	},
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.GetSharedContact(flags["url"].GetString())
		if err != nil {
			log.Fatalf("Error getting shared contact: %v", err)
		}
		if flags["json"].GetBool() {
			gsmhelpers.StreamOutput(result, "json", compressOutput)
		} else {
			gsmhelpers.StreamOutput(result, "xml", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(sharedContactsCmd, sharedContactsGetCmd, sharedContactFlags)
}
