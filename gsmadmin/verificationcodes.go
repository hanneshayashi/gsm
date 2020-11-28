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
	"gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// GenerateVerificationCodes generates new backup verification codes for the user.
func GenerateVerificationCodes(userKey string) (bool, error) {
	srv := getVerificationCodesService()
	c := srv.Generate(userKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey), func() error {
		return c.Do()
	})
	return result, err
}

// InvalidateVerificationCodes invalidates the current backup verification codes for the user.
func InvalidateVerificationCodes(userKey string) (bool, error) {
	srv := getVerificationCodesService()
	c := srv.Invalidate(userKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey), func() error {
		return c.Do()
	})
	return result, err
}

// ListVerificationCodes returns the current set of valid backup verification codes for the specified user.
func ListVerificationCodes(userKey, fields string) ([]*admin.VerificationCode, error) {
	srv := getVerificationCodesService()
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
	r, _ := result.(*admin.VerificationCodes)
	return r.Items, nil
}
