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
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteGroup deletes a group.
func DeleteGroup(groupKey string) (bool, error) {
	srv := getGroupsService()
	c := srv.Delete(groupKey)
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(groupKey), err)
	}
	return true, nil
}

// GetGroup retrieves a group's properties.
func GetGroup(groupKey, fields string) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Get(groupKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(groupKey), err)
	}
	return r, nil
}

// InsertGroup creates a group.
func InsertGroup(group *admin.Group, fields string) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Insert(group)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(group.Email), err)
	}
	return r, nil
}

// ListGroups retrieve all groups of a domain or of a user given a userKey (paginated)
func ListGroups(filter, userKey, domain, customer, fields string) iter.Seq2[*admin.Group, error] {
	srv := getGroupsService()
	c := srv.List().MaxResults(200)
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
	return func(yield func(*admin.Group, error) bool) {
		e := c.Pages(context.Background(), func(response *admin.Groups) error {
			for i := range response.Groups {
				if !yield(response.Groups[i], nil) {
					return errIterStopped
				}
			}
			return nil
		})
		if e != nil && !errors.Is(e, errIterStopped) {
			yield(nil, e)
		}
	}
}

// PatchGroup updates a group's properties. This method supports patch semantics.
func PatchGroup(groupKey, fields string, Group *admin.Group) (*admin.Group, error) {
	srv := getGroupsService()
	c := srv.Patch(groupKey, Group)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(groupKey), err)
	}
	return r, nil
}
