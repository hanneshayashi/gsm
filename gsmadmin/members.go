/*
Copyright © 2020-2023 Hannes Hayashi

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
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteMember removes a member from a group.
func DeleteMember(groupKey, memberKey string) (bool, error) {
	srv := getMembersService()
	c := srv.Delete(groupKey, memberKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(groupKey, memberKey), func() error {
		return c.Do()
	})
	return result, err
}

// GetMember retrieves a group member's properties.
func GetMember(groupKey, memberKey, fields string) (*admin.Member, error) {
	srv := getMembersService()
	c := srv.Get(groupKey, memberKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey, memberKey), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Member)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// HasMember checks whether the given user is a member of the group. Membership can be direct or nested.
func HasMember(groupKey, memberKey string) (bool, error) {
	srv := getMembersService()
	c := srv.HasMember(groupKey, memberKey)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey, memberKey), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	r, ok := result.(*admin.MembersHasMember)
	if !ok {
		return false, fmt.Errorf("result unknown")
	}
	return r.IsMember, nil
}

// InsertMember adds a user to the specified group.
func InsertMember(groupKey, fields string, member *admin.Member) (*admin.Member, error) {
	srv := getMembersService()
	c := srv.Insert(groupKey, member)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey, member.Email), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Member)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListMembers retrieves a paginated list of all members in a group.
func ListMembers(groupKey, roles, fields string, includeDerivedMembership bool, cap int) (<-chan *admin.Member, <-chan error) {
	srv := getMembersService()
	c := srv.List(groupKey).IncludeDerivedMembership(includeDerivedMembership).MaxResults(200)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if roles != "" {
		c = c.Roles(roles)
	}
	ch := make(chan *admin.Member, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *admin.Members) error {
			for i := range response.Members {
				ch <- response.Members[i]
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

// PatchMember updates the membership properties of a user in the specified group. This method supports patch semantics.
func PatchMember(groupKey, memberKey, fields string, member *admin.Member) (*admin.Member, error) {
	srv := getMembersService()
	c := srv.Patch(groupKey, memberKey, member)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(groupKey, memberKey), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Member)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
