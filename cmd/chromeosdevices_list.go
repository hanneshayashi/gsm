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
	"log"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// chromeOsDevicesListCmd represents the list command
var chromeOsDevicesListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Retrieves a paginated list of Chrome OS devices within an account.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/chromeosdevices/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.ListChromeOsDevices(flags["customerId"].GetString(), flags["query"].GetString(), flags["orgUnitPath"].GetString(), flags["fields"].GetString(), flags["projection"].GetString())
		if err != nil {
			log.Fatalf("Error listing Chrome OS devices: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(chromeOsDevicesCmd, chromeOsDevicesListCmd, chromeOsDeviceFlags)
}
