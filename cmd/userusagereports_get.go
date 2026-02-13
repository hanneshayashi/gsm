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
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmreports"

	"github.com/spf13/cobra"
)

// userUsageReportsGetCmd represents the get command
var userUsageReportsGetCmd = &cobra.Command{
	Use: "get",
	Short: `Retrieves a report which is a collection of properties and statistics for a set of users with the account.
For more information, see the User Usage Report guide.
For more information about the user report's parameters, see the Users Usage parameters reference guides.`,
	Long:              "Implements the API documented at https://developers.google.com/workspace/admin/reports/reference/rest/v1/userUsageReport/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		gsmhelpers.StreamOrCollect(gsmreports.GetUserUsageReport(flags["userKey"].GetString(), flags["date"].GetString(), flags["customerId"].GetString(), flags["filters"].GetString(), flags["orgUnitId"].GetString(), flags["parameters"].GetString(), flags["groupIdFilter"].GetString(), flags["fields"].GetString()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(userUsageReportsCmd, userUsageReportsGetCmd, userUsageReportFlags)
}
