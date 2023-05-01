/*
Copyright © 2020-2023 Hannes Hayashi

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

// devicesCancelWipeCmd represents the cancelWipe command
var devicesCancelWipeCmd = &cobra.Command{
	Use:   "cancelWipe",
	Short: "Cancels an unfinished device wipe.",
	Long: `This operation can be used to cancel device wipe in the gap between the wipe operation returning success and the device being wiped.
This operation is possible when the device is in a "pending wipe" state.
The device enters the "pending wipe" state when a wipe device command is issued, but has not yet been sent to the device.
The cancel wipe will fail if the wipe command has already been issued to the device.
Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/devices/cancelWipe`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cancelWipeRequest, err := mapToCancelWipeDeviceRequest(flags)
		if err != nil {
			log.Fatalf("Error building cancelWipeRequest object: %v", err)
		}
		result, err := gsmci.CancelDeviceWipe(flags["name"].GetString(), flags["fields"].GetString(), cancelWipeRequest)
		if err != nil {
			log.Fatalf("Error cancelling device wipe: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(devicesCmd, devicesCancelWipeCmd, deviceFlags)
}
