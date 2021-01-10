/*
Package gsmpeople implements the People API
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
package gsmpeople

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

func listPeopleConnections(c *people.PeopleConnectionsListCall, ch chan *people.Person, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*people.ListConnectionsResponse)
	for _, i := range r.Connections {
		ch <- i
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listPeopleConnections(c, ch, errKey)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListPeopleConnections provides a list of the authenticated user's contacts.
func ListPeopleConnections(resourceName, personFields, sources, fields string, cap int) (<-chan *people.Person, <-chan error) {
	srv := getPeopleConnectionsService()
	c := srv.List(resourceName)
	if personFields != "" {
		c.PersonFields(personFields)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *people.Person, cap)
	err := make(chan error, 1)
	go func() {
		e := listPeopleConnections(c, ch, gsmhelpers.FormatErrorKey(resourceName))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}
