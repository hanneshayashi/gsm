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

package gsmadmin

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteRole deleteRole deletes a role.
func DeleteRole(customer, roleID string) (bool, error) {
	srv := getRolesService()
	c := srv.Delete(customer, roleID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customer, roleID), func() error {
		return c.Do()
	})
	return result, err
}

// GetRole retrieves a role.
func GetRole(customer, roleID, fields string) (*admin.Role, error) {
	srv := getRolesService()
	c := srv.Get(customer, roleID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, roleID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Role)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
	}
	return r, nil
}

// InsertRole creates a role.
func InsertRole(customer, fields string, role *admin.Role) (*admin.Role, error) {
	srv := getRolesService()
	c := srv.Insert(customer, role)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, role.RoleName), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Role)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
	}
	return r, nil
}

// ListRoles retrieves a paginated list of all the roles in a domain.
func ListRoles(customer, fields string, cap int) (<-chan *admin.Role, <-chan error) {
	srv := getRolesService()
	c := srv.List(customer).MaxResults(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *admin.Role, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *admin.Roles) error {
			for i := range response.Items {
				ch <- response.Items[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// PatchRole updates a role. This method supports patch semantics.
func PatchRole(customer, roleID, fields string, role *admin.Role) (*admin.Role, error) {
	srv := getRolesService()
	c := srv.Patch(customer, roleID, role)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, roleID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Role)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
	}
	return r, nil
}
