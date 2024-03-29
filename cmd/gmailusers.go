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

// gmailUsersCmd represents the gmailUsers command
var gmailUsersCmd = &cobra.Command{
	Use:               "gmailUsers",
	Short:             "Gmail User Profiles (Part of Gmail API)",
	Long:              "Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var gmailUserFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"getProfile"},
		Type:         "string",
		Description:  "The user's email address. The special value \"me\" can be used to indicate the authenticated user.",
		Defaults:     map[string]any{"getProfile": "me"},
	},
	"fields": {
		AvailableFor: []string{"getProfile"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(gmailUsersCmd)
}
