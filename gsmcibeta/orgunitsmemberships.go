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

package gsmcibeta

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	cibeta "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/googleapi"
)

// ListOrgUnitMemberships lists OrgMembership resources in an OrgUnit treated as 'parent'.
func ListOrgUnitMemberships(parent, customer, filter, fields string, cap int) (chan *cibeta.OrgMembership, chan error) {
	srv := getOrgUnitsMembershipsService()
	c := srv.List(parent).Customer(customer).PageSize(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c.Filter(filter)
	}
	ch := make(chan *cibeta.OrgMembership, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *cibeta.ListOrgMembershipsResponse) error {
			for i := range response.OrgMemberships {
				ch <- response.OrgMemberships[i]
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

// MoveOrgUnitMembership moves an OrgMembership to a new OrgUnit.
func MoveOrgUnitMemberships(name, fields string, moveOrgMembershipRequest *cibeta.MoveOrgMembershipRequest) (*googleapi.RawMessage, error) {
	srv := getOrgUnitsMembershipsService()
	c := srv.Move(name, moveOrgMembershipRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*cibeta.Operation)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return &r.Response, nil
}
