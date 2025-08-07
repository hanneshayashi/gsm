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
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

// Gets the user capabilities.
func GetCapabilities(name, customer, fields string) (*drivelabels.GoogleAppsDriveLabelsV2UserCapabilities, error) {
	srv := getUsersService()
	c := srv.GetCapabilities(name)
	if customer != "" {
		c.Customer(customer)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drivelabels.GoogleAppsDriveLabelsV2UserCapabilities)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
