/*
Copyright © 2020-2023 Hannes Hayashi

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

package gsmcalendar

import (
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// QueryFreeBusy returns free/busy information for a set of calendars.
func QueryFreeBusy(freeBusyRequest *calendar.FreeBusyRequest, fields string) (*calendar.FreeBusyResponse, error) {
	srv := getFreebusyService()
	c := srv.Query(freeBusyRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey("Free/Busy Query"), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*calendar.FreeBusyResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
