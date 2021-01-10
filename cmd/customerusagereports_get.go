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

// customerUsageReportsGetCmd represents the get command
var customerUsageReportsGetCmd = &cobra.Command{
	Use: "get",
	Short: `Retrieves a report which is a collection of properties and statistics for a specific customer's account.
For more information, see the Customers Usage Report guide.
For more information about the customer report's parameters, see the Customers Usage parameters reference guides.`,
	Long:              "https://developers.google.com/admin-sdk/reports/reference/rest/v1/customerUsageReports/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmreports.GetCustomerUsageReport(flags["date"].GetString(), flags["customerId"].GetString(), flags["parameters"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*reports.UsageReport{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error getting Customer Usage Reports: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(customerUsageReportsCmd, customerUsageReportsGetCmd, customerUsageReportFlags)
}
