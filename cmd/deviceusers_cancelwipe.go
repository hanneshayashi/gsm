/*
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

// deviceUsersCancelWipeCmd represents the cancelWipe command
var deviceUsersCancelWipeCmd = &cobra.Command{
	Use:   "cancelWipe",
	Short: "Cancels an unfinished user account wipe.",
	Long: `This operation can be used to cancel device wipe in the gap between the wipe operation returning success and the device being wiped.
Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers/cancelWipe`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cancelWipeDeviceUserRequest, err := mapToCancelWipeDeviceUserRequest(flags)
		if err != nil {
			log.Fatalf("Error building cancelWipeDeviceUserRequest object: %v", err)
		}
		result, err := gsmci.CancelDeviceUserWipe(flags["name"].GetString(), flags["fields"].GetString(), cancelWipeDeviceUserRequest)
		if err != nil {
			log.Fatalf("Error cancelling device user wipe: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(deviceUsersCmd, deviceUsersCancelWipeCmd, deviceUserFlags)
}
