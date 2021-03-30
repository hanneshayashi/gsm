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
	"errors"
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// chromeOsDevicesCmd represents the chromeOsDevices command
var chromeOsDevicesCmd = &cobra.Command{
	Use:               "chromeOsDevices",
	Short:             "Managed Chrome OS Devices (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/chromeosdevices",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var chromeOsDeviceFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customerId": {
		AvailableFor: []string{"action", "get", "list", "moveToOU", "patch"},
		Type:         "string",
		Description: `The unique ID for the customer's Workspace account.
As an account administrator, you can also use the my_customer alias to represent your account's customerId.
The customerId is also returned as part of the Users resource.`,
		Defaults: map[string]interface{}{"action": "my_customer", "get": "my_customer", "list": "my_customer", "moveToOU": "my_customer", "patch": "my_customer"},
		Required: []string{"moveToOU"},
	},
	"resourceId": {
		AvailableFor: []string{"action"},
		Type:         "string",
		Description: `The unique ID of the device.
The resourceIds are returned in the response from the chromeosdevices.list method.`,
		Required:       []string{"action"},
		ExcludeFromAll: true,
	},
	"action": {
		AvailableFor: []string{"action"},
		Type:         "string",
		Description: `Action to be taken on the Chrome OS device

Acceptable values are:
deprovision  - Remove a device from management that is no longer active, being resold, or is being submitted for return / repair, use the deprovision action to dissociate it from management.
disable      - If you believe a device in your organization has been lost or stolen, you can disable the device so that no one else can use it.
               When a device is disabled, all the user can see when turning on the Chrome device is a screen telling them that it has been disabled, and your desired contact information of where to return the device.
               Note: Configuration of the message to appear on a disabled device must be completed within the admin console.
 reenable    - Re-enable a disabled device when a misplaced device is found or a lost device is returned. You can also use this feature if you accidentally mark a Chrome device as disabled.
               Note: The re-enable action can only be performed on devices marked as disabled.`,
		Required: []string{"action"},
	},
	"deprovisionReason": {
		AvailableFor: []string{"action"},
		Type:         "string",
		Description: `Only used when the action is deprovision. With the deprovision action, this field is required.

Note: The deprovision reason is audited because it might have implications on licenses for perpetual subscription customers.

Acceptable values are:
different_model_replacement  - Use if you're upgrading or replacing your device with a newer model of the same device.
retiring_device              - Use if you're reselling, donating, or permanently removing the device from use.
same_model_replacement       - Use if a hardware issue was encountered on a device and it is being replaced with the same model or a like-model replacement from a repair vendor / manufacturer.
upgrade_transfer             - Use if you're replacing your Cloud Ready devices with Chromebooks within one year.`,
		Required: []string{"action"},
	},
	"deviceId": {
		AvailableFor: []string{"get", "patch"},
		Type:         "string",
		Description: `The unique ID of the device.
The deviceIds are returned in the response from the chromeosdevices.list method.`,
		Defaults:       map[string]interface{}{"action": "my_customer", "get": "my_customer", "list": "my_customer", "moveToOU": "my_customer", "patch": "my_customer"},
		Required:       []string{"get", "patch"},
		ExcludeFromAll: true,
	},
	"projection": {
		AvailableFor: []string{"get", "list", "patch"},
		Type:         "string",
		Description: `Determines whether the response contains the full list of properties or only a subset.

Acceptable values are:
BASIC  - Excludes the model, meid, orderNumber, willAutoRenew, osVersion, platformVersion, firmwareVersion, macAddress, and bootMode properties.
FULL   - Includes all metadata fields.`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Device property to use for sorting results.

Acceptable values are:
annotatedLocation  - Chrome device location as annotated by the administrator.
annotatedUser      - Chrome device user as annotated by the administrator.
lastSync           - The date and time the Chrome device was last synchronized with the policy settings in the Admin console.
notes              - Chrome device notes as annotated by the administrator.
serialNumber       - The Chrome device serial number entered when the device was enabled.
status             - Chrome device status. For more information, see the chromeosdevices resource.
supportEndDate     - Chrome device support end date. This is applicable only for devices purchased directly from Google.`,
	},
	"orgUnitPath": {
		AvailableFor: []string{"list", "moveToOU", "patch"},
		Type:         "string",
		Description:  `The full path of the organizational unit or its unique ID.`,
		Required:     []string{"moveToOU"},
	},
	"query": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Search string in the format provided by List query operators
(https://developers.google.com/admin-sdk/directory/v1/list-query-operators).`,
	},
	"sortOrder": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Whether to return results in ascending or descending order. Must be used with the orderBy parameter.

