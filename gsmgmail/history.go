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
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

func makeListHistoryCallAndAppend(c *gmail.UsersHistoryListCall, history []*gmail.History) ([]*gmail.History, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, h := range r.History {
		history = append(history, h)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		history, err = makeListHistoryCallAndAppend(c, history)
	}
	return history, err
}

// ListHistory lists the history of all changes to the given mailbox. History results are returned in chronological order (increasing historyId).
func ListHistory(userID, labelID, fields string, startHistoryID uint64, historyTypes ...string) ([]*gmail.History, error) {
	srv := getUsersHistoryService()
	c := srv.List(userID).StartHistoryId(startHistoryID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if labelID != "" {
		c = c.LabelId(labelID)
	}
	if historyTypes != nil {
		c = c.HistoryTypes(historyTypes...)
	}
	var history []*gmail.History
	history, err := makeListHistoryCallAndAppend(c, history)
	if err != nil {
		return nil, err
	}
	return history, nil
}
