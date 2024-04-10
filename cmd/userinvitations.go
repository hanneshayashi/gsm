/*
Copyright © 2020-2024 Hannes Hayashi

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

	"github.com/spf13/cobra"
	ci "google.golang.org/api/cloudidentity/v1"
)

// userInvitationsCmd represents the userInvitations command
var userInvitationsCmd = &cobra.Command{
	Use:               "userInvitations",
	Short:             "Manage user invitations for unmanaged accounts (Part of Cloud Identity Beta API)",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/customers.userinvitations",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var userInvitationFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor:   []string{"cancel", "get", "isInvitableUser", "send"},
		Type:           "string",
		Description:    `UserInvitation name in the format customers/{customer}/userinvitations/{user_email_address}`,
		Required:       []string{"cancel", "get", "isInvitableUser", "send"},
		ExcludeFromAll: true,
	},
	"parent": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `The customer ID of the Google Workspace or Cloud Identity account the UserInvitation resources are associated with.`,
	},
	"filter": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `A query string for filtering UserInvitation results by their current state, in the format: "state=='invited'".`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The sort order of the list results.

You can sort the results in descending order based on either email or last update timestamp but not both, using orderBy="email desc". Currently, sorting is supported for updateTime asc, updateTime desc, email asc, and email desc.

If not specified, results will be returned based on email asc order.`,
	},
	"fields": {
		AvailableFor: []string{"get", "list", "send"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var userInvitationFlagsALL = gsmhelpers.GetAllFlags(userInvitationFlags)

func init() {
	rootCmd.AddCommand(userInvitationsCmd)
}

func mapToCancelUserInvitationRequest(_ map[string]*gsmhelpers.Value) (*ci.CancelUserInvitationRequest, error) {
	cancelUserInvitationRequest := &ci.CancelUserInvitationRequest{}
	return cancelUserInvitationRequest, nil
}

func mapToSendUserInvitationRequest(_ map[string]*gsmhelpers.Value) (*ci.SendUserInvitationRequest, error) {
	sendUserInvitationRequest := &ci.SendUserInvitationRequest{}
	return sendUserInvitationRequest, nil
}
