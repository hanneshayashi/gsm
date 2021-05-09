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

package gsmadmin

import (
	"context"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// TakeActionOnChromeOsDevice takes an action that affects a Chrome OS Device. This includes deprovisioning, disabling, and re-enabling devices.
func TakeActionOnChromeOsDevice(customerID, deviceID string, action *admin.ChromeOsDeviceAction) (bool, error) {
	srv := getChromeosdevicesService()
	c := srv.Action(customerID, deviceID, action)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID, deviceID), func() error {
		return c.Do()
	})
	return result, err
}

// GetChromeOsDevice retrieves a Chrome OS device's properties.
func GetChromeOsDevice(customerID, deviceID, fields, projection string) (*admin.ChromeOsDevice, error) {
	srv := getChromeosdevicesService()
	c := srv.Get(customerID, deviceID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if projection != "" {
		c.Projection(projection)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, deviceID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.ChromeOsDevice)
	return r, nil
}

// ListChromeOsDevices retrieves a paginated list of Chrome OS devices within an account.
func ListChromeOsDevices(customerID, query, orgUnitPath, fields, projection string, cap int) (<-chan *admin.ChromeOsDevice, <-chan error) {
	srv := getChromeosdevicesService()
	c := srv.List(customerID).MaxResults(10000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if projection != "" {
		c.Projection(projection)
	}
	if query != "" {
		c = c.Query(query)
	}
	if orgUnitPath != "" {
		c = c.OrgUnitPath(orgUnitPath)
	}
	ch := make(chan *admin.ChromeOsDevice, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *admin.ChromeOsDevices) error {
			for i := range response.Chromeosdevices {
				ch <- response.Chromeosdevices[i]
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

// MoveChromeOSDevicesToOU moves or inserts multiple Chrome OS devices to an organizational unit. You can move up to 50 devices at once.
func MoveChromeOSDevicesToOU(customerID, orgUnitPath string, devicesToMove *admin.ChromeOsMoveDevicesToOu) (bool, error) {
	srv := getChromeosdevicesService()
	c := srv.MoveDevicesToOu(customerID, orgUnitPath, devicesToMove)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID), func() error {
		return c.Do()
	})
	return result, err
}

// PatchChromeOsDevice updates a device's updatable properties, such as annotatedUser, annotatedLocation, notes, orgUnitPath, or annotatedAssetId. This method supports patch semantics.
func PatchChromeOsDevice(customerID, deviceID, fields, projection string, chromeOsDevice *admin.ChromeOsDevice) (*admin.ChromeOsDevice, error) {
	srv := getChromeosdevicesService()
	c := srv.Patch(customerID, deviceID, chromeOsDevice)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if projection != "" {
		c.Projection(projection)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, deviceID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.ChromeOsDevice)
	return r, nil
}
