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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

func makeListPeopleConnectionsCallAndAppend(c *people.PeopleConnectionsListCall, ps []*people.Person, errKey string) ([]*people.Person, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.ListConnectionsResponse)
	ps = append(ps, r.Connections...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		ps, err = makeListPeopleConnectionsCallAndAppend(c, ps, errKey)
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}

// ListPeopleConnections provides a list of the authenticated user's contacts.
func ListPeopleConnections(resourceName, personFields, sources, fields string) ([]*people.Person, error) {
	srv := getPeopleConnectionsService()
	c := srv.List(resourceName)
	if personFields != "" {
		c.PersonFields(personFields)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var people []*people.Person
	people, err := makeListPeopleConnectionsCallAndAppend(c, people, gsmhelpers.FormatErrorKey(resourceName))
	return people, err
}
