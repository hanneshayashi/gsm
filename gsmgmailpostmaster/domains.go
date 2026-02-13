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

package gsmgmailpostmaster

import (
	"errors"
	"context"
	"fmt"
	"iter"

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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r, nil
}

// ListDomains lists the domains that have been registered by the client.
// The order of domains in the response is unspecified and non-deterministic.
// Newly created domains will not necessarily be added to the end of this list.
func ListDomains(fields string) iter.Seq2[*gmailpostmastertools.Domain, error] {
	srv := getDomainsService()
	c := srv.List()
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*gmailpostmastertools.Domain, error) bool) {
		e := c.Pages(context.Background(), func(response *gmailpostmastertools.ListDomainsResponse) error {
			for i := range response.Domains {
				if !yield(response.Domains[i], nil) {
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
