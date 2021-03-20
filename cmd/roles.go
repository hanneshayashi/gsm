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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// rolesCmd represents the roles command
var rolesCmd = &cobra.Command{
	Use:               "roles",
	Short:             "Manage roles (Part of Admin SDK)",
	Long:              "http://developers.google.com/admin-sdk/directory/v1/reference/roles",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var roleFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customer": {
		AvailableFor: []string{"delete", "get", "insert", "list", "patch"},
		Type:         "string",
		Description:  `Immutable ID of the Workspace account.`,
		Defaults:     map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer", "patch": "my_customer"},
	},
	"roleId": {
		AvailableFor:   []string{"delete", "get", "patch"},
		Type:           "string",
		Description:    `Immutable ID of the role.`,
		ExcludeFromAll: true,
	},
	"roleName": {
		AvailableFor:   []string{"insert", "patch"},
		Type:           "string",
		Description:    `Name of the role.`,
		ExcludeFromAll: true,
	},
	"rolePrivileges": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description:  `The set of privileges that are granted to this role.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var roleFlagsALL = gsmhelpers.GetAllFlags(roleFlags)

func init() {
	rootCmd.AddCommand(rolesCmd)
}

func mapToRole(flags map[string]*gsmhelpers.Value) (*admin.Role, error) {
	role := &admin.Role{}
	if flags["roleName"].IsSet() {
		role.RoleName = flags["roleName"].GetString()
		if role.RoleName == "" {
			role.ForceSendFields = append(role.ForceSendFields, "RoleName")
		}
	}
	if flags["rolePrivileges"].IsSet() {
		role.RolePrivileges = []*admin.RoleRolePrivileges{}
		rolePrivileges := flags["rolePrivileges"].GetStringSlice()
		if len(rolePrivileges) > 0 {
			for i := range rolePrivileges {
				role.RolePrivileges = append(role.RolePrivileges, &admin.RoleRolePrivileges{PrivilegeName: rolePrivileges[i]})
			}
		} else {
			role.ForceSendFields = append(role.ForceSendFields, "RolePrivileges")
		}
	}
	return role, nil
}
