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
	"os"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	ci "google.golang.org/api/cloudidentity/v1"
)

// ssoProfileCredentialsCmd represents the inboundSamlSsoProfileCredentials command
var ssoProfileCredentialsCmd = &cobra.Command{
	Use:               "ssoProfileCredentials",
	Short:             "Manage inbound SAML SSO profile IdP Credentials (Part of Cloud Identity API)",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/inboundSamlSsoProfiles.idpCredentials",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var ssoProfileCredentialFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"add", "list"},
		Type:         "string",
		Description: `The InboundSamlSsoProfile that owns the IdpCredential(s). Format: inboundSamlSsoProfiles/{sso_profile_id}
If you don't specify the "inboundSamlSsoProfiles/" prefix, GSM will automatically prepend it for you.`,
		Required: []string{"add", "list"},
	},
	"pemFile": {
		AvailableFor: []string{"add"},
		Type:         "string",
		Description:  `The file path to a PEM encoded x509 certificate containing the public key for verifying IdP signatures.`,
		Required:     []string{"add"},
	},
	"name": {
		AvailableFor: []string{"delete", "get"},
		Type:         "string",
		Description: `The resource name of the IdpCredential.
Format: inboundSamlSsoProfiles/{sso_profile_id}/idpCredentials/{idp_credential_id}`,
		Required: []string{"delete", "get"},
	},
	"fields": {
		AvailableFor: []string{"add", "get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var ssoProfileCredentialFlagsALL = gsmhelpers.GetAllFlags(ssoProfileCredentialFlags)

func init() {
	rootCmd.AddCommand(ssoProfileCredentialsCmd)
}

func mapToAddIdpCredentialRequest(flags map[string]*gsmhelpers.Value) (*ci.AddIdpCredentialRequest, error) {
	request := &ci.AddIdpCredentialRequest{}
	if flags["pemFile"].IsSet() {
		content, err := os.ReadFile(flags["pemFile"].GetString())
		if err != nil {
			return nil, err
		}
		request.PemData = string(content)
	}
	return request, nil
}
