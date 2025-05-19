/*
Copyright © 2020-2025 Hannes Hayashi

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
	"google.golang.org/api/gmail/v1"
)

var draftFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"create", "delete", "get", "list", "send", "update"},
		Type:         "string",
		Description:  "The user's email address. The special value \"me\" can be used to indicate the authenticated user.",
		Defaults:     map[string]any{"create": "me", "delete": "me", "get": "me", "list": "me", "send": "me", "update": "me"},
	},
	"subject": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Subject of the (draft) message",
	},
	"html": {
		AvailableFor: []string{"create", "update"},
		Type:         "bool",
		Description:  "Send the body as HTML",
	},
	"to": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Recipient of the (draft) message",
	},
	"cc": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Copy (Cc)",
	},
	"bcc": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Blind Copy (Bcc)",
	},
	"body": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Body or content of the (draft) message",
	},
	"id": {
		AvailableFor:   []string{"delete", "get", "send", "update"},
		Type:           "string",
		Description:    "The ID of the draft.",
		Required:       []string{"delete", "get", "send", "update"},
		ExcludeFromAll: true,
	},
	"format": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The format to return the draft in.
"[MINIMAL|FULL|RAW|METADATA].
MINIMAL   - Returns only email message ID and labels; does not return the email headers, body, or payload.
FULL      - Returns the full email message data with body content parsed in the payload field; the raw field is not used. Format cannot be used when accessing the api using the gmail.metadata scope.
RAW       - Returns the full email message data with body content in the raw field as a base64url encoded string; the payload field is not used. Format cannot be used when accessing the api using the gmail.metadata scope.
METADATA  - Returns only email message ID, labels, and email headers.`,
		Defaults: map[string]any{"get": "MINIMAL"},
	},
	"q": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Only return draft messages matching the specified query.
Supports the same query format as the Gmail search box.
For example, "from:someuser@example.com rfc822msgid:<somemsgid@example.com> is:unread".`,
	},
	"includeSpamTrash": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Include drafts from SPAM and TRASH in the results.`,
	},
	"attachment": {
		AvailableFor: []string{"create", "update"},
		Type:         "stringSlice",
		Description: `Path to a file that should be attached to the message.
Can be used multiple times.`,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var draftFlagsALL = gsmhelpers.GetAllFlags(draftFlags)

// draftsCmd represents the drafts command
var draftsCmd = &cobra.Command{
	Use:               "drafts",
	Short:             "Manage Drafts (Part of Gmail API)",
	Long:              "Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users.drafts",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(draftsCmd)
}

func mapToDraft(flags map[string]*gsmhelpers.Value) (*gmail.Draft, error) {
	draft := &gmail.Draft{}
	var err error
	draft.Message, err = mapToMessage(flags)
	return draft, err
}
