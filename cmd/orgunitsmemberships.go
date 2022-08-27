/*
Copyright © 2020-2022 Hannes Hayashi

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
	cibeta "google.golang.org/api/cloudidentity/v1beta1"

	"github.com/spf13/cobra"
)

// orgUnitsMembershipsCmd represents the orgUnitsMemberships command
var orgUnitsMembershipsCmd = &cobra.Command{
	Use:               "orgUnitsMemberships",
	Short:             "Manage the memberships of Shared Drives in organizational units (OUs) (Part of Cloud Identity Beta API)",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1beta1/orgUnits.memberships",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var orgUnitsMembershipFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `OrgUnit which is queried for a list of memberships.
Format: orgUnits/{$orgUnitId}
where $orgUnitId is the orgUnitId from the Admin SDK OrgUnit resource.
If you don't specify the "orgUnits/" prefix, GSM will automatically prepend it to the request.`,
		Required: []string{"list"},
	},
	"name": {
		AvailableFor: []string{"move"},
		Type:         "string",
		Description: `Use "driveId" instead if you just want to specifify the driveId!
The resource name of the OrgMembership.
Format: orgUnits/{$orgUnitId}/memberships/{$membership}
The $orgUnitId is the orgUnitId from the Admin SDK OrgUnit resource. To manage a Membership without specifying source orgUnitId, this API also supports the wildcard character '-' for $orgUnitId per https://google.aip.dev/159.
The $membership shall be of the form {$entityType};{$memberId}, where $entityType is the enum value of OrgMembership.EntityType, and memberId is the id from Drive API (V3) Drive resource for OrgMembership.EntityType.SHARED_DRIVE.`,
		ExcludeFromAll: true,
	},
	"driveId": {
		AvailableFor: []string{"move"},
		Type:         "string",
		Description: `The driveId of the Shared Drive to be moved.
Use this instead of the name, if you just want to specifiy the driveId.
GSM will construct the following "name" for you:
orgUnits/-/memberships/shared_drive;{$driveId}`,
		ExcludeFromAll: true,
	},
	"destinationOrgUnit": {
		AvailableFor: []string{"move"},
		Type:         "string",
		Description: `OrgUnit where the membership will be moved to.
Format: orgUnits/{$orgUnitId}
where $orgUnitId is the orgUnitId from the Admin SDK OrgUnit resource.
If you don't specify the "orgUnits/" prefix, GSM will automatically prepend it to the request.`,
		Required: []string{"move"},
	},
	"customer": {
		AvailableFor: []string{"list", "move"},
		Type:         "string",
		Description: `Customer that this OrgMembership belongs to.
All authorization will happen on the role assignments of this customer.
Format: customers/{$customerId}
where $customerId is the id from the Admin SDK Customer resource.
You may also use customers/my_customer to specify your own organization.`,
		Defaults: map[string]any{"list": "customers/my_customer", "move": "customers/my_customer"},
	},
	"filter": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The search query.
Must be specified in Common Expression Language.
May only contain equality operators on the type (e.g., type == 'shared_drive').`,
	},
	"fields": {
		AvailableFor: []string{"list", "move"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var orgUnitsMembershipFlagsALL = gsmhelpers.GetAllFlags(orgUnitsMembershipFlags)

func init() {
	rootCmd.AddCommand(orgUnitsMembershipsCmd)
}

func mapToOrgMembershipMoveRequest(flags map[string]*gsmhelpers.Value) (*cibeta.MoveOrgMembershipRequest, error) {
	moveOrgMembershipRequest := &cibeta.MoveOrgMembershipRequest{}
	moveOrgMembershipRequest.Customer = flags["customer"].GetString()
	if moveOrgMembershipRequest.Customer == "" {
		moveOrgMembershipRequest.ForceSendFields = append(moveOrgMembershipRequest.ForceSendFields, "Customer")
	}
	if flags["destinationOrgUnit"].IsSet() {
		destinationOrgUnit := flags["destinationOrgUnit"].GetString()
		if destinationOrgUnit == "" {
			moveOrgMembershipRequest.ForceSendFields = append(moveOrgMembershipRequest.ForceSendFields, "DestinationOrgUnit")
		} else {
			moveOrgMembershipRequest.DestinationOrgUnit = gsmhelpers.EnsurePrefix(destinationOrgUnit, "orgUnits/")
		}
	}
	return moveOrgMembershipRequest, nil
}
