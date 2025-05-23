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

package gsmdrive

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// GetAccessProposal retrieves an AccessProposal by ID.
func GetAccessProposal(filedId, proposalId, fields string) (*drive.AccessProposal, error) {
	srv := getAccessProposalsService()
	c := srv.Get(filedId, proposalId)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(filedId, proposalId), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.AccessProposal)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListAccessProposals list the AccessProposals on a file.
// Note: Only approvers are able to list AccessProposals on a file. If the user is not an approver, returns a 403.
func ListAccessProposals(fileId, fields string, cap int) (<-chan *drive.AccessProposal, <-chan error) {
	srv := getAccessProposalsService()
	c := srv.List(fileId)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *drive.AccessProposal, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drive.ListAccessProposalsResponse) error {
			for i := range response.AccessProposals {
				ch <- response.AccessProposals[i]
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

// ResolveAccessProposal is used to approve or deny an Access Proposal.
func ResolveAccessProposal(filedId, proposalId string, request *drive.ResolveAccessProposalRequest) (bool, error) {
	srv := getAccessProposalsService()
	c := srv.Resolve(filedId, proposalId, request)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(filedId, proposalId), func() error {
		return c.Do()
	})
	return result, err
}
