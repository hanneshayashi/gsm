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
	return r, err
}

// DeleteReply deletes a reply.
func DeleteReply(fileID, commentID, replyID string) (bool, error) {
	srv := getRepliesService()
	err := srv.Delete(fileID, commentID, replyID).Do()
	if err != nil {
		return false, err
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
	return r, err
}

func makeListRepliesCallAndAppend(c *drive.RepliesListCall, replies []*drive.Reply) ([]*drive.Reply, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, p := range r.Replies {
		replies = append(replies, p)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		replies, err = makeListRepliesCallAndAppend(c, replies)
	}
	return replies, err
}

// ListReplies Lists a comment's replies.
func ListReplies(fileID, commentID, fields string, includeDeleted bool) ([]*drive.Reply, error) {
	srv := getRepliesService()
	c := srv.List(fileID, commentID).IncludeDeleted(includeDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var replies []*drive.Reply
	replies, err := makeListRepliesCallAndAppend(c, replies)
	return replies, err
}

// UpdateReply updates a reply with patch semantics.
func UpdateReply(fileID, commentID, replyID, fields string, reply *drive.Reply) (*drive.Reply, error) {
	srv := getRepliesService()
	c := srv.Update(fileID, commentID, replyID, reply)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
