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

package cmd

import (
	"log"
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	ci "google.golang.org/api/cloudidentity/v1"

	"github.com/spf13/cobra"
)

// clientStatesCmd represents the clientStates command
var clientStatesCmd = &cobra.Command{
	Use:               "clientStates",
	Short:             "Manage client states (Part of Cloud Identity API)",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers.clientStates",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var clientStateFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"get", "patch"},
		Type:         "string",
		Description: `Resource name of the ClientState in format: devices/{device_id}/deviceUsers/{device_user_id}/clientStates/{partner_id},
where device_id is the unique ID assigned to the Device, device_user_id is the unique ID assigned to the User and partner_id identifies the partner storing the data.
To get the client state for devices belonging to your own organization, the partnerId is in the format: customerId-*anystring*.
Where the customerId is your organization's customer ID and anystring is any suffix.
This suffix is used in setting up Custom Access Levels in Context-Aware Access.
You may use my_customer instead of the customer ID for devices managed by your own organization.
You may specify - in place of the {device_id}, so the ClientState resource name can be: devices/-/deviceUsers/{device_user_resource_id}/clientStates/{partner_id}.`,
		Required:       []string{"get", "patch"},
		ExcludeFromAll: true,
	},
	"customer": {
		AvailableFor: []string{"get", "list", "patch"},
		Type:         "string",
		Description: `Resource name of the customer.
If you're using this API for your own organization, use customers/my_customer.
If you're using this API to manage another organization, use customers/{customer_id}, where customer_id is the customer to whom the device belongs.`,
		Defaults: map[string]any{"get": "customers/my_customer", "list": "customers/my_customer", "patch": "customers/my_customer"},
	},
	"parent": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `To list all ClientStates, set this to "devices/-/deviceUsers/-".
To list all ClientStates owned by a DeviceUser, set this to the resource name of the DeviceUser.
Format: devices/{device}/deviceUsers/{deviceUser}`,
		Required: []string{"list"},
	},
	"filter": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `Additional restrictions when fetching list of client states.`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `Order specification for client states in the response.`,
	},
	"etag": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The token that needs to be passed back for concurrency control in updates.
Token needs to be passed back in UpdateRequest`,
		Required: []string{"patch"},
	},
	"customId": {
		AvailableFor:   []string{"patch"},
		Type:           "string",
		Description:    `This field may be used to store a unique identifier for the API resource within which these CustomAttributes are a field.`,
		Required:       []string{"patch"},
		ExcludeFromAll: true,
	},
	"assetTags": {
		AvailableFor:   []string{"patch"},
		Type:           "stringSlice",
		Description:    `The caller can specify asset tags for this resource`,
		Required:       []string{"patch"},
		ExcludeFromAll: true,
	},
	"keyValuePairs": {
		AvailableFor: []string{"patch"},
		Type:         "stringSlice",
		Description: `The map of key-value attributes stored by callers specific to a device.
The total serialized length of this map may not exceed 10KB.
No limit is placed on the number of attributes in a map.

Can be used multiple times in the format: "key=name;value=wrench;type=string"
where type may be one of the following
string
bool
number`,
		Required: []string{"patch"},
	},
	"fields": {
		AvailableFor: []string{"get", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var clientStateFlagsALL = gsmhelpers.GetAllFlags(clientStateFlags)

func init() {
	rootCmd.AddCommand(clientStatesCmd)
}

func mapToClientState(flags map[string]*gsmhelpers.Value) (*ci.GoogleAppsCloudidentityDevicesV1ClientState, error) {
	clientState := &ci.GoogleAppsCloudidentityDevicesV1ClientState{}
	if flags["etag"].IsSet() {
		clientState.Etag = flags["etag"].GetString()
		if clientState.Etag == "" {
			clientState.ForceSendFields = append(clientState.ForceSendFields, "Etag")
		}
	}
	if flags["customId"].IsSet() {
		clientState.CustomId = flags["customId"].GetString()
		if clientState.CustomId == "" {
			clientState.ForceSendFields = append(clientState.ForceSendFields, "CustomId")
		}
	}
	if flags["assetTags"].IsSet() {
		clientState.AssetTags = flags["assetTags"].GetStringSlice()
		if len(clientState.AssetTags) == 0 {
			clientState.ForceSendFields = append(clientState.ForceSendFields, "AssetTags")
		}
	}
	if flags["keyValuePairs"].IsSet() {
		keyValuePairs := flags["keyValuePairs"].GetStringSlice()
		if len(keyValuePairs) > 0 {
			clientState.KeyValuePairs = make(map[string]ci.GoogleAppsCloudidentityDevicesV1CustomAttributeValue)
			for i := range keyValuePairs {
				m := gsmhelpers.FlagToMap(keyValuePairs[i])
				value := ci.GoogleAppsCloudidentityDevicesV1CustomAttributeValue{}
				switch m["type"] {
				case "string":
					value.StringValue = m["value"]
					if value.StringValue == "" {
						value.ForceSendFields = append(value.ForceSendFields, "StringValue")
					}
				case "bool":
					value.BoolValue, _ = strconv.ParseBool(m["value"])
					if !value.BoolValue {
						value.ForceSendFields = append(value.ForceSendFields, "BoolValue")
					}
				case "number":
					value.NumberValue, _ = strconv.ParseFloat(m["value"], 64)
					if value.NumberValue == 0 {
						value.ForceSendFields = append(value.ForceSendFields, "NumberValue")
					}
				}
				clientState.KeyValuePairs[m["key"]] = value
			}
		} else {
			clientState.ForceSendFields = append(clientState.ForceSendFields, "KeyValuePairs")
		}
	}
	return clientState, nil
}