Acceptable values are:
ASCENDING   - Ascending order.
DESCENDING  - Descending order.`,
	},
	"deviceIds": {
		AvailableFor: []string{"moveToOU"},
		Type:         "stringSlice",
		Description:  `Chrome OS devices to be moved to OU`,
		Required:     []string{"moveToOU"},
	},
	"annotatedAssetId": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `The asset identifier as noted by an administrator or specified during enrollment.`,
	},
	"annotatedLocation": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The address or location of the device as noted by the administrator.
Maximum length is 200 characters. Empty values are allowed.`,
	},
	"annotatedUser": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The user of the device as noted by the administrator.
Maximum length is 100 characters. Empty values are allowed.`,
	},
	"notes": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Notes about this device added by the administrator.
This property can be searched with the list method's query parameter.
Maximum length is 500 characters. Empty values are allowed.`,
	},
	"fields": {
		AvailableFor: []string{"get", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var chromeOsDeviceFlagsALL = gsmhelpers.GetAllFlags(chromeOsDeviceFlags)

func init() {
	rootCmd.AddCommand(chromeOsDevicesCmd)
}

func mapToChromeOsDeviceAction(flags map[string]*gsmhelpers.Value) (*admin.ChromeOsDeviceAction, error) {
	action := flags["action"].GetString()
	deprovisionReason := flags["deprovisionReason"].GetString()
	flags["action"].GetString()
	if action == "deprovision" && deprovisionReason == "" {
		return nil, errors.New("reason must be specified with --deprovisionReason when deprovisioning a Chrome OS device")
	}
	if action != "deprovision" && deprovisionReason != "" {
		return nil, errors.New("--deprovisionReason may only be used when action is \"deprovision\"")
	}
	chromeOsDeviceAction := &admin.ChromeOsDeviceAction{
		Action:            action,
		DeprovisionReason: deprovisionReason,
	}
	return chromeOsDeviceAction, nil
}

func mapToChromeOsMoveDevicesToOu(flags map[string]*gsmhelpers.Value) (*admin.ChromeOsMoveDevicesToOu, error) {
	chromeOsMoveDevicesToOu := &admin.ChromeOsMoveDevicesToOu{}
	chromeOsMoveDevicesToOu.DeviceIds = flags["deviceId"].GetStringSlice()
	return chromeOsMoveDevicesToOu, nil
}

func mapToChromeOsDevice(flags map[string]*gsmhelpers.Value) (*admin.ChromeOsDevice, error) {
	chromeOsDevice := &admin.ChromeOsDevice{}
	if flags["annotatedUser"].IsSet() {
		chromeOsDevice.AnnotatedUser = flags["annotatedUser"].GetString()
		if chromeOsDevice.AnnotatedUser == "" {
			chromeOsDevice.ForceSendFields = append(chromeOsDevice.ForceSendFields, "AnnotatedUser")
		}
	}
	if flags["annotatedLocation"].IsSet() {
		chromeOsDevice.AnnotatedLocation = flags["annotatedLocation"].GetString()
		if chromeOsDevice.AnnotatedLocation == "" {
			chromeOsDevice.ForceSendFields = append(chromeOsDevice.ForceSendFields, "AnnotatedLocation")
		}
	}
	if flags["notes"].IsSet() {
		chromeOsDevice.Notes = flags["notes"].GetString()
		if chromeOsDevice.Notes == "" {
			chromeOsDevice.ForceSendFields = append(chromeOsDevice.ForceSendFields, "Notes")
		}
	}
	if flags["annotatedAssetID"].IsSet() {
		chromeOsDevice.AnnotatedAssetId = flags["annotatedAssetID"].GetString()
		if chromeOsDevice.AnnotatedAssetId == "" {
			chromeOsDevice.ForceSendFields = append(chromeOsDevice.ForceSendFields, "AnnotatedAssetId")
		}
	}
	if flags["orgUnitPath"].IsSet() {
		chromeOsDevice.OrgUnitPath = flags["orgUnitPath"].GetString()
		if chromeOsDevice.OrgUnitPath == "" {
			chromeOsDevice.ForceSendFields = append(chromeOsDevice.ForceSendFields, "OrgUnitPath")
		}
	}
	return chromeOsDevice, nil
}
