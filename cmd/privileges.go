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
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// privilegesCmd represents the privileges command
var privilegesCmd = &cobra.Command{
	Use:               "privileges",
	Short:             "Manage (list) Privileges (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/privileges",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var privilegeFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customer": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  "Immutable ID of the Workspace account.",
		Defaults:     map[string]interface{}{"list": "my_customer"},
	},
	"fields": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(privilegesCmd)
}
