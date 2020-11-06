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
package gsmadmin

import (
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteToken deletes all access tokens issued by a user for an application.rolesPatchCmd
func DeleteToken(userKey, clientID string) (bool, error) {
	srv := getTokensService()
	err := srv.Delete(userKey, clientID).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetToken gets information about an access token issued by a user.rolesPatchCmd
func GetToken(userKey, clientID, fields string) (*admin.Token, error) {
	srv := getTokensService()
	c := srv.Get(userKey, clientID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// ListTokens returns the set of tokens specified user has issued to 3rd party applications.rolesPatchCmd
func ListTokens(userKey, fields string) ([]*admin.Token, error) {
	srv := getTokensService()
	c := srv.List(userKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.Items, nil
}
