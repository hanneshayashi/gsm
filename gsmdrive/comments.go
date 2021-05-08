/*
Copyright © 2020-2021 Hannes Hayashi

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

func listComments(c *drive.CommentsListCall, ch chan *drive.Comment, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*drive.CommentList)
	for i := range r.Comments {
		ch <- r.Comments[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listComments(c, ch, errKey)
	}
	return err
}

// ListComments lists a file's comments.
func ListComments(fileID, startModifiedTime, fields string, includeDeleted bool, cap int) (<-chan *drive.Comment, <-chan error) {
	srv := getCommentsService()
	c := srv.List(fileID).IncludeDeleted(includeDeleted).PageSize(10000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if startModifiedTime != "" {
		c = c.StartModifiedTime(startModifiedTime)
	}
	ch := make(chan *drive.Comment, cap)
	err := make(chan error, 1)
	go func() {
		e := listComments(c, ch, gsmhelpers.FormatErrorKey(fileID))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
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
