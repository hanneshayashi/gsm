/*
Package gsmgmail implements the Gmail APIs
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
	"gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// DeleteThread Immediately and permanently deletes the specified thread.
// This operation cannot be undone. Prefer TrashThread instead.
func DeleteThread(userID, id string) (bool, error) {
	srv := getUsersThreadsService()
	c := srv.Delete(userID, id)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userID, id), func() error {
		return c.Do()
	})
	return result, err
}

// GetThread gets the specified thread.
func GetThread(userID, id, format, metadataHeaders, fields string) (*gmail.Thread, error) {
	srv := getUsersThreadsService()
	c := srv.Get(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if format != "" {
		c = c.Format(format)
		if format == "METADATA" {
			c = c.MetadataHeaders(metadataHeaders)
		}
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Thread)
	return r, nil
}

func makeListThreadsCallAndAppend(c *gmail.UsersThreadsListCall, threads []*gmail.Thread, errKey string) ([]*gmail.Thread, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.ListThreadsResponse)
	threads = append(threads, r.Threads...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		threads, err = makeListThreadsCallAndAppend(c, threads, errKey)
		if err != nil {
			return nil, err
		}
	}
	return threads, nil
}

// ListThreads lists the threads in the user's mailbox.
func ListThreads(userID, q, fields string, labelIDs []string, includeSpamTrash bool) ([]*gmail.Thread, error) {
	srv := getUsersThreadsService()
	c := srv.List(userID).IncludeSpamTrash(includeSpamTrash)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if q != "" {
		c = c.Q(q)
	}
	if len(labelIDs) > 0 {
		c = c.LabelIds(labelIDs...)
	}
	var threads []*gmail.Thread
	threads, err := makeListThreadsCallAndAppend(c, threads, gsmhelpers.FormatErrorKey(userID))
	if err != nil {
		return nil, err
	}
	return threads, nil
}

// ModifyThread modifies the labels applied to the thread. This applies to all messages in the thread.
func ModifyThread(userID, id, fields string, addLabelIds, removeLabelIds []string) (*gmail.Thread, error) {
	srv := getUsersThreadsService()
	c := srv.Modify(userID, id, &gmail.ModifyThreadRequest{AddLabelIds: addLabelIds, RemoveLabelIds: removeLabelIds})
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Thread)
	return r, nil
}

// TrashThread moves the specified thread to the trash.
func TrashThread(userID, id, fields string) (*gmail.Thread, error) {
	srv := getUsersThreadsService()
	c := srv.Trash(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Thread)
	return r, nil
}

// UntrashThread removes the specified thread from the trash.
func UntrashThread(userID, id, fields string) (*gmail.Thread, error) {
	srv := getUsersThreadsService()
	c := srv.Untrash(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, id), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Thread)
	return r, nil
}
