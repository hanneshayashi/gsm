/*
Package gsmpeople implements the People API
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
package gsmpeople

import (
	"gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/people/v1"
)

// ModifyContactGroupMembers modifies the members of a contact group owned by the authenticated user.
func ModifyContactGroupMembers(resourceName, fields string, modifyContactGroupMembersRequest *people.ModifyContactGroupMembersRequest) (*people.ModifyContactGroupMembersResponse, error) {
	srv := getContactGroupsMembersService()
	c := srv.Modify(resourceName, modifyContactGroupMembersRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(resourceName), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*people.ModifyContactGroupMembersResponse)
	return r, nil
}
