/*
Copyright © 2020-2022 Hannes Hayashi

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

package gsmadmin

import (
	"fmt"
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// GetCommand gets command data a specific command issued to the device.
func GetCommand(customerID, deviceID, fields string, commandID int64) (*admin.DirectoryChromeosdevicesCommand, error) {
	srv := getCustomerDevicesChromeosCommandsService()
	c := srv.Get(customerID, deviceID, commandID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, deviceID, strconv.FormatInt(commandID, 10)), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.DirectoryChromeosdevicesCommand)
	if !ok {
		return nil, fmt.Errorf("Result unknown")
	}
	return r, nil
}
