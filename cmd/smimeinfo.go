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
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// smimeInfoCmd represents the smimeInfo command
var smimeInfoCmd = &cobra.Command{
	Use:               "smimeInfo",
	Short:             "Manage users' S/MIME configs for send-as aliases",
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.settings.sendAs.smimeInfo",
	DisableAutoGenTag: true,
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
	"encryptedKeyPassword": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description:  `Encrypted key password, when key is encrypted.`,
	},
	"pkcs12": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description: `Path to a PKCS#12 format file containing a single private/public key pair and certificate chain.
This format is only accepted from client for creating a new SmimeInfo and is never returned, because the private key is not intended to be exported.
PKCS#12 may be encrypted, in which case encryptedKeyPassword should be set appropriately.`,
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
	if flags["encryptedKeyPassword"].IsSet() {
		smimeInfo.EncryptedKeyPassword = flags["encryptedKeyPassword"].GetString()
		if smimeInfo.EncryptedKeyPassword == "" {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "EncryptedKeyPassword")
		}
	}
	if flags["pkcs12"].IsSet() {
		f, err := os.Open(flags["pkcs12"].GetString())
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		smimeInfo.Pkcs12 = base64.URLEncoding.EncodeToString(b)
		if smimeInfo.Pkcs12 == "" {
			smimeInfo.ForceSendFields = append(smimeInfo.ForceSendFields, "Pkcs12")
		}
	}
	return smimeInfo, nil
}
