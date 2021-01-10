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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmreports"
	reports "google.golang.org/api/admin/reports/v1"

	"github.com/spf13/cobra"
)

// activitiesListCmd represents the list command
var activitiesListCmd = &cobra.Command{
	Use: "list",
	Short: `Retrieves a list of activities for a specific customer's account and application such as the Admin console application or the Google Drive application.
For more information, see the guides for administrator and Google Drive activity reports.
For more information about the activity report's parameters, see the activity parameters reference guides.`,
	Long:              "https://developers.google.com/admin-sdk/reports/reference/rest/v1/activities/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmreports.ListActivities(flags["userKey"].GetString(), flags["applicationName"].GetString(), flags["actorIpAddress"].GetString(), flags["customerId"].GetString(), flags["endTime"].GetString(), flags["eventName"].GetString(), flags["filters"].GetString(), flags["groupIdFilter"].GetString(), flags["orgUnitId"].GetString(), flags["startTime"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*reports.Activity{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing activities: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(activitiesCmd, activitiesListCmd, activityFlags)
}
