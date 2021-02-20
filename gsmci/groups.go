/*
Package gsmci implements the Cloud Identity API
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
package gsmci

import (
	"encoding/json"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// CreateGroup creates a group.
func CreateGroup(group *ci.Group, initialGroupConfig, fields string) (map[string]interface{}, error) {
	srv := getGroupsService()
	c := srv.Create(group).InitialGroupConfig(initialGroupConfig)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(group.GroupKey.Id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// DeleteGroup deletes a group.
func DeleteGroup(name string) (bool, error) {
	srv := getGroupsService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// PatchGroup updates a group using patch semantics.
func PatchGroup(name, updateMask, fields string, group *ci.Group) (map[string]interface{}, error) {
	srv := getGroupsService()
	c := srv.Patch(name, group).UpdateMask(updateMask)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetGroup retrieves a group.
func GetGroup(name, fields string) (*ci.Group, error) {
	srv := getGroupsService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(email), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return "", err
	}
	r, _ := result.(*ci.LookupGroupNameResponse)
	return r.Name, nil
}

func listGroups(c *ci.GroupsListCall, ch chan *ci.Group, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*ci.ListGroupsResponse)
	for _, i := range r.Groups {
		ch <- i
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listGroups(c, ch, errKey)
	}
	return err
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
		e := listGroups(c, ch, gsmhelpers.FormatErrorKey(parent))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

func searchGroups(c *ci.GroupsSearchCall, ch chan *ci.Group, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*ci.SearchGroupsResponse)
	for _, i := range r.Groups {
		ch <- i
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = searchGroups(c, ch, errKey)
	}
	return err
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
		e := searchGroups(c, ch, gsmhelpers.FormatErrorKey(query))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}
