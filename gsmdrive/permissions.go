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

package gsmdrive

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// CreatePermission creates a permission for a file or shared drive.
func CreatePermission(fileID, emailMessage, fields string, useDomainAdminAccess, sendNotificationEmail, transferOwnership, moveToNewOwnersRoot bool, permission *drive.Permission) (*drive.Permission, error) {
	srv := getPermissionsService()
	c := srv.Create(fileID, permission).UseDomainAdminAccess(useDomainAdminAccess).SendNotificationEmail(sendNotificationEmail).SupportsAllDrives(true).TransferOwnership(transferOwnership).MoveToNewOwnersRoot(moveToNewOwnersRoot)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if emailMessage != "" {
		c = c.EmailMessage(emailMessage)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Permission)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// DeletePermission deletes a permission.
func DeletePermission(fileID, permissionID string, useDomainAdminAccess bool, enforceExpansiveAccess bool) (bool, error) {
	srv := getPermissionsService()
	c := srv.Delete(fileID, permissionID).UseDomainAdminAccess(useDomainAdminAccess).EnforceExpansiveAccess(enforceExpansiveAccess).SupportsAllDrives(true)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(fileID, permissionID), func() error {
		return c.Do()
	})
	return result, err
}

// GetPermission gets a permission by ID.
func GetPermission(fileID, permissionID, fields string, useDomainAdminAccess bool) (*drive.Permission, error) {
	srv := getPermissionsService()
	c := srv.Get(fileID, permissionID).SupportsAllDrives(true).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, permissionID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Permission)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListPermissions lists a file's or shared drive's permissions.
func ListPermissions(fileID, includePermissionsForView, fields string, useDomainAdminAccess bool, cap int) (<-chan *drive.Permission, <-chan error) {
	srv := getPermissionsService()
	c := srv.List(fileID).SupportsAllDrives(true).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	ch := make(chan *drive.Permission, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drive.PermissionList) error {
			for i := range response.Permissions {
				ch <- response.Permissions[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// UpdatePermission updates a permission with patch semantics.
func UpdatePermission(fileID, permissionID, fields string, useDomainAdminAccess, removeExpiration, enforceExpansiveAccess bool, permission *drive.Permission) (*drive.Permission, error) {
	srv := getPermissionsService()
	permission.EmailAddress = ""
	permission.Domain = ""
	c := srv.Update(fileID, permissionID, permission).SupportsAllDrives(true).UseDomainAdminAccess(useDomainAdminAccess).RemoveExpiration(removeExpiration).EnforceExpansiveAccess(enforceExpansiveAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if permission.Role == "owner" {
		c.TransferOwnership(true)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, permissionID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Permission)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
