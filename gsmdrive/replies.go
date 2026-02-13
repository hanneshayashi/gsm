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

// CreateReply creates a new reply to a comment.
func CreateReply(fileID, commentID, fields string, reply *drive.Reply) (*drive.Reply, error) {
	srv := getRepliesService()
	c := srv.Create(fileID, commentID, reply)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID, commentID), err)
	}
	return r, nil
}

// DeleteReply deletes a reply.
func DeleteReply(fileID, commentID, replyID string) (bool, error) {
	srv := getRepliesService()
	c := srv.Delete(fileID, commentID, replyID)
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID, commentID, replyID), err)
	}
	return true, nil
}

// GetReply gets a reply by ID.
func GetReply(fileID, commentID, replyID, fields string, includeDeleted bool) (*drive.Reply, error) {
	srv := getRepliesService()
	c := srv.Get(fileID, commentID, replyID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID, commentID, replyID), err)
	}
	return r, nil
}

// ListReplies Lists a comment's replies.
func ListReplies(fileID, commentID, fields string, includeDeleted bool) iter.Seq2[*drive.Reply, error] {
	srv := getRepliesService()
	c := srv.List(fileID, commentID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*drive.Reply, error) bool) {
		e := c.Pages(context.Background(), func(response *drive.ReplyList) error {
			for i := range response.Replies {
				if !yield(response.Replies[i], nil) {
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

// UpdateReply updates a reply with patch semantics.
func UpdateReply(fileID, commentID, replyID, fields string, reply *drive.Reply) (*drive.Reply, error) {
	srv := getRepliesService()
	c := srv.Update(fileID, commentID, replyID, reply)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(fileID, commentID, replyID), err)
	}
	return r, nil
}
