/*
Package gsmgmail implements the Gmail APIs
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
package gsmgmail

import (
	"gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// GetUserProfile Ggts the specified user's Gmail profile.
func GetUserProfile(userID, fields string) (*gmail.Profile, error) {
	srv := getUsersService()
	c := srv.GetProfile(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Profile)
	return r, nil
}
