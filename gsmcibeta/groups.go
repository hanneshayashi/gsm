/*
Copyright © 2020-2022 Hannes Hayashi

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
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	cibeta "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/googleapi"
)

// GetSecuritySettings returns the security settings of a group.
func GetSecuritySettings(name, readMask, fields string) (*cibeta.SecuritySettings, error) {
	srv := getGroupsService()
	c := srv.GetSecuritySettings(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if readMask != "" {
		c.ReadMask(readMask)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*cibeta.SecuritySettings)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
	}
	return r, nil
}

// UpdateSecuritySettings updates the security settings of a group.
func UpdateSecuritySettings(name, updateMask, fields string, securitysettings *cibeta.SecuritySettings) (*googleapi.RawMessage, error) {
	srv := getGroupsService()
	c := srv.UpdateSecuritySettings(name, securitysettings)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if updateMask != "" {
		c.UpdateMask(updateMask)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*cibeta.Operation)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
	}
	return &r.Response, nil
}
