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
	"strings"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	ci "google.golang.org/api/cloudidentity/v1"
)

// ssoProfilesCmd represents the ssoProfiles command
var ssoProfilesCmd = &cobra.Command{
	Use:               "ssoProfiles",
	Short:             "Manage inbound SAML SSO profiles (Part of Cloud Identity API)",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/inboundSamlSsoProfiles",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var ssoProfileFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"delete", "get", "patch"},
		Type:         "string",
		Description: `The resource name of the InboundSamlSsoProfile to delete.
Format: inboundSamlSsoProfiles/{sso_profile_id}.
If you don't specify the "inboundSamlSsoProfiles/" prefix, GSM will automatically prepend it for you.`,
		Required: []string{"delete", "get", "patch"},
	},
	"filter": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A Common Expression Language expression to filter the results.
The only supported filter is filtering by customer. For example: customer=="customers/C0123abc".
Omitting the filter or specifying a filter of customer=="customers/my_customer" will return the profiles for the customer that the caller (authenticated user) belongs to.`,
	},
	"customer": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `The customer.
For example: customers/C0123abc.`,
		Defaults: map[string]any{"create": "customers/my_customer", "patch": "customers/my_customer"},
	},
	"displayName": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  `Human-readable name of the SAML SSO profile.`,
	},
	"entityId": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  `The SAML Entity ID of the identity provider.`,
		Required:     []string{"create"},
	},
	"singleSignOnServiceUri": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `The SingleSignOnService endpoint location (sign-in page URL) of the identity provider.
This is the URL where the AuthnRequest will be sent.
Must use HTTPS.
Assumed to accept the HTTP-Redirect binding.`,
		Required: []string{"create"},
	},
	"logoutRedirectUri": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `The Logout Redirect URL (sign-out page URL) of the identity provider.
When a user clicks the sign-out link on a Google page, they will be redirected to this URL.
This is a pure redirect with no attached SAML LogoutRequest i.e. SAML single logout is not supported.
Must use HTTPS.`,
	},
	"changePasswordUri": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `The Change Password URL of the identity provider.
Users will be sent to this URL when changing their passwords at myaccount.google.com.
This takes precedence over the change password URL configured at customer-level.
Must use HTTPS.`,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var ssoProfileFlagsALL = gsmhelpers.GetAllFlags(ssoProfileFlags)

func init() {
	rootCmd.AddCommand(ssoProfilesCmd)
}

func mapToSsoProfile(flags map[string]*gsmhelpers.Value) (*ci.InboundSamlSsoProfile, string, error) {
	updateMask := []string{}
	profile := &ci.InboundSamlSsoProfile{
		Customer:  flags["customer"].GetString(),
		IdpConfig: &ci.SamlIdpConfig{},
	}
	if flags["displayName"].IsSet() {
		profile.DisplayName = flags["displayName"].GetString()
		if profile.DisplayName == "" {
			profile.ForceSendFields = append(profile.ForceSendFields, "DisplayName")
		}
		updateMask = append(updateMask, "displayName")
	}
	if flags["entityId"].IsSet() {
		profile.IdpConfig.EntityId = flags["entityId"].GetString()
		if profile.IdpConfig.EntityId == "" {
			profile.IdpConfig.ForceSendFields = append(profile.IdpConfig.ForceSendFields, "EntityId")
		}
		updateMask = append(updateMask, "idpConfig.entityId")
	}
	if flags["singleSignOnServiceUri"].IsSet() {
		profile.IdpConfig.SingleSignOnServiceUri = flags["singleSignOnServiceUri"].GetString()
		if profile.IdpConfig.SingleSignOnServiceUri == "" {
			profile.IdpConfig.ForceSendFields = append(profile.IdpConfig.ForceSendFields, "SingleSignOnServiceUri")
		}
		updateMask = append(updateMask, "idpConfig.singleSignOnServiceUri")
	}
	if flags["logoutRedirectUri"].IsSet() {
		profile.IdpConfig.LogoutRedirectUri = flags["logoutRedirectUri"].GetString()
		if profile.IdpConfig.LogoutRedirectUri == "" {
			profile.IdpConfig.ForceSendFields = append(profile.IdpConfig.ForceSendFields, "LogoutRedirectUri")
		}
		updateMask = append(updateMask, "idpConfig.logoutRedirectUri")
	}
	if flags["changePasswordUri"].IsSet() {
		profile.IdpConfig.ChangePasswordUri = flags["changePasswordUri"].GetString()
		if profile.IdpConfig.ChangePasswordUri == "" {
			profile.IdpConfig.ForceSendFields = append(profile.IdpConfig.ForceSendFields, "ChangePasswordUri")
		}
		updateMask = append(updateMask, "idpConfig.changePasswordUri")
	}
	return profile, strings.Join(updateMask, ","), nil
}
