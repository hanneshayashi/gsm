/*
Copyright © 2020-2023 Hannes Hayashi

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

// ssoAssignmentsCmd represents the ssoAssignments command
var ssoAssignmentsCmd = &cobra.Command{
	Use:               "ssoAssignments",
	Short:             "Manage inbound SAML SSO assignments (Part of Cloud Identity API)",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/inboundSsoAssignments",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var ssoAssignmentFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customer": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `The customer.
For example: customers/C0123abc.`,
		Defaults: map[string]any{"create": "customers/my_customer"},
	},
	"rank": {
		AvailableFor: []string{"create", "patch"},
		Type:         "int64",
		Description:  `Must be zero (which is the default value so it can be omitted) for assignments with targetOrgUnit set and must be greater-than-or-equal-to one for assignments with targetGroup set.`,
	},
	"ssoMode": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `Inbound SSO behaviors.
May be one of the following:
- SSO_OFF                      - Disable SSO for the targeted users.
- SAML_SSO                     - Use an external SAML Identity Provider for SSO for the targeted users.
- DOMAIN_WIDE_SAML_IF_ENABLED  - Use the domain-wide SAML Identity Provider for the targeted users if one is configured; otherwise, this is equivalent to SSO_OFF.
                                 Note that this will also be equivalent to SSO_OFF if/when support for domain-wide SAML is removed.
                                 Google may disallow this mode at that point and existing assignments with this mode may be automatically changed to SSO_OFF.`,
	},
	"inboundSamlSsoProfile": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `Name of the InboundSamlSsoProfile to use.
Must be of the form inboundSamlSsoProfiles/{inboundSamlSsoProfile}.`,
		Required: []string{"create"},
	},
	"targetGroup": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `Must be of the form groups/{group}
Only ONE of --targetGroup and --targetOrgUnit may be specified.`,
	},
	"targetOrgUnit": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `Must be of the form orgUnits/{orgUnit}.
Only ONE of --targetGroup and --targetOrgUnit may be specified.`,
	},
	"name": {
		AvailableFor: []string{"delete", "get", "patch"},
		Type:         "string",
		Description: `The resource name of the InboundSsoAssignment.
Format: inboundSsoAssignments/{assignment}`,
		Required: []string{"delete", "get", "patch"},
	},
	"filter": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A CEL expression to filter the results.
The only supported filter is filtering by customer. For example: customer==customers/C0123abc.
Omitting the filter or specifying a filter of customer==customers/my_customer will return the assignments for the customer that the caller (authenticated user) belongs to.`,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var ssoAssignmentFlagsALL = gsmhelpers.GetAllFlags(ssoAssignmentFlags)

func init() {
	rootCmd.AddCommand(ssoAssignmentsCmd)
}

func mapToSsoAssignment(flags map[string]*gsmhelpers.Value) (*ci.InboundSsoAssignment, string, error) {
	updateMask := []string{}
	assignment := &ci.InboundSsoAssignment{
		Customer: flags["customer"].GetString(),
	}
	if flags["targetGroup"].IsSet() {
		assignment.TargetGroup = flags["targetGroup"].GetString()
		if assignment.TargetGroup == "" {
			assignment.ForceSendFields = append(assignment.ForceSendFields, "TargetGroup")
		}
		updateMask = append(updateMask, "targetGroup")
	}
	if flags["targetOrgUnit"].IsSet() {
		assignment.TargetOrgUnit = flags["targetOrgUnit"].GetString()
		if assignment.TargetOrgUnit == "" {
			assignment.ForceSendFields = append(assignment.ForceSendFields, "TargetOrgUnit")
		}
		updateMask = append(updateMask, "targetOrgUnit")
	}
	if flags["rank"].IsSet() {
		assignment.Rank = flags["rank"].GetInt64()
		if assignment.Rank == 0 {
			assignment.ForceSendFields = append(assignment.ForceSendFields, "Rank")
		}
		updateMask = append(updateMask, "rank")
	}
	if flags["ssoMode"].IsSet() {
		assignment.SsoMode = flags["ssoMode"].GetString()
		if assignment.SsoMode == "" {
			assignment.ForceSendFields = append(assignment.ForceSendFields, "SsoMode")
		}
		updateMask = append(updateMask, "ssoMode")
	}
	if flags["inboundSamlSsoProfile"].IsSet() {
		assignment.SamlSsoInfo = &ci.SamlSsoInfo{
			InboundSamlSsoProfile: flags["inboundSamlSsoProfile"].GetString(),
		}
		if assignment.SamlSsoInfo.InboundSamlSsoProfile == "" {
			assignment.SamlSsoInfo.ForceSendFields = append(assignment.SamlSsoInfo.ForceSendFields, "InboundSamlSsoProfile")
		}
		updateMask = append(updateMask, "samlSsoInfo.inboundSamlSsoProfile")
	}
	if flags["redirectCondition"].IsSet() {
		assignment.SignInBehavior = &ci.SignInBehavior{
			RedirectCondition: flags["redirectCondition"].GetString(),
		}
		if assignment.SignInBehavior.RedirectCondition == "" {
			assignment.SignInBehavior.ForceSendFields = append(assignment.SignInBehavior.ForceSendFields, "RedirectCondition")
		}
		updateMask = append(updateMask, "signInBehavior.redirectCondition")
	}
	return assignment, strings.Join(updateMask, ","), nil
}
