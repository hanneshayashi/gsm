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

package gsmci

import (
	"context"
	"fmt"

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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(group.GroupKey.Id), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
	}
	return &r.Response, nil
}

// DeleteGroup deletes a group.
func DeleteGroup(name string) (bool, error) {
	srv := getGroupsService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Group)
	return r, nil
}

// LookupGroup looks up a group via its email address and returns its resourceName
func LookupGroup(email string) (string, error) {
	srv := getGroupsService()
	c := srv.Lookup().GroupKeyId(email)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(email), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return "", err
	}
	r, _ := result.(*ci.LookupGroupNameResponse)
	return r.Name, nil
}

// ListGroups retrieves a list of groups
func ListGroups(parent, view, fields string, cap int) (<-chan *ci.Group, <-chan error) {
	srv := getGroupsService()
	c := srv.List().Parent(parent).PageSize(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if view != "" {
		c.View(view)
	}
	ch := make(chan *ci.Group, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.ListGroupsResponse) error {
			for i := range response.Groups {
				ch <- response.Groups[i]
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

// SearchGroups searches for Groups matching a specified query.
func SearchGroups(query, view, fields string, cap int) (<-chan *ci.Group, <-chan error) {
	srv := getGroupsService()
	c := srv.Search().Query(query).PageSize(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if view != "" {
		c.View(view)
	}
	ch := make(chan *ci.Group, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.SearchGroupsResponse) error {
			for i := range response.Groups {
				ch <- response.Groups[i]
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
