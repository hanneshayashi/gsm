/*
Package gsmdrive implements the Drive API
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Comment)
	return r, nil
}

// DeleteComment deletes a comment.
func DeleteComment(fileID, commentID string) (bool, error) {
	srv := getCommentsService()
	c := srv.Delete(fileID, commentID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(fileID, commentID), func() error {
		return c.Do()
	})
	return result, err
}

// GetComment gets a comment by ID.
func GetComment(fileID, commentID, fields string, includeDeleted bool) (*drive.Comment, error) {
	srv := getCommentsService()
	c := srv.Get(fileID, commentID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, commentID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Comment)
	return r, nil
}

func makeListCommentsCallAndAppend(c *drive.CommentsListCall, comments []*drive.Comment, errKey string) ([]*drive.Comment, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.CommentList)
	comments = append(comments, r.Comments...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		comments, err = makeListCommentsCallAndAppend(c, comments, errKey)
	}
	return comments, err
}

// ListComments lists a file's comments.
func ListComments(fileID, startModifiedTime, fields string, includeDeleted bool) ([]*drive.Comment, error) {
	srv := getCommentsService()
	c := srv.List(fileID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if startModifiedTime != "" {
		c = c.StartModifiedTime(startModifiedTime)
	}
	var comments []*drive.Comment
	comments, err := makeListCommentsCallAndAppend(c, comments, gsmhelpers.FormatErrorKey(fileID))
	return comments, err
}

// UpdateComment updates a comment with patch semantics.
func UpdateComment(fileID, commentID, fields string, comment *drive.Comment) (*drive.Comment, error) {
	srv := getCommentsService()
	c := srv.Update(fileID, commentID, comment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, commentID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Comment)
	return r, nil
}
