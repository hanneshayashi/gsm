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

// CreateSsoProfile creates an InboundSamlSsoProfile for a customer.
func CreateSsoProfile(fields string, profile *ci.InboundSamlSsoProfile) (*googleapi.RawMessage, error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.Create(profile)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(profile.DisplayName), err)
	}
	return &r.Response, nil
}

// DeleteSsoProfile deletes an InboundSamlSsoProfile.
func DeleteSsoProfile(name string) (bool, error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.Delete(name)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r, nil
}

// ListSsoProfiles retrieves a list of InboundSamlSsoProfile resources.
func ListSsoProfiles(filter, fields string) iter.Seq2[*ci.InboundSamlSsoProfile, error] {
	srv := getInboundSamlSsoProfilesService()
	c := srv.List().PageSize(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c.Filter(fields)
	}
	return func(yield func(*ci.InboundSamlSsoProfile, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.ListInboundSamlSsoProfilesResponse) error {
			for i := range response.InboundSamlSsoProfiles {
				if !yield(response.InboundSamlSsoProfiles[i], nil) {
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

// PatchSsoProfile updates an InboundSamlSsoProfile.
func PatchSsoProfile(name, updateMask, fields string, profile *ci.InboundSamlSsoProfile) (*googleapi.RawMessage, error) {
	srv := getInboundSamlSsoProfilesService()
	c := srv.Patch(name, profile).UpdateMask(updateMask)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(profile.DisplayName), err)
	}
	return &r.Response, nil
}
