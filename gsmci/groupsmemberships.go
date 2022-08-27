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

package gsmci

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// ListMembers lists the members of a group
func ListMembers(parent, fields, view string, cap int) (<-chan *ci.Membership, <-chan error) {
	srv := getGroupsMembershipsService()
	c := srv.List(parent).PageSize(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if view != "" {
		c.View(view)
	}
	ch := make(chan *ci.Membership, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.ListMembershipsResponse) error {
			for i := range response.Memberships {
				ch <- response.Memberships[i]
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

// CheckTransitiveMembership checks a potential member for membership in a group.
func CheckTransitiveMembership(parent, query string) (bool, error) {
	srv := getGroupsMembershipsService()
	c := srv.CheckTransitiveMembership(parent).Query(query)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, query), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	r, ok := result.(*ci.CheckTransitiveMembershipResponse)
	if !ok {
		return false, fmt.Errorf("result unknown")
	}
	return r.HasMembership, nil
}

// CreateMembership creates a Membership.
func CreateMembership(parent, fields string, membership *ci.Membership) (*googleapi.RawMessage, error) {
	srv := getGroupsMembershipsService()
	c := srv.Create(parent, membership)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, membership.PreferredMemberKey.Id), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return &r.Response, nil
}

// DeleteMembership deletes a Membership.
func DeleteMembership(name string) (bool, error) {
	srv := getGroupsMembershipsService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetMembership retrieves a Membership.
func GetMembership(name, fields string) (*ci.Membership, error) {
	srv := getGroupsMembershipsService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Membership)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// GetMembershipGraph gets a membership graph of just a member or both a member and a group.
// Given a member, the response will contain all membership paths from the member.
// Given both a group and a member, the response will contain all membership paths between the group and the member.
func GetMembershipGraph(parent, query, fields string) (*googleapi.RawMessage, error) {
	srv := getGroupsMembershipsService()
	c := srv.GetMembershipGraph(parent).Query(query)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, query), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.Operation)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return &r.Response, nil
}

// LookupMembership looks up the resource name of a Membership by its EntityKey.
func LookupMembership(parent, memberKeyID, memberKeyNamespace string) (string, error) {
	srv := getGroupsMembershipsService()
	c := srv.Lookup(parent).MemberKeyId(memberKeyID).MemberKeyNamespace(memberKeyNamespace)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, memberKeyID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return "", err
	}
	r, ok := result.(*ci.LookupMembershipNameResponse)
	if !ok {
		return "", fmt.Errorf("result unknown")
	}
	return r.Name, nil
}

// ModifyMembershipRoles modifies the MembershipRoles of a Membership.
func ModifyMembershipRoles(name, fields string, modifyMembershipRolesRequest *ci.ModifyMembershipRolesRequest) (*ci.Membership, error) {
	srv := getGroupsMembershipsService()
	c := srv.ModifyMembershipRoles(name, modifyMembershipRolesRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.ModifyMembershipRolesResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r.Membership, nil
}

// SearchTransitiveGroups searches transitive groups of a member.
func SearchTransitiveGroups(parent, query, fields string, cap int) (<-chan *ci.GroupRelation, <-chan error) {
	srv := getGroupsMembershipsService()
	c := srv.SearchTransitiveGroups(parent).Query(query).PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *ci.GroupRelation, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.SearchTransitiveGroupsResponse) error {
			for i := range response.Memberships {
				ch <- response.Memberships[i]
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

// SearchTransitiveMemberships search transitive memberships of a group.
func SearchTransitiveMemberships(parent, fields string, cap int) (<-chan *ci.MemberRelation, <-chan error) {
	srv := getGroupsMembershipsService()
	c := srv.SearchTransitiveMemberships(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *ci.MemberRelation, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.SearchTransitiveMembershipsResponse) error {
			for i := range response.Memberships {
				ch <- response.Memberships[i]
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
