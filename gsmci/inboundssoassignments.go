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

package gsmci

import (
	"context"
	"fmt"

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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(assignment.TargetGroup, assignment.TargetOrgUnit, assignment.SamlSsoInfo.InboundSamlSsoProfile), func() (any, error) {
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

// DeleteSsoAssignment deletes an InboundSsoAssignment.
func DeleteSsoAssignment(name string) (bool, error) {
	srv := getInboundSsoAssignmentsService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*ci.InboundSsoAssignment)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListSsoAssignment lists the InboundSsoAssignments for a Customer.
func ListSsoAssignment(filter, fields string, cap int) (<-chan *ci.InboundSsoAssignment, <-chan error) {
	srv := getInboundSsoAssignmentsService()
	c := srv.List().PageSize(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c.Filter(fields)
	}
	ch := make(chan *ci.InboundSsoAssignment, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *ci.ListInboundSsoAssignmentsResponse) error {
			for i := range response.InboundSsoAssignments {
				ch <- response.InboundSsoAssignments[i]
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

// PatchSsoAssignment patchs an InboundSsoAssignment for users and devices in a Customer under a given Group or OrgUnit.
func PatchSsoAssignment(name, updateMask, fields string, assignment *ci.InboundSsoAssignment) (*googleapi.RawMessage, error) {
	srv := getInboundSsoAssignmentsService()
	c := srv.Patch(name, assignment).UpdateMask(updateMask)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
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
