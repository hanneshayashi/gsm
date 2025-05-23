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

// DeleteUserAlias removes an alias.
func DeleteUserAlias(userKey, alias string) (bool, error) {
	srv := getUsersAliasesService()
	c := srv.Delete(userKey, alias)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey, alias), func() error {
		return c.Do()
	})
	return result, err
}

// InsertUserAlias adds an alias.
func InsertUserAlias(userKey, fields string, alias *admin.Alias) (*admin.Alias, error) {
	srv := getUsersAliasesService()
	c := srv.Insert(userKey, alias)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey, alias.Alias), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Alias)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListUserAliases lists all aliases for a user.
func ListUserAliases(userKey, fields string) ([]any, error) {
	srv := getUsersAliasesService()
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
	r, ok := result.(*admin.Aliases)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r.Aliases, nil
}
