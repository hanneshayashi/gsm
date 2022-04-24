/*
Copyright © 2020-2022 Hannes Hayashi

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
	"time"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// customerUsageReportsCmd represents the customerUsageReports command
var customerUsageReportsCmd = &cobra.Command{
	Use:               "customerUsageReports",
	Short:             "Manage (get) Customer Usage Reports (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/reports/reference/rest/v1/customerUsageReports",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var customerUsageReportFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"date": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Represents the date the usage occurred.
The timestamp is in the ISO 8601 format, yyyy-mm-dd.
We recommend you use your account's time zone for this.`,
		Defaults: map[string]any{"get": time.Now().Format("2006-01-02")},
	},
	"customerId": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  `The unique ID of the customer to retrieve data for.`,
	},
	"parameters": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The parameters query string is a comma-separated list of event parameters that refine a report's results.
The parameter is associated with a specific application.
The application values for the Customers usage report include accounts, app_maker, apps_scripts, calendar, classroom, cros, docs, gmail, gplus, device_management, meet, and sites.
A parameters query string is in the CSV form of app_name1:param_name1, app_name2:param_name2.

Note: The API doesn't accept multiple values of a parameter.
If a particular parameter is supplied more than once in the API request, the API only accepts the last value of that request parameter.
In addition, if an invalid request parameter is supplied in the API request, the API ignores that request parameter and returns the response corresponding to the remaining valid request parameters.
An example of an invalid request parameter is one that does not belong to the application.
If no parameters are requested, all parameters are returned.`,
	},
	"fields": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var customerUsageReportFlagsALL = gsmhelpers.GetAllFlags(customerUsageReportFlags)

func init() {
	rootCmd.AddCommand(customerUsageReportsCmd)
}
