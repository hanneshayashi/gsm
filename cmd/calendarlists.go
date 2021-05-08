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
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// calendarListsCmd represents the calendarLists command
var calendarListsCmd = &cobra.Command{
	Use:               "calendarLists",
	Short:             "Manage entries in users' calendar list (Part of Calendar API)",
	Long:              "https://developers.google.com/calendar/v3/reference/calendarList",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var calendarListFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"calendarId": {
		AvailableFor: []string{"delete", "get", "list", "patch"},
		Type:         "string",
		Description: `Calendar identifier. To retrieve calendar IDs call the calendarList.list method.
If you want to access the primary calendar of the currently logged in user, use the "primary" keyword.`,
		Defaults:       map[string]interface{}{"delete": "primary", "get": "primary", "list": "primary", "patch": "primary"},
		ExcludeFromAll: true,
	},
	"colorRgbFormat": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description: `Whether to use the foregroundColor and backgroundColor fields to write the calendar colors (RGB).
If this feature is used, the index-based colorId field will be set to the best matching option automatically.`,
	},
	"id": {
		AvailableFor:   []string{"insert", "patch"},
		Type:           "string",
		Description:    `Identifier of the calendar.`,
		ExcludeFromAll: true,
	},
	"backgroundColor": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The main color of the calendar in the hexadecimal format "#0088aa".
This property supersedes the index-based colorId property.
To set or change this property, you need to specify colorRgbFormat=true in the parameters of the insert, update and patch methods.`,
	},
	"colorId": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The color of the calendar.
This is an ID referring to an entry in the calendar section of the colors definition (see the colors endpoint).
This property is superseded by the backgroundColor and foregroundColor properties and can be ignored when using these properties.`,
	},
	"defaultReminders": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The default reminders that the authenticated user has for this calendar.
Must be given in the form of '--defaultReminders "method=[popup|email];minutes[0-40320]
Where
email  - Reminders are sent via email.
popup  - Reminders are sent via a UI popup.`,
	},
	"foregroundColor": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The foreground color of the calendar in the hexadecimal format "#ffffff".
This property supersedes the index-based colorId property.
To set or change this property, you need to specify colorRgbFormat=true in the parameters of the insert, update and patch methods.`,
	},
	"hidden": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Whether the calendar has been hidden from the list.`,
	},
	"notificationsType": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The type of notification.
[eventCreation|eventChange|eventCancellation|eventResponse|agenda]
eventCreation      - Notification sent when a new event is put on the calendar.
eventChange        - Notification sent when an event is changed.
eventCancellation  - Notification sent when an event is cancelled.
eventResponse      - Notification sent when an attendee responds to the event invitation.
agenda             - An agenda with the events of the day (sent out in the morning).
Note that all notifications are sent via email ("method" is always set to "email" atomatically)`,
	},
	"selected": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Whether the calendar content shows up in the calendar UI`,
	},
	"summaryOverride": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `The summary that the authenticated user has set for this calendar.`,
	},
	"minAccessRole": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The minimum access role for the user in the returned entries.
Optional. The default is no restriction.
[freeBusyReader|owner|reader|writer]
freeBusyReader  - The user can read free/busy information.
owner           - The user can read and modify events and access control lists.
reader          - The user can read events that are not private.
writer          - The user can read and modify events.`,
	},
	"showDeleted": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether to include deleted calendar list entries in the result.`,
	},
	"showHidden": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether to show hidden entries.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var calendarListFlagsALL = gsmhelpers.GetAllFlags(calendarListFlags)

func init() {
	rootCmd.AddCommand(calendarListsCmd)
}
func mapToCalendarListEntry(flags map[string]*gsmhelpers.Value) (*calendar.CalendarListEntry, error) {
	calendarListEntry := &calendar.CalendarListEntry{}
	// string fields
	if flags["id"].IsSet() {
		calendarListEntry.Id = flags["id"].GetString()
		if calendarListEntry.Id == "" {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "Id")
		}
	}
	if flags["backgroundColor"].IsSet() {
		calendarListEntry.BackgroundColor = flags["backgroundColor"].GetString()
		if calendarListEntry.BackgroundColor == "" {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "BackgroundColor")
		}
	}
	if flags["colorId"].IsSet() {
		calendarListEntry.ColorId = flags["colorId"].GetString()
		if calendarListEntry.ColorId == "" {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "ColorId")
		}
	}
	if flags["foregroundColor"].IsSet() {
		calendarListEntry.ForegroundColor = flags["foregroundColor"].GetString()
		if calendarListEntry.ForegroundColor == "" {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "ForegroundColor")
		}
	}
	if flags["summaryOverride"].IsSet() {
		calendarListEntry.SummaryOverride = flags["summaryOverride"].GetString()
		if calendarListEntry.SummaryOverride == "" {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "SummaryOverride")
		}
	}
	// bool fields
	if flags["hidden"].IsSet() {
		calendarListEntry.Hidden = flags["hidden"].GetBool()
		if !calendarListEntry.Hidden {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "Hidden")
		}
	}
	if flags["selected"].IsSet() {
		calendarListEntry.Selected = flags["selected"].GetBool()
		if !calendarListEntry.Selected {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "Selected")
		}
	}
	// stringSlice fields
	if flags["defaultReminders"].IsSet() {
		calendarListEntry.DefaultReminders = []*calendar.EventReminder{}
		defaultReminders := flags["defaultReminders"].GetStringSlice()
		if len(defaultReminders) > 0 {
			for i := range defaultReminders {
				m := gsmhelpers.FlagToMap(defaultReminders[i])
				if m["minutes"] == "" {
					continue
				}
				minutes, err := strconv.ParseInt(m["minutes"], 10, 64)
				if err != nil {
					return nil, err
				}
				calendarListEntry.DefaultReminders = append(calendarListEntry.DefaultReminders, &calendar.EventReminder{Method: m["method"], Minutes: minutes})
			}
		} else {
			calendarListEntry.ForceSendFields = append(calendarListEntry.ForceSendFields, "DefaultReminders")
		}
	}
	if flags["notificationsType"].IsSet() {
		calendarListEntry.NotificationSettings = &calendar.CalendarListEntryNotificationSettings{}
		notifications := flags["notificationsType"].GetStringSlice()
		if len(notifications) > 0 {
			for i := range notifications {
				calendarListEntry.NotificationSettings.Notifications = append(calendarListEntry.NotificationSettings.Notifications, &calendar.CalendarNotification{Method: "email", Type: notifications[i]})
			}
		} else {
			calendarListEntry.NotificationSettings.ForceSendFields = append(calendarListEntry.NotificationSettings.ForceSendFields, "Notifications")
		}
	}
	return calendarListEntry, nil
}
