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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// calendarsCmd represents the calendars command
var calendarsCmd = &cobra.Command{
	Use:               "calendars",
	Short:             "Manage users' calendars (Part of Calendar API)",
	Long:              "https://developers.google.com/calendar/v3/reference/calendars",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var calendarFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"calendarId": {
		AvailableFor: []string{"clear", "delete", "get", "patch"},
		Type:         "string",
		Description: `Calendar identifier. To retrieve calendar IDs call the calendarList.list method.
If you want to access the primary calendar of the currently logged in user, use the "primary" keyword.`,
		Required:       []string{"clear", "delete", "patch"},
		Defaults:       map[string]interface{}{"get": "primary"},
		ExcludeFromAll: true,
	},
	"summary": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Title of the calendar.`,
		Required:     []string{"insert"},
	},
	"description": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Description of the calendar.`,
	},
	"location": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Geographic location of the calendar as free-form text.`,
	},
	"timeZone": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The time zone of the calendar.
(Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich").`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var calendarFlagsALL = gsmhelpers.GetAllFlags(calendarFlags)

func init() {
	rootCmd.AddCommand(calendarsCmd)
}

func mapToCalendar(flags map[string]*gsmhelpers.Value) (*calendar.Calendar, error) {
	calendar := &calendar.Calendar{}
	if flags["summary"].IsSet() {
		calendar.Summary = flags["summary"].GetString()
		if calendar.Summary == "" {
			calendar.ForceSendFields = append(calendar.ForceSendFields, "Summary")
		}
	}
	if flags["description"].IsSet() {
		calendar.Description = flags["description"].GetString()
		if calendar.Description == "" {
			calendar.ForceSendFields = append(calendar.ForceSendFields, "Description")
		}
	}
	if flags["location"].IsSet() {
		calendar.Location = flags["location"].GetString()
		if calendar.Location == "" {
			calendar.ForceSendFields = append(calendar.ForceSendFields, "Location")
		}
	}
	if flags["timeZone"].IsSet() {
		calendar.TimeZone = flags["timeZone"].GetString()
		if calendar.TimeZone == "" {
			calendar.ForceSendFields = append(calendar.ForceSendFields, "TimeZone")
		}
	}
	return calendar, nil
}
