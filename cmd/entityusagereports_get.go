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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmreports"
	reports "google.golang.org/api/admin/reports/v1"

	"github.com/spf13/cobra"
)

// entityUsageReportsGetCmd represents the get command
var entityUsageReportsGetCmd = &cobra.Command{
	Use: "get",
	Short: `Retrieves a report which is a collection of properties and statistics for entities used by users within the account.
For more information, see the Entities Usage Report guide.
For more information about the entities report's parameters, see the Entities Usage parameters reference guides.`,
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/reports/reference/rest/v1/entityUsageReports/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmreports.GetEntityUsageReport(flags["entityType"].GetString(), flags["entityKey"].GetString(), flags["date"].GetString(), flags["customerId"].GetString(), flags["filters"].GetString(), flags["parameters"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*reports.UsageReport{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error getting Entity Usage Reports: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(entityUsageReportsCmd, entityUsageReportsGetCmd, entityUsageReportFlags)
}
