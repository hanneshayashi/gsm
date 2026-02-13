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

package gsmci

import (
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// CreateGroup creates a group.
func CreateGroup(group *ci.Group, initialGroupConfig, fields string) (*googleapi.RawMessage, error) {
	srv := getGroupsService()
	c := srv.Create(group).InitialGroupConfig(initialGroupConfig)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(group.GroupKey.Id), err)
	}
	return &r.Response, nil
}

// DeleteGroup deletes a group.
func DeleteGroup(name string) (bool, error) {
	srv := getGroupsService()
	c := srv.Delete(name)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return true, nil
}

// PatchGroup updates a group using patch semantics.
func PatchGroup(name, updateMask, fields string, group *ci.Group) (*googleapi.RawMessage, error) {
	srv := getGroupsService()
	c := srv.Patch(name, group).UpdateMask(updateMask)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return &r.Response, nil
}

// GetGroup retrieves a group.
func GetGroup(name, fields string) (*ci.Group, error) {
	srv := getGroupsService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r, nil
}

// LookupGroup looks up a group via its email address and returns its resourceName
func LookupGroup(email string) (string, error) {
	srv := getGroupsService()
	c := srv.Lookup().GroupKeyId(email)
	r, err := c.Do()
	if err != nil {
		return "", fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(email), err)
	}
	return r.Name, nil
}

// ListGroups retrieves a list of groups
func ListGroups(parent, view, fields string) iter.Seq2[*ci.Group, error] {
	srv := getGroupsService()
	c := srv.List().Parent(parent).PageSize(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if view != "" {
		c.View(view)
	}
	return func(yield func(*ci.Group, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.ListGroupsResponse) error {
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

// SearchGroups searches for Groups matching a specified query.
func SearchGroups(query, view, fields string) iter.Seq2[*ci.Group, error] {
	srv := getGroupsService()
	c := srv.Search().Query(query).PageSize(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if view != "" {
		c.View(view)
	}
	return func(yield func(*ci.Group, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.SearchGroupsResponse) error {
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

// GetSecuritySettings returns the security settings of a group.
func GetSecuritySettings(name, readMask, fields string) (*ci.SecuritySettings, error) {
	srv := getGroupsService()
	c := srv.GetSecuritySettings(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if readMask != "" {
		c.ReadMask(readMask)
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r, nil
}

// UpdateSecuritySettings updates the security settings of a group.
func UpdateSecuritySettings(name, updateMask, fields string, securitysettings *ci.SecuritySettings) (*googleapi.RawMessage, error) {
	srv := getGroupsService()
	c := srv.UpdateSecuritySettings(name, securitysettings)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if updateMask != "" {
		c.UpdateMask(updateMask)
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return &r.Response, nil
}
