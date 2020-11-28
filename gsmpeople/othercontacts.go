/*
Package gsmpeople implements the People API
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
	"gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

// CopyOtherContactToMyContactsGroup copies an "Other contact" to a new contact in the user's "myContacts" group
func CopyOtherContactToMyContactsGroup(resourceName, fields string, sources []string, copyOtherContactToMyContactsGroupRequest *people.CopyOtherContactToMyContactsGroupRequest) (*people.Person, error) {
	srv := getOtherContactsService()
	c := srv.CopyOtherContactToMyContactsGroup(resourceName, copyOtherContactToMyContactsGroupRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(resourceName), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.Person)
	return r, nil
}

func makeListOtherContactsCallAndAppend(c *people.OtherContactsListCall, otherContacts []*people.Person, errKey string) ([]*people.Person, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.ListOtherContactsResponse)
	for _, o := range r.OtherContacts {
		otherContacts = append(otherContacts, o)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		otherContacts, err = makeListOtherContactsCallAndAppend(c, otherContacts, errKey)
		if err != nil {
			return nil, err
		}
	}
	return otherContacts, nil
}

// ListOtherContacts lists all "Other contacts", that is contacts that are not in a contact group.
// "Other contacts" are typically auto created contacts from interactions.
func ListOtherContacts(readMask, fields string) ([]*people.Person, error) {
	srv := getOtherContactsService()
	c := srv.List().ReadMask(readMask)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var otherContacts []*people.Person
	otherContacts, err := makeListOtherContactsCallAndAppend(c, otherContacts, gsmhelpers.FormatErrorKey("List other contacts"))
	return otherContacts, err
}
