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
package gsmcalendar

import (
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// GetColors returns the color definitions for calendars and events.
func GetColors(fields string) (*calendar.Colors, error) {
	srv := getColorsService()
	c := srv.Get()
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
