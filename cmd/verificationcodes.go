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
)

// verificationCodesCmd represents the verificationCodes command
var verificationCodesCmd = &cobra.Command{
	Use:               "verificationCodes",
	Short:             "Manage backup Verification Codes for Users (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/verificationCodes",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var verificationCodeFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userKey": {
		AvailableFor: []string{"generate", "invalidate", "list"},
		Type:         "string",
		Description: `Identifies the user in the API request.
The value can be the user's primary email address, alias email address, or unique user ID.`,
		Required:       []string{"generate", "invalidate", "list"},
		ExcludeFromAll: true,
	},
	"fields": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Recursive: []string{"list"},
	},
}
var verificationCodeFlagsALL = gsmhelpers.GetAllFlags(verificationCodeFlags)

func init() {
	rootCmd.AddCommand(verificationCodesCmd)
}
