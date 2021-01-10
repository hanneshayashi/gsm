/*
Package gsmadmin implements the Admin SDK APIs
Copyright © 2020-2021 Hannes Hayashi

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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteGroupAlias removes an alias.
func DeleteGroupAlias(groupKey, alias string) (bool, error) {
	srv := getGroupsAliasesService()
	c := srv.Delete(groupKey, alias)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(groupKey), func() error {
		return c.Do()
	})
	return result, err
}

// InsertGroupAlias adds an alias for the group.
func InsertGroupAlias(groupKey, fields string, alias *admin.Alias) (*admin.Alias, error) {
	srv := getGroupsAliasesService()
	c := srv.Insert(groupKey, alias)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey, alias.Alias), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Alias)
	return r, nil
}

// ListGroupAliases lists all aliases for a group.
func ListGroupAliases(groupKey, fields string) ([]interface{}, error) {
	srv := getGroupsAliasesService()
	c := srv.List(groupKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Aliases)
	return r.Aliases, nil
}
