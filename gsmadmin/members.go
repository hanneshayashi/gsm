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
package gsmadmin

import (
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteMember removes a member from a group.
func DeleteMember(groupKey, memberKey string) (bool, error) {
	srv := getMembersService()
	err := srv.Delete(groupKey, memberKey).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetMember retrieves a group member's properties.
func GetMember(groupKey, memberKey, fields string) (*admin.Member, error) {
	srv := getMembersService()
	c := srv.Get(groupKey, memberKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// HasMember checks whether the given user is a member of the group. Membership can be direct or nested.
func HasMember(groupKey, memberKey string) (bool, error) {
	srv := getMembersService()
	r, err := srv.HasMember(groupKey, memberKey).Do()
	if err != nil {
		return false, err
	}
	return r.IsMember, nil
}

// InsertMember adds a user to the specified group.
func InsertMember(key, fields string, Member *admin.Member) (*admin.Member, error) {
	srv := getMembersService()
	c := srv.Insert(key, Member)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

func makeListMembersCallAndAppend(c *admin.MembersListCall, members []*admin.Member) ([]*admin.Member, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, m := range r.Members {
		members = append(members, m)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		members, err = makeListMembersCallAndAppend(c, members)
	}
	return members, err
}

// ListMembers retrieves a paginated list of all members in a group.
func ListMembers(key, roles, fields string, includeDerivedMembership bool) ([]*admin.Member, error) {
	srv := getMembersService()
	c := srv.List(key).IncludeDerivedMembership(includeDerivedMembership)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if roles != "" {
		c = c.Roles(roles)
	}
	var members []*admin.Member
	members, err := makeListMembersCallAndAppend(c, members)
	return members, err
}

// PatchMember updates the membership properties of a user in the specified group. This method supports patch semantics.
func PatchMember(groupKey, memberKey, fields string, member *admin.Member) (*admin.Member, error) {
	srv := getMembersService()
	c := srv.Patch(groupKey, memberKey, member)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
