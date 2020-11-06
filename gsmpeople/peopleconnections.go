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
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

func makeListPeopleConnectionsCallAndAppend(c *people.PeopleConnectionsListCall, people []*people.Person) ([]*people.Person, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, c := range r.Connections {
		people = append(people, c)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		people, err = makeListPeopleConnectionsCallAndAppend(c, people)
		if err != nil {
			return nil, err
		}
	}
	return people, nil
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
	people, err := makeListPeopleConnectionsCallAndAppend(c, people)
	return people, err
}
