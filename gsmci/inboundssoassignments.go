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

// CreateSsoAssignment creates an InboundSsoAssignment for users and devices in a Customer under a given Group or OrgUnit.
func CreateSsoAssignment(fields string, assignment *ci.InboundSsoAssignment) (*googleapi.RawMessage, error) {
	srv := getInboundSsoAssignmentsService()
	c := srv.Create(assignment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(assignment.TargetGroup, assignment.TargetOrgUnit, assignment.SamlSsoInfo.InboundSamlSsoProfile), err)
	}
	return &r.Response, nil
}

// DeleteSsoAssignment deletes an InboundSsoAssignment.
func DeleteSsoAssignment(name string) (bool, error) {
	srv := getInboundSsoAssignmentsService()
	c := srv.Delete(name)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return true, nil
}

// GetSsoAssignment gets an InboundSsoAssignment.
func GetSsoAssignment(name, fields string) (*ci.InboundSsoAssignment, error) {
	srv := getInboundSsoAssignmentsService()
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

// ListSsoAssignment lists the InboundSsoAssignments for a Customer.
func ListSsoAssignment(filter, fields string) iter.Seq2[*ci.InboundSsoAssignment, error] {
	srv := getInboundSsoAssignmentsService()
	c := srv.List().PageSize(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c.Filter(fields)
	}
	return func(yield func(*ci.InboundSsoAssignment, error) bool) {
		e := c.Pages(context.Background(), func(response *ci.ListInboundSsoAssignmentsResponse) error {
			for i := range response.InboundSsoAssignments {
				if !yield(response.InboundSsoAssignments[i], nil) {
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

// PatchSsoAssignment patches an InboundSsoAssignment for users and devices in a Customer under a given Group or OrgUnit.
func PatchSsoAssignment(name, updateMask, fields string, assignment *ci.InboundSsoAssignment) (*googleapi.RawMessage, error) {
	srv := getInboundSsoAssignmentsService()
	c := srv.Patch(name, assignment).UpdateMask(updateMask)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return &r.Response, nil
}
