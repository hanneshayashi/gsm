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

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Information about the user, the user's Drive, and system capabilities (Part of Drive API)",
	Long: `This API only works with the currently authenticated user!
Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/about`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var aboutFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fields": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Defaults: map[string]any{"get": "user(displayName, emailAddress, permissionId)"},
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
