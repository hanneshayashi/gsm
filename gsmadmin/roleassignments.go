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
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteRoleAssignment deletes a role assignment.
func DeleteRoleAssignment(customer, roleAssignmentID string) (bool, error) {
	srv := getRoleAssignmentsService()
	c := srv.Delete(customer, roleAssignmentID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customer, roleAssignmentID), func() error {
		return c.Do()
	})
	return result, err
}

// GetRoleAssignment retrieve a role assignment.
func GetRoleAssignment(customer, roleAssignmentID, fields string) (*admin.RoleAssignment, error) {
	srv := getRoleAssignmentsService()
	c := srv.Get(customer, roleAssignmentID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, roleAssignmentID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.RoleAssignment)
	return r, nil
}

// InsertRoleAssignment creates a role assignment.
func InsertRoleAssignment(customer, fields string, roleAssignment *admin.RoleAssignment) (*admin.RoleAssignment, error) {
	srv := getRoleAssignmentsService()
	c := srv.Insert(customer, roleAssignment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, strconv.FormatInt(roleAssignment.RoleId, 10), roleAssignment.AssignedTo), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.RoleAssignment)
	return r, nil
}

func listRoleAssignments(c *admin.RoleAssignmentsListCall, ch chan *admin.RoleAssignment, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*admin.RoleAssignments)
	for _, i := range r.Items {
		ch <- i
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listRoleAssignments(c, ch, errKey)
	}
	return err
}

// ListRoleAssignments retrieves a paginated list of all roleAssignments.
func ListRoleAssignments(customer, roleID, userKey, fields string, cap int) (<-chan *admin.RoleAssignment, <-chan error) {
	srv := getRoleAssignmentsService()
	c := srv.List(customer).MaxResults(200)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if roleID != "" {
		c = c.RoleId(roleID)
	}
	if userKey != "" {
		c = c.UserKey(userKey)
	}
	ch := make(chan *admin.RoleAssignment, cap)
	err := make(chan error, 1)
	go func() {
		e := listRoleAssignments(c, ch, gsmhelpers.FormatErrorKey(customer, roleID, userKey))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}
