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
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

// CreateLabel creates a new Label.
func CreateLabel(label *drivelabels.GoogleAppsDriveLabelsV2Label, languageCode, fields string, useAdminAccess bool) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	srv := getLabelsService()
	c := srv.Create(label).UseAdminAccess(useAdminAccess)
	if languageCode != "" {
		c.LanguageCode(languageCode)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(label.Name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2Label)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// DeleteLabel permanently deletes a Label and related metadata on Drive Items.
// Once deleted, the Label and related Drive item metadata will be deleted.
// Only draft Labels, and disabled Labels may be deleted.
func DeleteLabel(name, requiredRevisionId string, useAdminAccess bool) (bool, error) {
	srv := getLabelsService()
	c := srv.Delete(name).UseAdminAccess(useAdminAccess)
	if requiredRevisionId != "" {
		c.WriteControlRequiredRevisionId(requiredRevisionId)
	}
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(name), func() error {
		_, err := c.Do()
		return err
	})
	return result, err
}

// Delta updates a single Label by applying a set of update requests resulting in a new draft revision.
// The batch update is all-or-nothing: If any of the update requests are invalid, no changes are applied.
// The resulting draft revision must be published before the changes may be used with Drive Items.
func Delta(name, fields string, request *drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelResponse, error) {
	srv := getLabelsService()
	c := srv.Delta(name, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// Disable disables a published Label.
// Disabling a Label will result in a new disabled published revision based on the current published revision.
// If there is a draft revision, a new disabled draft revision will be created based on the latest draft revision.
// Older draft revisions will be deleted.
func Disable(name, fields string, request *drivelabels.GoogleAppsDriveLabelsV2DisableLabelRequest) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	srv := getLabelsService()
	c := srv.Disable(name, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2Label)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// Enable enables a disabled Label and restore it to its published state.
// This will result in a new published revision based on the current disabled published revision.
// If there is an existing disabled draft revision, a new revision will be created based on that draft and will be enabled.
func Enable(name, fields string, request *drivelabels.GoogleAppsDriveLabelsV2EnableLabelRequest) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	srv := getLabelsService()
	c := srv.Enable(name, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2Label)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// Publish all draft changes to the Label.
// Once published, the Label may not return to its draft state.
// See google.apps.drive.labels.v2.Lifecycle for more information.
// Publishing a Label will result in a new published revision.
// All previous draft revisions will be deleted.
// Previous published revisions will be kept but are subject to automated deletion as needed.
// Once published, some changes are no longer permitted.
// Generally, any change that would invalidate or cause new restrictions on existing metadata related to the Label will be rejected.
func Publish(name, fields string, request *drivelabels.GoogleAppsDriveLabelsV2PublishLabelRequest) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	srv := getLabelsService()
	c := srv.Publish(name, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2Label)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// Publish all draft changes to the Label.
// Once published, the Label may not return to its draft state.
// See google.apps.drive.labels.v2.Lifecycle for more information.
// Publishing a Label will result in a new published revision.
// All previous draft revisions will be deleted.
// Previous published revisions will be kept but are subject to automated deletion as needed.
// Once published, some changes are no longer permitted.
// Generally, any change that would invalidate or cause new restrictions on existing metadata related to the Label will be rejected.
func UpdateLabelCopyMode(name, fields string, request *drivelabels.GoogleAppsDriveLabelsV2UpdateLabelCopyModeRequest) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	srv := getLabelsService()
	c := srv.UpdateLabelCopyMode(name, request)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2Label)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// UpdatePermissions updates a Label's permissions.
// If a permission for the indicated principal doesn't exist, a new Label Permission is created, otherwise the existing permission is updated.
// Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.
func UpdatePermissions(name, fields string, useAdminAccess bool, request *drivelabels.GoogleAppsDriveLabelsV2LabelPermission) (*drivelabels.GoogleAppsDriveLabelsV2LabelPermission, error) {
	srv := getLabelsService()
	c := srv.UpdatePermissions(name, request).UseAdminAccess(useAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
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

// GetLabel gets a label by its resource name. Resource name may be any of:
// labels/{id} - See labels/{id}@latest
// labels/{id}@latest - Gets the latest revision of the label.
// labels/{id}@published - Gets the current published revision of the label.
// labels/{id}@{revisionId} - Gets the label at the specified revision ID.
func GetLabel(name, languageCode, view, fields string, useAdminAccess bool) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	srv := getLabelsService()
	c := srv.Get(name).UseAdminAccess(useAdminAccess)
	if languageCode != "" {
		c.LanguageCode(languageCode)
	}
	if view != "" {
		c.View(view)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2Label)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListLabels list labels.
func ListLabels(languageCode, view, minimumRole, fields string, useAdminAccess, publishedOnly bool, cap int) (<-chan *drivelabels.GoogleAppsDriveLabelsV2Label, <-chan error) {
	srv := getLabelsService()
	c := srv.List().PublishedOnly(publishedOnly).UseAdminAccess(useAdminAccess)
	if languageCode != "" {
		c.LanguageCode(languageCode)
	}
	if view != "" {
		c.View(view)
	}
	if minimumRole != "" {
		c.MinimumRole(minimumRole)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *drivelabels.GoogleAppsDriveLabelsV2Label, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drivelabels.GoogleAppsDriveLabelsV2ListLabelsResponse) error {
			for i := range response.Labels {
				ch <- response.Labels[i]
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
