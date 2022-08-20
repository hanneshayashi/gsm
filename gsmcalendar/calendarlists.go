/*
Copyright © 2020-2022 Hannes Hayashi

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
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// DeleteCalendarListEntry removes a calendar from the user's calendar list.
func DeleteCalendarListEntry(calendarID string) (bool, error) {
	srv := getCalendarListService()
	c := srv.Delete(calendarID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(calendarID), func() error {
		return c.Do()
	})
	return result, err
}

// GetCalendarListEntry returns a calendar from the user's calendar list.
func GetCalendarListEntry(calendarID, fields string) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
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
	r, ok := result.(*calendar.CalendarListEntry)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// InsertCalendarListEntry inserts an existing calendar into the user's calendar list.
func InsertCalendarListEntry(calendarListEntry *calendar.CalendarListEntry, colorRgbFormat bool, fields string) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.Insert(calendarListEntry).ColorRgbFormat(colorRgbFormat)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarListEntry.Id), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*calendar.CalendarListEntry)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListCalendarListEntries returns the calendars on the user's calendar list.
func ListCalendarListEntries(minAccessRole, fields string, showHidden, showDeleted bool, cap int) (<-chan *calendar.CalendarListEntry, <-chan error) {
	srv := getCalendarListService()
	c := srv.List().ShowDeleted(showDeleted).ShowHidden(showHidden).MaxResults(250)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if minAccessRole != "" {
		c = c.MinAccessRole(minAccessRole)
	}
	ch := make(chan *calendar.CalendarListEntry, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *calendar.CalendarList) error {
			for i := range response.Items {
				ch <- response.Items[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// PatchCalendarListEntry updates an existing calendar on the user's calendar list. This method supports patch semantics.
func PatchCalendarListEntry(calendarID, fields string, calendarListEntry *calendar.CalendarListEntry, colorRgbFormat bool) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.Patch(calendarID, calendarListEntry).ColorRgbFormat(colorRgbFormat)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*calendar.CalendarListEntry)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
