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

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// permissionsCmd represents the permissions command
var permissionsCmd = &cobra.Command{
	Use:               "permissions",
	Short:             "Manage file and drive permissions (Part of Drive API)",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/permissions",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var permissionFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fileId": {
		AvailableFor: []string{"create", "delete", "get", "list", "update"},
		Type:         "string",
		Description:  "Id of the file or drive",
		Required:     []string{"create", "delete", "get", "list", "update"},
	},
	"permissionId": {
		AvailableFor:   []string{"delete", "get", "update"},
		Type:           "string",
		Description:    "The ID of the permission.",
		ExcludeFromAll: true,
		Recursive:      []string{"delete", "update"},
	},
	"emailAddress": {
		AvailableFor:   []string{"create", "delete", "get", "update"},
		Type:           "string",
		Description:    "The email address of the user or group to which this permission refers.",
		ExcludeFromAll: true,
		Recursive:      []string{"create", "delete", "update"},
	},
	"moveToNewOwnersRoot": {
		AvailableFor: []string{"create", "update"},
		Type:         "bool",
		Description: `This parameter only takes effect if the item is not in a shared drive and the request is attempting to transfer the ownership of the item.
When set to true, the item is moved to the new owner's My Drive root folder and all prior parents removed.
however, the file will be added to the new owner's My Drive root folder, unless it is already in the new owner's My Drive.`,
		Recursive: []string{"create", "update"},
	},
	"transferOwnership": {
		AvailableFor: []string{"create", "update"},
		Type:         "bool",
		Description: `Whether to transfer ownership to the specified user and downgrade the current owner to a writer.
This parameter is required as an acknowledgement of the side effect.`,
		Recursive: []string{"create", "update"},
	},
	"type": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The type of the grantee.
[user|group|domain|anyone].
When creating a permission, if type is user or group, you must provide an emailAddress for the user or group.
When type is domain, you must provide a domain.
There isn't extra information required for a anyone type.`,
		Required:  []string{"create"},
		Recursive: []string{"create", "update"},
	},
	"domain": {
		AvailableFor: []string{"create", "delete", "get", "update"},
		Type:         "string",
		Description:  "The domain to which this permission refers.",
		Recursive:    []string{"create", "delete", "update"},
	},
	"allowFileDiscovery": {
		AvailableFor: []string{"create", "update"},
		Type:         "bool",
		Description: `Whether the permission allows the file to be discovered through search.
This is only applicable for permissions of type domain or anyone.`,
		Recursive: []string{"create", "update"},
	},
	"role": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The role granted by this permission.
While new values may be supported in the future, the following are currently allowed:
[owner|organizer|fileOrganizer|writer|commenter|reader]`,
		Required:  []string{"create"},
		Recursive: []string{"create", "update"},
	},
	"emailMessage": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "A plain text custom message to include in the notification email",
		Recursive:    []string{"create", "update"},
	},
	"useDomainAdminAccess": {
		AvailableFor: []string{"create", "delete", "get", "list", "update"},
		Type:         "bool",
		Description:  "Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID parameter refers to a shared drive and the requester is an administrator of the domain to which the shared drive belongs.",
		Recursive:    []string{"create", "delete", "list", "update"},
	},
	"sendNotificationEmail": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description: `Whether to send a notification email when sharing to users or groups.
This defaults to true for users and groups, and is not allowed for other requests.
It must not be disabled for ownership transfers.`,
		Recursive: []string{"create"},
	},
	"view": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `Indicates the view for this permission.
Only populated for permissions that belong to a view. published is the only supported value.`,
	},
	"includePermissionsForView": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Specifies which additional view's permissions to include in the response.
Only 'published' is supported.`,
	},
	"removeExpiration": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description:  `Whether to remove the expiration date.`,
		Recursive:    []string{"update"},
	},
	"expirationTime": {
		AvailableFor: []string{"update"},
		Type:         "string",
		Description: `The time at which this permission will expire (RFC 3339 date-time). Expiration times have the following restrictions:
They can only be set on user and group permissions
The time must be in the future
The time cannot be more than a year in the future`,
		Recursive: []string{"update"},
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Recursive: []string{"create", "list", "update"},
	},
}
var permissionFlagsALL = gsmhelpers.GetAllFlags(permissionFlags)

func init() {
	rootCmd.AddCommand(permissionsCmd)
}

func mapToPermission(flags map[string]*gsmhelpers.Value) (*drive.Permission, error) {
	permission := &drive.Permission{}
	if flags["emailAddress"].IsSet() {
		permission.EmailAddress = flags["emailAddress"].GetString()
		if permission.EmailAddress == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "EmailAddress")
		}
	}
	if flags["type"].IsSet() {
		permission.Type = flags["type"].GetString()
		if permission.Type == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Type")
		}
	}
	if flags["domain"].IsSet() {
		permission.Domain = flags["domain"].GetString()
		if permission.Domain == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Domain")
		}
	}
	if flags["role"].IsSet() {
		permission.Role = flags["role"].GetString()
		if permission.Role == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "Role")
		}
	}
	if flags["allowFileDiscovery"].IsSet() {
		permission.AllowFileDiscovery = flags["allowFileDiscovery"].GetBool()
		if !permission.AllowFileDiscovery {
			permission.ForceSendFields = append(permission.ForceSendFields, "AllowFileDiscovery")
		}
	}
	if flags["expirationTime"].IsSet() {
		permission.ExpirationTime = flags["expirationTime"].GetString()
		if permission.ExpirationTime == "" {
			permission.ForceSendFields = append(permission.ForceSendFields, "ExpirationTime")
		}
	}
	return permission, nil
}
