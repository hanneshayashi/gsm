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

// ListPeopleConnections provides a list of the authenticated user's contacts.
func ListPeopleConnections(resourceName, personFields, sources, sortOrder, fields string, cap int) (<-chan *people.Person, <-chan error) {
	srv := getPeopleConnectionsService()
	c := srv.List(resourceName)
	if personFields != "" {
		c.PersonFields(personFields)
	}
	if sortOrder != "" {
		c.SortOrder(sortOrder)
	}
	if sources != "" {
		c.Sources(sources)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *people.Person, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *people.ListConnectionsResponse) error {
			for i := range response.Connections {
				ch <- response.Connections[i]
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
