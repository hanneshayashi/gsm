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
package gsmgmail

import (
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// CreateLabel creates a new label.
func CreateLabel(userID, fields string, label *gmail.Label) (*gmail.Label, error) {
	srv := getUsersLabelsService()
	c := srv.Create(userID, label)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// DeleteLabel immediately and permanently deletes the specified label and removes it from any messages and threads that it is applied to.
func DeleteLabel(userID, id string) (bool, error) {
	srv := getUsersLabelsService()
	err := srv.Delete(userID, id).Do()
	if err != nil {
		return false, err
	}
	return true, err
}

// GetLabel gets the specified label.
func GetLabel(userID, id, fields string) (*gmail.Label, error) {
	srv := getUsersLabelsService()
	c := srv.Get(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// ListLabels lists all labels in the user's mailbox.
func ListLabels(userID, fields string) ([]*gmail.Label, error) {
	srv := getUsersLabelsService()
	c := srv.List(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.Labels, nil
}

// PatchLabel PATCHes the specified label.
func PatchLabel(userID, id, fields string, label *gmail.Label) (*gmail.Label, error) {
	srv := getUsersLabelsService()
	c := srv.Patch(userID, id, label)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
