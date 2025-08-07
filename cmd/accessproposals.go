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

package cmd

import (
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/drive/v3"

	"github.com/spf13/cobra"
)

// accessProposalsCmd represents the accessProposals command
var accessProposalsCmd = &cobra.Command{
	Use:   "accessProposals",
	Short: "Manage Access Proposals on a file (Part of Drive API)",
	Long: `An access proposal is a proposal from a requester to an approver to grant a recipient access to a Google Drive item.
See https://developers.google.com/workspace/drive/api/guides/pending-access#:~:text=An%20access%20proposal%20is%20a,access%20proposals%20across%20Drive%20files.

This API only works with the currently authenticated user!
Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/accessproposals`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var accessProposalsFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fields": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
	"fileId": {
		AvailableFor: []string{"get", "list", "resolve"},
		Type:         "string",
		Description:  `The id of the item the request is on.`,
		Required:     []string{"get", "list", "resolve"},
	},
	"proposalId": {
		AvailableFor: []string{"get", "resolve"},
		Type:         "string",
		Description:  `The id of the access proposal.`,
		Required:     []string{"get", "resolve"},
	},
	"role": {
		AvailableFor: []string{"resolve"},
		Type:         "stringSlice",
		Description: `The roles the approver has allowed, if any. Note: This field is required for the ACCEPT action.
This flag can be used multiple times.`,
	},
	"view": {
		AvailableFor: []string{"resolve"},
		Type:         "string",
		Description: `Indicates the view for this access proposal.
This should only be set when the proposal belongs to a view.
'published' is the only supported value.`,
	},
	"action": {
		AvailableFor: []string{"resolve"},
		Type:         "string",
		Description: `The action to take on the AccessProposal.
Must be one of the following:
ACCEPT  The user accepts the proposal. Note: If this action is used, the role field must have at least one value.
DENY    The user denies the proposal`,
		Required: []string{"resolve"},
	},
	"sendNotification": {
		AvailableFor: []string{"resolve"},
		Type:         "bool",
		Description:  `Whether to send an email to the requester when the AccessProposal is denied or accepted.`,
	},
}

func init() {
	rootCmd.AddCommand(accessProposalsCmd)
}

func mapToResolveAccessProposalRequest(flags map[string]*gsmhelpers.Value) (*drive.ResolveAccessProposalRequest, error) {
	request := &drive.ResolveAccessProposalRequest{}
	if flags["role"].IsSet() {
		roles := flags["role"].GetStringSlice()
		if len(roles) > 0 {
			request.Role = roles
		} else {
			request.ForceSendFields = append(request.ForceSendFields, "Role")
		}
	}
	if flags["view"].IsSet() {
		request.View = flags["view"].GetString()
		if request.View == "" {
			request.ForceSendFields = append(request.ForceSendFields, "View")
		}
	}
	if flags["action"].IsSet() {
		request.Action = flags["action"].GetString()
		if request.Action == "" {
			request.ForceSendFields = append(request.ForceSendFields, "Action")
		}
	}
	if flags["sendNotification"].IsSet() {
		request.SendNotification = flags["sendNotification"].GetBool()
		if !request.SendNotification {
			request.ForceSendFields = append(request.ForceSendFields, "SendNotification")
		}
	}
	return request, nil
}
