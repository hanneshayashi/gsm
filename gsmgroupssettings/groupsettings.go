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

package gsmgroupssettings

import (
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/groupssettings/v1"
)

// GetGroupSettings retrieves a group's settings identified by the group email address.
func GetGroupSettings(groupUniqueID, fields string) (*groupssettings.Groups, error) {
	srv := getGroupssettingsService()
	c := srv.Groups.Get(groupUniqueID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupUniqueID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*groupssettings.Groups)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// PatchGroupSettings updates an existing resource. This method supports patch semantics.
func PatchGroupSettings(groupUniqueID, fields string, groups *groupssettings.Groups) (*groupssettings.Groups, error) {
	srv := getGroupssettingsService()
	c := srv.Groups.Patch(groupUniqueID, groups)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupUniqueID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*groupssettings.Groups)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
