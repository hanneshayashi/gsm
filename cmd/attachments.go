/*
Package cmd contains the commands available to the end user
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
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// attachmentsCmd represents the attachments command
var attachmentsCmd = &cobra.Command{
	Use:   "attachments",
	Short: "Manage (get..) message attachements (Part of Gmail API)",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.messages.attachments",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var attachmentFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The user's email address.
The special value me can be used to indicate the authenticated user.`,
		Defaults: map[string]interface{}{"get": "me"},
	},
	"messageId": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  "The ID of the message containing the attachment.",
		Required:     []string{"get"},
	},
	"id": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  "The ID of the attachment.",
		Required:     []string{"get"},
	},
	"fields": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var attachmentFlagsALL = gsmhelpers.GetAllFlags(attachmentFlags)

func init() {
	rootCmd.AddCommand(attachmentsCmd)
}
