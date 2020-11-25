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
	"io"

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
	for _, m := range media {
		c = c.Media(m)
	}
	r, err := c.Do()
	return r, err
}

// DeleteDraft immediately and permanently deletes the specified draft. Does not simply trash it.
func DeleteDraft(userID, id string) (bool, error) {
	srv := getUsersDraftsService()
	err := srv.Delete(userID, id).Do()
	if err != nil {
		return false, err
	}
	return true, nil
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
	r, err := c.Do()
	return r, err
}

func makeListDraftsCallAndAppend(c *gmail.UsersDraftsListCall, drafts []*gmail.Draft) ([]*gmail.Draft, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, d := range r.Drafts {
		drafts = append(drafts, d)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		drafts, err = makeListDraftsCallAndAppend(c, drafts)
		if err != nil {
			return nil, err
		}
	}
	return drafts, err
}

// ListDrafts lists the drafts in the user's mailbox.
func ListDrafts(userID, q, fields string, includeSpamTrash bool) ([]*gmail.Draft, error) {
	srv := getUsersDraftsService()
	c := srv.List(userID).IncludeSpamTrash(includeSpamTrash)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if q != "" {
		c = c.Q(q)
	}
	var drafts []*gmail.Draft
	drafts, err := makeListDraftsCallAndAppend(c, drafts)
	if err != nil {
		return nil, err
	}
	return drafts, nil
}

// SendDraft sends the specified, existing draft to the recipients in the To, Cc, and Bcc headers.
func SendDraft(userID string, draft *gmail.Draft, media ...io.Reader) (*gmail.Message, error) {
	srv := getUsersDraftsService()
	c := srv.Send(userID, draft)
	for _, m := range media {
		c = c.Media(m)
	}
	r, err := c.Do()
	return r, err
}

// UpdateDraft replaces a draft's content.
func UpdateDraft(userID, id, fields string, draft *gmail.Draft, media ...io.Reader) (*gmail.Draft, error) {
	srv := getUsersDraftsService()
	c := srv.Update(userID, id, draft)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	for _, m := range media {
		c = c.Media(m)
	}
	r, err := c.Do()
	return r, err
}
