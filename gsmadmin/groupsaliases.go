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
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteGroupAlias removes an alias.
func DeleteGroupAlias(groupKey, alias string) (bool, error) {
	srv := getGroupsAliasesService()
	c := srv.Delete(groupKey, alias)
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(groupKey), err)
	}
	return true, nil
}

// InsertGroupAlias adds an alias for the group.
func InsertGroupAlias(groupKey, fields string, alias *admin.Alias) (*admin.Alias, error) {
	srv := getGroupsAliasesService()
	c := srv.Insert(groupKey, alias)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(groupKey, alias.Alias), err)
	}
	return r, nil
}

// ListGroupAliases lists all aliases for a group.
func ListGroupAliases(groupKey, fields string) ([]any, error) {
	srv := getGroupsAliasesService()
	c := srv.List(groupKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(groupKey), err)
	}
	return r.Aliases, nil
}
