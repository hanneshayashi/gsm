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
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// CancelInvitation cancels a UserInvitation that was already sent.
func CancelInvitation(name string, cancelUserInvitationRequest *ci.CancelUserInvitationRequest) (*googleapi.RawMessage, error) {
	srv := getCustomersUserinvitationsService()
	c := srv.Cancel(name, cancelUserInvitationRequest)
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return &r.Response, nil
}

// GetInvitation retrieves a UserInvitation resource.
func GetInvitation(name, fields string) (*ci.UserInvitation, error) {
	srv := getCustomersUserinvitationsService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r, nil
}

// IsInvitableUser verifies whether a user account is eligible to receive a UserInvitation (is an unmanaged account).
// Eligibility is based on the following criteria:
//   - the email address is a consumer account and it's the primary email address of the account, and
//   - the domain of the email address matches an existing verified Google Workspace or Cloud Identity domain
//
// If both conditions are met, the user is eligible.
func IsInvitableUser(name string) (bool, error) {
	srv := getCustomersUserinvitationsService()
	c := srv.IsInvitableUser(name)
	r, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r.IsInvitableUser, nil
}

// ListUserInvitations retrieves a list of UserInvitation resources.
func ListUserInvitations(parent, filter, orderBy, fields string) iter.Seq2[*ci.UserInvitation, error] {
	srv := getCustomersUserinvitationsService()
	c := srv.List(parent).PageSize(200)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c.Filter(fields)
	}
	if orderBy != "" {
		c.OrderBy(orderBy)
	}
	return func(yield func(*ci.UserInvitation, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.ListUserInvitationsResponse) error {
			for i := range response.UserInvitations {
				if !yield(response.UserInvitations[i], nil) {
					return errIterStopped
				}
			}
			return nil
		})
		if e != nil && !errors.Is(e, errIterStopped) {
			yield(nil, e)
		}
	}
}

// SendInvitation sends a UserInvitation to email.
// If the UserInvitation does not exist for this request and it is a valid request, the request creates a UserInvitation.
func SendInvitation(name, fields string, sendUserInvitationRequest *ci.SendUserInvitationRequest) (*googleapi.RawMessage, error) {
	srv := getCustomersUserinvitationsService()
	c := srv.Send(name, sendUserInvitationRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return &r.Response, nil
}
