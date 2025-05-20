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
	"google.golang.org/api/calendar/v3"
)

// freeBusyCmd represents the freeBusy command
var freeBusyCmd = &cobra.Command{
	Use:               "freeBusy",
	Short:             "Query free/busy information (Part of Calendar API)",
	Long:              "Implements the API documented at https://developers.google.com/calendar/api/v3/reference/freebusy",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var freeBusyFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"timeMin": {
		AvailableFor: []string{"query"},
		Type:         "string",
		Description:  `The start of the interval for the query formatted as per RFC3339.`,
		Required:     []string{"query"},
	},
	"timeMax": {
		AvailableFor: []string{"query"},
		Type:         "string",
		Description:  `The end of the interval for the query formatted as per RFC3339.`,
		Required:     []string{"query"},
	},
	"timeZone": {
		AvailableFor: []string{"query"},
		Type:         "string",
		Description: `Time zone used in the response.
Optional. The default is UTC.`,
	},
	"groupExpansionMax": {
		AvailableFor: []string{"query"},
		Type:         "int64",
		Description: `Maximal number of calendar identifiers to be provided for a single group.
Optional. An error is returned for a group with more members than this value. Maximum value is 100.`,
	},
	"calendarExpansionMax": {
		AvailableFor: []string{"query"},
		Type:         "int64",
		Description: `Maximal number of calendars for which FreeBusy information is to be provided.
Optional. Maximum value is 50.`,
	},
	"id": {
		AvailableFor: []string{"query"},
		Type:         "stringSlice",
		Description:  `The identifier of a calendar or a group.`,
	},
	"fields": {
		AvailableFor: []string{"query"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(freeBusyCmd)
}

func mapToFreeBusyRequest(flags map[string]*gsmhelpers.Value) (*calendar.FreeBusyRequest, error) {
	freeBusyRequest := &calendar.FreeBusyRequest{}
	if flags["timeMin"].IsSet() {
		freeBusyRequest.TimeMin = flags["timeMin"].GetString()
		if freeBusyRequest.TimeMin == "" {
			freeBusyRequest.ForceSendFields = append(freeBusyRequest.ForceSendFields, "TimeMin")
		}
	}
	if flags["timeMax"].IsSet() {
		freeBusyRequest.TimeMax = flags["timeMax"].GetString()
		if freeBusyRequest.TimeMax == "" {
			freeBusyRequest.ForceSendFields = append(freeBusyRequest.ForceSendFields, "TimeMax")
		}
	}
	if flags["groupExpansionMax"].IsSet() {
		freeBusyRequest.GroupExpansionMax = flags["groupExpansionMax"].GetInt64()
		if freeBusyRequest.GroupExpansionMax == 0 {
			freeBusyRequest.ForceSendFields = append(freeBusyRequest.ForceSendFields, "GroupExpansionMax")
		}
	}
	if flags["calendarExpansionMax"].IsSet() {
		freeBusyRequest.CalendarExpansionMax = flags["calendarExpansionMax"].GetInt64()
		if freeBusyRequest.CalendarExpansionMax == 0 {
			freeBusyRequest.ForceSendFields = append(freeBusyRequest.ForceSendFields, "CalendarExpansionMax")
		}
	}
	if flags["id"].IsSet() {
		freeBusyRequest.Items = []*calendar.FreeBusyRequestItem{}
		ids := flags["id"].GetStringSlice()
		if len(ids) > 0 {
			for i := range ids {
				freeBusyRequest.Items = append(freeBusyRequest.Items, &calendar.FreeBusyRequestItem{Id: ids[i]})
			}
		} else {
			freeBusyRequest.ForceSendFields = append(freeBusyRequest.ForceSendFields, "Items")
		}
	}
	return freeBusyRequest, nil
}
