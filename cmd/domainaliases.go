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
	admin "google.golang.org/api/admin/directory/v1"
)

// domainAliasesCmd represents the domainAliases command
var domainAliasesCmd = &cobra.Command{
	Use:               "domainAliases",
	Short:             "Manage Domain Aliases (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/domainAliases",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var domainAliasFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customer": {
		AvailableFor: []string{"delete", "get", "insert", "list"},
		Type:         "string",
		Description:  "Immutable ID of the Workspace account.",
		Defaults:     map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer"},
	},
	"domainAliasName": {
		AvailableFor:   []string{"delete", "get", "insert"},
		Type:           "string",
		Description:    "Name of domain alias.",
		Required:       []string{"delete", "get", "insert"},
		ExcludeFromAll: true,
	},
	"parentDomainName": {
		AvailableFor: []string{"insert", "list"},
		Type:         "string",
		Description:  "Name of domain alias.",
		Required:     []string{"insert"},
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var domainAliasFlagsALL = gsmhelpers.GetAllFlags(domainAliasFlags)

func init() {
	rootCmd.AddCommand(domainAliasesCmd)
}

func mapToDomainAliases(flags map[string]*gsmhelpers.Value) (*admin.DomainAlias, error) {
	domainAlias := &admin.DomainAlias{}
	domainAlias.DomainAliasName = flags["domainAliasName"].GetString()
	domainAlias.ParentDomainName = flags["parentDomainName"].GetString()
	domainAlias.Kind = "admin#directory#domainAlias"
	return domainAlias, nil
}
