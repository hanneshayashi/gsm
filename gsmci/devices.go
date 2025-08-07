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

package gsmci

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// CancelDeviceWipe cancels an unfinished device wipe.
// This operation can be used to cancel device wipe in the gap between the wipe operation returning success and the device being wiped.
// This operation is possible when the device is in a "pending wipe" state.
// The device enters the "pending wipe" state when a wipe device command is issued, but has not yet been sent to the device.
// The cancel wipe will fail if the wipe command has already been issued to the device.
func CancelDeviceWipe(name, fields string, cancelWipeDeviceRequest *ci.GoogleAppsCloudidentityDevicesV1CancelWipeDeviceRequest) (*googleapi.RawMessage, error) {
	srv := getDevicesService()
	c := srv.CancelWipe(name, cancelWipeDeviceRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(cancelWipeDeviceRequest.Customer, name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return &r.Response, nil
}

// CreateDevice creates a device.
// Only company-owned device may be created.
func CreateDevice(customer, fields string, device *ci.GoogleAppsCloudidentityDevicesV1Device) (*googleapi.RawMessage, error) {
	srv := getDevicesService()
	c := srv.Create(device)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if customer != "" {
		c.Customer(customer)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, device.SerialNumber), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return &r.Response, nil
}

// DeleteDevice deletes the specified device.
func DeleteDevice(name, customer string) (*googleapi.RawMessage, error) {
	srv := getDevicesService()
	c := srv.Delete(name)
	if customer != "" {
		c.Customer(customer)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return &r.Response, nil
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.GoogleAppsCloudidentityDevicesV1Device)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
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
	if filter != "" {
		c.Filter(filter)
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
		e := c.Pages(context.Background(), func(response *ci.GoogleAppsCloudidentityDevicesV1ListDevicesResponse) error {
			for i := range response.Devices {
				ch <- response.Devices[i]
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

// WipeDevice wipes all data on the specified device.
func WipeDevice(name, fields string, wipeDeviceRequest *ci.GoogleAppsCloudidentityDevicesV1WipeDeviceRequest) (*googleapi.RawMessage, error) {
	srv := getDevicesService()
	c := srv.Wipe(name, wipeDeviceRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(wipeDeviceRequest.Customer, name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return &r.Response, nil
}
