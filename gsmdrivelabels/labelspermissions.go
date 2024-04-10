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

package gsmdrivelabels

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

// Lists the LabelPermissions on a Label.
func ListLabelPermissions(parent, fields string, useAdminAccess bool, cap int) (<-chan *drivelabels.GoogleAppsDriveLabelsV2LabelPermission, <-chan error) {
	srv := getLabelsPermissionsService()
	c := srv.List(parent).PageSize(200).UseAdminAccess(useAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *drivelabels.GoogleAppsDriveLabelsV2LabelPermission, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drivelabels.GoogleAppsDriveLabelsV2ListLabelPermissionsResponse) error {
			for i := range response.LabelPermissions {
				ch <- response.LabelPermissions[i]
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

// CreateLabelPermission updates a Label's permissions.
// If a permission for the indicated principal doesn't exist, a new Label Permission is created, otherwise the existing permission is updated.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func CreateLabelPermission(parent, fields string, useAdminAccess bool, permission *drivelabels.GoogleAppsDriveLabelsV2LabelPermission) (*drivelabels.GoogleAppsDriveLabelsV2LabelPermission, error) {
	srv := getLabelsPermissionsService()
	c := srv.Create(parent, permission).UseAdminAccess(useAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2LabelPermission)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// DeleteLabelPermission deletes a principal's permission on a Label.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func DeleteLabelPermission(name string, useAdminAccess bool) (bool, error) {
	srv := getLabelsPermissionsService()
	c := srv.Delete(name).UseAdminAccess(useAdminAccess)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(name), func() error {
		_, err := c.Do()
		return err
	})
	return result, err
}

// DeleteLabelPermission deletes a principal's permission on a Label.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func BatchDeleteLabelPermissions(parent string, request *drivelabels.GoogleAppsDriveLabelsV2BatchDeleteLabelPermissionsRequest) (bool, error) {
	srv := getLabelsPermissionsService()
	c := srv.BatchDelete(parent, request)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(parent), func() error {
		_, err := c.Do()
		return err
	})
	return result, err
}

// BatchDeleteLabelPermissions updatess Label permissions.
// If a permission for the indicated principal doesn't exist, a new Label Permission is created, otherwise the existing permission is updated.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func BatchUpdateLabelPermissions(parent, fields string, request *drivelabels.GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsRequest) (*drivelabels.GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsResponse, error) {
	srv := getLabelsPermissionsService()
	c := srv.BatchUpdate(parent, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
