/*
Copyright © 2020-2025 Hannes Hayashi

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
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// ClearCalendar clears a primary calendar.
// This operation deletes all events associated with the primary calendar of an account.
func ClearCalendar(calendarID string) (bool, error) {
	srv := getCalendarsService()
	c := srv.Clear(calendarID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(calendarID), func() error {
		return c.Do()
	})
	return result, err
}

// DeleteCalendar deletes a secondary calendar.
// Use calendars.clear for clearing all events on primary calendars.
func DeleteCalendar(calendarID string) (bool, error) {
	srv := getCalendarsService()
	c := srv.Delete(calendarID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(calendarID), func() error {
		return c.Do()
	})
	return result, err
}

// GetCalendar returns metadata for a calendar.
func GetCalendar(calendarID, fields string) (*calendar.Calendar, error) {
	srv := getCalendarsService()
	c := srv.Get(calendarID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*calendar.Calendar)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// InsertCalendar creates a secondary calendar.
func InsertCalendar(cal *calendar.Calendar, fields string) (*calendar.Calendar, error) {
	srv := getCalendarsService()
	c := srv.Insert(cal)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(cal.Summary), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*calendar.Calendar)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// PatchCalendar updates metadata for a calendar.
// This method supports patch semantics.
func PatchCalendar(calendarID, fields string, cal *calendar.Calendar) (*calendar.Calendar, error) {
	srv := getCalendarsService()
	c := srv.Patch(calendarID, cal)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*calendar.Calendar)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
