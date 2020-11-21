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
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// roleAssignmentsCmd represents the roleAssignments command
var roleAssignmentsCmd = &cobra.Command{
	Use:   "roleAssignments",
	Short: "Manage Role Assignments (Part of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/roleAssignments",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var roleAssignmentFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customer": {
		AvailableFor: []string{"delete", "get", "insert", "list"},
		Type:         "string",
		Description:  `Immutable ID of the G Suite account.`,
		Defaults:     map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer"},
	},
	"roleAssignmentId": {
		AvailableFor: []string{"delete", "get"},
		Type:         "string",
		Description:  `Immutable ID of the role assignment.`,
		Required:     []string{"delete", "get"},
	},
	"assignedTo": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description:  `The unique ID of the user this role is assigned to.`,
		Required:     []string{"insert"},
	},
	"orgUnitId": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description:  `If the role is restricted to an organization unit, this contains the ID for the organization unit the exercise of this role is restricted to.`,
	},
	"roleId": {
		AvailableFor: []string{"insert", "list"},
		Type:         "int64",
		Description:  `The ID of the role that is assigned.`,
		Required:     []string{"insert"},
	},
	"scopeType": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description: `The scope in which this role is assigned.
Acceptable values are:
"CUSTOMER"
"ORG_UNIT"`,
	},
	"userKey": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The user's primary email address, alias email address, or unique user ID.
If included in the request, returns role assignments only for this user.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var roleAssignmentFlagsALL = gsmhelpers.GetAllFlags(roleAssignmentFlags)

func init() {
	rootCmd.AddCommand(roleAssignmentsCmd)
}

func mapToRoleAssignment(flags map[string]*gsmhelpers.Value) (*admin.RoleAssignment, error) {
	roleAssignment := &admin.RoleAssignment{}
	if flags["assignedTo"].IsSet() {
		roleAssignment.AssignedTo = flags["assignedTo"].GetString()
		if roleAssignment.AssignedTo == "" {
			roleAssignment.ForceSendFields = append(roleAssignment.ForceSendFields, "AssignedTo")
		}
	}
	if flags["orgUnitId"].IsSet() {
		roleAssignment.OrgUnitId = flags["orgUnitId"].GetString()
		if roleAssignment.OrgUnitId == "" {
			roleAssignment.ForceSendFields = append(roleAssignment.ForceSendFields, "orgUnitId")
		}
	}
	if flags["scopeType"].IsSet() {
		roleAssignment.ScopeType = flags["scopeType"].GetString()
		if roleAssignment.ScopeType == "" {
			roleAssignment.ForceSendFields = append(roleAssignment.ForceSendFields, "scopeType")
		}
	}
	if flags["roleId"].IsSet() {
		roleAssignment.RoleId = flags["roleId"].GetInt64()
		if roleAssignment.RoleId == 0 {
			roleAssignment.ForceSendFields = append(roleAssignment.ForceSendFields, "roleId")
		}
	}
	return roleAssignment, nil
}
