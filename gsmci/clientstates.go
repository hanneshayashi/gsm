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
	"context"
	"encoding/json"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// GetClientState gets the client state for the device user
func GetClientState(name, customer, fields string) (*ci.GoogleAppsCloudidentityDevicesV1ClientState, error) {
	srv := getDevicesDeviceUsersClientStatesService()
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
	r, _ := result.(*ci.GoogleAppsCloudidentityDevicesV1ClientState)
	return r, nil
}

// ListClientStates lists the client states for the given search query.
func ListClientStates(parent, customer, filter, orderBy, fields string, cap int) (<-chan *ci.GoogleAppsCloudidentityDevicesV1ClientState, <-chan error) {
	srv := getDevicesDeviceUsersClientStatesService()
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
	if filter != "" {
		c.Filter(filter)
	}
	ch := make(chan *ci.GoogleAppsCloudidentityDevicesV1ClientState, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.GoogleAppsCloudidentityDevicesV1ListClientStatesResponse) error {
			for i := range response.ClientStates {
				ch <- response.ClientStates[i]
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

// PatchClientState updates the client state for the device user
func PatchClientState(name, customer, updateMask, fields string, clientState *ci.GoogleAppsCloudidentityDevicesV1ClientState) (map[string]interface{}, error) {
	srv := getDevicesDeviceUsersClientStatesService()
	c := srv.Patch(name, clientState)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if customer != "" {
		c.Customer(customer)
	}
	if updateMask != "" {
		c.UpdateMask(updateMask)
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
