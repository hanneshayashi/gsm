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
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// CreateComment creates a new comment on a file.
func CreateComment(fileID, fields string, comment *drive.Comment) (*drive.Comment, error) {
	srv := getCommentsService()
	c := srv.Create(fileID, comment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID), err)
	}
	return r, nil
}

// DeleteComment deletes a comment.
func DeleteComment(fileID, commentID string) (bool, error) {
	srv := getCommentsService()
	c := srv.Delete(fileID, commentID)
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID, commentID), err)
	}
	return true, nil
}

// GetComment gets a comment by ID.
func GetComment(fileID, commentID, fields string, includeDeleted bool) (*drive.Comment, error) {
	srv := getCommentsService()
	c := srv.Get(fileID, commentID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID, commentID), err)
	}
	return r, nil
}

// ListComments lists a file's comments.
func ListComments(fileID, startModifiedTime, fields string, includeDeleted bool) iter.Seq2[*drive.Comment, error] {
	srv := getCommentsService()
	c := srv.List(fileID).IncludeDeleted(includeDeleted).PageSize(10000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if startModifiedTime != "" {
		c = c.StartModifiedTime(startModifiedTime)
	}
	return func(yield func(*drive.Comment, error) bool) {
		e := c.Pages(context.Background(), func(response *drive.CommentList) error {
			for i := range response.Comments {
				if !yield(response.Comments[i], nil) {
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

// UpdateComment updates a comment with patch semantics.
func UpdateComment(fileID, commentID, fields string, comment *drive.Comment) (*drive.Comment, error) {
	srv := getCommentsService()
	c := srv.Update(fileID, commentID, comment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID, commentID), err)
	}
	return r, nil
}
