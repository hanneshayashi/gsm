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

package cmd

import (
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// chromeOsCmd represents the chromeOs command
var chromeOsCmd = &cobra.Command{
	Use:               "chromeOs",
	Short:             "Issue Commands to Chrome OS Devices (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/customer.devices.chromeos",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var chromeOsFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customerId": {
		AvailableFor: []string{"issueCommand"},
		Type:         "string",
		Description: `The unique ID for the customer's Workspace account.
As an account administrator, you can also use the my_customer alias to represent your account's customerId.
The customerId is also returned as part of the Users resource.`,
		Defaults: map[string]any{"issueCommand": "my_customer"},
	},
	"deviceId": {
		AvailableFor:   []string{"issueCommand"},
		Type:           "string",
		Description:    `Immutable ID of Chrome OS Device.`,
		Required:       []string{"issueCommand"},
		ExcludeFromAll: true,
	},
	"commandType": {
		AvailableFor: []string{"issueCommand"},
		Type:         "string",
		Description: `The type of command.

Acceptable values are:
REBOOT             - Reboot the device.
                     Can only be issued to Kiosk and managed guest session devices.
TAKE_A_SCREENSHOT  - Take a screenshot of the device.
                     Only available if the device is in Kiosk Mode.
SET_VOLUME         - Set the volume of the device.
                     Can only be issued to Kiosk and managed guest session devices.
WIPE_USERS         - Wipe all the users off of the device.
                     Executing this command in the device will remove all user profile data, but it will keep device policy and enrollment.
REMOTE_POWERWASH   - Wipes the device by performing a power wash.
					 Executing this command in the device will remove all data including user policies, device policies and enrollment policies.
					 Warning: This will revert the device back to a factory state with no enrollment unless the device is subject to forced or auto enrollment.
					 Use with caution, as this is an irreversible action!`,
		Required: []string{"issueCommand"},
	},
	"payload": {
		AvailableFor: []string{"issueCommand"},
		Type:         "string",
		Description:  `The payload for the command, provide it only if command supports it. The following commands support adding payload: - SET_VOLUME: Payload is a stringified JSON object in the form: { "volume": 50 }. The volume has to be an integer in the range [0,100].`,
	},
}
var chromeOsFlagsALL = gsmhelpers.GetAllFlags(chromeOsDeviceFlags)

func init() {
	rootCmd.AddCommand(chromeOsCmd)
}

func mapToDirectoryChromeosdevicesIssueCommandRequest(flags map[string]*gsmhelpers.Value) (*admin.DirectoryChromeosdevicesIssueCommandRequest, error) {
	directoryChromeosdevicesIssueCommandRequest := &admin.DirectoryChromeosdevicesIssueCommandRequest{}
	if flags["commandType"].IsSet() {
		directoryChromeosdevicesIssueCommandRequest.CommandType = flags["commandType"].GetString()
		if directoryChromeosdevicesIssueCommandRequest.CommandType == "" {
			directoryChromeosdevicesIssueCommandRequest.ForceSendFields = append(directoryChromeosdevicesIssueCommandRequest.ForceSendFields, "CommandType")
		}
	}
	if flags["payload"].IsSet() {
		directoryChromeosdevicesIssueCommandRequest.Payload = flags["payload"].GetString()
		if directoryChromeosdevicesIssueCommandRequest.Payload == "" {
			directoryChromeosdevicesIssueCommandRequest.ForceSendFields = append(directoryChromeosdevicesIssueCommandRequest.ForceSendFields, "Payload")
		}
	}
	return directoryChromeosdevicesIssueCommandRequest, nil
}
