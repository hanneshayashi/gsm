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

// deviceUsersBlockCmd represents the block command
var deviceUsersBlockCmd = &cobra.Command{
	Use:               "block",
	Short:             "Blocks device from accessing user data",
	Long:              `https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers/block`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		blockDeviceUserRequest, err := mapToBlockDeviceUserRequest(flags)
		if err != nil {
			log.Fatalf("Error building blockDeviceUserRequest object: %v", err)
		}
		result, err := gsmci.BlockDeviceUser(flags["name"].GetString(), flags["fields"].GetString(), blockDeviceUserRequest)
		if err != nil {
			log.Fatalf("Error blocking device user: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(deviceUsersCmd, deviceUsersBlockCmd, deviceUserFlags)
}
