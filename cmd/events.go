/*
Package cmd contains the commands available to the end user
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
	"log"
	"strconv"
	"strings"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Manage events in users' calendars (Part of Calendar API)",
	Long: `This API only works in the user's context. Set the subject to the user's
email address to use this API!
https://developers.google.com/calendar/v3/reference/events`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var eventFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"eventId": {
		AvailableFor:   []string{"delete", "get", "import", "instances", "move", "patch"},
		Type:           "string",
		Description:    "Event identifier.",
		Required:       []string{"delete", "get", "import", "instances", "move", "patch"},
		ExcludeFromAll: true,
	},
	"calendarId": {
		AvailableFor: []string{"delete", "get", "import", "insert", "instances", "list", "move", "patch", "quickAdd"},
		Type:         "string",
		Description: `Calendar identifier. To retrieve calendar IDs call the calendarList.list method.
If you want to access the primary calendar of the currently logged in user, use the "primary" keyword.`,
		Defaults: map[string]interface{}{"delete": "primary", "get": "primary", "import": "primary", "insert": "primary", "instances": "primary", "list": "primary", "move": "primary", "patch": "primary", "quickAdd": "primary"},
	},
	"conferenceDataVersion": {
		AvailableFor: []string{"import", "insert", "patch"},
		Type:         "int64",
		Description: `Version number of conference data supported by the API client.
Version 0 assumes no conference data support and ignores conference data in the event's body.
Version 1 enables support for copying of ConferenceData as well as for creating new conferences using the createRequest field of conferenceData.`,
		Defaults: map[string]interface{}{"import": int64(1), "insert": int64(1), "patch": int64(1)},
	},
	"maxAttendees": {
		AvailableFor: []string{"get", "insert", "instances", "list", "patch"},
		Type:         "int64",
		Description: `The maximum number of attendees to include in the response.
If there are more than the specified number of attendees, only the participant is returned.`,
	},
	"sendUpdates": {
		AvailableFor: []string{"delete", "insert", "move", "patch", "quickAdd"},
		Type:         "string",
		Description: `Guests who should receive notifications about the event update (for example, title changes, etc.).
[all|externalOnly|none]
all           - Notifications are sent to all guests.
externalOnly  - Notifications are sent to non-Google Calendar guests only.
none          - No notifications are sent. This value should only be used for migration use cases (note that in most migration cases the import method should be used).`,
		Defaults: map[string]interface{}{"delete": "none", "insert": "none", "move": "none", "patch": "none", "quickAdd": "none"},
	},
	"supportsAttachments": {
		AvailableFor: []string{"import", "insert", "patch"},
		Type:         "bool",
		Description:  "Whether API client performing operation supports event attachments.",
	},
	"anyoneCanAddSelf": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  "Whether anyone can invite themselves to the event (currently works for Google+ events only).",
	},
	"fileUrl": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `URL link to the attachment.
For adding Google Drive file attachments use the same format as in alternateLink property of the Files resource in the Drive API.
Can be used multiple times to add more than one file.`,
	},
	"attendeesOmitted": {
		AvailableFor: []string{"patch"},
		Type:         "bool",
		Description: `Whether attendees may have been omitted from the event's representation.
When retrieving an event, this may be due to a restriction specified by the maxAttendee query parameter.
When updating an event, this can be used to only update the participant's response.`,
	},
	"attendees": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `Must be given in the following format: "--attendees "email=some.address@domain.com;resource=[true|false];optional=[true|false];responseStatus=accepted""
Can be used multiple times to invite more than one attendee.
If you batchPatch an event, rememder to specify ALL attendees (not just new ones)!`,
	},
	"colorId": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The color of the event.
This is an ID referring to an entry in the event section of the colors definition (see the colors endpoint).`,
	},
	"addConferenceData": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Whether to add a Meet conference to the event.`,
	},
	"description": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  "Description of the event. Can contain HTML.",
	},
	"endDate": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  "The date, in the format \"yyyy-mm-dd\", if this is an all-day event.",
	},
	"endDateTime": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The time, as a combined date-time value (formatted according to RFC3339).
A time zone offset is required unless a time zone is explicitly specified in timeZone.`,
	},
	"endTimeZone": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The time zone in which the time is specified.
(Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".)
For recurring events this field is required and specifies the time zone in which the recurrence is expanded.
For single events this field is optional and indicates a custom time zone for the event start/end.`,
	},
	"privateExtendedProperty": {
		AvailableFor: []string{"insert", "patch", "list"},
		Type:         "stringSlice",
		Description:  `Properties that are private to the copy of the event that appears on this calendar.`,
	},
	"sharedExtendedProperty": {
		AvailableFor: []string{"insert", "patch", "list"},
		Type:         "stringSlice",
		Description:  `Properties that are shared between copies of the event on other attendees' calendars.`,
	},
	"guestsCanInviteOthers": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Whether attendees other than the organizer can invite others to the event.`,
		Defaults:     map[string]interface{}{"import": true, "insert": true},
	},
	"guestsCanModify": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Whether attendees other than the organizer can modify the event.`,
	},
	"guestsCanSeeOtherGuests": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Whether attendees other than the organizer can see who the event's attendees are.`,
		Defaults:     map[string]interface{}{"import": true, "insert": true},
	},
	"id": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Opaque identifier of the event. When creating new single or recurring events, you can specify their IDs.
Provided IDs must follow these rules:
characters allowed in the ID are those used in base32hex encoding, i.e. lowercase letters a-v and digits 0-9, see section 3.1.2 in RFC2938
the length of the ID must be between 5 and 1024 characters
the ID must be unique per calendar
Due to the globally distributed nature of the system, we cannot guarantee that ID collisions will be detected at event creation time.
To minimize the risk of collisions we recommend using an established UUID algorithm such as one described in RFC4122.
If you do not specify an ID, it will be automatically generated by the server.

Note that the icalUID and the id are not identical and only one of them should be supplied at event creation time.
One difference in their semantics is that in recurring events, all occurrences of one event have different ids while they all share the same icalUIDs.`,
		ExcludeFromAll: true,
	},
	"location": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  "Geographic location of the event as free-form text.",
	},
	"recurrence": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `List of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545.
Note that DTSTART and DTEND lines are not allowed in this field; event start and end times are specified in the start and end fields.
This field is omitted for single events or instances of recurring events.`,
	},
	"reminderOverride": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `If the event doesn't use the default reminders, this lists the reminders specific to the event, or, if not set, indicates that no reminders are set for this event.
The maximum number of override reminders is 5.
Must be specified in the following format: "--reminderOverride "method=[email|popup];minutes=[0-40320]""
Can be used multiple times to specify more than one reminder override.`,
	},
	"useDefaultReminders": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Whether the default reminders of the calendar apply to the event.`,
		Defaults:     map[string]interface{}{"insert": true},
	},
	"sequence": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "int64",
		Description:  "Sequence number as per iCalendar.",
	},
	"startDate": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  "The date, in the format \"yyyy-mm-dd\", if this is an all-day event.",
	},
	"startDateTime": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The time, as a combined date-time value (formatted according to RFC3339).
A time zone offset is required unless a time zone is explicitly specified in timeZone.`,
	},
	"startTimeZone": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The time zone in which the time is specified.
(Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".)
For recurring events this field is required and specifies the time zone in which the recurrence is expanded.
For single events this field is optional and indicates a custom time zone for the event start/end.`,
	},
	"status": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Status of the event.
[confirmed|tentative|cancelled]
confirmed  - The event is confirmed. This is the default status.
tentative  - The event is tentatively confirmed.
cancelled  - The event is cancelled (deleted). The list method returns cancelled events only on incremental sync (when syncToken or updatedMin are specified) or if the showDeleted flag is set to true. The get method always returns them.
A cancelled status represents two different states depending on the event type:

Cancelled exceptions of an uncancelled recurring event indicate that this instance should no longer be presented to the user. Clients should store these events for the lifetime of the parent recurring event.
Cancelled exceptions are only guaranteed to have values for the id, recurringEventId and originalStartTime fields populated. The other fields might be empty.

All other cancelled events represent deleted events. Clients should remove their locally synced copies. Such cancelled events will eventually disappear, so do not rely on them being available indefinitely.
Deleted events are only guaranteed to have the id field populated.

On the organizer's calendar, cancelled events continue to expose event details (summary, location, etc.) so that they can be restored (undeleted). Similarly, the events to which the user was invited and that they manually removed continue to provide details. However, incremental sync requests with showDeleted set to false will not return these details.
If an event changes its organizer (for example via the move operation) and the original organizer is not on the attendee list, it will leave behind a cancelled event where only the id field is guaranteed to be populated.`,
		Defaults: map[string]interface{}{"import": "confirmed", "insert": "confirmed"},
	},
	"summary": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  "Title of the event.",
	},
	"transparency": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Whether the event blocks time on the calendar.
[opaque|transparent]
opaque       - Default value. The event does block time on the calendar. This is equivalent to setting Show me as to Busy in the Calendar UI.
transparent  - The event does not block time on the calendar. This is equivalent to setting Show me as to Available in the Calendar UI.`,
		Defaults: map[string]interface{}{"import": "opaque", "insert": "opaque"},
	},
	"visibility": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Visibility of the event.
[default|public|private|confidential]
default       - Uses the default visibility for events on the calendar. This is the default value.
public        - The event is public and event details are visible to all readers of the calendar.
private       - The event is private and only event attendees may view event details.
confidential  - The event is private. This value is provided for compatibility reasons.`,
		Defaults: map[string]interface{}{"import": "default", "insert": "default"},
	},
	"iCalUID": {
		AvailableFor:   []string{"list"},
		Type:           "string",
		Description:    `ICalUID sets the optional parameter "iCalUID": Specifies event ID in the iCalendar format to be included in the response.`,
		ExcludeFromAll: true,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The order of the events returned in the result. Optional. The default is an unspecified, stable order.
Acceptable values are:
startTime  - Order by the start date/time (ascending). This is only available when querying single events (i.e. the parameter singleEvents is True)
updated    - Order by last modification time (ascending).`,
	},
	"timeZone": {
		AvailableFor: []string{"get", "instances", "list"},
		Type:         "string",
		Description: `Time zone used in the response. Optional.
The default is the time zone of the calendar.`,
	},
	"updatedMin": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Lower bound for an event's last modification time (as a RFC3339 timestamp) to filter by.
When specified, entries deleted since this time will always be included regardless of showDeleted.`,
	},
	"timeMax": {
		AvailableFor: []string{"instances", "list"},
		Type:         "string",
		Description: `Upper bound (exclusive) for an event's start time to filter by.
Must be an RFC3339 timestamp with mandatory time zone offset, for example, 2011-06-03T10:00:00-07:00, 2011-06-03T10:00:00Z.
Milliseconds may be provided but are ignored. If timeMin is set, timeMax must be greater than timeMin.`,
	},
	"timeMin": {
		AvailableFor: []string{"instances", "list"},
		Type:         "string",
		Description: `Lower bound (exclusive) for an event's end time to filter by.
Must be an RFC3339 timestamp with mandatory time zone offset, for example, 2011-06-03T10:00:00-07:00, 2011-06-03T10:00:00Z.
Milliseconds may be provided but are ignored. If timeMax is set, timeMin must be smaller than timeMax.`,
	},
	"showDeleted": {
		AvailableFor: []string{"instances", "list"},
		Type:         "bool",
		Description: `Whether to include deleted events (with status equals "cancelled") in the result.
Cancelled instances of recurring events (but not the underlying recurring event) will still be included if showDeleted and singleEvents are both False.
If showDeleted and singleEvents are both True, only single instances of deleted events (but not the underlying recurring events) are returned.`,
	},
	"showHiddenInvitations": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether to include hidden invitations in the result.`,
	},
	"singleEvents": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether to expand recurring events into instances and only return single one-off events and instances of recurring events, but not the underlying recurring events themselves.`,
	},
	"q": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `Free text search terms to find events that match these terms in any field, except for extended properties.`,
	},
	"text": {
		AvailableFor: []string{"quickAdd"},
		Type:         "string",
		Description:  `The text describing the event to be created.`,
	},
	"fields": {
		AvailableFor: []string{"get", "import", "insert", "instances", "list", "move", "patch", "quickAdd"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
	"destination": {
		AvailableFor: []string{"move", "import"},
		Type:         "string",
		Description:  `Calendar identifier of the target calendar.`,
		Required:     []string{"move", "import"},
	},
	"originalStart": {
		AvailableFor: []string{"instances"},
		Type:         "string",
		Description:  `The original start time of the instance in the result.`,
	},
}
var eventFlagsALL = gsmhelpers.GetAllFlags(eventFlags)

func init() {
	rootCmd.AddCommand(eventsCmd)
}

func mapToEvent(flags map[string]*gsmhelpers.Value) (*calendar.Event, error) {
	event := &calendar.Event{}
	if flags["anyoneCanAddSelf"].IsSet() {
		event.AnyoneCanAddSelf = flags["anyoneCanAddSelf"].GetBool()
		if event.AnyoneCanAddSelf == false {
			event.ForceSendFields = append(event.ForceSendFields, "AnyoneCanAddSelf")
		}
	}
	if flags["fileUrl"].IsSet() {
		fileURLs := flags["fileUrl"].GetStringSlice()
		if len(fileURLs) > 0 {
			event.Attachments = []*calendar.EventAttachment{}
			for _, f := range fileURLs {
				event.Attachments = append(event.Attachments, &calendar.EventAttachment{FileUrl: f})
			}
		} else {
			event.ForceSendFields = append(event.ForceSendFields, "Attachments")
		}
	}
	if flags["attendees"].IsSet() {
		attendees := flags["attendees"].GetStringSlice()
		if len(attendees) > 0 {
			event.Attendees = []*calendar.EventAttendee{}
			for _, a := range attendees {
				m := gsmhelpers.FlagToMap(a)
				optional, err := strconv.ParseBool(m["optional"])
				if err != nil {
					log.Printf("Error parsing %v to bool: %v. Setting to false.", m["optional"], err)
				}
				resource, err := strconv.ParseBool(m["resource"])
				if err != nil {
					log.Printf("Error parsing %v to bool: %v. Setting to false.", m["resource"], err)
				}
				event.Attendees = append(event.Attendees, &calendar.EventAttendee{Email: m["email"], Optional: optional, Resource: resource, ResponseStatus: m["responseStatus"]})
			}
		} else {
			event.ForceSendFields = append(event.ForceSendFields, "Attendees")
		}
	}
	if flags["colorId"].IsSet() {
		event.ColorId = flags["colorId"].GetString()
		if event.ColorId == "" {
			event.ForceSendFields = append(event.ForceSendFields, "ColorId")
		}
	}
	addConferenceData := flags["addConferenceData"].GetBool()
	if addConferenceData {
		u, _ := uuid.NewRandom()
		event.ConferenceData = &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: u.String(),
			},
		}
	}
	if flags["description"].IsSet() {
		event.Description = flags["description"].GetString()
		if event.Description == "" {
			event.ForceSendFields = append(event.ForceSendFields, "Description")
		}
	}
	if flags["endDate"].IsSet() || flags["endDateTime"].IsSet() || flags["endTimeZone"].IsSet() {
		event.End = &calendar.EventDateTime{}
		event.End.Date = flags["endDate"].GetString()
		event.End.DateTime = flags["endDateTime"].GetString()
		event.End.TimeZone = flags["endTimeZone"].GetString()
	}
	if flags["privateExtendedProperty"].IsSet() || flags["sharedExtendedProperty"].IsSet() {
		event.ExtendedProperties = &calendar.EventExtendedProperties{}
		privateExtendedProperties := flags["privateExtendedProperty"].GetStringSlice()
		if len(privateExtendedProperties) > 0 {
			event.ExtendedProperties.Private = make(map[string]string)
			for _, pep := range privateExtendedProperties {
				split := strings.Split(pep, "=")
				event.ExtendedProperties.Private[split[0]] = split[1]
			}
		} else {
			event.ExtendedProperties.ForceSendFields = append(event.ExtendedProperties.ForceSendFields, "Private")
		}
		sharedExtendedProperties := flags["sharedExtendedProperty"].GetStringSlice()
		if len(sharedExtendedProperties) > 0 {
			event.ExtendedProperties.Shared = make(map[string]string)
			for _, sep := range sharedExtendedProperties {
				split := strings.Split(sep, "=")
				event.ExtendedProperties.Shared[split[0]] = split[1]
			}
		} else {
			event.ExtendedProperties.ForceSendFields = append(event.ExtendedProperties.ForceSendFields, "Shared")
		}
	}
	if flags["guestsCanInviteOthers"].IsSet() {
		guestsCanInviteOthers := flags["guestsCanInviteOthers"].GetBool()
		event.GuestsCanInviteOthers = &guestsCanInviteOthers
	}
	if flags["guestsCanModify"].IsSet() {
		event.GuestsCanModify = flags["guestsCanModify"].GetBool()
		if event.GuestsCanModify == false {
			event.ForceSendFields = append(event.ForceSendFields, "GuestsCanModify")
		}
	}
	if flags["guestsCanSeeOtherGuests"].IsSet() {
		guestsCanSeeOtherGuests := flags["guestsCanSeeOtherGuests"].GetBool()
		event.GuestsCanSeeOtherGuests = &guestsCanSeeOtherGuests
	}
	if flags["id"].IsSet() {
		event.Id = flags["id"].GetString()
		if event.Id == "" {
			event.ForceSendFields = append(event.ForceSendFields, "Id")
		}
	}
	if flags["location"].IsSet() {
		event.Location = flags["location"].GetString()
		if event.Location == "" {
			event.ForceSendFields = append(event.ForceSendFields, "Location")
		}
	}
	if flags["recurrence"].IsSet() {
		recurrences := flags["recurrence"].GetStringSlice()
		if len(recurrences) > 0 {
			event.Recurrence = append(event.Recurrence, recurrences...)
		} else {
			event.ForceSendFields = append(event.ForceSendFields, "Recurrence")
		}
	}
	if flags["reminderOverride"].IsSet() || flags["useDefaultReminders"].IsSet() {
		event.Reminders = &calendar.EventReminders{}
		if flags["reminderOverride"].IsSet() {
			reminderOverrides := flags["reminderOverride"].GetStringSlice()
			if len(reminderOverrides) > 0 {
				for _, ro := range reminderOverrides {
					m := gsmhelpers.FlagToMap(ro)
					if m["minutes"] == "" {
						continue
					}
					minutes, err := strconv.ParseInt(m["minutes"], 10, 64)
					if err != nil {
						return nil, err
					}
					event.Reminders.Overrides = append(event.Reminders.Overrides, &calendar.EventReminder{Method: m["method"], Minutes: minutes})
				}
			} else {
				event.Reminders.ForceSendFields = append(event.Reminders.ForceSendFields, "Overrides")
			}
		}
		if flags["useDefaultReminders"].IsSet() {
			event.Reminders.UseDefault = flags["useDefaultReminders"].GetBool()
			if !event.Reminders.UseDefault {
				event.Reminders.ForceSendFields = append(event.Reminders.ForceSendFields, "UseDefault")
			}
		}
	}
	if flags["sequence"].IsSet() {
		event.Sequence = flags["sequence"].GetInt64()
		if event.Sequence == 0 {
			event.ForceSendFields = append(event.ForceSendFields, "Sequence")
		}
	}
	if flags["startDate"].IsSet() || flags["startDateTime"].IsSet() || flags["startTimeZone"].IsSet() {
		event.Start = &calendar.EventDateTime{}
		event.Start.Date = flags["startDate"].GetString()
		event.Start.DateTime = flags["startDateTime"].GetString()
		event.Start.TimeZone = flags["startTimeZone"].GetString()

	}
	if flags["status"].IsSet() {
		event.Status = flags["status"].GetString()
		if event.Status == "" {
			event.ForceSendFields = append(event.ForceSendFields, "Status")
		}
	}
	if flags["summary"].IsSet() {
		event.Summary = flags["summary"].GetString()
		if event.Summary == "" {
			event.ForceSendFields = append(event.ForceSendFields, "Summary")
		}
	}
	if flags["transparency"].IsSet() {
		event.Transparency = flags["transparency"].GetString()
		if event.Transparency == "" {
			event.ForceSendFields = append(event.ForceSendFields, "Transparency")
		}
	}
	if flags["visibility"].IsSet() {
		event.Visibility = flags["visibility"].GetString()
		if event.Visibility == "" {
			event.ForceSendFields = append(event.ForceSendFields, "Visibility")
		}
	}
	return event, nil
}
