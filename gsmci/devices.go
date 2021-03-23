/*
Package gsmci implements the Cloud Identity API
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
package gsmci

import (
	"encoding/json"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// CancelDeviceWipe cancels an unfinished device wipe.
// This operation can be used to cancel device wipe in the gap between the wipe operation returning success and the device being wiped.
// This operation is possible when the device is in a "pending wipe" state.
// The device enters the "pending wipe" state when a wipe device command is issued, but has not yet been sent to the device.
// The cancel wipe will fail if the wipe command has already been issued to the device.
func CancelDeviceWipe(name, fields string, cancelWipeDeviceRequest *ci.GoogleAppsCloudidentityDevicesV1CancelWipeDeviceRequest) (map[string]interface{}, error) {
	srv := getDevicesService()
	c := srv.CancelWipe(name, cancelWipeDeviceRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(cancelWipeDeviceRequest.Customer, name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// CreateDevice creates a device.
// Only company-owned device may be created.
func CreateDevice(customer, fields string, device *ci.GoogleAppsCloudidentityDevicesV1Device) (map[string]interface{}, error) {
	srv := getDevicesService()
	c := srv.Create(device)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if customer != "" {
		c.Customer(customer)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, device.SerialNumber), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// DeleteDevice deletes the specified device.
func DeleteDevice(name, customer string) (map[string]interface{}, error) {
	srv := getDevicesService()
	c := srv.Delete(name)
	if customer != "" {
		c.Customer(customer)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetDevice retrieves the specified device.
func GetDevice(name, customer, fields string) (*ci.GoogleAppsCloudidentityDevicesV1Device, error) {
	srv := getDevicesService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if customer != "" {
		c.Customer(customer)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.GoogleAppsCloudidentityDevicesV1Device)
	return r, nil
}

func listDevices(c *ci.DevicesListCall, ch chan *ci.GoogleAppsCloudidentityDevicesV1Device, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*ci.GoogleAppsCloudidentityDevicesV1ListDevicesResponse)
	for i := range r.Devices {
		ch <- r.Devices[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listDevices(c, ch, errKey)
	}
	return err
}

// ListDevices lists/searches devices.
func ListDevices(customer, filter, orderBy, view, fields string, cap int) (<-chan *ci.GoogleAppsCloudidentityDevicesV1Device, <-chan error) {
	srv := getDevicesService()
	c := srv.List()
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if customer != "" {
		c.Customer(customer)
	}
	if orderBy != "" {
		c.OrderBy(orderBy)
	}
	if view != "" {
		c.View(view)
	}
	ch := make(chan *ci.GoogleAppsCloudidentityDevicesV1Device, cap)
	err := make(chan error, 1)
	go func() {
		e := listDevices(c, ch, gsmhelpers.FormatErrorKey(customer, filter))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// WipeDevice wipes all data on the specified device.
func WipeDevice(name, fields string, wipeDeviceRequest *ci.GoogleAppsCloudidentityDevicesV1WipeDeviceRequest) (map[string]interface{}, error) {
	srv := getDevicesService()
	c := srv.Wipe(name, wipeDeviceRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(wipeDeviceRequest.Customer, name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}