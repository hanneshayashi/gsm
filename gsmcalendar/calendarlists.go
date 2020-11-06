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
package gsmcalendar

import (
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// DeleteCalendarListEntry removes a calendar from the user's calendar list.
func DeleteCalendarListEntry(calendarID string) (bool, error) {
	srv := getCalendarListService()
	err := srv.Delete(calendarID).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetCalendarListEntry returns a calendar from the user's calendar list.
func GetCalendarListEntry(calendarID, fields string) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.Get(calendarID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertCalendarListEntry inserts an existing calendar into the user's calendar list.
func InsertCalendarListEntry(calendarListEntry *calendar.CalendarListEntry, colorRgbFormat bool, fields string) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.Insert(calendarListEntry).ColorRgbFormat(colorRgbFormat)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// ListCalendarListEntries returns the calendars on the user's calendar list.
func ListCalendarListEntries(minAccessRole, fields string) ([]*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.List()
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if minAccessRole != "" {
		c = c.MinAccessRole(minAccessRole)
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.Items, nil
}

// PatchCalendarListEntry updates an existing calendar on the user's calendar list. This method supports patch semantics.
func PatchCalendarListEntry(calendarID, fields string, calendarListEntry *calendar.CalendarListEntry, colorRgbFormat bool) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.Patch(calendarID, calendarListEntry).ColorRgbFormat(colorRgbFormat)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
