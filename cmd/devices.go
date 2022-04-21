/*
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
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	ci "google.golang.org/api/cloudidentity/v1"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:               "devices",
	Short:             "Manage Devices (Part of Cloud Identity API)",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/devices",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var deviceFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor:   []string{"cancelWipe", "delete", "get", "wipe"},
		Type:           "string",
		Description:    "Resource name of the Device in format: devices/{device_id}, where device_id is the unique ID assigned to the Device.",
		Required:       []string{"cancelWipe", "delete", "get", "wipe"},
		ExcludeFromAll: true,
	},
	"customer": {
		AvailableFor: []string{"cancelWipe", "create", "delete", "get", "list", "wipe"},
		Type:         "string",
		Description: `Resource name of the customer.
If you're using this API for your own organization, use customers/my_customer.
If you're using this API to manage another organization, use customers/{customer_id}, where customer_id is the customer to whom the device belongs.`,
		Defaults: map[string]any{"cancelWipe": "customers/my_customer", "create": "customers/my_customer", "delete": "customers/my_customer", "get": "customers/my_customer", "list": "customers/my_customer", "wipe": "customers/my_customer"},
	},
	"serialNumber": {
		AvailableFor:   []string{"create"},
		Type:           "string",
		Description:    `Serial Number of device. Example: HT82V1A01076.`,
		Required:       []string{"create"},
		ExcludeFromAll: true,
	},
	"assetTag": {
		AvailableFor:   []string{"create"},
		Type:           "string",
		Description:    `Asset tag of the device.`,
		ExcludeFromAll: true,
	},
	"deviceType": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `Type of device:
ANDROID      Device is an Android device
IOS          Device is an iOS device
GOOGLE_SYNC  Device is a Google Sync device.
WINDOWS      Device is a Windows device.
MAC_OS       Device is a MacOS device.
LINUX        Device is a Linux device.
CHROME_OS    Device is a ChromeOS device.`,
		Required: []string{"create"},
	},
	"filter": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Additional restrictions when fetching list of devices.
For a list of search fields, refer to https://developers.google.com/admin-sdk/directory/v1/search-operators.
Multiple search fields are separated by the space character.`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Order specification for devices in the response.
Only one of the following field names may be used to specify the order: createTime, lastSyncTime, model, osVersion, deviceType and serialNumber.
desc may be specified optionally at the end to specify results to be sorted in descending order.
Default order is ascending.`,
	},
	"view": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The view to use for the List request.
Possible values are:
COMPANY_INVENTORY      This view contains all devices imported by the company admin.
                       Each device in the response contains all information specified by the company admin when importing the device (i.e. asset tags).
                       This includes devices that may be unaassigned or assigned to users.
USER_ASSIGNED_DEVICES  This view contains all devices with at least one user registered on the device.
                       Each device in the response contains all device information, except for asset tags.`,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var deviceFlagsALL = gsmhelpers.GetAllFlags(deviceFlags)

func init() {
	rootCmd.AddCommand(devicesCmd)
}

func mapToWipeDeviceRequest(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1WipeDeviceRequest, error) {
	wipeDeviceRequest := &ci.GoogleAppsCloudidentityDevicesV1WipeDeviceRequest{}
	wipeDeviceRequest.Customer = flags["customer"].GetString()
	if wipeDeviceRequest.Customer == "" {
		wipeDeviceRequest.ForceSendFields = append(wipeDeviceRequest.ForceSendFields, "Customer")
	}
	return wipeDeviceRequest, nil
}

func mapToCancelWipeDeviceRequest(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1CancelWipeDeviceRequest, error) {
	cancelWipeRequest := &ci.GoogleAppsCloudidentityDevicesV1CancelWipeDeviceRequest{}
	cancelWipeRequest.Customer = flags["customer"].GetString()
	if cancelWipeRequest.Customer == "" {
		cancelWipeRequest.ForceSendFields = append(cancelWipeRequest.ForceSendFields, "Customer")
	}
	return cancelWipeRequest, nil
}

func mapToDevice(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1Device, error) {
	device := &ci.GoogleAppsCloudidentityDevicesV1Device{}
	if flags["serialNumber"].IsSet() {
		device.SerialNumber = flags["serialNumber"].GetString()
		if device.SerialNumber == "" {
			device.ForceSendFields = append(device.ForceSendFields, "SerialNumber")
		}
	}
	if flags["deviceType"].IsSet() {
		device.DeviceType = flags["deviceType"].GetString()
		if device.DeviceType == "" {
			device.ForceSendFields = append(device.ForceSendFields, "DeviceType")
		}
	}
	return device, nil
}
