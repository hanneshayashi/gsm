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
	"context"

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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, commentID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Reply)
	return r, nil
}

// DeleteReply deletes a reply.
func DeleteReply(fileID, commentID, replyID string) (bool, error) {
	srv := getRepliesService()
	c := srv.Delete(fileID, commentID, replyID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(fileID, commentID, replyID), func() error {
		return c.Do()
	})
	return result, err
}

// GetReply gets a reply by ID.
func GetReply(fileID, commentID, replyID, fields string, includeDeleted bool) (*drive.Reply, error) {
	srv := getRepliesService()
	c := srv.Get(fileID, commentID, replyID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, commentID, replyID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Reply)
	return r, nil
}

// ListReplies Lists a comment's replies.
func ListReplies(fileID, commentID, fields string, includeDeleted bool, cap int) (<-chan *drive.Reply, <-chan error) {
	srv := getRepliesService()
	c := srv.List(fileID, commentID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *drive.Reply, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drive.ReplyList) error {
			for i := range response.Replies {
				ch <- response.Replies[i]
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

// UpdateReply updates a reply with patch semantics.
func UpdateReply(fileID, commentID, replyID, fields string, reply *drive.Reply) (*drive.Reply, error) {
	srv := getRepliesService()
	c := srv.Update(fileID, commentID, replyID, reply)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(fileID, commentID, replyID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*drive.Reply)
	return r, nil
}
