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
	"iter"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// ListHistory lists the history of all changes to the given mailbox. History results are returned in chronological order (increasing historyId).
func ListHistory(userID, labelID, fields string, startHistoryID uint64, historyTypes []string) iter.Seq2[*gmail.History, error] {
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
	return func(yield func(*gmail.History, error) bool) {
		e := c.Pages(context.Background(), func(response *gmail.ListHistoryResponse) error {
			for i := range response.History {
				if !yield(response.History[i], nil) {
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
