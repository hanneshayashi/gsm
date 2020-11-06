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
package gsmadmin

import (
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// TakeActionOnChromeOsDevice takes an action that affects a Chrome OS Device. This includes deprovisioning, disabling, and re-enabling devices.
func TakeActionOnChromeOsDevice(customerID, deviceID string, action *admin.ChromeOsDeviceAction) (bool, error) {
	srv := getChromeosdevicesService()
	err := srv.Action(customerID, deviceID, action).Do()
	if err != nil {
		return false, err
	}
	return true, nil
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
	r, err := c.Do()
	return r, err
}

func makeListChromeOsDevicesCallAndAppend(c *admin.ChromeosdevicesListCall, chromeosDevices []*admin.ChromeOsDevice) ([]*admin.ChromeOsDevice, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, c := range r.Chromeosdevices {
		chromeosDevices = append(chromeosDevices, c)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		chromeosDevices, err = makeListChromeOsDevicesCallAndAppend(c, chromeosDevices)
	}
	return chromeosDevices, err
}

// ListChromeOsDevices retrieves a paginated list of Chrome OS devices within an account.
func ListChromeOsDevices(customerID, query, orgUnitPath, fields, projection string) ([]*admin.ChromeOsDevice, error) {
	srv := getChromeosdevicesService()
	c := srv.List(customerID)
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
	var chromeOsDevices []*admin.ChromeOsDevice
	chromeOsDevices, err := makeListChromeOsDevicesCallAndAppend(c, chromeOsDevices)
	return chromeOsDevices, err
}

// MoveChromeOSDevicesToOU moves or inserts multiple Chrome OS devices to an organizational unit. You can move up to 50 devices at once.
func MoveChromeOSDevicesToOU(customerID, orgUnitPath string, devicesToMove *admin.ChromeOsMoveDevicesToOu) (bool, error) {
	srv := getChromeosdevicesService()
	err := srv.MoveDevicesToOu(customerID, orgUnitPath, devicesToMove).Do()
	if err != nil {
		return false, err
	}
	return true, nil
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
	r, err := c.Do()
	return r, err
}
