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

	"github.com/spf13/cobra"
)

// activitiesCmd represents the activities command
var activitiesCmd = &cobra.Command{
	Use:               "activities",
	Short:             "Manage (list) activities (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/reports/reference/rest/v1/activities",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var activityFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userKey": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `Represents the profile ID or the user email for which the data should be filtered. Can be all for all information, or userKey for a user's unique Workspace profile ID or their primary email address.`,
		Defaults:     map[string]any{"list": "all"},
	},
	"applicationName": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Application name for which the events are to be retrieved.
The following values are accepted:
ACCESS_TRANSPARENCY   - The Workspace Access Transparency activity reports return information about different types of Access Transparency activity events.
ADMIN                 - The Admin console application's activity reports return account information about different types of administrator activity events.
CALENDAR              - The Workspace Calendar application's activity reports return information about various Calendar activity events.
CHAT                  - The Chat activity reports return information about various Chat activity events.
DRIVE                 - The Google Drive application's activity reports return information about various Google Drive activity events. The Drive activity report is only available for Workspace Business customers.
GCP                   - The Google Cloud Platform application's activity reports return information about various GCP activity events.
GPLUS                 - The Google+ application's activity reports return information about various Google+ activity events.
GROUPS                - The Google Groups application's activity reports return information about various Groups activity events.
GROUPS_ENTERPRISE     - The Enterprise Groups activity reports return information about various Enterprise group activity events.
JAMBOARD              - The Jamboard activity reports return information about various Jamboard activity events.
LOGIN                 - The Login application's activity reports return account information about different types of Login activity events.
MEET                  - The Meet Audit activity report return information about different types of Meet Audit activity events.
MOBILE                - The Mobile Audit activity report return information about different types of Mobile Audit activity events.
RULES                 - The Rules activity report return information about different types of Rules activity events.
SAML                  - The SAML activity report return information about different types of SAML activity events.
TOKEN                 - The Token application's activity reports return account information about different types of Token activity events.
USER_ACCOUNTS         - The User Accounts application's activity reports return account information about different types of User Accounts activity events.
CONTEXT_AWARE_ACCESS  - The Context-aware access activity reports return information about users' access denied events due to Context-aware access rules.
CHROME                - The Chrome activity reports return information about unsafe events reported in the context of the WebProtect features of BeyondCorp.
DATA_STUDIO           - The Data Studio activity reports return information about various types of Data Studio activity events.`,
		Required: []string{"list"},
	},
	"actorIpAddress": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The Internet Protocol (IP) Address of host where the event was performed.
This is an additional way to filter a report's summary using the IP address of the user whose activity is being reported.
This IP address may or may not reflect the user's physical location.
For example, the IP address can be the user's proxy server's address or a virtual private network (VPN) address.
This parameter supports both IPv4 and IPv6 address versions.`,
	},
	"customerId": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `The unique ID of the customer to retrieve data for.`,
	},
	"endTime": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Sets the end of the range of time shown in the report.
The date is in the RFC 3339 format, for example 2010-10-28T10:26:35.000Z.
The default value is the approximate time of the API request.
An API report has three basic time concepts:
  - Date of the API's request for a report: When the API created and retrieved the report.
  - Report's start time: The beginning of the timespan shown in the report.
    The startTime must be before the endTime (if specified) and the current time when the request is made, or the API returns an error.
  - Report's end time: The end of the timespan shown in the report.
	For example, the timespan of events summarized in a report can start in April and end in May.
    The report itself can be requested in August.
  - If the endTime is not specified, the report returns all activities from the startTime until the current time or the most recent 180 days if the startTime is more than 180 days in the past.`,
	},
	"eventName": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The name of the event being queried by the API.
Each eventName is related to a specific Workspace service or feature which the API organizes into types of events.
An example is the Google Calendar events in the Admin console application's reports.
The Calendar Settings type structure has all of the Calendar eventName activities reported by the API.
When an administrator changes a Calendar setting, the API reports this activity in the Calendar Settings type and eventName parameters.
For more information about eventName query strings and parameters, see the list of event names for various applications above in applicationName.`,
	},
	"filters": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The filters query string is a comma-separated list.
The list is composed of event parameters that are manipulated by relational operators.
Event parameters are in the form [parameter1 name][relational operator][parameter1 value],[parameter2 name][relational operator][parameter2 value],...

These event parameters are associated with a specific eventName. An empty report is returned if the filtered request's parameter does not belong to the eventName. For more information about eventName parameters, see the list of event names for various applications above in applicationName.

In the following Admin Activity example, the <> operator is URL-encoded in the request's query string (%3C%3E):

GET...&eventName=CHANGE_CALENDAR_SETTING
&filters=NEW_VALUE%3C%3EREAD_ONLY_ACCESS

In the following Drive example, the list can be a view or edit event's doc_id parameter with a value that is manipulated by an 'equal to' (==) or 'not equal to' (<>) relational operator. In the first example, the report returns each edited document's doc_id. In the second example, the report returns each viewed document's doc_id that equals the value 12345 and does not return any viewed document's which have a doc_id value of 98765. The <> operator is URL-encoded in the request's query string (%3C%3E):

GET...&eventName=edit&filters=doc_id
GET...&eventName=view&filters=doc_id==12345,doc_id%3C%3E98765
The relational operators include:

== - 'equal to'.
<> - 'not equal to'. It is URL-encoded (%3C%3E).
< - 'less than'. It is URL-encoded (%3C).
<= - 'less than or equal to'. It is URL-encoded (%3C=).
> - 'greater than'. It is URL-encoded (%3E).
>= - 'greater than or equal to'. It is URL-encoded (%3E=).

Note: The API doesn't accept multiple values of a parameter.
If a particular parameter is supplied more than once in the API request, the API only accepts the last value of that request parameter.
In addition, if an invalid request parameter is supplied in the API request, the API ignores that request parameter and returns the response corresponding to the remaining valid request parameters.
If no parameters are requested, all parameters are returned.`,
	},
	"orgUnitId": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `ID of the organizational unit to report on.
Activity records will be shown only for users who belong to the specified organizational unit.

Data before Dec 17, 2018 doesn't appear in the filtered results.`,
	},
	"startTime": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Sets the beginning of the range of time shown in the report.
The date is in the RFC 3339 format, for example 2010-10-28T10:26:35.000Z.
The report returns all activities from startTime until endTime.
The startTime must be before the endTime (if specified) and the current time when the request is made, or the API returns an error.`,
	},
	"groupIdFilter": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Comma separated group ids (obfuscated) on which user activities are filtered, i.e, the response will contain activities for only those users that are a part of at least one of the group ids mentioned here.
Format: "id:abc123,id:xyz456"`,
	},
	"fields": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var activityFlagsALL = gsmhelpers.GetAllFlags(activityFlags)

func init() {
	rootCmd.AddCommand(activitiesCmd)
}
