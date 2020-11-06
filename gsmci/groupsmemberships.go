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
	"encoding/json"

	ci "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/googleapi"
)

func makeListMembersCallAndAppend(c *ci.GroupsMembershipsListCall, members []*ci.Membership) ([]*ci.Membership, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, m := range r.Memberships {
		members = append(members, m)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		members, err = makeListMembersCallAndAppend(c, members)
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
	members, err := makeListMembersCallAndAppend(c, members)
	return members, err
}

// CheckTransitiveMembership checks a potential member for membership in a group.
func CheckTransitiveMembership(parent, query string) (bool, error) {
	srv := getGroupsMembershipsService()
	r, err := srv.CheckTransitiveMembership(parent).Query(query).Do()
	if err != nil {
		return false, err
	}
	return r.HasMembership, err
}

// CreateMembership creates a Membership.
func CreateMembership(parent, fields string, membership *ci.Membership) (googleapi.RawMessage, error) {
	srv := getGroupsMembershipsService()
	c := srv.Create(parent, membership)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.Response, err
}

// DeleteMembership deletes a Membership.
func DeleteMembership(name string) (bool, error) {
	srv := getGroupsMembershipsService()
	_, err := srv.Delete(name).Do()
	if err != nil {
		return false, err
	}
	return true, err
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
		return nil, err
	}
	return r, err
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
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal(r.Response, &m)
	if err != nil {
		return nil, err
	}
	return m, err
}

// LookupMembership looks up the resource name of a Membership by its EntityKey.
func LookupMembership(parent, memberKeyID, memberKeyNamespace string) (string, error) {
	srv := getGroupsMembershipsService()
	r, err := srv.Lookup(parent).MemberKeyId(memberKeyID).MemberKeyNamespace(memberKeyNamespace).Do()
	if err != nil {
		return "", err
	}
	return r.Name, err
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
		return nil, err
	}
	return r.Membership, err
}

func makeSearchTransitiveGroupsCallAndAppend(c *ci.GroupsMembershipsSearchTransitiveGroupsCall, members []*ci.GroupRelation) ([]*ci.GroupRelation, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, m := range r.Memberships {
		members = append(members, m)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		members, err = makeSearchTransitiveGroupsCallAndAppend(c, members)
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
	memberships, err := makeSearchTransitiveGroupsCallAndAppend(c, memberships)
	return memberships, err
}

func makeSearchTransitiveMembershipsCallAndAppend(c *ci.GroupsMembershipsSearchTransitiveMembershipsCall, members []*ci.MemberRelation) ([]*ci.MemberRelation, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, m := range r.Memberships {
		members = append(members, m)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		members, err = makeSearchTransitiveMembershipsCallAndAppend(c, members)
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
	memberships, err := makeSearchTransitiveMembershipsCallAndAppend(c, memberships)
	return memberships, err
}
