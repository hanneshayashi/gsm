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

// DeleteGroup deletes a group.
func DeleteGroup(groupKey string) (bool, error) {
	srv := getGroupsService()
	c := srv.Delete(groupKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(groupKey), func() error {
		return c.Do()
	})
	return result, err
}

// GetGroup retrieves a group's properties.
func GetGroup(groupKey, fields string) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Get(groupKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Group)
	return r, nil
}

// InsertGroup creates a group.
func InsertGroup(group *admin.Group, fields string) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Insert(group)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(group.Email), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Group)
	return r, nil
}

func makeListGroupsCallAndAppend(c *admin.GroupsListCall, groups []*admin.Group, errKey string) ([]*admin.Group, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Groups)
	for _, g := range r.Groups {
		groups = append(groups, g)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		groups, err = makeListGroupsCallAndAppend(c, groups, errKey)
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
	groups, err := makeListGroupsCallAndAppend(c, groups, gsmhelpers.FormatErrorKey(customer))
	return groups, err
}

// PatchGroup updates a group's properties. This method supports patch semantics.
func PatchGroup(groupKey, fields string, Group *admin.Group) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Patch(groupKey, Group)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Group)
	return r, nil
}
