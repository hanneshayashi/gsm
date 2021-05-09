/*
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

package gsmgmailpostmaster

import (
	"context"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmailpostmastertools/v1"
	"google.golang.org/api/googleapi"
)

// GetDomain Gets a specific domain registered by the client.
// Returns NOT_FOUND if the domain does not exist.
func GetDomain(name, fields string) (*gmailpostmastertools.Domain, error) {
	srv := getDomainsService()
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
	r, _ := result.(*gmailpostmastertools.Domain)
	return r, nil
}

// ListDomains lists the domains that have been registered by the client.
// The order of domains in the response is unspecified and non-deterministic.
// Newly created domains will not necessarily be added to the end of this list.
func ListDomains(fields string, cap int) (<-chan *gmailpostmastertools.Domain, <-chan error) {
	srv := getDomainsService()
	c := srv.List()
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *gmailpostmastertools.Domain, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *gmailpostmastertools.ListDomainsResponse) error {
			for i := range response.Domains {
				ch <- response.Domains[i]
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
