/*
Copyright © 2020-2025 Hannes Hayashi

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

// userAliasesCmd represents the userAliases command
var userAliasesCmd = &cobra.Command{
	Use:               "userAliases",
	Short:             "Manage user aliases, which are alternative email addresses (Part of Admin SDK - not Gmail API!)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/users.aliases",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var userAliasFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"alias": {
		AvailableFor:   []string{"delete", "insert"},
		Type:           "string",
		Description:    `The alias email address.`,
		Required:       []string{"delete", "insert"},
		ExcludeFromAll: true,
	},
	"userKey": {
		AvailableFor: []string{"delete", "insert", "list"},
		Type:         "string",
		Description: `Identifies the user in the API request.
The value can be the user's primary email address, alias email address, or unique user ID.`,
		Required: []string{"delete", "insert", "list"},
	},
	"fields": {
		AvailableFor: []string{"insert", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Recursive: []string{"list"},
	},
}
var userAliasFlagsALL = gsmhelpers.GetAllFlags(userAliasFlags)

func init() {
	rootCmd.AddCommand(userAliasesCmd)
}

func mapToUserAlias(flags map[string]*gsmhelpers.Value) (*admin.Alias, error) {
	alias := &admin.Alias{}
	if flags["alias"].IsSet() {
		alias.Alias = flags["alias"].GetString()
		if alias.Alias == "" {
			alias.ForceSendFields = append(alias.ForceSendFields, "Alias")
		}
	}
	return alias, nil
}
