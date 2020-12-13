/*
Package gsmadmin implements the Admin SDK APIs
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
package gsmadmin

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteResourcesCalendar deletes a calendar resource.
func DeleteResourcesCalendar(customer, calendarResourceID string) (bool, error) {
	srv := getResourcesCalendarsService()
	c := srv.Delete(customer, calendarResourceID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customer, calendarResourceID), func() error {
		return c.Do()
	})
	return result, err
}

// GetResourcesCalendar retrieves a calendar resource.
func GetResourcesCalendar(customer, calendarResourceID, fields string) (*admin.CalendarResource, error) {
	srv := getResourcesCalendarsService()
	c := srv.Get(customer, calendarResourceID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, calendarResourceID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.CalendarResource)
	return r, nil
}

// InsertResourcesCalendar inserts a calendar resource.
func InsertResourcesCalendar(customer, fields string, calendarResource *admin.CalendarResource) (*admin.CalendarResource, error) {
	srv := getResourcesCalendarsService()
	c := srv.Insert(customer, calendarResource)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, calendarResource.ResourceName), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.CalendarResource)
	return r, nil
}

func makeListResourceCalendarsCallAndAppend(c *admin.ResourcesCalendarsListCall, calendars []*admin.CalendarResource, errKey string) ([]*admin.CalendarResource, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.CalendarResources)
	calendars = append(calendars, r.Items...)
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		calendars, err = makeListResourceCalendarsCallAndAppend(c, calendars, errKey)
	}
	return calendars, err
}

// ListResourcesCalendars retrieves a list of calendar resources for an account.
func ListResourcesCalendars(customer, orderBy, query, fields string) ([]*admin.CalendarResource, error) {
	srv := getResourcesCalendarsService()
	c := srv.List(customer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if orderBy != "" {
		c = c.OrderBy(orderBy)
	}
	if query != "" {
		c = c.Query(query)
	}
	var calendars []*admin.CalendarResource
	calendars, err := makeListResourceCalendarsCallAndAppend(c, calendars, gsmhelpers.FormatErrorKey(customer))
	return calendars, err
}

// PatchResourcesCalendar updates a calendar resource. This method supports patch semantics.
func PatchResourcesCalendar(customer, calendarResourceID, fields string, calendar *admin.CalendarResource) (*admin.CalendarResource, error) {
	srv := getResourcesCalendarsService()
	c := srv.Patch(customer, calendarResourceID, calendar)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, calendarResourceID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.CalendarResource)
	return r, nil
}
