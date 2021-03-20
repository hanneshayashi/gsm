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

	"github.com/hanneshayashi/gsm/gsmgmailpostmaster"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/gmailpostmastertools/v1"

	"github.com/spf13/cobra"
)

// postmasterTrafficStatsListCmd represents the list command
var postmasterTrafficStatsListCmd = &cobra.Command{
	Use: "list",
	Short: `List traffic statistics for all available days.
Returns PERMISSION_DENIED if user does not have permission to access TrafficStats for the domain.`,
	Long:              "https://developers.google.com/gmail/postmaster/reference/rest/v1/domains.trafficStats/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent := gsmgmailpostmaster.GetPostmasterDomainName(flags["parent"].GetString())
		result, err := gsmgmailpostmaster.ListTrafficStats(parent, flags["fields"].GetString(), flags["startDateDay"].GetInt64(), flags["startDateMonth"].GetInt64(), flags["startDateYear"].GetInt64(), flags["endDateDay"].GetInt64(), flags["endDateMonth"].GetInt64(), flags["endDateYear"].GetInt64(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*gmailpostmastertools.TrafficStats{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing traffic stats: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(postmasterTrafficStatsCmd, postmasterTrafficStatsListCmd, postmasterTrafficStatFlags)
}
