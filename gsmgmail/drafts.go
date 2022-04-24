/*
Copyright © 2020-2022 Hannes Hayashi

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
	"context"
	"io"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// CreateDraft creates a new draft with the DRAFT label.
func CreateDraft(userID, fields string, draft *gmail.Draft, media ...io.Reader) (*gmail.Draft, error) {
	srv := getUsersDraftsService()
	c := srv.Create(userID, draft)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	for i := range media {
		c = c.Media(media[i])
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Draft)
	return r, nil
}

// DeleteDraft immediately and permanently deletes the specified draft. Does not simply trash it.
func DeleteDraft(userID, id string) (bool, error) {
	srv := getUsersDraftsService()
	c := srv.Delete(userID, id)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userID, id), func() error {
		return c.Do()
	})
	return result, err
}

// GetDraft gets the specified draft.
func GetDraft(userID, id, format, fields string) (*gmail.Draft, error) {
	srv := getUsersDraftsService()
	c := srv.Get(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if format != "" {
		c = c.Format(format)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, id), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Draft)
	return r, nil
}

// ListDrafts lists the drafts in the user's mailbox.
func ListDrafts(userID, q, fields string, includeSpamTrash bool, cap int) (<-chan *gmail.Draft, <-chan error) {
	srv := getUsersDraftsService()
	c := srv.List(userID).IncludeSpamTrash(includeSpamTrash).MaxResults(10000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if q != "" {
		c = c.Q(q)
	}
	ch := make(chan *gmail.Draft, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *gmail.ListDraftsResponse) error {
			for i := range response.Drafts {
				ch <- response.Drafts[i]
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

// SendDraft sends the specified, existing draft to the recipients in the To, Cc, and Bcc headers.
func SendDraft(userID string, draft *gmail.Draft, media ...io.Reader) (*gmail.Message, error) {
	srv := getUsersDraftsService()
	c := srv.Send(userID, draft)
	for i := range media {
		c = c.Media(media[i])
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Message)
	return r, nil
}

// UpdateDraft replaces a draft's content.
func UpdateDraft(userID, id, fields string, draft *gmail.Draft, media ...io.Reader) (*gmail.Draft, error) {
	srv := getUsersDraftsService()
	c := srv.Update(userID, id, draft)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	for i := range media {
		c = c.Media(media[i])
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, id), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.Draft)
	return r, nil
}
