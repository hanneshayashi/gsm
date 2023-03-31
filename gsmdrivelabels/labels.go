/*
Copyright © 2020-2022 Hannes Hayashi

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

// GetLabel gets a label by its resource name. Resource name may be any of:
// labels/{id} - See labels/{id}@latest
// labels/{id}@latest - Gets the latest revision of the label.
// labels/{id}@published - Gets the current published revision of the label.
// labels/{id}@{revisionId} - Gets the label at the specified revision ID.
func GetLabel(fileID, languageCode, view, fields string, useAdminAccess bool) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	srv := getLabelsService()
	c := srv.Get(fileID).UseAdminAccess(useAdminAccess)
	if languageCode != "" {
		c.LanguageCode(languageCode)
	}
	if view != "" {
		c.View(view)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (any, error) {
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
	c := srv.List().PublishedOnly(publishedOnly)
	if useAdminAccess {
		c.UseAdminAccess(useAdminAccess)
	}
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
