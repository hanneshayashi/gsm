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

// TakeActionOnMobileDevice takes an action that affects a mobile device. For example, remotely wiping a device.
func TakeActionOnMobileDevice(customerID, resourceID string, action *admin.MobileDeviceAction) (bool, error) {
	srv := getMobiledevicesService()
	c := srv.Action(customerID, resourceID, action)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID, resourceID), func() error {
		return c.Do()
	})
	return result, err
}

// DeleteMobileDevice removes a mobile device.
func DeleteMobileDevice(customerID, resourceID string) (bool, error) {
	srv := getMobiledevicesService()
	c := srv.Delete(customerID, resourceID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID, resourceID), func() error {
		return c.Do()
	})
	return result, err
}

// GetMobileDevice retrieves a mobile device's properties.
func GetMobileDevice(customerID, resourceID, fields, projection string) (*admin.MobileDevice, error) {
	srv := getMobiledevicesService()
	c := srv.Get(customerID, resourceID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if projection != "" {
		c.Projection(projection)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, resourceID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.MobileDevice)
	return r, nil
}

func listMobileDevices(c *admin.MobiledevicesListCall, ch chan *admin.MobileDevice, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*admin.MobileDevices)
	for _, i := range r.Mobiledevices {
		ch <- i
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		err = listMobileDevices(c, ch, errKey)
	}
	return err
}

// ListMobileDevices retrieves a paginated list of all mobile devices for an account.
func ListMobileDevices(customerID, query, fields, projection, orderBy, sortOrder string, cap int) (<-chan *admin.MobileDevice, <-chan error) {
	srv := getMobiledevicesService()
	c := srv.List(customerID).MaxResults(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if query != "" {
		c = c.Query(query)
	}
	if projection != "" {
		c = c.Projection(projection)
	}
	if orderBy != "" {
		c = c.OrderBy(orderBy)
	}
	if sortOrder != "" {
		c = c.SortOrder(sortOrder)
	}
	ch := make(chan *admin.MobileDevice, cap)
	err := make(chan error, 1)
	go func() {
		e := listMobileDevices(c, ch, gsmhelpers.FormatErrorKey(customerID))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	return ch, err
}
