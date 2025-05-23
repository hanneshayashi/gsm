/*
Copyright © 2020-2023 Hannes Hayashi

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
)

// driveLabelUsersCmd represents the driveLabelUsers command
var driveLabelUsersCmd = &cobra.Command{
	Use:               "driveLabelUsers",
	Short:             "Manages Drive Label Users (Part of Drive Labels API)",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/labels/reference/rest/v2/users",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var driveLabelUserFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"getCapabilities"},
		Type:         "string",
		Description: `The resource name of the user.
Only "users/me/capabilities" is supported. (Default)`,
		Defaults: map[string]any{"getCapabilities": "users/me/capabilities"},
	},
	"customer": {
		AvailableFor: []string{"getCapabilities"},
		Type:         "string",
		Description: `The customer to scope this request to. For example: "customers/abcd1234".
If unset, will return settings within the current customer.`,
	},
	"fields": {
		AvailableFor: []string{"getCapabilities"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var labelUserFlagsALL = gsmhelpers.GetAllFlags(driveLabelUserFlags)

func init() {
	rootCmd.AddCommand(driveLabelUsersCmd)
}
