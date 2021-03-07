/*
Package gsmgmail implements the Gmail APIs
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
package gsmgmail

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

func listHistory(c *gmail.UsersHistoryListCall, ch chan *gmail.History, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*gmail.ListHistoryResponse)
	for i := range r.History {
		ch <- r.History[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listHistory(c, ch, errKey)
	}
	return err
}

// ListHistory lists the history of all changes to the given mailbox. History results are returned in chronological order (increasing historyId).
func ListHistory(userID, labelID, fields string, startHistoryID uint64, historyTypes []string, cap int) (<-chan *gmail.History, <-chan error) {
	srv := getUsersHistoryService()
	c := srv.List(userID).StartHistoryId(startHistoryID).MaxResults(10000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if labelID != "" {
		c = c.LabelId(labelID)
	}
	if historyTypes != nil {
		c = c.HistoryTypes(historyTypes...)
	}
	ch := make(chan *gmail.History, cap)
	err := make(chan error, 1)
	go func() {
		e := listHistory(c, ch, gsmhelpers.FormatErrorKey(userID))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}
