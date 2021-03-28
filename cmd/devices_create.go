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

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// devicesCreateCmd represents the create command
var devicesCreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Creates a device. Only company-owned device may be created.",
	Long:              `https://cloud.google.com/identity/docs/reference/rest/v1/devices/create`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		device, err := mapToDevice(flags)
		if err != nil {
			log.Fatalf("Error building device object: %v", err)
		}
		result, err := gsmci.CreateDevice(flags["customer"].GetString(), flags["fields"].GetString(), device)
		if err != nil {
			log.Fatalf("Error creating device: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(devicesCmd, devicesCreateCmd, deviceFlags)
}
