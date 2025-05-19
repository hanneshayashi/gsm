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

	"github.com/hanneshayashi/gsm/gsmgmailpostmaster"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// postmasterTrafficStatsGetCmd represents the get command
var postmasterTrafficStatsGetCmd = &cobra.Command{
	Use: "get",
	Short: `Get traffic statistics for a domain on a specific date.
Returns PERMISSION_DENIED if user does not have permission to access TrafficStats for the domain.`,
	Long:              "Implements the API documented at https://developers.google.com/gmail/postmaster/reference/rest/v1/domains.trafficStats/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmailpostmaster.GetTrafficStats(flags["name"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting traffic stat: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(postmasterTrafficStatsCmd, postmasterTrafficStatsGetCmd, postmasterTrafficStatFlags)
}
