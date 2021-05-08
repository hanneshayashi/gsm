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
package gsmci

import (
	"encoding/json"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// ApproveDeviceUser approves device to access user data.
func ApproveDeviceUser(name, fields string, approveDeviceRequest *ci.GoogleAppsCloudidentityDevicesV1ApproveDeviceUserRequest) (map[string]interface{}, error) {
	srv := getDevicesDeviceUsersService()
	c := srv.Approve(name, approveDeviceRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(approveDeviceRequest.Customer, name), func() (interface{}, error) {
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

// BlockDeviceUser blocks device from accessing user data
func BlockDeviceUser(name, fields string, blockDeviceRequest *ci.GoogleAppsCloudidentityDevicesV1BlockDeviceUserRequest) (map[string]interface{}, error) {
	srv := getDevicesDeviceUsersService()
	c := srv.Block(name, blockDeviceRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(blockDeviceRequest.Customer, name), func() (interface{}, error) {
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

// CancelDeviceUserWipe cancels an unfinished user account wipe.
// This operation can be used to cancel device wipe in the gap between the wipe operation returning success and the device being wiped.
func CancelDeviceUserWipe(name, fields string, cancelWipeRequest *ci.GoogleAppsCloudidentityDevicesV1CancelWipeDeviceUserRequest) (map[string]interface{}, error) {
	srv := getDevicesDeviceUsersService()
	c := srv.CancelWipe(name, cancelWipeRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(cancelWipeRequest.Customer, name), func() (interface{}, error) {
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

// DeleteDeviceUser deletes the specified DeviceUser.
// This also revokes the user's access to device data.
func DeleteDeviceUser(name, customer string) (map[string]interface{}, error) {
	srv := getDevicesDeviceUsersService()
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

// GetDeviceUser retrieves the specified DeviceUser
func GetDeviceUser(name, customer, fields string) (*ci.GoogleAppsCloudidentityDevicesV1DeviceUser, error) {
	srv := getDevicesDeviceUsersService()
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
	r, _ := result.(*ci.GoogleAppsCloudidentityDevicesV1DeviceUser)
	return r, nil
}

func listDeviceUsers(c *ci.DevicesDeviceUsersListCall, ch chan *ci.GoogleAppsCloudidentityDevicesV1DeviceUser, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*ci.GoogleAppsCloudidentityDevicesV1ListDeviceUsersResponse)
	for i := range r.DeviceUsers {
		ch <- r.DeviceUsers[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listDeviceUsers(c, ch, errKey)
	}
	return err
}

// ListDeviceUsers lists/searches DeviceUsers.
func ListDeviceUsers(parent, customer, filter, orderBy, fields string, cap int) (<-chan *ci.GoogleAppsCloudidentityDevicesV1DeviceUser, <-chan error) {
	srv := getDevicesDeviceUsersService()
	c := srv.List(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if customer != "" {
		c.Customer(customer)
	}
	if orderBy != "" {
		c.OrderBy(orderBy)
	}
	ch := make(chan *ci.GoogleAppsCloudidentityDevicesV1DeviceUser, cap)
	err := make(chan error, 1)
	go func() {
		e := listDeviceUsers(c, ch, gsmhelpers.FormatErrorKey(customer, filter))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

func lookupDeviceUsers(c *ci.DevicesDeviceUsersLookupCall, ch chan *ci.GoogleAppsCloudidentityDevicesV1DeviceUser, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*ci.GoogleAppsCloudidentityDevicesV1ListDeviceUsersResponse)
	for i := range r.DeviceUsers {
		ch <- r.DeviceUsers[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = lookupDeviceUsers(c, ch, errKey)
	}
	return err
}

// LookupDeviceUsers looks up resource names of the DeviceUsers associated with the caller's credentials, as well as the properties provided in the request.
// This method must be called with end-user credentials with the scope: https://www.googleapis.com/auth/cloud-identity.devices.lookup
// If multiple properties are provided, only DeviceUsers having all of these properties are considered as matches - i.e. the query behaves like an AND.
// Different platforms require different amounts of information from the caller to ensure that the DeviceUser is uniquely identified.
//  - iOS: No properties need to be passed, the caller's credentials are sufficient to identify the corresponding DeviceUser.
//  - Android: Specifying the 'androidId' field is required.
//  - Desktop: Specifying the 'rawResourceId' field is required.
func LookupDeviceUsers(parent, androidID, rawResourceID, userID, fields string, cap int) (<-chan *ci.GoogleAppsCloudidentityDevicesV1DeviceUser, <-chan error) {
	srv := getDevicesDeviceUsersService()
	c := srv.Lookup(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if androidID != "" {
		c.AndroidId(androidID)
	}
	if rawResourceID != "" {
		c.RawResourceId(rawResourceID)
	}
	if userID != "" {
		c.UserId(userID)
	}
	ch := make(chan *ci.GoogleAppsCloudidentityDevicesV1DeviceUser, cap)
	err := make(chan error, 1)
	go func() {
		e := lookupDeviceUsers(c, ch, gsmhelpers.FormatErrorKey(parent, androidID, rawResourceID, userID))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// WipeDeviceUser wipes the user's account on a device.
// Other data on the device that is not associated with the user's work account is not affected.
// For example, if a Gmail app is installed on a device that is used for personal and work purposes,
// and the user is logged in to the Gmail app with their personal account as well as their work account,
// wiping the "deviceUser" by their work administrator will not affect their personal account within Gmail or other apps such as Photos.
func WipeDeviceUser(name, fields string, wipeRequest *ci.GoogleAppsCloudidentityDevicesV1WipeDeviceUserRequest) (map[string]interface{}, error) {
	srv := getDevicesDeviceUsersService()
	c := srv.Wipe(name, wipeRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(wipeRequest.Customer, name), func() (interface{}, error) {
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
