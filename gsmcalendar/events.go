/*
Package gsmcalendar implements the Calendar API
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
package gsmcalendar

import (
	"gsm/gsmhelpers"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// DeleteEvent deletes an event.
func DeleteEvent(calendarID, eventID, sendUpdates string) (bool, error) {
	srv := getEventsService()
	c := srv.Delete(calendarID, eventID).SendUpdates(sendUpdates)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(calendarID, eventID), func() error {
		return c.Do()
	})
	return result, err
}

// GetEvent returns an event.
func GetEvent(calendarID, eventID, timeZone, fields string, maxAttendees int64) (*calendar.Event, error) {
	srv := getEventsService()
	c := srv.Get(calendarID, eventID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if timeZone != "" {
		c = c.TimeZone(timeZone)
	}
	if maxAttendees != 0 {
		c = c.MaxAttendees(maxAttendees)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, eventID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Event)
	return r, nil
}

// ImportEvent imports an event. This operation is used to add a private copy of an existing event to a calendar.
func ImportEvent(calendarID, fields string, event *calendar.Event, conferenceDataVersion int64, supportsAttachments bool) (*calendar.Event, error) {
	srv := getEventsService()
	c := srv.Import(calendarID, event).ConferenceDataVersion(conferenceDataVersion).SupportsAttachments(supportsAttachments)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, event.Id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Event)
	return r, nil
}

// InsertEvent creates an event.
func InsertEvent(calendarID, sendUpdates, fields string, event *calendar.Event, conferenceDataVersion, maxAttendees int64, supportsAttachments bool) (*calendar.Event, error) {
	srv := getEventsService()
	c := srv.Insert(calendarID, event).ConferenceDataVersion(conferenceDataVersion).SupportsAttachments(supportsAttachments).ConferenceDataVersion(conferenceDataVersion).SendUpdates(sendUpdates)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if maxAttendees != 0 {
		c = c.MaxAttendees(maxAttendees)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, event.Id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Event)
	return r, nil
}

func makeListInstancesCallAndAppend(c *calendar.EventsInstancesCall, eventInstances []*calendar.Event, errKey string) ([]*calendar.Event, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Events)
	eventInstances = append(eventInstances, r.Items...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		eventInstances, err = makeListInstancesCallAndAppend(c, eventInstances, errKey)
	}
	return eventInstances, err
}

// ListInstances returns instances of the specified recurring event.
func ListInstances(calendarID, eventID, originalStart, timeZone, timeMax, timeMin, fields string, maxAttendees int64, showDeleted bool) ([]*calendar.Event, error) {
	srv := getEventsService()
	c := srv.Instances(calendarID, eventID).ShowDeleted(showDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if timeZone != "" {
		c = c.TimeZone(timeZone)
	}
	if timeMax != "" {
		c = c.TimeMax(timeMax)
	}
	if timeMin != "" {
		c = c.TimeMin(timeMin)
	}
	if maxAttendees != 0 {
		c = c.MaxAttendees(maxAttendees)
	}
	var eventInstances []*calendar.Event
	eventInstances, err := makeListInstancesCallAndAppend(c, eventInstances, gsmhelpers.FormatErrorKey(calendarID, eventID))
	return eventInstances, err
}

func makeListEventsCallAndAppend(c *calendar.EventsListCall, events []*calendar.Event, errKey string) ([]*calendar.Event, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Events)
	events = append(events, r.Items...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		events, err = makeListEventsCallAndAppend(c, events, errKey)
	}
	return events, err
}

// ListEvents returns events on the specified calendar.
func ListEvents(calendarID, iCalUID, orderBy, q, timeZone, timeMax, timeMin, updatedMin, fields string, privateExtendedProperties, sharedExtendedProperties []string, maxAttendees int64, showDeleted, showHiddenInvitations, singleEvents bool) ([]*calendar.Event, error) {
	srv := getEventsService()
	c := srv.List(calendarID).ShowDeleted(showDeleted).ShowHiddenInvitations(showHiddenInvitations).SingleEvents(singleEvents)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if timeZone != "" {
		c = c.TimeZone(timeZone)
	}
	if timeMax != "" {
		c = c.TimeMax(timeMax)
	}
	if timeMin != "" {
		c = c.TimeMin(timeMin)
	}
	if maxAttendees != 0 {
		c = c.MaxAttendees(maxAttendees)
	}
	if iCalUID != "" {
		c = c.ICalUID(iCalUID)
	}
	if orderBy != "" {
		c = c.OrderBy(orderBy)
	}
	if q != "" {
		c = c.Q(q)
	}
	for _, pep := range privateExtendedProperties {
		c = c.PrivateExtendedProperty(pep)
	}
	for _, sep := range sharedExtendedProperties {
		c = c.SharedExtendedProperty(sep)
	}
	if updatedMin != "" {
		c = c.UpdatedMin(updatedMin)
	}
	var events []*calendar.Event
	events, err := makeListEventsCallAndAppend(c, events, gsmhelpers.FormatErrorKey(calendarID))
	return events, err
}

// MoveEvent moves an event to another calendar, i.e. changes an event's organizer.
func MoveEvent(calendarID, eventID, destination, sendUpdates, fields string) (*calendar.Event, error) {
	srv := getEventsService()
	c := srv.Move(calendarID, eventID, destination).SendUpdates(sendUpdates)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, eventID, destination), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Event)
	return r, nil
}

// PatchEvent Updates an event. This method supports patch semantics.
// The field values you specify replace the existing values. Fields that you don’t specify in the request remain unchanged.
// Array fields, if specified, overwrite the existing arrays; this discards any previous array elements.
func PatchEvent(calendarID, eventID, sendUpdates, fields string, event *calendar.Event, conferenceDataVersion, maxAttendees int64, supportsAttachments bool) (*calendar.Event, error) {
	srv := getEventsService()
	c := srv.Patch(calendarID, eventID, event).SendUpdates(sendUpdates).SupportsAttachments(supportsAttachments).ConferenceDataVersion(conferenceDataVersion)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if maxAttendees != 0 {
		c = c.MaxAttendees(maxAttendees)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, eventID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Event)
	return r, nil
}

// QuickAddEvent creates an event based on a simple text string.
func QuickAddEvent(calendarID, text, sendUpdates, fields string) (*calendar.Event, error) {
	srv := getEventsService()
	c := srv.QuickAdd(calendarID, text).SendUpdates(sendUpdates)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Event)
	return r, nil
}
