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

package gsmci

import (
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/googleapi"
)

// ListMembers lists the members of a group
func ListMembers(parent, fields, view string) iter.Seq2[*ci.Membership, error] {
	srv := getGroupsMembershipsService()
	c := srv.List(parent).PageSize(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if view != "" {
		c.View(view)
	}
	return func(yield func(*ci.Membership, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.ListMembershipsResponse) error {
			for i := range response.Memberships {
				if !yield(response.Memberships[i], nil) {
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

// CheckTransitiveMembership checks a potential member for membership in a group.
func CheckTransitiveMembership(parent, query string) (bool, error) {
	srv := getGroupsMembershipsService()
	c := srv.CheckTransitiveMembership(parent).Query(query)
	r, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent, query), err)
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent, membership.PreferredMemberKey.Id), err)
	}
	return &r.Response, nil
}

// DeleteMembership deletes a Membership.
func DeleteMembership(name string) (bool, error) {
	srv := getGroupsMembershipsService()
	c := srv.Delete(name)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent, query), err)
	}
	return &r.Response, nil
}

// LookupMembership looks up the resource name of a Membership by its EntityKey.
func LookupMembership(parent, memberKeyID, memberKeyNamespace string) (string, error) {
	srv := getGroupsMembershipsService()
	c := srv.Lookup(parent).MemberKeyId(memberKeyID).MemberKeyNamespace(memberKeyNamespace)
	r, err := c.Do()
	if err != nil {
		return "", fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent, memberKeyID), err)
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r.Membership, nil
}

// SearchTransitiveGroups searches transitive groups of a member.
func SearchTransitiveGroups(parent, query, fields string) iter.Seq2[*ci.GroupRelation, error] {
	srv := getGroupsMembershipsService()
	c := srv.SearchTransitiveGroups(parent).Query(query).PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*ci.GroupRelation, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.SearchTransitiveGroupsResponse) error {
			for i := range response.Memberships {
				if !yield(response.Memberships[i], nil) {
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

// SearchTransitiveMemberships search transitive memberships of a group.
func SearchTransitiveMemberships(parent, fields string) iter.Seq2[*ci.MemberRelation, error] {
	srv := getGroupsMembershipsService()
	c := srv.SearchTransitiveMemberships(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*ci.MemberRelation, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.SearchTransitiveMembershipsResponse) error {
			for i := range response.Memberships {
				if !yield(response.Memberships[i], nil) {
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
