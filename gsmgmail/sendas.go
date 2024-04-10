/*
Copyright © 2020-2024 Hannes Hayashi

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
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// CreateSendAs creates a custom "from" send-as alias.
// If an SMTP MSA is specified, Gmail will attempt to connect to the SMTP service to validate the configuration before creating the alias.
// If ownership verification is required for the alias, a message will be sent to the email address and the resource's verification status will be set to pending;
// otherwise, the resource will be created with verification status set to accepted.
// If a signature is provided, Gmail will sanitize the HTML before saving it with the alias.
func CreateSendAs(userID, fields string, sendAs *gmail.SendAs) (*gmail.SendAs, error) {
	srv := getUsersSettingsSendAsService()
	c := srv.Create(userID, sendAs)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, sendAs.SendAsEmail), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*gmail.SendAs)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// DeleteSendAs deletes the specified send-as alias.
// Revokes any verification that may have been required for using it.
func DeleteSendAs(userID, sendAsEmail string) (bool, error) {
	srv := getUsersSettingsSendAsService()
	c := srv.Delete(userID, sendAsEmail)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userID, sendAsEmail), func() error {
		return c.Do()
	})
	return result, err
}

// GetSendAs gets the specified send-as alias.
// Fails with an HTTP 404 error if the specified address is not a member of the collection.
func GetSendAs(userID, sendAsEmail, fields string) (*gmail.SendAs, error) {
	srv := getUsersSettingsSendAsService()
	c := srv.Get(userID, sendAsEmail)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, sendAsEmail), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*gmail.SendAs)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListSendAs lists the send-as aliases for the specified account.
// The result includes the primary send-as address associated with the account as well as any custom "from" aliases.
func ListSendAs(userID, fields string) ([]*gmail.SendAs, error) {
	srv := getUsersSettingsSendAsService()
	c := srv.List(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*gmail.ListSendAsResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r.SendAs, nil
}

// PatchSendAs PATCHes the specified send-as alias.
func PatchSendAs(userID, sendAsEmail, fields string, sendAs *gmail.SendAs) (*gmail.SendAs, error) {
	srv := getUsersSettingsSendAsService()
	c := srv.Patch(userID, sendAsEmail, sendAs)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, sendAsEmail), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*gmail.SendAs)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// VerifySendAs sends a verification email to the specified send-as alias address.
// The verification status must be pending.
func VerifySendAs(userID, sendAsEmail string) (bool, error) {
	srv := getUsersSettingsSendAsService()
	c := srv.Verify(userID, sendAsEmail)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userID, sendAsEmail), func() error {
		return c.Do()
	})
	return result, err
}
