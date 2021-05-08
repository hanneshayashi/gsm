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

// aspsCmd represents the asps command
var aspsCmd = &cobra.Command{
	Use:   "asps",
	Short: "Manage ASPs (application-specific password) for a user (Part of Admin SDK)",
	Long: `An application-specific password (ASP) is used with applications
that do not accept a verification code when logging into the application on
certain devices. The ASP access code is used instead of the login and password
you commonly use when accessing an application through a browser. For more
information about ASPs and how to create one, see the
https://http//support.google.com/a/bin/answer.py?amp;answer=1032419.

https://developers.google.com/admin-sdk/directory/v1/reference/asps`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var aspFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"codeId": {
		AvailableFor:   []string{"delete", "get"},
		Type:           "int64",
		Description:    "The unique ID of the ASP",
		Required:       []string{"delete", "get"},
		ExcludeFromAll: true,
	},
	"userKey": {
		AvailableFor: []string{"delete", "get", "list"},
		Type:         "string",
		Description: `Identifies the user in the API request.
The value can be the user's primary email address, alias email address, or unique user ID.`,
		Required: []string{"delete", "get", "list"},
	},
	"fields": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Recursive: []string{"list"},
	},
}
var aspFlagsALL = gsmhelpers.GetAllFlags(aspFlags)

func init() {
	rootCmd.AddCommand(aspsCmd)
}
