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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, roleID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Role)
	return r, nil
}

// InsertRole creates a role.
func InsertRole(customer, fields string, role *admin.Role) (*admin.Role, error) {
	srv := getRolesService()
	c := srv.Insert(customer, role)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, role.RoleName), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Role)
	return r, nil
}

func makeListRolesCallAndAppend(c *admin.RolesListCall, roles []*admin.Role, errKey string) ([]*admin.Role, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Roles)
	for _, r := range r.Items {
		roles = append(roles, r)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		roles, err = makeListRolesCallAndAppend(c, roles, errKey)
	}
	return roles, err
}

// ListRoles retrieves a paginated list of all the roles in a domain.
func ListRoles(customer, fields string) ([]*admin.Role, error) {
	srv := getRolesService()
	c := srv.List(customer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var roles []*admin.Role
	roles, err := makeListRolesCallAndAppend(c, roles, customer)
	return roles, err
}

// PatchRole updates a role. This method supports patch semantics.
func PatchRole(customer, roleID, fields string, role *admin.Role) (*admin.Role, error) {
	srv := getRolesService()
	c := srv.Patch(customer, roleID, role)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, roleID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Role)
	return r, nil
}
