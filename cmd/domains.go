/*
Copyright © 2020-2024 Hannes Hayashi

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

// domainsCmd represents the domains command
var domainsCmd = &cobra.Command{
	Use:               "domains",
	Short:             "Manage Domains (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/domains",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var domainFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customer": {
		AvailableFor: []string{"delete", "get", "insert", "list"},
		Type:         "string",
		Description:  "Immutable ID of the Workspace account.",
		Defaults:     map[string]any{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer"},
	},
	"domainName": {
		AvailableFor:   []string{"delete", "get", "insert"},
		Type:           "string",
		Description:    "Name of domain .",
		Required:       []string{"delete", "get", "insert"},
		ExcludeFromAll: true,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var domainFlagsALL = gsmhelpers.GetAllFlags(domainFlags)

func init() {
	rootCmd.AddCommand(domainsCmd)
}

func mapToDomain(flags map[string]*gsmhelpers.Value) (*admin.Domains, error) {
	domain := &admin.Domains{}
	domain.DomainName = flags["domainName"].GetString()
	domain.Kind = "admin#directory#domain"
	return domain, nil
}
