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

// DeleteGroup deletes a group.
func DeleteGroup(groupkey string) (bool, error) {
	srv := getGroupsService()
	err := srv.Delete(groupkey).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetGroup retrieves a group's properties.
func GetGroup(groupkey, fields string) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Get(groupkey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertGroup creates a group.
func InsertGroup(Group *admin.Group, fields string) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Insert(Group)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

func makeListGroupsCallAndAppend(c *admin.GroupsListCall, groups []*admin.Group) ([]*admin.Group, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, g := range r.Groups {
		groups = append(groups, g)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		groups, err = makeListGroupsCallAndAppend(c, groups)
	}
	return groups, err
}

// ListGroups retrieve all groups of a domain or of a user given a userKey (paginated)
func ListGroups(filter, userKey, domain, customer, fields string) ([]*admin.Group, error) {
	srv := getGroupsService()
	c := srv.List()
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if userKey != "" {
		c = c.UserKey(userKey)
	} else {
		c = c.Customer(customer)
	}
	if filter != "" {
		c = c.Query(filter)
	}
	if domain != "" {
		c = c.Domain(domain)
	}
	var groups []*admin.Group
	groups, err := makeListGroupsCallAndAppend(c, groups)
	return groups, err
}

// PatchGroup updates a group's properties. This method supports patch semantics.
func PatchGroup(groupkey, fields string, Group *admin.Group) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Patch(groupkey, Group)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
