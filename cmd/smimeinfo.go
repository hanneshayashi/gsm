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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// smimeInfoCmd represents the smimeInfo command
var smimeInfoCmd = &cobra.Command{
	Use:   "smimeInfo",
	Short: "Manage users' S/MIME configs for send-as aliases",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.settings.sendAs.smimeInfo",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var smimeInfoFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"delete", "get", "insert", "list", "setDefault"},
		Type:         "string",
		Description:  "The user's email address. The special value me can be used to indicate the authenticated user.",
		Defaults:     map[string]interface{}{"delete": "me", "get": "me", "insert": "me", "list": "me", "setDefault": "me"},
	},
	"sendAsEmail": {
		AvailableFor: []string{"delete", "get", "insert", "list", "setDefault"},
		Type:         "string",
		Description:  `The email address that appears in the "From:" header for mail sent using this alias.`,
		Required:     []string{"delete", "get", "insert", "list", "setDefault"},
	},
	"id": {
		AvailableFor:   []string{"delete", "get", "setDefault"},
		Type:           "string",
		Description:    `The immutable ID for the SmimeInfo.`,
		Required:       []string{"delete", "get", "setDefault"},
		ExcludeFromAll: true,
	},
	"issuerCn": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description:  `The S/MIME certificate issuer's common name.`,
	},
	"isDefault": {
		AvailableFor: []string{"insert"},
		Type:         "bool",
		Description:  `The S/MIME certificate issuer's common name.`,
	},
	"expiration": {
		AvailableFor: []string{"insert"},
		Type:         "bool",
		Description:  `When the certificate expires (in milliseconds since epoch).`,
	},
	"encryptedKeyPassword": {
		AvailableFor: []string{"insert"},
		Type:         "bool",
		Description:  `Encrypted key password, when key is encrypted.`,
	},
	"pem": {
		AvailableFor: []string{"insert"},
		Type:         "bool",
		Description: `PEM formatted X509 concatenated certificate string (standard base64 encoding).
Format used for returning key, which includes public key as well as certificate chain (not private key).`,
	},
	"pkcs12": {
		AvailableFor: []string{"insert"},
		Type:         "bool",
		Description: `PKCS#12 format containing a single private/public key pair and certificate chain.
This format is only accepted from client for creating a new SmimeInfo and is never returned, because the private key is not intended to be exported.
PKCS#12 may be encrypted, in which case encryptedKeyPassword should be set appropriately.

A base64-encoded string.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var smimeInfoFlagsALL = gsmhelpers.GetAllFlags(smimeInfoFlags)

func init() {
	rootCmd.AddCommand(smimeInfoCmd)
}

func mapToSmimeInfo(flags map[string]*gsmhelpers.Value) (*gmail.SmimeInfo, error) {
	smimeInfo := &gmail.SmimeInfo{}
	if flags["issuerCn"].IsSet() {
		smimeInfo.IssuerCn = flags["issuerCn"].GetString()
		if smimeInfo.IssuerCn == "" {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "IssuerCn")
		}
	}
	if flags["encryptedKeyPassword"].IsSet() {
		smimeInfo.EncryptedKeyPassword = flags["encryptedKeyPassword"].GetString()
		if smimeInfo.EncryptedKeyPassword == "" {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "EncryptedKeyPassword")
		}
	}
	if flags["isDefault"].IsSet() {
		smimeInfo.IsDefault = flags["isDefault"].GetBool()
		if !smimeInfo.IsDefault {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "IsDefault")
		}
	}
	if flags["expiration"].IsSet() {
		smimeInfo.Expiration = flags["expiration"].GetInt64()
		if smimeInfo.Expiration == 0 {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "Expiration")
		}
	}
	if flags["pem"].IsSet() {
		smimeInfo.Pem = flags["pem"].GetString()
		if smimeInfo.Pem == "" {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "Pem")
		}
	}
	if flags["pkcs12"].IsSet() {
		smimeInfo.Pkcs12 = flags["pkcs12"].GetString()
		if smimeInfo.Pkcs12 == "" {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "Pkcs12")
		}
	}
	return smimeInfo, nil
}
