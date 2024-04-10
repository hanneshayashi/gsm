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
	"fmt"
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/drivelabels/v2"

	"github.com/spf13/cobra"
)

// driveLabelPermissionsCmd represents the driveLabelPermissions command
var driveLabelPermissionsCmd = &cobra.Command{
	Use:   "driveLabelPermissions",
	Short: "Manages Drive Label Permissions (Part of Drive Labels API)",
	Long: `Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels.permissions
Use ONE of the following to specify the target label:
name
email
audience
person
group`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var driveLabelPermissionFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"list", "create", "batchDelete", "update", "batchUpdate"},
		Type:         "string",
		Description: `The parent Label resource name. Format: labels/{label}
If you don't specify the "labels/" prefix, GSM will automatically prepend it to the request.`,
		Required: []string{"list", "create", "batchDelete", "update", "batchUpdate"},
	},
	"name": {
		AvailableFor: []string{"delete", "update"},
		Type:         "string",
		Description: `Resource name of this permission. Format: labels/{label}
If you don't specify the "labels/" prefix, GSM will automatically prepend it to the request.`,
		Required: []string{"delete", "update"},
	},
	"permissionName": {
		AvailableFor: []string{"batchDelete"},
		Type:         "stringSlice",
		Description: `Label Permission resource name. Format: labels/{label}
May be used multiple times to delete multiple permissions at once.
If you don't specify the "labels/" prefix, GSM will automatically prepend it to the request.`,
		Required: []string{"batchDelete"},
	},
	"permission": {
		AvailableFor: []string{"batchUpdate"},
		Type:         "stringSlice",
		Description: `A permission.
In order to update an existing permission use the following format:
"name=...;role=..."

In order to create a new permissions use one of the following depending on your use case:
"email=...;role=..."
"group=...;role=..."
"person=...;role=..."
"audience=...;role=..."

May be used multiple times to update multiple permissions at once.
If you don't specify the "labels/" prefix, GSM will automatically prepend it to the request.`,
		Required: []string{"batchUpdate"},
	},
	"email": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `Specifies the email address for a user or group pricinpal.
Not populated for audience principals.
User and Group permissions may only be inserted using email address.
On update requests, if email address is specified, no principal should be specified.`,
	},
	"role": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The role the principal should have.
"[READER|APPLIER|ORGANIZER|EDITOR].
READER     - A reader can read the label and associated metadata applied to Drive items.
APPLIER    - An applier can write associated metadata on Drive items in which they also have write access to. Implies READER.
ORGANIZER  - An organizer can pin this label in shared drives they manage and add new appliers to the label.
EDITOR     - Editors can make any update including deleting the label which also deletes the associated Drive item metadata. Implies APPLIER.`,
		Required: []string{"create", "update"},
	},
	"audience": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `Audience to grant a role to.
The magic value of audiences/default may be used to apply the role to the default audience in the context of the organization that owns the Label.`,
	},
	"useAdminAccess": {
		AvailableFor: []string{"list", "create", "delete", "update", "batchDelete", "batchUpdate"},
		Type:         "bool",
		Description: `Set to true in order to use the user's admin credentials.
The server verifies that the user is an admin for the label before allowing access.`,
	},
	"fields": {
		AvailableFor: []string{"list", "create", "update", "batchUpdate"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var labelPermissionFlagsALL = gsmhelpers.GetAllFlags(driveLabelPermissionFlags)

func init() {
	rootCmd.AddCommand(driveLabelPermissionsCmd)
}

func mapToDriveLabelCreatePermission(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2LabelPermission, error) {
	permission := &drivelabels.GoogleAppsDriveLabelsV2LabelPermission{}
	if flags["role"].IsSet() {
		permission.Role = flags["role"].GetString()
		if permission.Role == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Role")
		}
	}
	if flags["email"].IsSet() {
		permission.Email = flags["email"].GetString()
		if permission.Email == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Email")
		}
	}
	if flags["audience"].IsSet() {
		permission.Audience = flags["audience"].GetString()
		if permission.Audience == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Audience")
		}
	}
	return permission, nil
}

func mapToDriveLabelUpdatePermission(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2LabelPermission, error) {
	permission := &drivelabels.GoogleAppsDriveLabelsV2LabelPermission{}
	if flags["role"].IsSet() {
		permission.Role = flags["role"].GetString()
		if permission.Role == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Role")
		}
	}
	if flags["name"].IsSet() {
		permission.Name = flags["name"].GetString()
		if permission.Name == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Name")
		}
	}
	return permission, nil
}

func mapToBatchDeleteDriveLabelsRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2BatchDeleteLabelPermissionsRequest, error) {
	names := flags["permissionName"].GetStringSlice()
	request := &drivelabels.GoogleAppsDriveLabelsV2BatchDeleteLabelPermissionsRequest{
		Requests:       make([]*drivelabels.GoogleAppsDriveLabelsV2DeleteLabelPermissionRequest, len(names)),
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
	}
	for i := range names {
		request.Requests[i] = &drivelabels.GoogleAppsDriveLabelsV2DeleteLabelPermissionRequest{
			Name: gsmhelpers.EnsurePrefix(names[i], "labels/"),
		}
	}
	return request, nil
}

func mapToBatchUpdateDriveLabelsRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsRequest, error) {
	permissions := flags["permission"].GetStringSlice()
	request := &drivelabels.GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsRequest{
		Requests:       make([]*drivelabels.GoogleAppsDriveLabelsV2UpdateLabelPermissionRequest, len(permissions)),
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
	}
	for i := range permissions {
		m := gsmhelpers.FlagToMap(permissions[i])
		request.Requests[i] = &drivelabels.GoogleAppsDriveLabelsV2UpdateLabelPermissionRequest{
			LabelPermission: &drivelabels.GoogleAppsDriveLabelsV2LabelPermission{},
		}
		role, ok := m["role"]
		if !ok {
			return nil, fmt.Errorf("role is missing")
		}
		request.Requests[i].LabelPermission.Role = role
		if request.Requests[i].LabelPermission.Role == "" {
			request.Requests[i].LabelPermission.ForceSendFields = append(request.Requests[i].LabelPermission.ForceSendFields, "Role")
		}
		email, ok := m["email"]
		if ok {
			request.Requests[i].LabelPermission.Email = email
			if request.Requests[i].LabelPermission.Email == "" {
				request.Requests[i].LabelPermission.ForceSendFields = append(request.Requests[i].LabelPermission.ForceSendFields, "Email")
			}
		}
		audience, ok := m["audience"]
		if ok {
			request.Requests[i].LabelPermission.Audience = audience
			if request.Requests[i].LabelPermission.Audience == "" {
				request.Requests[i].LabelPermission.ForceSendFields = append(request.Requests[i].LabelPermission.ForceSendFields, "Audience")
			}
		}
	}
	return request, nil
}
