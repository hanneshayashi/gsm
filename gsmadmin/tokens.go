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

package gsmadmin

import (
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteToken deletes all access tokens issued by a user for an application.rolesPatchCmd
func DeleteToken(userKey, clientID string) (bool, error) {
	srv := getTokensService()
	c := srv.Delete(userKey, clientID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey, clientID), func() error {
		return c.Do()
	})
	return result, err
}

// GetToken gets information about an access token issued by a user.rolesPatchCmd
func GetToken(userKey, clientID, fields string) (*admin.Token, error) {
	srv := getTokensService()
	c := srv.Get(userKey, clientID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey, clientID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Token)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListTokens returns the set of tokens specified user has issued to 3rd party applications.rolesPatchCmd
func ListTokens(userKey, fields string) ([]*admin.Token, error) {
	srv := getTokensService()
	c := srv.List(userKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Tokens)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r.Items, nil
}
