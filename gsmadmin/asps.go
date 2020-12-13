/*
Package gsmadmin implements the Admin SDK APIs
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
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteAsp deletes an ASP issued by a user.
func DeleteAsp(userKey string, codeID int64) (bool, error) {
	srv := getAspsService()
	c := srv.Delete(userKey, codeID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey), func() error {
		return c.Do()
	})
	return result, err
}

// GetAsp gets information about an ASP issued by a user.
func GetAsp(userKey, fields string, codeID int64) (*admin.Asp, error) {
	srv := getAspsService()
	c := srv.Get(userKey, codeID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey, strconv.FormatInt(codeID, 10)), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Asp)
	return r, nil
}

// ListAsps lists the ASPs issued by a user.
func ListAsps(userKey, fields string) ([]*admin.Asp, error) {
	srv := getAspsService()
	c := srv.List(userKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Asps)
	return r.Items, nil
}
