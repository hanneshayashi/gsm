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
)

// chromeOsCommandsCmd represents the chromeOsCommands command
var chromeOsCommandsCmd = &cobra.Command{
	Use:               "chromeOsCommands",
	Short:             "Get information about commands issued to Chrome OS Devices (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/reference/rest/v1/customer.devices.chromeos.commands",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var chromeOsCommandFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customerId": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The unique ID for the customer's Workspace account.
As an account administrator, you can also use the my_customer alias to represent your account's customerId.
The customerId is also returned as part of the Users resource.`,
		Defaults: map[string]interface{}{"get": "my_customer"},
	},
	"deviceId": {
		AvailableFor:   []string{"get"},
		Type:           "string",
		Description:    `Immutable ID of Chrome OS Device.`,
		Required:       []string{"get"},
		ExcludeFromAll: true,
	},
	"commandId": {
		AvailableFor: []string{"get"},
		Type:         "int64",
		Description:  `Immutable ID of Chrome OS Device Command.`,
		Required:     []string{"get"},
	},
	"fields": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var chromeOsCommandFlagsALL = gsmhelpers.GetAllFlags(chromeOsDeviceFlags)

func init() {
	rootCmd.AddCommand(chromeOsCmd)
}
