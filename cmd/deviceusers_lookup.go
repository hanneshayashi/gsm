/*
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
	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// deviceUsersLookupCmd represents the lookup command
var deviceUsersLookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Looks up resource names of the DeviceUsers associated with the caller's credentials, as well as the properties provided in the request.",
	Long: `This method must be called with end-user credentials with the scope: https://www.googleapis.com/auth/cloud-identity.devices.lookup

If multiple properties are provided, only DeviceUsers having all of these properties are considered as matches - i.e. the query behaves like an AND.
Different platforms require different amounts of information from the caller to ensure that the DeviceUser is uniquely identified.
 - iOS: No properties need to be passed, the caller's credentials are sufficient to identify the corresponding DeviceUser.
 - Android: Specifying the 'androidId' field is required.
 - Desktop: Specifying the 'rawResourceId' field is required.
 Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers/lookup`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		gsmhelpers.StreamOrCollect(gsmci.LookupDeviceUsers(flags["parent"].GetString(), flags["androidId"].GetString(), flags["rawResourceId"].GetString(), flags["userId"].GetString(), flags["fields"].GetString()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(deviceUsersCmd, deviceUsersLookupCmd, deviceUserFlags)
}
