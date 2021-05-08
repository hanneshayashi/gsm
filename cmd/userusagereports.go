/*
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
	"time"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// userUsageReportsCmd represents the userUsageReports command
var userUsageReportsCmd = &cobra.Command{
	Use:               "userUsageReports",
	Short:             "Manage (get) User Usage Reports (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/reports/reference/rest/v1/userUsageReports",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var userUsageReportFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userKey": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Represents the profile ID or the user email for which the data should be filtered.
Can be "all" for all information, or userKey for a user's unique Workspace profile ID or their primary email address.`,
		Defaults: map[string]interface{}{"get": "all"},
	},
	"date": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Represents the date the usage occurred.
The timestamp is in the ISO 8601 format, yyyy-mm-dd.
We recommend you use your account's time zone for this.`,
		Defaults: map[string]interface{}{"get": time.Now().Format("2006-01-02")},
	},
	"customerId": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  `The unique ID of the customer to retrieve data for.`,
	},
	"filters": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The filters query string is a comma-separated list of an application's event parameters where the parameter's value is manipulated by a relational operator.
The filters query string includes the name of the application whose usage is returned in the report.
The application values for the Entities usage report include accounts, docs, and gmail.

Filters are in the form [application name]:[parameter name][relational operator][parameter value],....

In this example, the <> 'not equal to' operator is URL-encoded in the request's query string (%3C%3E):

GET
https://www.googleapis.com/admin/reports/v1/usage/gplus_communities/all/dates/2017-12-01
?parameters=gplus:community_name,gplus:num_total_members
&filters=gplus:num_total_members%3C%3E0 
The relational operators include:

== - 'equal to'.
<> - 'not equal to'. It is URL-encoded (%3C%3E).
< - 'less than'. It is URL-encoded (%3C).
<= - 'less than or equal to'. It is URL-encoded (%3C=).
> - 'greater than'. It is URL-encoded (%3E).
>= - 'greater than or equal to'. It is URL-encoded (%3E=).
Filters can only be applied to numeric parameters.`,
	},
	"orgUnitId": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `ID of the organizational unit to report on.
Activity records will be shown only for users who belong to the specified organizational unit.

Data before Dec 17, 2018 doesn't appear in the filtered results.`,
	},
	"parameters": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The parameters query string is a comma-separated list of event parameters that refine a report's results.
The parameter is associated with a specific application.
The application values for the Entities usage report are only gplus.
A parameter query string is in the CSV form of [app_name1:param_name1], [app_name2:param_name2]....

Note: The API doesn't accept multiple values of a parameter.
If a particular parameter is supplied more than once in the API request, the API only accepts the last value of that request parameter.
In addition, if an invalid request parameter is supplied in the API request, the API ignores that request parameter and returns the response corresponding to the remaining valid request parameters.

An example of an invalid request parameter is one that does not belong to the application. If no parameters are requested, all parameters are returned.`,
	},
	"groupIdFilter": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Comma separated group ids (obfuscated) on which user activities are filtered, i.e, the response will contain activities for only those users that are a part of at least one of the group ids mentioned here.
Format: "id:abc123,id:xyz456"`,
	},
	"fields": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var userUsageReportFlagsALL = gsmhelpers.GetAllFlags(userUsageReportFlags)

func init() {
	rootCmd.AddCommand(userUsageReportsCmd)
}
