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

// BatchDeleteMessages deletes many messages by message ID. Provides no guarantees that messages were not already deleted or even existed at all.
func BatchDeleteMessages(userID string, ids []string) (bool, error) {
	srv := getUsersMessagesService()
	err := srv.BatchDelete(userID, &gmail.BatchDeleteMessagesRequest{Ids: ids}).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// BatchModifyMessages modifies the labels on the specified messages.
func BatchModifyMessages(userID string, ids, addLabelIds, removeLabelIds []string) (bool, error) {
	srv := getUsersMessagesService()
	err := srv.BatchModify(userID, &gmail.BatchModifyMessagesRequest{Ids: ids, AddLabelIds: addLabelIds, RemoveLabelIds: removeLabelIds}).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteMessage immediately and permanently deletes the specified message.
// This operation cannot be undone. Prefer messages.trash instead.
func DeleteMessage(userID, id string) (bool, error) {
	srv := getUsersMessagesService()
	err := srv.Delete(userID, id).Do()
	if err != nil {
		return false, err
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
	return r, err
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
	return r, err
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
	return r, err
}

func makeListMessagesCallAndAppend(c *gmail.UsersMessagesListCall, messages []*gmail.Message) ([]*gmail.Message, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, m := range r.Messages {
		messages = append(messages, m)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		messages, err = makeListMessagesCallAndAppend(c, messages)
	}
	return messages, err
}

// ListMessages lists the messages in the user's mailbox.
func ListMessages(userID, q, fields string, labelIds []string, includeSpamTrash bool) ([]*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.List(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if labelIds != nil {
		c = c.LabelIds(labelIds...)
	}
	if q != "" {
		c = c.Q(q)
	}
	var messages []*gmail.Message
	messages, err := makeListMessagesCallAndAppend(c, messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// ModifyMessages modifies the labels on the specified messages.
func ModifyMessages(userID, id, fields string, addLabelIds, removeLabelIds []string) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Modify(userID, id, &gmail.ModifyMessageRequest{AddLabelIds: addLabelIds, RemoveLabelIds: removeLabelIds})
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// SendMessage sends the specified message to the recipients in the To, Cc, and Bcc headers.
func SendMessage(userID, fields string, message *gmail.Message) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Send(userID, message)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// TrashMessage moves the specified message to the trash.
func TrashMessage(userID, id, fields string) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Trash(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// UntrashMessage removes the specified message from the trash.
func UntrashMessage(userID, id, fields string) (*gmail.Message, error) {
	srv := getUsersMessagesService()
	c := srv.Untrash(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
