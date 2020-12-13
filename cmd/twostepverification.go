/*
Package cmd contains the commands available to the end user
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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

var twoStepVerificationFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userKey": {
		AvailableFor: []string{"turnOff"},
		Type:         "string",
		Description: `Identifies the user in the API request.
The value can be the user's primary email address, alias email address, or unique user ID.`,
		Required:       []string{"turnOff"},
		ExcludeFromAll: true,
	},
}
var twoStepVerificationFlagsALL = gsmhelpers.GetAllFlags(userFlags)

// usersCmd represents the users command
var twoStepVerificationCmd = &cobra.Command{
	Use:               "twoStepVerification",
	Short:             "Manage Two Step Verification for users (Park of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/reference/rest/v1/twoStepVerification/turnOff",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(twoStepVerificationCmd)
}
