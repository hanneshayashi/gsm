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

package gsmpeople

import (
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

// CreateContact creates a new contact and returns the person resource for that contact.
func CreateContact(person *people.Person, personFields, sources, fields string) (*people.Person, error) {
	srv := getpService()
	c := srv.CreateContact(person).PersonFields(personFields)
	if sources != "" {
		c.Sources(sources)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey("Create Contact"), err)
	}
	return r, nil
}

// DeleteContact deletes a contact person. Any non-contact data will not be deleted.
func DeleteContact(resourceName string) (bool, error) {
	srv := getpService()
	c := srv.DeleteContact(resourceName)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(resourceName), err)
	}
	return true, nil
}

// DeleteContactPhoto deletes a contact's photo.
func DeleteContactPhoto(resourceName, personFields, sources, fields string) (bool, error) {
	srv := getpService()
	c := srv.DeleteContactPhoto(resourceName)
	if personFields != "" {
		c.PersonFields(personFields)
	}
	if sources != "" {
		c.Sources(sources)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(resourceName), err)
	}
	return true, nil
}

// GetContact provides information about a person by specifying a resource name.
// Use people/me to indicate the authenticated user.
func GetContact(resourceName, personFields, sources, fields string) (*people.Person, error) {
	srv := getpService()
	c := srv.Get(resourceName)
	if personFields != "" {
		c.PersonFields(personFields)
	}
	if sources != "" {
		c.Sources(sources)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(resourceName), err)
	}
	return r, nil
}

// GetContactsBatch provides information about a list of specific people by specifying a list of requested resource names.
// Use people/me to indicate the authenticated user.
func GetContactsBatch(resourceNames []string, personFields, sources, fields string) (*people.GetPeopleResponse, error) {
	srv := getpService()
	c := srv.GetBatchGet().PersonFields(personFields).ResourceNames(resourceNames...)
	if sources != "" {
		c.Sources(sources)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(resourceNames...), err)
	}
	return r, nil
}

// ListDirectoryPeople provides a list of domain profiles and domain contacts in the authenticated user's domain directory.
func ListDirectoryPeople(readMask, sources, fields string, mergeSources []string) iter.Seq2[*people.Person, error] {
	srv := getpService()
	c := srv.ListDirectoryPeople().ReadMask(readMask).Sources(sources)
	if mergeSources != nil {
		c.MergeSources(mergeSources...)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*people.Person, error) bool) {
		e := c.Pages(context.Background(), func(response *people.ListDirectoryPeopleResponse) error {
			for i := range response.People {
				if !yield(response.People[i], nil) {
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

// SearchDirectoryPeople provides a list of domain profiles and domain contacts in the authenticated user's domain directory that match the search query.
func SearchDirectoryPeople(readMask, sources, query, fields string, mergeSources []string) iter.Seq2[*people.Person, error] {
	srv := getpService()
	c := srv.SearchDirectoryPeople().ReadMask(readMask).Sources(sources).Query(query).PageSize(500)
	if mergeSources != nil {
		c.MergeSources(mergeSources...)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*people.Person, error) bool) {
		e := c.Pages(context.Background(), func(response *people.SearchDirectoryPeopleResponse) error {
			for i := range response.People {
				if !yield(response.People[i], nil) {
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

// UpdateContact updates a new contact and returns the person resource for that contact.
func UpdateContact(resourceName, updatePersonFields, personFields, sources, fields string, person *people.Person) (*people.Person, error) {
	srv := getpService()
	c := srv.UpdateContact(resourceName, person).UpdatePersonFields(updatePersonFields)
	if personFields != "" {
		c.PersonFields(personFields)
	}
	if sources != "" {
		c.Sources(sources)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(resourceName), err)
	}
	return r, nil
}

// UpdateContactPhoto updates a contact's photo.
func UpdateContactPhoto(resourceName, fields string, updateContactPhotoRequest *people.UpdateContactPhotoRequest) (*people.Person, error) {
	srv := getpService()
	c := srv.UpdateContactPhoto(resourceName, updateContactPhotoRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(resourceName), err)
	}
	return r.Person, nil
}
