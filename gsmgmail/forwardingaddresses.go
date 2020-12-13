/*
Package gsmgmail implements the Gmail APIs
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
package gsmgmail

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// CreateForwardingAddress creates a forwarding address.
// If ownership verification is required, a message will be sent to the recipient and the resource's verification status will be set to pending;
// otherwise, the resource will be created with verification status set to accepted.
func CreateForwardingAddress(userID, fields string, forwardingAddress *gmail.ForwardingAddress) (*gmail.ForwardingAddress, error) {
	srv := getUsersSettingsForwardingAddressesService()
	c := srv.Create(userID, forwardingAddress)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, forwardingAddress.ForwardingEmail), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.ForwardingAddress)
	return r, nil
}

// DeleteForwardingAddress deletes the specified forwarding address and revokes any verification that may have been required.
func DeleteForwardingAddress(userID, forwardingEmail string) (bool, error) {
	srv := getUsersSettingsForwardingAddressesService()
	c := srv.Delete(userID, forwardingEmail)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userID, forwardingEmail), func() error {
		return c.Do()
	})
	return result, err
}

// GetForwardingAddress gets the specified forwarding address.
func GetForwardingAddress(userID, forwardingEmail, fields string) (*gmail.ForwardingAddress, error) {
	srv := getUsersSettingsForwardingAddressesService()
	c := srv.Get(userID, forwardingEmail)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, forwardingEmail), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.ForwardingAddress)
	return r, nil
}

// ListForwardingAddresses lists the forwarding addresses for the specified account.
func ListForwardingAddresses(userID, fields string) ([]*gmail.ForwardingAddress, error) {
	srv := getUsersSettingsForwardingAddressesService()
	c := srv.List(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.ListForwardingAddressesResponse)
	return r.ForwardingAddresses, nil
}
