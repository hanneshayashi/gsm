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

	"github.com/spf13/cobra"
)

// threadsCmd represents the threads command
var threadsCmd = &cobra.Command{
	Use:               "threads",
	Short:             "Manage threads in users' mailboxes (Part of Gmail API)",
	Long:              "Implements the API documented at https://developers.google.com/workspace/gmail/api/reference/rest/v1/users.threads",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var threadFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"delete", "get", "list", "modify", "trash", "untrash"},
		Type:         "string",
		Description:  "The user's email address. The special value me can be used to indicate the authenticated user.",
		Defaults:     map[string]any{"delete": "me", "get": "me", "list": "me", "modify": "me", "trash": "me", "untrash": "me"},
	},
	"id": {
		AvailableFor:   []string{"delete", "get", "modify", "trash", "untrash"},
		Type:           "string",
		Description:    "ID of the Thread.",
		Required:       []string{"delete", "get", "modify", "trash", "untrash"},
		ExcludeFromAll: true,
	},
	"format": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The format to return the message in.
[MINIMAL|FULL|RAW|METADATA]
MINIMAL   - Returns only email message ID and labels; does not return the email headers, body, or payload.
FULL      - Returns the full email message data with body content parsed in the payload field; the raw field is not used. Format cannot be used when accessing the api using the gmail.metadata scope.
RAW       - Returns the full email message data with body content in the raw field as a base64url encoded string; the payload field is not used. Format cannot be used when accessing the api using the gmail.metadata scope.
METADATA  - Returns only email message ID, labels, and email headers.`,
	},
	"metadataHeaders": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  `When given and format is METADATA, only include headers specified.`,
	},
	"q": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Only return threads matching the specified query.
Supports the same query format as the Gmail search box.
For example, "from:someuser@example.com rfc822msgid:<somemsgid@example.com> is:unread".
Parameter cannot be used when accessing the api using the gmail.metadata scope.`,
	},
	"labelIds": {
		AvailableFor: []string{"list"},
		Type:         "stringSlice",
		Description:  `Only return threads with labels that match all of the specified label IDs.`,
	},
	"includeSpamTrash": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Include threads from SPAM and TRASH in the results.`,
	},
	"addLabelIds": {
		AvailableFor: []string{"modify"},
		Type:         "stringSlice",
		Description:  `A list of label IDs to add to threads.`,
	},
	"removeLabelIds": {
		AvailableFor: []string{"modify"},
		Type:         "stringSlice",
		Description:  `A list of label IDs to remove from threads.`,
	},
	"fields": {
		AvailableFor: []string{"get", "list", "modify", "trash", "untrash"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var threadFlagsALL = gsmhelpers.GetAllFlags(threadFlags)

func init() {
	rootCmd.AddCommand(threadsCmd)
}
