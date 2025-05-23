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
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// CreateFilter creates a filter.
func CreateFilter(userID, fields string, settingsfilter *gmail.Filter) (*gmail.Filter, error) {
	srv := getUsersSettingsFiltersService()
	c := srv.Create(userID, settingsfilter)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*gmail.Filter)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// DeleteFilter deletes a filter.
func DeleteFilter(userID, id string) (bool, error) {
	srv := getUsersSettingsFiltersService()
	c := srv.Delete(userID, id)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userID, id), func() error {
		return c.Do()
	})
	return result, err
}

// GetFilter gets a filter.
func GetFilter(userID, id, fields string) (*gmail.Filter, error) {
	srv := getUsersSettingsFiltersService()
	c := srv.Get(userID, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID, id), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*gmail.Filter)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListFilters lists the message filters of a Gmail user.
func ListFilters(userID, fields string) ([]*gmail.Filter, error) {
	srv := getUsersSettingsFiltersService()
	c := srv.List(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*gmail.ListFiltersResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r.Filter, nil
}
