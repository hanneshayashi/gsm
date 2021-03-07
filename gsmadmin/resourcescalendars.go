/*
Package gsmadmin implements the Admin SDK APIs
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
package gsmadmin

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteCalendarResource deletes a calendar resource.
func DeleteCalendarResource(customer, calendarResourceID string) (bool, error) {
	srv := getResourcesCalendarsService()
	c := srv.Delete(customer, calendarResourceID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customer, calendarResourceID), func() error {
		return c.Do()
	})
	return result, err
}

// GetCalendarResource retrieves a calendar resource.
func GetCalendarResource(customer, calendarResourceID, fields string) (*admin.CalendarResource, error) {
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

// InsertCalendarResource inserts a calendar resource.
func InsertCalendarResource(customer, fields string, calendarResource *admin.CalendarResource) (*admin.CalendarResource, error) {
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

func listCalendarResources(c *admin.ResourcesCalendarsListCall, ch chan *admin.CalendarResource, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*admin.CalendarResources)
	for i := range r.Items {
		ch <- r.Items[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listCalendarResources(c, ch, errKey)
	}
	return err
}

// ListCalendarResources retrieves a list of calendar resources for an account.
func ListCalendarResources(customer, orderBy, query, fields string, cap int) (<-chan *admin.CalendarResource, <-chan error) {
	srv := getResourcesCalendarsService()
	c := srv.List(customer).MaxResults(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if orderBy != "" {
		c = c.OrderBy(orderBy)
	}
	if query != "" {
		c = c.Query(query)
	}
	ch := make(chan *admin.CalendarResource, cap)
	err := make(chan error, 1)
	go func() {
		e := listCalendarResources(c, ch, gsmhelpers.FormatErrorKey(customer))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// PatchCalendarResource updates a calendar resource. This method supports patch semantics.
func PatchCalendarResource(customer, calendarResourceID, fields string, calendar *admin.CalendarResource) (*admin.CalendarResource, error) {
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
