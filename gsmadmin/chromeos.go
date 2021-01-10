/*
Package gsmadmin implements the Admin SDK APIs
Copyright © 2020-2021 Hannes Hayashi

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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
)

// IssueCommand issues a command for the device to execute.
func IssueCommand(customerID, deviceID string, issueCommandRequest *admin.DirectoryChromeosdevicesIssueCommandRequest) (int64, error) {
	srv := getCustomerDevicesChromeosService()
	c := srv.IssueCommand(customerID, deviceID, issueCommandRequest)
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, deviceID, issueCommandRequest.CommandType, issueCommandRequest.Payload), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return 0, err
	}
	r, _ := result.(*admin.DirectoryChromeosdevicesIssueCommandResponse)
	return r.CommandId, nil
}
