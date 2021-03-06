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

package gsmcibeta

import (
	"context"
	"encoding/json"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	cibeta "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/googleapi"
)

// CancelInvitation cancels a UserInvitation that was already sent.
func CancelInvitation(name string, cancelUserInvitationRequest *cibeta.CancelUserInvitationRequest) (map[string]interface{}, error) {
	srv := getCustomersUserinvitationsServiceService()
	c := srv.Cancel(name, cancelUserInvitationRequest)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*cibeta.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetInvitation retrieves a UserInvitation resource.
func GetInvitation(name, fields string) (*cibeta.UserInvitation, error) {
	srv := getCustomersUserinvitationsServiceService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*cibeta.UserInvitation)
	return r, nil
}

// IsInvitableUser verifies whether a user account is eligible to receive a UserInvitation (is an unmanaged account).
// Eligibility is based on the following criteria:
//  - the email address is a consumer account and it's the primary email address of the account, and
//  - the domain of the email address matches an existing verified Google Workspace or Cloud Identity domain
// If both conditions are met, the user is eligible.
func IsInvitableUser(name string) (bool, error) {
	srv := getCustomersUserinvitationsServiceService()
	c := srv.IsInvitableUser(name)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	r, _ := result.(*cibeta.IsInvitableUserResponse)
	return r.IsInvitableUser, nil
}

// ListUserInvitations retrieves a list of UserInvitation resources.
func ListUserInvitations(parent, filter, orderBy, fields string, cap int) (<-chan *cibeta.UserInvitation, <-chan error) {
	srv := getCustomersUserinvitationsServiceService()
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
	ch := make(chan *cibeta.UserInvitation, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *cibeta.ListUserInvitationsResponse) error {
			for i := range response.UserInvitations {
				ch <- response.UserInvitations[i]
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

// SendInvitation sends a UserInvitation to email.
// If the UserInvitation does not exist for this request and it is a valid request, the request creates a UserInvitation.
func SendInvitation(name, fields string, sendUserInvitationRequest *cibeta.SendUserInvitationRequest) (map[string]interface{}, error) {
	srv := getCustomersUserinvitationsServiceService()
	c := srv.Send(name, sendUserInvitationRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*cibeta.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
