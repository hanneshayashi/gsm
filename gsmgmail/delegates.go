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
package gsmgmail

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// GetDelegate gets the specified delegate.
// Note that a delegate user must be referred to by their primary email address, and not an email alias.
// This method is only available to service account clients that have been delegated domain-wide authority.
func GetDelegate(userID, delegateEmail, fields string) (*gmail.Delegate, error) {
	srv := getUsersSettingsDelegatesService()
	c := srv.Get(userID, delegateEmail)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, delegateEmail), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Delegate)
	return r, nil
}

// DeleteDelegate removes the specified delegate (which can be of any verification status), and revokes any verification that may have been required for using it.
// Note that a delegate user must be referred to by their primary email address, and not an email alias.
// This method is only available to service account clients that have been delegated domain-wide authority.
func DeleteDelegate(userID, delegateEmail string) (bool, error) {
	srv := getUsersSettingsDelegatesService()
	c := srv.Delete(userID, delegateEmail)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userID, delegateEmail), func() error {
		return c.Do()
	})
	return result, err
}

// ListDelegates lists the delegates for the specified account.
// This method is only available to service account clients that have been delegated domain-wide authority.
func ListDelegates(userID, fields string) ([]*gmail.Delegate, error) {
	srv := getUsersSettingsDelegatesService()
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
	r, _ := result.(*gmail.ListDelegatesResponse)
	return r.Delegates, nil
}

// CreateDelegate adds a delegate with its verification status set directly to accepted, without sending any verification email.
// The delegate user must be a member of the same Workspace organization as the delegator user.
// Gmail imposes limitations on the number of delegates and delegators each user in a Workspace organization can have.
// These limits depend on your organization, but in general each user can have up to 25 delegates and up to 10 delegators.
// Note that a delegate user must be referred to by their primary email address, and not an email alias.
// Also note that when a new delegate is created, there may be up to a one minute delay before the new delegate is available for use.
// This method is only available to service account clients that have been delegated domain-wide authority.
func CreateDelegate(userID, fields string, delegate *gmail.Delegate) (*gmail.Delegate, error) {
	srv := getUsersSettingsDelegatesService()
	c := srv.Create(userID, delegate)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, delegate.DelegateEmail), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Delegate)
	return r, nil
}
