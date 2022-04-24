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

package gsmpeople

import (
	"context"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

// BatchGetContactGroups gets a list of contact groups owned by the authenticated user by specifying a list of contact group resource names.
func BatchGetContactGroups(resourceNames []string, maxMembers int64, fields string) (*people.BatchGetContactGroupsResponse, error) {
	srv := getContactGroupsService()
	c := srv.BatchGet().ResourceNames(resourceNames...)
	if maxMembers != 0 {
		c.MaxMembers(maxMembers)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(resourceNames...), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.BatchGetContactGroupsResponse)
	return r, nil
}

// CreateContactGroup creates a new contact group owned by the authenticated user.
func CreateContactGroup(createContactGroupRequest *people.CreateContactGroupRequest, fields string) (*people.ContactGroup, error) {
	srv := getContactGroupsService()
	c := srv.Create(createContactGroupRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(createContactGroupRequest.ContactGroup.Name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.ContactGroup)
	return r, nil
}

// DeleteContactGroup deletes an existing contact group owned by the authenticated user by specifying a contact group resource name.
func DeleteContactGroup(resourceName string, deleteContacts bool) (bool, error) {
	srv := getContactGroupsService()
	c := srv.Delete(resourceName).DeleteContacts(deleteContacts)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(resourceName), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetContactGroup gets a specific contact group owned by the authenticated user by specifying a contact group resource name.
func GetContactGroup(resourceName, fields string, maxMembers int64) (*people.ContactGroup, error) {
	srv := getContactGroupsService()
	c := srv.Get(resourceName)
	if maxMembers != 0 {
		c.MaxMembers(maxMembers)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(resourceName), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.ContactGroup)
	return r, nil
}

// ListContactGroups lists all contact groups owned by the authenticated user.
// Members of the contact groups are not populated.
func ListContactGroups(fields string, cap int) (<-chan *people.ContactGroup, <-chan error) {
	srv := getContactGroupsService()
	c := srv.List().PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *people.ContactGroup, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *people.ListContactGroupsResponse) error {
			for i := range response.ContactGroups {
				ch <- response.ContactGroups[i]
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

// UpdateContactGroup updates a new contact group owned by the authenticated user.
func UpdateContactGroup(resourceName, fields string, updateContactGroupRequest *people.UpdateContactGroupRequest) (*people.ContactGroup, error) {
	srv := getContactGroupsService()
	c := srv.Update(resourceName, updateContactGroupRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(resourceName), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.ContactGroup)
	return r, nil
}
