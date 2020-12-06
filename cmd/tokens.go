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
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// tokensCmd represents the tokens command
var tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "Managed OAuth access tokens for users (Part of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/tokens",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var tokenFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"clientId": {
		AvailableFor: []string{"delete", "get"},
		Type:         "string",
		Description:  `The Client ID of the application the token is issued to.`,
		Required:     []string{"delete", "get"},
		Recursive:    []string{"delete"},
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
var tokenFlagsALL = gsmhelpers.GetAllFlags(tokenFlags)

func init() {
	rootCmd.AddCommand(tokensCmd)
}
