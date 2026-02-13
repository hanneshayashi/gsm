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
	"iter"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

// ListPeopleConnections provides a list of the authenticated user's contacts.
func ListPeopleConnections(resourceName, personFields, sources, sortOrder, fields string) iter.Seq2[*people.Person, error] {
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
	return func(yield func(*people.Person, error) bool) {
		e := c.Pages(context.Background(), func(response *people.ListConnectionsResponse) error {
			for i := range response.Connections {
				if !yield(response.Connections[i], nil) {
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
