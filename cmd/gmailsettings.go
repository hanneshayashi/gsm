/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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
	"fmt"
	"log"

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// gmailSettingsCmd represents the gmailSettings command
var gmailSettingsCmd = &cobra.Command{
	Use:               "gmailSettings",
	Short:             "Manage Gmail settings for users (Part of Gmail API)",
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.settings",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var gmailSettingFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"getAutoForwarding", "getImap", "getLanguage", "getPop", "getVacation", "updateAutoForwarding", "updateImap", "updateLanguage", "updatePop", "updateVacation"},
		Type:         "string",
		Description:  "The user's email address. The special value \"me\" can be used to indicate the authenticated user.",
		Defaults:     map[string]interface{}{"getAutoForwarding": "me", "getImap": "me", "getLanguage": "me", "getPop": "me", "getVacation": "me", "updateAutoForwarding": "me", "updateImap": "me", "updateLanguage": "me", "updatePop": "me", "updateVacation": "me"},
	},
	"enabled": {
		AvailableFor: []string{"updateAutoForwarding", "updateImap"},
		Type:         "bool",
		Description:  `Whether the setting is enabled`,
	},
	"emailAddress": {
		AvailableFor: []string{"updateAutoForwarding"},
		Type:         "string",
		Description: `Email address to which all incoming messages are forwarded.
This email address must be a verified member of the forwarding addresses.`,
	},
	"disposition": {
		AvailableFor: []string{"updateAutoForwarding", "updatePop"},
		Type:         "string",
		Description: `The state that a message should be left in after it has been forwarded.
[LEAVE_IN_INBOX|ARCHIVE|TRASH|MARK_READ]
LEAVE_IN_INBOX  - Leave the message in the INBOX.
ARCHIVE         - Archive the message.
TRASH           - Move the message to the TRASH.
MARK_READ       - Leave the message in the INBOX and mark it as read.`,
	},
	"autoExpunge": {
		AvailableFor: []string{"updateAutoForwarding", "updateImap"},
		Type:         "boolean",
		Description: `If this value is true, Gmail will immediately expunge a message when it is marked as deleted in IMAP.
Otherwise, Gmail will wait for an update from the client before expunging messages marked as deleted.`,
	},
	"expungeBehavior": {
		AvailableFor: []string{"updateImap"},
		Type:         "string",
		Description: `The action that will be executed on a message when it is marked as deleted and expunged from the last visible IMAP folder.
[ARCHIVE|TRASH|DELETE_FOREVER]
ARCHIVE - Archive messages marked as deleted.
TRASH - Move messages marked as deleted to the trash.
DELETE_FOREVER - Immediately and permanently delete messages marked as deleted. The expunged messages cannot be recovered.`,
	},
	"maxFolderSize": {
		AvailableFor: []string{"updateImap"},
		Type:         "int64",
		Description: `An optional limit on the number of messages that an IMAP folder may contain.
Legal values are 0, 1000, 2000, 5000 or 10000. A value of zero is interpreted to mean that there is no limit.`,
	},
	"displayLanguage": {
		AvailableFor: []string{"updateLanguage"},
		Type:         "string",
		Description: `The language to display Gmail in, formatted as an RFC 3066 Language Tag (for example en-GB, fr or ja for British English, French, or Japanese respectively).

The set of languages supported by Gmail evolves over time, so please refer to the "Language" dropdown in the Gmail settings for all available options, as described in the language settings help article. A table of sample values is also provided in the Managing Language Settings guide

Not all Gmail clients can display the same set of languages. In the case that a user's display language is not available for use on a particular client, said client automatically chooses to display in the closest supported variant (or a reasonable default).`,
	},
	"accessWindow": {
		AvailableFor: []string{"updatePop"},
		Type:         "string",
		Description: `The range of messages which are accessible via POP.
[DISABLED|FROM_NOW_ON|ALL_MAIL]
DISABLED     - Indicates that no messages are accessible via POP.
FROM_NOW_ON  - Indicates that unfetched messages received after some past point in time are accessible via POP.
ALL_MAIL     - Indicates that all unfetched messages are accessible via POP.`,
	},
	"enableAutoReply": {
		AvailableFor: []string{"updateVacation"},
		Type:         "bool",
		Description:  `Flag that controls whether Gmail automatically replies to messages.`,
	},
	"responseSubject": {
		AvailableFor: []string{"updateVacation"},
		Type:         "string",
		Description: `Optional text to prepend to the subject line in vacation responses.
In order to enable auto-replies, either the response subject or the response body must be nonempty.`,
	},
	"responseBodyPlainText": {
		AvailableFor: []string{"updateVacation"},
		Type:         "string",
		Description: `Response body in plain text format.
If both responseBodyPlainText and responseBodyHtml are specified, responseBodyHtml will be used.`,
	},
	"responseBodyHtml": {
		AvailableFor: []string{"updateVacation"},
		Type:         "string",
		Description: `Response body in HTML format. Gmail will sanitize the HTML before storing it.
If both responseBodyPlainText and responseBodyHtml are specified, responseBodyHtml will be used.`,
	},
	"restrictToContacts": {
		AvailableFor: []string{"updateVacation"},
		Type:         "bool",
		Description:  `Flag that determines whether responses are sent to recipients who are not in the user's list of contacts.`,
	},
	"restrictToDomain": {
		AvailableFor: []string{"updateVacation"},
		Type:         "bool",
		Description: `Flag that determines whether responses are sent to recipients who are outside of the user's domain.
This feature is only available for Workspace users.`,
	},
	"startTime": {
		AvailableFor: []string{"updateVacation"},
		Type:         "int64",
		Description: `An optional start time for sending auto-replies (epoch ms).
When this is specified, Gmail will automatically reply only to messages that it receives after the start time.
If both startTime and endTime are specified, startTime must precede endTime.`,
	},
	"endTime": {
		AvailableFor: []string{"updateVacation"},
		Type:         "int64",
		Description: `An optional end time for sending auto-replies (epoch ms).
When this is specified, Gmail will automatically reply only to messages that it receives before the end time.
If both startTime and endTime are specified, startTime must precede endTime.`,
	},
	"fields": {
		AvailableFor: []string{"getAutoForwarding", "getImap", "getLanguage", "getPop", "getVacation", "updateAutoForwarding", "updateImap", "updateLanguage", "updatePop", "updateVacation"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(gmailSettingsCmd)
}

func mapToAutoforwarding(flags map[string]*gsmhelpers.Value) (*gmail.AutoForwarding, error) {
	autoForwarding := &gmail.AutoForwarding{}
	if flags["enabled"].IsSet() {
		autoForwarding.Enabled = flags["enabled"].GetBool()
		if !autoForwarding.Enabled {
			autoForwarding.ForceSendFields = append(autoForwarding.ForceSendFields, "Enabled")
		}
	}
	if flags["emailAddress"].IsSet() {
		autoForwarding.EmailAddress = flags["emailAddress"].GetString()
		if autoForwarding.EmailAddress == "" {
			autoForwarding.ForceSendFields = append(autoForwarding.ForceSendFields, "EmailAddress")
		}
	}
	if flags["disposition"].IsSet() {
		autoForwarding.Disposition = flags["disposition"].GetString()
		if !gsmgmail.DispositionIsValid(autoForwarding.Disposition) {
			return nil, fmt.Errorf("%s is not a valid value for disposition", autoForwarding.Disposition)
		}
		if autoForwarding.Disposition == "" {
			autoForwarding.ForceSendFields = append(autoForwarding.ForceSendFields, "Disposition")
		}
	}
	return autoForwarding, nil
}

func mapToImapSettings(flags map[string]*gsmhelpers.Value) (*gmail.ImapSettings, error) {
	imapSettings := &gmail.ImapSettings{}
	if flags["enabled"].IsSet() {
		imapSettings.Enabled = flags["enabled"].GetBool()
		if !imapSettings.Enabled {
			imapSettings.ForceSendFields = append(imapSettings.ForceSendFields, "Enabled")
		}
	}
	if flags["expungeBehavior"].IsSet() {
		imapSettings.ExpungeBehavior = flags["expungeBehavior"].GetString()
		if !gsmgmail.ExpungeBehaviourIsValid(imapSettings.ExpungeBehavior) {
			return nil, fmt.Errorf("%s is not a valid value for expunge behavior", imapSettings.ExpungeBehavior)
		}
		if imapSettings.ExpungeBehavior == "" {
			imapSettings.ForceSendFields = append(imapSettings.ForceSendFields, "ExpungeBehavior")
		}
	}
	if flags["autoExpunge"].IsSet() {
		imapSettings.AutoExpunge = flags["autoExpunge"].GetBool()
		if !imapSettings.AutoExpunge {
			imapSettings.ForceSendFields = append(imapSettings.ForceSendFields, "AutoExpunge")
		}
	}
	if flags["maxFolderSize"].IsSet() {
		imapSettings.MaxFolderSize = flags["maxFolderSize"].GetInt64()
		if imapSettings.MaxFolderSize == 0 {
			imapSettings.ForceSendFields = append(imapSettings.ForceSendFields, "MaxFolderSize")
		}
	}
	return imapSettings, nil
}

func mapToLanguageSettings(flags map[string]*gsmhelpers.Value) (*gmail.LanguageSettings, error) {
	languageSettings := &gmail.LanguageSettings{}
	if flags["displayLanguage"].IsSet() {
		languageSettings.DisplayLanguage = flags["displayLanguage"].GetString()
		if languageSettings.DisplayLanguage == "" {
			languageSettings.ForceSendFields = append(languageSettings.ForceSendFields, "DisplayLanguage")
		}
	}
	return languageSettings, nil
}

func mapToPopSettings(flags map[string]*gsmhelpers.Value) (*gmail.PopSettings, error) {
	popSettings := &gmail.PopSettings{}
	if flags["accessWindow"].IsSet() {
		popSettings.AccessWindow = flags["accessWindow"].GetString()
		if !gsmgmail.AccessWindowIsValid(popSettings.AccessWindow) {
			return nil, fmt.Errorf("%s is not a valid value for access window", popSettings.AccessWindow)
		}
		if popSettings.AccessWindow == "" {
			popSettings.ForceSendFields = append(popSettings.ForceSendFields, "AccessWindow")
		}
	}
	if flags["disposition"].IsSet() {
		popSettings.Disposition = flags["disposition"].GetString()
		if !gsmgmail.DispositionIsValid(popSettings.Disposition) {
			return nil, fmt.Errorf("%s is not a valid value for disposition", popSettings.Disposition)
		}
		if popSettings.Disposition == "" {
			popSettings.ForceSendFields = append(popSettings.ForceSendFields, "Disposition")
		}
	}
	return popSettings, nil
}

func mapToVacationSettings(flags map[string]*gsmhelpers.Value) (*gmail.VacationSettings, error) {
	vacationSettings := &gmail.VacationSettings{}
	if flags["enableAutoReply"].IsSet() {
		vacationSettings.EnableAutoReply = flags["enableAutoReply"].GetBool()
		if !vacationSettings.EnableAutoReply {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "EnableAutoReply")
		}
	}
	if flags["responseSubject"].IsSet() {
		vacationSettings.ResponseSubject = flags["responseSubject"].GetString()
		if vacationSettings.ResponseSubject == "" {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "ResponseSubject")
		}
	}
	if flags["responseBodyPlainText"].IsSet() {
		vacationSettings.ResponseBodyPlainText = flags["responseBodyPlainText"].GetString()
		if vacationSettings.ResponseBodyPlainText == "" {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "ResponseBodyPlainText")
		}
	}
	if flags["responseBodyHtml"].IsSet() {
		vacationSettings.ResponseBodyHtml = flags["responseBodyHtml"].GetString()
		if vacationSettings.ResponseBodyHtml == "" {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "ResponseBodyHtml")
		}
	}
	if flags["restrictToContacts"].IsSet() {
		vacationSettings.RestrictToContacts = flags["restrictToContacts"].GetBool()
		if !vacationSettings.RestrictToContacts {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "RestrictToContacts")
		}
	}
	if flags["restrictToDomain"].IsSet() {
		vacationSettings.RestrictToDomain = flags["restrictToDomain"].GetBool()
		if !vacationSettings.RestrictToDomain {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "RestrictToDomain")
		}
	}
	if flags["startTime"].IsSet() {
		vacationSettings.StartTime = flags["startTime"].GetInt64()
		if vacationSettings.StartTime == 0 {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "StartTime")
		}
	}
	if flags["endTime"].IsSet() {
		vacationSettings.EndTime = flags["endTime"].GetInt64()
		if vacationSettings.EndTime == 0 {
			vacationSettings.ForceSendFields = append(vacationSettings.ForceSendFields, "EndTime")
		}
	}
	return vacationSettings, nil
}
