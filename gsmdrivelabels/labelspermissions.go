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

package gsmdrivelabels

import (
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

// Lists the LabelPermissions on a Label.
func ListLabelPermissions(parent, fields string, useAdminAccess bool) iter.Seq2[*drivelabels.GoogleAppsDriveLabelsV2LabelPermission, error] {
	srv := getLabelsPermissionsService()
	c := srv.List(parent).PageSize(200).UseAdminAccess(useAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*drivelabels.GoogleAppsDriveLabelsV2LabelPermission, error) bool) {
		e := c.Pages(context.Background(), func(response *drivelabels.GoogleAppsDriveLabelsV2ListLabelPermissionsResponse) error {
			for i := range response.LabelPermissions {
				if !yield(response.LabelPermissions[i], nil) {
					return errIterStopped
				}
			}
			return nil
		})
		if e != nil && !errors.Is(e, errIterStopped) {
			yield(nil, e)
		}
	}
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent), err)
	}
	return r, nil
}

// DeleteLabelPermission deletes a principal's permission on a Label.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func DeleteLabelPermission(name string, useAdminAccess bool) (bool, error) {
	srv := getLabelsPermissionsService()
	c := srv.Delete(name).UseAdminAccess(useAdminAccess)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return true, nil
}

// DeleteLabelPermission deletes a principal's permission on a Label.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func BatchDeleteLabelPermissions(parent string, request *drivelabels.GoogleAppsDriveLabelsV2BatchDeleteLabelPermissionsRequest) (bool, error) {
	srv := getLabelsPermissionsService()
	c := srv.BatchDelete(parent, request)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent), err)
	}
	return true, nil
}

// BatchDeleteLabelPermissions updates Label permissions.
// If a permission for the indicated principal doesn't exist, a new Label Permission is created, otherwise the existing permission is updated.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func BatchUpdateLabelPermissions(parent, fields string, request *drivelabels.GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsRequest) (*drivelabels.GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsResponse, error) {
	srv := getLabelsPermissionsService()
	c := srv.BatchUpdate(parent, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent), err)
	}
	return r, nil
}
