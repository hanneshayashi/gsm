/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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

// chromeOsDevicesActionCmd represents the action command
var chromeOsDevicesActionCmd = &cobra.Command{
	Use:               "action",
	Short:             "Takes an action that affects a Chrome OS Device. This includes deprovisioning, disabling, and re-enabling devices.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/chromeosdevices/action",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		a, err := mapToChromeOsDeviceAction(flags)
		if err != nil {
			log.Fatalf("Error building chromeOsDeviceAction object: %v", err)
		}
		result, err := gsmadmin.TakeActionOnChromeOsDevice(flags["customerId"].GetString(), flags["resourceId"].GetString(), a)
		if err != nil {
			log.Fatalf("Error taking action on Chrome OS device: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(chromeOsDevicesCmd, chromeOsDevicesActionCmd, chromeOsDeviceFlags)
}
