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
)

// postmasterDomainsCmd represents the postmasterDomains command
var postmasterDomainsCmd = &cobra.Command{
	Use:               "postmasterDomains",
	Short:             "Use Gmail Postmaster Tools to manage domain (Part of Gmail Postmaster API)",
	Long:              `You need to set up your domain(s) at https://postmaster.google.com/u/1/managedomains first.
https://developers.google.com/gmail/postmaster/reference/rest/v1/domains`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var postmasterDomainFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  `Fully qualified domain name.`,
		Required:     []string{"get"},
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var postmasterDomainFlagsALL = gsmhelpers.GetAllFlags(postmasterDomainFlags)

func init() {
	rootCmd.AddCommand(postmasterDomainsCmd)
}
