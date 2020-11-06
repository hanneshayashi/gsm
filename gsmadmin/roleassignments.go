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
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteRoleAssignment deletes a role assignment.
func DeleteRoleAssignment(customer, roleAssignmentID string) (bool, error) {
	srv := getRoleAssignmentsService()
	err := srv.Delete(customer, roleAssignmentID).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetRoleAssignment retrieve a role assignment.
func GetRoleAssignment(customer, roleAssignmentID, fields string) (*admin.RoleAssignment, error) {
	srv := getRoleAssignmentsService()
	c := srv.Get(customer, roleAssignmentID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertRoleAssignment creates a role assignment.
func InsertRoleAssignment(customer, fields string, roleAssignment *admin.RoleAssignment) (*admin.RoleAssignment, error) {
	srv := getRoleAssignmentsService()
	c := srv.Insert(customer, roleAssignment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

func makeListRoleAssignmentsCallAndAppend(c *admin.RoleAssignmentsListCall, roleAssignments []*admin.RoleAssignment) ([]*admin.RoleAssignment, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, r := range r.Items {
		roleAssignments = append(roleAssignments, r)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		roleAssignments, err = makeListRoleAssignmentsCallAndAppend(c, roleAssignments)
	}
	return roleAssignments, err
}

// ListRoleAssignments retrieves a paginated list of all roleAssignments.
func ListRoleAssignments(customer, roleID, userKey, fields string) ([]*admin.RoleAssignment, error) {
	srv := getRoleAssignmentsService()
	c := srv.List(customer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if roleID != "" {
		c = c.RoleId(roleID)
	}
	if userKey != "" {
		c = c.UserKey(userKey)
	}
	var roleAssignments []*admin.RoleAssignment
	roleAssignments, err := makeListRoleAssignmentsCallAndAppend(c, roleAssignments)
	return roleAssignments, err
}
