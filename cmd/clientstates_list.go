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
	ci "google.golang.org/api/cloudidentity/v1"

	"github.com/spf13/cobra"
)

// clientStatesListCmd represents the list command
var clientStatesListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists the client states for the given search query.",
	Long:              `https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers.clientStates/list`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmci.ListClientStates(flags["parent"].GetString(), flags["customer"].GetString(), flags["filter"].GetString(), flags["orderBy"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*ci.GoogleAppsCloudidentityDevicesV1ClientState{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing client states: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(clientStatesCmd, clientStatesListCmd, clientStateFlags)
}