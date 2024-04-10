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

package gsmci

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// AddSsoProfileIdpCredential adds an IdpCredential.
// Up to 2 credentials are allowed.
func AddSsoProfileIdpCredential(parent, fields string, request *ci.AddIdpCredentialRequest) (*googleapi.RawMessage, error) {
	srv := getInboundSamlSsoProfilesIdpCredentialsService()
	c := srv.Add(parent, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent), func() (any, error) {
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

// DeleteSsoProfileIdpCredential deletes an IdpCredential.
func DeleteSsoProfileIdpCredential(name string) (bool, error) {
	srv := getInboundSamlSsoProfilesIdpCredentialsService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetSsoProfileIdpCredential gets an IdpCredential.
func GetSsoProfileIdpCredential(parent, fields string) (*ci.IdpCredential, error) {
	srv := getInboundSamlSsoProfilesIdpCredentialsService()
	c := srv.Get(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.IdpCredential)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListSsoProfileIdpCredential returns a list of IdpCredentials in an InboundSamlSsoProfile.
func ListSsoProfileIdpCredential(parent, fields string, cap int) (<-chan *ci.IdpCredential, <-chan error) {
	srv := getInboundSamlSsoProfilesIdpCredentialsService()
	c := srv.List(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *ci.IdpCredential, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.ListIdpCredentialsResponse) error {
			for i := range response.IdpCredentials {
				ch <- response.IdpCredentials[i]
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
