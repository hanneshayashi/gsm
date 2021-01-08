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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.CalendarListEntry)
	return r, nil
}

// InsertCalendarListEntry inserts an existing calendar into the user's calendar list.
func InsertCalendarListEntry(calendarListEntry *calendar.CalendarListEntry, colorRgbFormat bool, fields string) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.Insert(calendarListEntry).ColorRgbFormat(colorRgbFormat)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarListEntry.Id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.CalendarListEntry)
	return r, nil
}

func listCalendarListEntries(c *calendar.CalendarListListCall, ch chan *calendar.CalendarListEntry, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*calendar.CalendarList)
	for _, i := range r.Items {
		ch <- i
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		err = listCalendarListEntries(c, ch, errKey)
	}
	return err
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
		e := listCalendarListEntries(c, ch, gsmhelpers.FormatErrorKey("List calendar list entries"))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	return ch, err
}

// PatchCalendarListEntry updates an existing calendar on the user's calendar list. This method supports patch semantics.
func PatchCalendarListEntry(calendarID, fields string, calendarListEntry *calendar.CalendarListEntry, colorRgbFormat bool) (*calendar.CalendarListEntry, error) {
	srv := getCalendarListService()
	c := srv.Patch(calendarID, calendarListEntry).ColorRgbFormat(colorRgbFormat)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.CalendarListEntry)
	return r, nil
}
