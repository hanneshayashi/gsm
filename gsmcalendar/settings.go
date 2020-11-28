/*
Package gsmcalendar implements the Calendar API
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
package gsmcalendar

import (
	"gsm/gsmhelpers"

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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(setting), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Setting)
	return r, nil
}

func makeListSettingsCallAndAppend(c *calendar.SettingsListCall, settings []*calendar.Setting, errKey string) ([]*calendar.Setting, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Settings)
	for _, s := range r.Items {
		settings = append(settings, s)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		settings, err = makeListSettingsCallAndAppend(c, settings, errKey)
	}
	return settings, err
}

// ListSettings returns all user settings for the authenticated user.
func ListSettings(fields string) ([]*calendar.Setting, error) {
	srv := getSettingsService()
	c := srv.List()
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var settings []*calendar.Setting
	settings, err := makeListSettingsCallAndAppend(c, settings, gsmhelpers.FormatErrorKey("List settings"))
	return settings, err
}
