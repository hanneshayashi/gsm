/*
Copyright © 2020-2025 Hannes Hayashi

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
	"context"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

// Lists the LabelLocks on a Label.
func ListLabelLocks(parent, fields string, cap int) (<-chan *drivelabels.GoogleAppsDriveLabelsV2LabelLock, <-chan error) {
	srv := getLabelsLocksService()
	c := srv.List(parent).PageSize(200)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *drivelabels.GoogleAppsDriveLabelsV2LabelLock, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drivelabels.GoogleAppsDriveLabelsV2ListLabelLocksResponse) error {
			for i := range response.LabelLocks {
				ch <- response.LabelLocks[i]
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
