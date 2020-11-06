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
package gsmgmail

import (
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// CreateForwardingAddress creates a forwarding address.
// If ownership verification is required, a message will be sent to the recipient and the resource's verification status will be set to pending;
// otherwise, the resource will be created with verification status set to accepted.
func CreateForwardingAddress(userID, fields string, settingsforwardingaddress *gmail.ForwardingAddress) (*gmail.ForwardingAddress, error) {
	srv := getUsersSettingsForwardingAddressesService()
	c := srv.Create(userID, settingsforwardingaddress)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// DeleteForwardingAddress deletes the specified forwarding address and revokes any verification that may have been required.
func DeleteForwardingAddress(userID, forwardingEmail string) (bool, error) {
	srv := getUsersSettingsForwardingAddressesService()
	err := srv.Delete(userID, forwardingEmail).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetForwardingAddress gets the specified forwarding address.
func GetForwardingAddress(userID, forwardingEmail, fields string) (*gmail.ForwardingAddress, error) {
	srv := getUsersSettingsForwardingAddressesService()
	c := srv.Get(userID, forwardingEmail)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// ListForwardingAddresses lists the forwarding addresses for the specified account.
func ListForwardingAddresses(userID, fields string) ([]*gmail.ForwardingAddress, error) {
	srv := getUsersSettingsForwardingAddressesService()
	c := srv.List(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.ForwardingAddresses, err
}
