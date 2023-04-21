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

package gsmadmin

import (
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteUserPhoto removes the user's photo.
func DeleteUserPhoto(userKey string) (bool, error) {
	srv := getUsersPhotosService()
	c := srv.Delete(userKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey), func() error {
		return c.Do()
	})
	return result, err
}

// GetUserPhoto retrieves the user's photo.
func GetUserPhoto(userKey, fields string) (*admin.UserPhoto, error) {
	srv := getUsersPhotosService()
	c := srv.Get(userKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.UserPhoto)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// UpdateUserPhoto adds a photo for the user.
func UpdateUserPhoto(userKey, fields string, userPhoto *admin.UserPhoto) (*admin.UserPhoto, error) {
	srv := getUsersPhotosService()
	c := srv.Update(userKey, userPhoto)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.UserPhoto)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
