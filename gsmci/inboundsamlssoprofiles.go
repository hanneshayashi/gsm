/*
Copyright © 2020-2023 Hannes Hayashi

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

// CreateSsoProfile creates an InboundSamlSsoProfile for a customer.
func CreateSsoProfile(fields string, profile *ci.InboundSamlSsoProfile) (*googleapi.RawMessage, error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.Create(profile)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(profile.DisplayName), func() (any, error) {
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

// DeleteSsoProfile deletes an InboundSamlSsoProfile.
func DeleteSsoProfile(name string) (bool, error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetSsoProfile gets an InboundSamlSsoProfile.
func GetSsoProfile(name, fields string) (*ci.InboundSamlSsoProfile, error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.InboundSamlSsoProfile)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListSsoProfiles retrieves a list of InboundSamlSsoProfile resources.
func ListSsoProfiles(filter, fields string, cap int) (<-chan *ci.InboundSamlSsoProfile, <-chan error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.List().PageSize(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c.Filter(fields)
	}
	ch := make(chan *ci.InboundSamlSsoProfile, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.ListInboundSamlSsoProfilesResponse) error {
			for i := range response.InboundSamlSsoProfiles {
				ch <- response.InboundSamlSsoProfiles[i]
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

// PatchSsoProfile updates an InboundSamlSsoProfile.
func PatchSsoProfile(name, updateMask, fields string, profile *ci.InboundSamlSsoProfile) (*googleapi.RawMessage, error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.Patch(name, profile).UpdateMask(updateMask)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(profile.DisplayName), func() (any, error) {
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
