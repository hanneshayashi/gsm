/*
Package gsmci implements the Cloud Identity (Beta) API
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
	"encoding/json"
	"gsm/gsmhelpers"

	ci "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/googleapi"
)

func makeListMembersCallAndAppend(c *ci.GroupsMembershipsListCall, members []*ci.Membership, errKey string) ([]*ci.Membership, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.ListMembershipsResponse)
	for _, m := range r.Memberships {
		members = append(members, m)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		members, err = makeListMembersCallAndAppend(c, members, errKey)
	}
	return members, err
}

// ListMembers lists the members of a group
func ListMembers(parent, fields, view string) ([]*ci.Membership, error) {
	srv := getGroupsMembershipsService()
	c := srv.List(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if view != "" {
		c.View(view)
	}
	var members []*ci.Membership
	members, err := makeListMembersCallAndAppend(c, members, gsmhelpers.FormatErrorKey(parent))
	return members, err
}

// CheckTransitiveMembership checks a potential member for membership in a group.
func CheckTransitiveMembership(parent, query string) (bool, error) {
	srv := getGroupsMembershipsService()
	c := srv.CheckTransitiveMembership(parent).Query(query)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, query), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	r, _ := result.(*ci.CheckTransitiveMembershipResponse)
	return r.HasMembership, nil
}

// CreateMembership creates a Membership.
func CreateMembership(parent, fields string, membership *ci.Membership) (googleapi.RawMessage, error) {
	srv := getGroupsMembershipsService()
	c := srv.Create(parent, membership)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, membership.MemberKey.Id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	return r.Response, nil
}

// DeleteMembership deletes a Membership.
func DeleteMembership(name string) (bool, error) {
	srv := getGroupsMembershipsService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Membership)
	return r, nil
}

// GetMembershipGraph gets a membership graph of just a member or both a member and a group.
// Given a member, the response will contain all membership paths from the member.
// Given both a group and a member, the response will contain all membership paths between the group and the member.
func GetMembershipGraph(parent, query, fields string) (map[string]interface{}, error) {
	srv := getGroupsMembershipsService()
	c := srv.GetMembershipGraph(parent).Query(query)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, query), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.Operation)
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// LookupMembership looks up the resource name of a Membership by its EntityKey.
func LookupMembership(parent, memberKeyID, memberKeyNamespace string) (string, error) {
	srv := getGroupsMembershipsService()
	c := srv.Lookup(parent).MemberKeyId(memberKeyID).MemberKeyNamespace(memberKeyNamespace)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent, memberKeyID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return "", err
	}
	r, _ := result.(*ci.LookupMembershipNameResponse)
	return r.Name, nil
}

// ModifyMembershipRoles modifies the MembershipRoles of a Membership.
func ModifyMembershipRoles(name, fields string, modifyMembershipRolesRequest *ci.ModifyMembershipRolesRequest) (*ci.Membership, error) {
	srv := getGroupsMembershipsService()
	c := srv.ModifyMembershipRoles(name, modifyMembershipRolesRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.ModifyMembershipRolesResponse)
	return r.Membership, nil
}

func makeSearchTransitiveGroupsCallAndAppend(c *ci.GroupsMembershipsSearchTransitiveGroupsCall, members []*ci.GroupRelation, errKey string) ([]*ci.GroupRelation, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.SearchTransitiveGroupsResponse)
	for _, m := range r.Memberships {
		members = append(members, m)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		members, err = makeSearchTransitiveGroupsCallAndAppend(c, members, errKey)
	}
	return members, err
}

// SearchTransitiveGroups searches transitive groups of a member.
func SearchTransitiveGroups(parent, query, fields string) ([]*ci.GroupRelation, error) {
	srv := getGroupsMembershipsService()
	c := srv.SearchTransitiveGroups(parent).Query(query)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var memberships []*ci.GroupRelation
	memberships, err := makeSearchTransitiveGroupsCallAndAppend(c, memberships, gsmhelpers.FormatErrorKey(parent, query))
	return memberships, err
}

func makeSearchTransitiveMembershipsCallAndAppend(c *ci.GroupsMembershipsSearchTransitiveMembershipsCall, members []*ci.MemberRelation, errKey string) ([]*ci.MemberRelation, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*ci.SearchTransitiveMembershipsResponse)
	for _, m := range r.Memberships {
		members = append(members, m)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		members, err = makeSearchTransitiveMembershipsCallAndAppend(c, members, errKey)
	}
	return members, err
}

// SearchTransitiveMemberships search transitive memberships of a group.
func SearchTransitiveMemberships(parent, fields string) ([]*ci.MemberRelation, error) {
	srv := getGroupsMembershipsService()
	c := srv.SearchTransitiveMemberships(parent)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var memberships []*ci.MemberRelation
	memberships, err := makeSearchTransitiveMembershipsCallAndAppend(c, memberships, gsmhelpers.FormatErrorKey(parent))
	return memberships, err
}
