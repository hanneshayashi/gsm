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
	"errors"

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// sendAsCmd represents the sendAs command
var sendAsCmd = &cobra.Command{
	Use:   "sendAs",
	Short: "Manage send-as settings for users (Part of Gmail API)",
	Long: `Settings associated with a send-as alias, which can be either the primary login address associated with the account or a custom "from" address.
Send-as aliases correspond to the "Send Mail As" feature in the web interface.
https://developers.google.com/gmail/api/reference/rest/v1/users.settings.sendAs`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var sendAsFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"create", "delete", "get", "list", "patch", "verify"},
		Type:         "string",
		Description:  "The user's email address. The special value me can be used to indicate the authenticated user.",
		Defaults:     map[string]interface{}{"create": "me", "delete": "me", "get": "me", "list": "me", "patch": "me", "verify": "me"},
	},
	"sendAsEmail": {
		AvailableFor: []string{"create", "delete", "get", "patch", "verify"},
		Type:         "string",
		Description:  `The email address that appears in the "From:" header for mail sent using this alias.`,
		Required:     []string{"create", "delete", "patch", "verify"},
	},
	"displayName": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `A name that appears in the "From:" header for mail sent using this alias.
For custom "from" addresses, when this is empty, Gmail will populate the "From:" header with the name that is used for the primary address associated with the account.
If the admin has disabled the ability for users to update their name format, requests to update this field for the primary login will silently fail.`,
	},
	"replyToAddress": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `An optional email address that is included in a "Reply-To:" header for mail sent using this alias.
If this is empty, Gmail will not generate a "Reply-To:" header.`,
	},
	"signature": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "An optional HTML signature that is included in messages composed with this alias in the Gmail web UI.",
	},
	"isDefault": {
		AvailableFor: []string{"create", "patch"},
		Type:         "bool",
		Description: `Whether this address is selected as the default "From:" address in situations such as composing a new message or sending a vacation auto-reply.
Every Gmail account has exactly one default send-as address, so the only legal value that clients may write to this field is true.
Changing this from false to true for an address will result in this field becoming false for the other previous default address.`,
	},
	"treatAsAlias": {
		AvailableFor: []string{"create", "patch"},
		Type:         "bool",
		Description: `Whether Gmail should treat this address as an alias for the user's primary email address.
This setting only applies to custom "from" aliases. See https://support.google.com/a/answer/1710338`,
	},
	"host": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "The hostname of the SMTP service. Required for SMTP.",
	},
	"port": {
		AvailableFor: []string{"create", "patch"},
		Type:         "int64",
		Description:  "The port of the SMTP service. Required for SMTP.",
	},
	"username": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "The username that will be used for authentication with the SMTP service.",
	},
	"password": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "The password that will be used for authentication with the SMTP service.",
	},
	"securityMode": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `The protocol that will be used to secure communication with the SMTP service. Required for SMTP.
[NONE|SSL|STARTTLS]
NONE      - Communication with the remote SMTP service is unsecured. Requires port 25.
SSL       - Communication with the remote SMTP service is secured using SSL.
STARTTLS  - Communication with the remote SMTP service is secured using STARTTLS.`,
		Defaults: map[string]interface{}{"create": "NONE", "patch": "NONE"},
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var sendAsFlagsALL = gsmhelpers.GetAllFlags(sendAsFlags)

func init() {
	rootCmd.AddCommand(sendAsCmd)
}

func mapToSendAs(flags map[string]*gsmhelpers.Value) (*gmail.SendAs, error) {
	sendAs := &gmail.SendAs{}
	if flags["sendAsEmail"].IsSet() {
		sendAs.SendAsEmail = flags["sendAsEmail"].GetString()
		if sendAs.SendAsEmail == "" {
			sendAs.ForceSendFields = append(sendAs.ForceSendFields, "SendAsEmail")
		}
	}
	if flags["displayName"].IsSet() {
		sendAs.DisplayName = flags["displayName"].GetString()
		if sendAs.DisplayName == "" {
			sendAs.ForceSendFields = append(sendAs.ForceSendFields, "DisplayName")
		}
	}
	if flags["replyToAddress"].IsSet() {
		sendAs.ReplyToAddress = flags["replyToAddress"].GetString()
		if sendAs.ReplyToAddress == "" {
			sendAs.ForceSendFields = append(sendAs.ForceSendFields, "ReplyToAddress")
		}
	}
	if flags["signature"].IsSet() {
		sendAs.Signature = flags["signature"].GetString()
		if sendAs.Signature == "" {
			sendAs.ForceSendFields = append(sendAs.ForceSendFields, "Signature")
		}
	}
	if flags["isPrimary"].IsSet() {
		sendAs.IsPrimary = flags["isPrimary"].GetBool()
		if !sendAs.IsPrimary {
			sendAs.ForceSendFields = append(sendAs.ForceSendFields, "IsPrimary")
		}
	}
	if flags["isDefault"].IsSet() {
		sendAs.IsDefault = flags["isDefault"].GetBool()
		if !sendAs.IsDefault {
			sendAs.ForceSendFields = append(sendAs.ForceSendFields, "IsDefault")
		}
	}
	if flags["treatAsAlias"].IsSet() {
		sendAs.TreatAsAlias = flags["treatAsAlias"].GetBool()
		if !sendAs.TreatAsAlias {
			sendAs.ForceSendFields = append(sendAs.ForceSendFields, "TreatAsAlias")
		}
	}
	if flags["host"].IsSet() || flags["port"].IsSet() || flags["username"].IsSet() || flags["password"].IsSet() || flags["securityMode"].IsSet() {
		securityMode := flags["securityMode"].GetString()
		if !gsmgmail.SecurityModeIsValid(securityMode) {
			return nil, errors.New("%s is not a valid value for security mode")
		}
		sendAs.SmtpMsa = &gmail.SmtpMsa{}
		if flags["host"].IsSet() {
			sendAs.SmtpMsa.Host = flags["host"].GetString()
			if sendAs.SmtpMsa.Host == "" {
				sendAs.SmtpMsa.ForceSendFields = append(sendAs.SmtpMsa.ForceSendFields, "Host")
			}
		}
		if flags["password"].IsSet() {
			sendAs.SmtpMsa.Password = flags["password"].GetString()
			if sendAs.SmtpMsa.Password == "" {
				sendAs.SmtpMsa.ForceSendFields = append(sendAs.SmtpMsa.ForceSendFields, "Password")
			}
		}
		if flags["port"].IsSet() {
			sendAs.SmtpMsa.Port = flags["port"].GetInt64()
			if sendAs.SmtpMsa.Port == 0 {
				sendAs.SmtpMsa.ForceSendFields = append(sendAs.SmtpMsa.ForceSendFields, "Port")
			}
		}
		if flags["securityMode"].IsSet() {
			sendAs.SmtpMsa.SecurityMode = flags["securityMode"].GetString()
			if sendAs.SmtpMsa.SecurityMode == "" {
				sendAs.SmtpMsa.ForceSendFields = append(sendAs.SmtpMsa.ForceSendFields, "SecurityMode")
			}
		}
		if flags["username"].IsSet() {
			sendAs.SmtpMsa.Username = flags["username"].GetString()
			if sendAs.SmtpMsa.Username == "" {
				sendAs.SmtpMsa.ForceSendFields = append(sendAs.SmtpMsa.ForceSendFields, "Username")
			}
		}
	}
	return sendAs, nil
}
