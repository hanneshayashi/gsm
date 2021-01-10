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

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// mobileDevicesCmd represents the mobileDevices command
var mobileDevicesCmd = &cobra.Command{
	Use:               "mobileDevices",
	Short:             "Manage Mobile Devices (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/mobiledevices",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var mobileDeviceFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customerId": {
		AvailableFor: []string{"action", "delete", "get", "list"},
		Type:         "string",
		Description: `The unique ID for the customer's Workspace account.
As an account administrator, you can also use the my_customer alias to represent your account's customerId.
The customerId is also returned as part of the Users resource.`,
		Defaults: map[string]interface{}{"action": "my_customer", "delete": "my_customer", "get": "my_customer", "list": "my_customer"},
	},
	"resourceId": {
		AvailableFor:   []string{"action", "delete", "get"},
		Type:           "string",
		Description:    `The unique ID the API service uses to identify the mobile device.`,
		ExcludeFromAll: true,
	},
	"action": {
		AvailableFor: []string{"action"},
		Type:         "string",
		Description: `The action to be performed on the device.
[admin_account_wipe|admin_remote_wipe|approve|approve|block|cancel_remote_wipe_then_activate|cancel_remote_wipe_then_block]
admin_account_wipe                - Remotely wipes only Workspace data from the device. See the administration help center for more information.
admin_remote_wipe                 - Remotely wipes all data on the device. See the administration help center for more information.
approve                           - Approves the device. If you've selected Enable device activation, devices that register after the device activation setting is enabled will need to be approved before they can start syncing with your domain. Enabling device activation forces the device user to install the Device Policy app to sync with Workspace.
block                             - Blocks access to Workspace data (mail, calendar, and contacts) on the device. The user can still access their mail, calendar, and contacts from a desktop computer or mobile browser.
cancel_remote_wipe_then_activate  - Cancels a remote wipe of the device and then reactivates it.
cancel_remote_wipe_then_block     - Cancels a remote wipe of the device and then blocks it.`,
	},
	"projection": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Restrict information returned to a set of selected fields.
Acceptable values are:
BASIC  - Includes only the basic metadata fields (e.g., deviceId, model, status, type, and status)
FULL   - Includes all metadata fields`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Device property to use for sorting results.
Acceptable values are:
deviceId  - The serial number for a Google Sync mobile device. For Android devices, this is a software generated unique identifier.
email     - The device owner's email address.
lastSync  - Last policy settings sync date time of the device.
model     - The mobile device's model.
name      - The device owner's user name.
os        - The device's operating system.
status    - The device status.
type      - Type of the device.`,
	},
	"query": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Search string in the format provided by List query operators.
See https://developers.google.com/admin-sdk/directory/v1/list-query-operators`,
	},
	"sortOrder": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Whether to return results in ascending or descending order. Must be used with the orderBy parameter.
Acceptable values are:
ASCENDING   - Ascending order.
DESCENDING  - Descending order.`,
	},
	"fields": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var mobileDeviceFlagsALL = gsmhelpers.GetAllFlags(mobileDeviceFlags)

func init() {
	rootCmd.AddCommand(mobileDevicesCmd)
}

func mapToMobileDeviceAction(flags map[string]*gsmhelpers.Value) (*admin.MobileDeviceAction, error) {
	mobileDeviceAction := &admin.MobileDeviceAction{}
	if flags["action"].IsSet() {
		mobileDeviceAction.Action = flags["action"].GetString()
		if mobileDeviceAction.Action == "" {
			mobileDeviceAction.ForceSendFields = append(mobileDeviceAction.ForceSendFields, "Action")
		}
	}
	return nil, nil
}
