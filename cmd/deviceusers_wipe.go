/*
Copyright © 2020-2025 Hannes Hayashi

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

// deviceUsersWipeCmd represents the wipe command
var deviceUsersWipeCmd = &cobra.Command{
	Use:   "wipe",
	Short: "Wipes the user's account on a device.",
	Long: `Other data on the device that is not associated with the user's work account is not affected.
For example, if a Gmail app is installed on a device that is used for personal and work purposes, and the user is logged in to the Gmail app with their personal account as well as their work account, wiping the "deviceUser" by their work administrator will not affect their personal account within Gmail or other apps such as Photos.
Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers/wipe`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		wipeDeviceUserRequest, err := mapToWipeDeviceUserRequest(flags)
		if err != nil {
			log.Fatalf("Error building wipeDeviceUserRequest object: %v", err)
		}
		result, err := gsmci.WipeDeviceUser(flags["name"].GetString(), flags["fields"].GetString(), wipeDeviceUserRequest)
		if err != nil {
			log.Fatalf("Error wiping device user: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(deviceUsersCmd, deviceUsersWipeCmd, deviceUserFlags)
}
