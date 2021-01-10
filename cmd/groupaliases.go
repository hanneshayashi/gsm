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
	admin "google.golang.org/api/admin/directory/v1"
)

// groupAliasesCmd represents the groupAliases command
var groupAliasesCmd = &cobra.Command{
	Use:               "groupAliases",
	Short:             "Manage group aliases, which are alternative email addresses (Part of Admin SDK - not Gmail API!)",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/groups/aliases",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var groupAliasFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"alias": {
		AvailableFor: []string{"delete", "insert"},
		Type:         "string",
		Description:  `The alias.`,
		Required:     []string{"delete", "insert"},
	},
	"groupKey": {
		AvailableFor: []string{"delete", "insert", "list"},
		Type:         "string",
		Description: `Identifies the group in the API request.
The value can be the group's email address, group alias, or the unique group ID.`,
		Required:       []string{"delete", "insert", "list"},
		ExcludeFromAll: true,
	},
	"fields": {
		AvailableFor: []string{"insert", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var groupAliasFlagsALL = gsmhelpers.GetAllFlags(groupAliasFlags)

func init() {
	rootCmd.AddCommand(groupAliasesCmd)
}

func mapToGroupAlias(flags map[string]*gsmhelpers.Value) (*admin.Alias, error) {
	alias := &admin.Alias{}
	if flags["alias"].IsSet() {
		alias.Alias = flags["alias"].GetString()
		if alias.Alias == "" {
			alias.ForceSendFields = append(alias.ForceSendFields, "Alias")
		}
	}
	return alias, nil
}
