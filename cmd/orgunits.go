/*
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
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// orgUnitsCmd represents the orgUnits command
var orgUnitsCmd = &cobra.Command{
	Use:               "orgUnits",
	Short:             "Manage Organizational Unit (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/v1/reference/orgunits",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var orgUnitFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customerId": {
		AvailableFor: []string{"delete", "get", "insert", "list", "patch"},
		Type:         "string",
		Description: `The unique ID for the customer's Workspace account.
As an account administrator, you can also use the my_customer alias to represent your account's customerId.
The customerId is also returned as part of the Users resource.`,
		Defaults: map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer", "patch": "my_customer"},
	},
	"orgUnitPath": {
		AvailableFor:   []string{"delete", "get", "list", "patch"},
		Type:           "string",
		Description:    `The full path of the organizational unit or its unique ID.`,
		ExcludeFromAll: true,
	},
	"name": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The organizational unit's path name.
For example, an organizational unit's name within the /corp/support/sales_support parent path is sales_support.`,
		Required: []string{"patch"},
	},
	"blockInheritance": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description: `Determines if a sub-organizational unit can inherit the settings of the parent organization.
The default value is false, meaning a sub-organizational unit inherits the settings of the nearest parent organizational unit.
For more information on inheritance and users in an organization structure, see the administration help center.`,
	},
	"description": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Description of the organizational unit.`,
	},
	"parentOrgUnitId": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The unique ID of the parent organizational unit.
Required, unless parentOrgUnitPath is set.`,
	},
	"parentOrgUnitPath": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The organizational unit's parent path.
For example, /corp/sales is the parent path for /corp/sales/sales_support organizational unit.
Required, unless parentOrgUnitId is set.`,
	},
	"type": {
		AvailableFor: []string{"insert", "list", "patch"},
		Type:         "string",
		Description: `Whether to return all sub-organizations or just immediate children.
Acceptable values are:
all       - All sub-organizational units.
children  - Immediate children only (default).`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var orgUnitFlagsALL = gsmhelpers.GetAllFlags(orgUnitFlags)

func init() {
	rootCmd.AddCommand(orgUnitsCmd)
}

func mapToOrgUnit(flags map[string]*gsmhelpers.Value) (*admin.OrgUnit, error) {
	orgUnit := &admin.OrgUnit{}
	if flags["name"].IsSet() {
		orgUnit.Name = flags["name"].GetString()
		if orgUnit.Name == "" {
			orgUnit.ForceSendFields = append(orgUnit.ForceSendFields, "Name")
		}
	}
	if flags["blockInheritance"].IsSet() {
		orgUnit.BlockInheritance = flags["blockInheritance"].GetBool()
		if !orgUnit.BlockInheritance {
			orgUnit.ForceSendFields = append(orgUnit.ForceSendFields, "BlockInheritance")
		}
	}
	if flags["description"].IsSet() {
		orgUnit.Description = flags["description"].GetString()
		if orgUnit.Description == "" {
			orgUnit.ForceSendFields = append(orgUnit.ForceSendFields, "Description")
		}
	}
	if flags["orgUnitPath"].IsSet() {
		orgUnit.OrgUnitPath = flags["orgUnitPath"].GetString()
		if orgUnit.OrgUnitPath == "" {
			orgUnit.ForceSendFields = append(orgUnit.ForceSendFields, "OrgUnitPath")
		}
	}
	if flags["parentOrgUnitId"].IsSet() {
		orgUnit.ParentOrgUnitId = flags["parentOrgUnitId"].GetString()
		if orgUnit.ParentOrgUnitId == "" {
			orgUnit.ForceSendFields = append(orgUnit.ForceSendFields, "ParentOrgUnitId")
		}
	}
	if flags["parentOrgUnitPath"].IsSet() {
		orgUnit.ParentOrgUnitPath = flags["parentOrgUnitPath"].GetString()
		if orgUnit.ParentOrgUnitPath == "" {
			orgUnit.ForceSendFields = append(orgUnit.ForceSendFields, "ParentOrgUnitPath")
		}
	}
	return orgUnit, nil
}
