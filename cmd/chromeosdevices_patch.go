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

// chromeOsDevicesPatchCmd represents the patch command
var chromeOsDevicesPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates a device's updatable properties, such as annotatedUser, annotatedLocation, notes, orgUnitPath, or annotatedAssetId. This method supports patch semantics",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/chromeosdevices/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToChromeOsDevice(flags)
		if err != nil {
			log.Fatalf("Error building chromeOsDevice object: %v", err)
		}
		result, err := gsmadmin.PatchChromeOsDevice(flags["customerId"].GetString(), flags["deviceId"].GetString(), flags["fields"].GetString(), flags["projection"].GetString(), c)
		if err != nil {
			log.Fatalf("Error patching Chrome OS device: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(chromeOsDevicesCmd, chromeOsDevicesPatchCmd, chromeOsDeviceFlags)
}
