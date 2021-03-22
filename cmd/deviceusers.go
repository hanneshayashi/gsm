/*
Package cmd contains the commands available to the end user
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
package cmd

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"
	ci "google.golang.org/api/cloudidentity/v1"

	"github.com/spf13/cobra"
)

// deviceUsersCmd represents the deviceUsers command
var deviceUsersCmd = &cobra.Command{
	Use:               "deviceUsers",
	Short:             "Manage device users (Part of Cloud Identity API)",
	Long:              "https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var deviceUserFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor:   []string{"approve", "block"},
		Type:           "string",
		Description:    `Resource name of the Device in format: devices/{device_id}/deviceUsers/{device_user_id}, where device_id is the unique ID assigned to the Device, and device_user_id is the unique ID assigned to the User.`,
		Required:       []string{"approve", "block"},
		ExcludeFromAll: true,
	},
	"customer": {
		AvailableFor: []string{"cancelWipe", "create", "delete", "get", "list", "wipe"},
		Type:         "string",
		Description: `Resource name of the customer.
If you're using this API for your own organization, use customers/my_customer.
If you're using this API to manage another organization, use customers/{customer_id}, where customer_id is the customer to whom the device belongs.`,
		Defaults: map[string]interface{}{"cancelWipe": "customers/my_customer", "create": "customers/my_customer", "delete": "customers/my_customer", "get": "customers/my_customer", "list": "customers/my_customer", "wipe": "customers/my_customer"},
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var deviceUserFlagsALL = gsmhelpers.GetAllFlags(deviceUserFlags)

func init() {
	rootCmd.AddCommand(deviceUsersCmd)
}

func mapToApproveDeviceUserRequest(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1ApproveDeviceUserRequest, error) {
	approveDeviceUserRequest := &ci.GoogleAppsCloudidentityDevicesV1ApproveDeviceUserRequest{}
	if flags["customer"].IsSet() {
		approveDeviceUserRequest.Customer = flags["customer"].GetString()
		if approveDeviceUserRequest.Customer == "" {
			approveDeviceUserRequest.ForceSendFields = append(approveDeviceUserRequest.ForceSendFields, "Customer")
		}
	}
	return approveDeviceUserRequest, nil
}

func mapToBlockDeviceUserRequest(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1BlockDeviceUserRequest, error) {
	blockDeviceUserRequest := &ci.GoogleAppsCloudidentityDevicesV1BlockDeviceUserRequest{}
	if flags["customer"].IsSet() {
		blockDeviceUserRequest.Customer = flags["customer"].GetString()
		if blockDeviceUserRequest.Customer == "" {
			blockDeviceUserRequest.ForceSendFields = append(blockDeviceUserRequest.ForceSendFields, "Customer")
		}
	}
	return blockDeviceUserRequest, nil
}

func mapToCancelWipeDeviceUserRequest(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1CancelWipeDeviceUserRequest, error) {
	cancelWipeRequest := &ci.GoogleAppsCloudidentityDevicesV1CancelWipeDeviceUserRequest{}
	if flags["customer"].IsSet() {
		cancelWipeRequest.Customer = flags["customer"].GetString()
		if cancelWipeRequest.Customer == "" {
			cancelWipeRequest.ForceSendFields = append(cancelWipeRequest.ForceSendFields, "Customer")
		}
	}
	return cancelWipeRequest, nil
}

func mapToWipeDeviceUserRequest(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1WipeDeviceUserRequest, error) {
	wipeRequest := &ci.GoogleAppsCloudidentityDevicesV1WipeDeviceUserRequest{}
	if flags["customer"].IsSet() {
		wipeRequest.Customer = flags["customer"].GetString()
		if wipeRequest.Customer == "" {
			wipeRequest.ForceSendFields = append(wipeRequest.ForceSendFields, "Customer")
		}
	}
	return wipeRequest, nil
}
