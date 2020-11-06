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
package gsmdrive

import (
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// CreatePermission creates a permission for a file or shared drive.
func CreatePermission(fileID, emailMessage, fields string, useDomainAdminAccess, sendNotificationEmail bool, permission *drive.Permission) (*drive.Permission, error) {
	srv := getPermissionsService()
	c := srv.Create(fileID, permission).UseDomainAdminAccess(useDomainAdminAccess).SendNotificationEmail(sendNotificationEmail).EnforceSingleParent(true).SupportsAllDrives(true)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if emailMessage != "" {
		c = c.EmailMessage(emailMessage)
	}
	if permission.Role == "owner" {
		c.TransferOwnership(true).SendNotificationEmail(true)
	}
	r, err := c.Do()
	return r, err
}

// DeletePermission deletes a permission.
func DeletePermission(fileID, permissionID string, useDomainAdminAccess bool) (bool, error) {
	srv := getPermissionsService()
	err := srv.Delete(fileID, permissionID).UseDomainAdminAccess(useDomainAdminAccess).SupportsAllDrives(true).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetPermission gets a permission by ID.
func GetPermission(fileID, permissionID, fields string, useDomainAdminAccess bool) (*drive.Permission, error) {
	srv := getPermissionsService()
	c := srv.Get(fileID, permissionID).SupportsAllDrives(true).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

func makeListPermissionsCallAndAppend(c *drive.PermissionsListCall, permissions []*drive.Permission) ([]*drive.Permission, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, p := range r.Permissions {
		permissions = append(permissions, p)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		permissions, err = makeListPermissionsCallAndAppend(c, permissions)
	}
	return permissions, err
}

// ListPermissions lists a file's or shared drive's permissions.
func ListPermissions(fileID, includePermissionsForView, fields string, useDomainAdminAccess bool) ([]*drive.Permission, error) {
	srv := getPermissionsService()
	c := srv.List(fileID).SupportsAllDrives(true).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	var permissions []*drive.Permission
	permissions, err := makeListPermissionsCallAndAppend(c, permissions)
	return permissions, err
}

// UpdatePermission updates a permission with patch semantics.
func UpdatePermission(fileID, permissionID, fields string, useDomainAdminAccess, removeExpiration bool, permission *drive.Permission) (*drive.Permission, error) {
	srv := getPermissionsService()
	c := srv.Update(fileID, permissionID, permission).SupportsAllDrives(true).UseDomainAdminAccess(useDomainAdminAccess).RemoveExpiration(removeExpiration)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if permission.Role == "owner" {
		c.TransferOwnership(true)
	}
	r, err := c.Do()
	return r, err
}
