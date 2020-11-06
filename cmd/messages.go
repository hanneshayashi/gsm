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
	"encoding/base64"
	"fmt"
	"gsm/gsmhelpers"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// messagesCmd represents the messages command
var messagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "Manage users' messages (Part of Gmail API)",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.messages",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var messageFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{ //TODO
	"userId": {
		AvailableFor: []string{"delete", "modify", "get", "import", "insert", "list", "send", "trash", "untrash"},
		Type:         "string",
		Description:  `The user's email address. The special value \"me\" can be used to indicate the authenticated user.`,
		Defaults:     map[string]interface{}{"delete": "me", "modify": "me", "get": "me", "import": "me", "insert": "me", "list": "me", "send": "me", "trash": "me", "untrash": "me"},
	},
	"ids": {
		AvailableFor: []string{"batchDelete"},
		Type:         "stringSlice",
		Description:  `The IDs of the messages. There is a limit of 1000 ids per request.`,
		Required:     []string{"batchDelete"},
	},
	"addLabelIds": {
		AvailableFor: []string{"modify"},
		Type:         "stringSlice",
		Description:  `A list of label IDs to add to messages.`,
	},
	"removeLabelIds": {
		AvailableFor: []string{"modify"},
		Type:         "stringSlice",
		Description:  `A list of label IDs to remove from messages.`,
	},
	"id": {
		AvailableFor: []string{"delete", "get", "modify", "trash", "untrash"},
		Type:         "string",
		Description:  `The ID of the message.`,
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
		Defaults: map[string]interface{}{"get": "MINIMAL"},
	},
	"metadataHeaders": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  `When given and format is METADATA, only include headers specified.`,
	},
	"eml": {
		AvailableFor: []string{"insert", "import"},
		Type:         "string",
		Description:  `Path to the local .eml file`,
		Required:     []string{"insert", "import"},
	},
	"internalDateSource": {
		AvailableFor: []string{"insert", "import"},
		Type:         "string",
		Description:  `Source for Gmail's internal date of the message. [DATE_HEADER|RECEIVED_TIME]`,
		Defaults:     map[string]interface{}{"insert": "DATE_HEADER", "import": "DATE_HEADER"},
	},
	"deleted": {
		AvailableFor: []string{"insert", "import"},
		Type:         "bool",
		Description: `Mark the email as permanently deleted (not TRASH) and only visible in Google Vault to a Vault administrator.
Only used for G Suite accounts.`,
	},
	"neverMarkSpam": {
		AvailableFor: []string{"import"},
		Type:         "bool",
		Description:  `Ignore the Gmail spam classifier decision and never mark this email as SPAM in the mailbox.`,
	},
	"processForCalendar": {
		AvailableFor: []string{"import"},
		Type:         "bool",
		Description:  `Process calendar invites in the email and add any extracted meetings to the Google Calendar for this user.`,
	},
	"q": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Only return messages matching the specified query.
Supports the same query format as the Gmail search box.
For example, "from:someuser@example.com rfc822msgid:<somemsgid@example.com> is:unread".
Parameter cannot be used when accessing the api using the gmail.metadata scope.`,
	},
	"labelIds": {
		AvailableFor: []string{"list"},
		Type:         "stringSlice",
		Description:  `Only return messages with labels that match all of the specified label IDs.`,
	},
	"includeSpamTrash": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Include messages from SPAM and TRASH in the results.`,
	},
	"subject": {
		AvailableFor: []string{"send"},
		Type:         "string",
		Description:  "Subject of the (draft) message",
	},
	"to": {
		AvailableFor: []string{"send"},
		Type:         "string",
		Description:  "Recipient of the (draft) message",
	},
	"cc": {
		AvailableFor: []string{"send"},
		Type:         "string",
		Description:  "Copy (Cc)",
	},
	"bcc": {
		AvailableFor: []string{"send"},
		Type:         "string",
		Description:  "Blind Copy (Bcc)",
	},
	"body": {
		AvailableFor: []string{"send"},
		Type:         "string",
		Description:  "Body or content of the (draft) message",
	},
	"fields": {
		AvailableFor: []string{"get", "import", "insert", "list", "modify", "send", "trash", "untrash"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(messagesCmd)
}

func mapToMessage(flags map[string]*gsmhelpers.Value) (*gmail.Message, error) {
	message := &gmail.Message{}
	header := make(map[string]string)
	if flags["to"].IsSet() {
		header["To"] = flags["to"].GetString()
	}
	if flags["cc"].IsSet() {
		header["Cc"] = flags["cc"].GetString()
	}
	if flags["bcc"].IsSet() {
		header["Bcc"] = flags["bcc"].GetString()
	}
	header["Subject"] = flags["subject"].GetString()
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = `text/html; charset="utf-8"`
	header["Content-Transfer-Encoding"] = "base64"
	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\n", k, v)
	}
	body := flags["body"].GetString()
	msg += "\n" + body
	message.Raw = base64.URLEncoding.EncodeToString([]byte(msg))
	return message, nil
}

func emlToMessage(eml string) (*gmail.Message, error) {
	file, err := os.Open(eml)
	if err != nil {
		return nil, fmt.Errorf("Error opening %s: %v", eml, err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %v", eml, err)
	}
	message := &gmail.Message{
		Raw: base64.URLEncoding.EncodeToString(bytes),
	}
	return message, nil
}
