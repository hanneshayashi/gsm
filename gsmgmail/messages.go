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
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// BatchDeleteMessages deletes many messages by message ID. Provides no guarantees that messages were not already deleted or even existed at all.
func BatchDeleteMessages(userID string, ids []string) (bool, error) {
	srv := getUsersMessagesService()
	c := srv.BatchDelete(userID, &gmail.BatchDeleteMessagesRequest{Ids: ids})
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID), err)
	}
	return true, nil
}

// BatchModifyMessages modifies the labels on the specified messages.
func BatchModifyMessages(userID string, ids, addLabelIds, removeLabelIds []string) (bool, error) {
	srv := getUsersMessagesService()
	c := srv.BatchModify(userID, &gmail.BatchModifyMessagesRequest{Ids: ids, AddLabelIds: addLabelIds, RemoveLabelIds: removeLabelIds})
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID), err)
	}
	return true, nil
}

// DeleteMessage immediately and permanently deletes the specified message.
// This operation cannot be undone. Prefer messages.trash instead.
func DeleteMessage(userID, id string) (bool, error) {
	srv := getUsersMessagesService()
	c := srv.Delete(userID, id)
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID, id), err)
	}
	return true, nil
}

// GetMessage gets the specified message.
func GetMessage(userID, id, format, metadataHeaders, fields string) (*gmail.Message, error) {
	srv := getUsersMessagesService()
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID, id), err)
	}
	return r, nil
}

// ImportMessage imports a message into only this user's mailbox, with standard email delivery scanning and classification similar to receiving via SMTP.
// Does not send a message.
func ImportMessage(userID, internalDateSource, fields string, message *gmail.Message, deleted, neverMarkSpam, processForCalendar bool) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Import(userID, message).Deleted(deleted).InternalDateSource(internalDateSource).NeverMarkSpam(neverMarkSpam).ProcessForCalendar(processForCalendar)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID, message.Id), err)
	}
	return r, nil
}

// InsertMessage directly inserts a message into only this user's mailbox similar to IMAP APPEND, bypassing most scanning and classification.
// Does not send a message.
func InsertMessage(userID, internalDateSource, fields string, message *gmail.Message, deleted bool) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Insert(userID, message).Deleted(deleted).InternalDateSource(internalDateSource)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID, message.Id), err)
	}
	return r, nil
}

// ListMessages lists the messages in the user's mailbox.
func ListMessages(userID, q, fields string, labelIds []string, includeSpamTrash bool) iter.Seq2[*gmail.Message, error] {
	srv := getUsersMessagesService()
	c := srv.List(userID).MaxResults(10000).IncludeSpamTrash(includeSpamTrash)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if labelIds != nil {
		c = c.LabelIds(labelIds...)
	}
	if q != "" {
		c = c.Q(q)
	}
	return func(yield func(*gmail.Message, error) bool) {
		e := c.Pages(context.Background(), func(response *gmail.ListMessagesResponse) error {
			for i := range response.Messages {
				if !yield(response.Messages[i], nil) {
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

// ModifyMessage modifies the labels on the specified message.
func ModifyMessage(userID, id, fields string, addLabelIds, removeLabelIds []string) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Modify(userID, id, &gmail.ModifyMessageRequest{AddLabelIds: addLabelIds, RemoveLabelIds: removeLabelIds})
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID, id), err)
	}
	return r, nil
}

// SendMessage sends the specified message to the recipients in the To, Cc, and Bcc headers.
func SendMessage(userID, fields string, message *gmail.Message) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Send(userID, message)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID), err)
	}
	return r, nil
}

// TrashMessage moves the specified message to the trash.
func TrashMessage(userID, id, fields string) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Trash(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID, id), err)
	}
	return r, nil
}

// UntrashMessage removes the specified message from the trash.
func UntrashMessage(userID, id, fields string) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Untrash(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(userID, id), err)
	}
	return r, nil
}
