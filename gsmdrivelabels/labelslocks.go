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

package gsmdrivelabels

import (
	"errors"
	"context"
	"iter"
	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

// Lists the LabelLocks on a Label.
func ListLabelLocks(parent, fields string) iter.Seq2[*drivelabels.GoogleAppsDriveLabelsV2LabelLock, error] {
	srv := getLabelsLocksService()
	c := srv.List(parent).PageSize(200)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*drivelabels.GoogleAppsDriveLabelsV2LabelLock, error) bool) {
		e := c.Pages(context.Background(), func(response *drivelabels.GoogleAppsDriveLabelsV2ListLabelLocksResponse) error {
			for i := range response.LabelLocks {
				if !yield(response.LabelLocks[i], nil) {
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
