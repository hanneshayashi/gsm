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
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

// CopyOtherContactToMyContactsGroup copies an "Other contact" to a new contact in the user's "myContacts" group
func CopyOtherContactToMyContactsGroup(resourceName, fields string, copyOtherContactToMyContactsGroupRequest *people.CopyOtherContactToMyContactsGroupRequest) (*people.Person, error) {
	srv := getOtherContactsService()
	c := srv.CopyOtherContactToMyContactsGroup(resourceName, copyOtherContactToMyContactsGroupRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(resourceName), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*people.Person)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListOtherContacts lists all "Other contacts", that is contacts that are not in a contact group.
// "Other contacts" are typically auto created contacts from interactions.
func ListOtherContacts(readMask, fields string, cap int) (<-chan *people.Person, <-chan error) {
	srv := getOtherContactsService()
	c := srv.List().ReadMask(readMask).PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *people.Person, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *people.ListOtherContactsResponse) error {
			for i := range response.OtherContacts {
				ch <- response.OtherContacts[i]
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
