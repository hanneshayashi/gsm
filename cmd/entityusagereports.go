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

// entityUsageReportsCmd represents the entityUsageReports command
var entityUsageReportsCmd = &cobra.Command{
	Use:               "entityUsageReports",
	Short:             "Manage (get) Entity Usage Reports (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/reports/reference/rest/v1/entityUsageReports",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var entityUsageReportFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"entityType": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Represents the type of entity for the report.
Accepted values are:
GPLUS_COMMUNITIES  - Returns a report on Google+ communities.`,
		Defaults: map[string]interface{}{"get": "GPLUS_COMMUNITIES"},
	},
	"entityKey": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Represents the key of the object to filter the data with.
Accepted values are:
ALL         - Returns activity events for all users.
ENTITY_KEY  - Represents an app-specific identifier for the entity.
			  For details on how to obtain the entityKey for a particular entityType, see the https://developers.google.com/admin-sdk/reports/v1/reference/usage-ref-appendix-a/entities`,
		Defaults: map[string]interface{}{"get": "ALL"},
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
	"fields": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var entityUsageReportFlagsALL = gsmhelpers.GetAllFlags(entityUsageReportFlags)

func init() {
	rootCmd.AddCommand(entityUsageReportsCmd)
}
