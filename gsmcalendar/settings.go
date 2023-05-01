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
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// GetSetting returns a single user setting.
func GetSetting(setting, fields string) (*calendar.Setting, error) {
	srv := getSettingsService()
	c := srv.Get(setting)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(setting), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*calendar.Setting)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListSettings returns all user settings for the authenticated user.
func ListSettings(fields string, cap int) (<-chan *calendar.Setting, <-chan error) {
	srv := getSettingsService()
	c := srv.List().MaxResults(250)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *calendar.Setting, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *calendar.Settings) error {
			for i := range response.Items {
				ch <- response.Items[i]
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
