/*
Copyright © 2020-2022 Hannes Hayashi

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

// chromeOsDevicesMoveToOUCmd represents the moveToOU command
var chromeOsDevicesMoveToOUCmd = &cobra.Command{
	Use:               "moveToOU",
	Short:             "Move or insert multiple Chrome OS devices to an organizational unit. You can move up to 50 devices at once.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/chromeosdevices/moveDevicesToOu",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		d, err := mapToChromeOsMoveDevicesToOu(flags)
		if err != nil {
			log.Fatalf("Error building ChromeOsMoveDevicesToOu object: %v", err)
		}
		result, err := gsmadmin.MoveChromeOSDevicesToOU(flags["customerId"].GetString(), flags["orgUnitPath"].GetString(), d)
		if err != nil {
			log.Fatalf("Error moving Chrome OS device: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(chromeOsDevicesCmd, chromeOsDevicesMoveToOUCmd, chromeOsDeviceFlags)
}
