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
	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// chromeOsCommandsGetCmd represents the get command
var chromeOsCommandsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets command data a specific command issued to the device.",
	Long:  "https://developers.google.com/admin-sdk/directory/reference/rest/v1/customer.devices.chromeos.commands/get",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.GetCommand(flags["customerId"].GetString(), flags["deviceId"].GetString(), flags["fields"].GetString(), flags["commandId"].GetInt64())
		if err != nil {
			log.Fatalf("Error getting command: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(chromeOsCommandsCmd, chromeOsCommandsGetCmd, chromeOsCommandFlags)
}
