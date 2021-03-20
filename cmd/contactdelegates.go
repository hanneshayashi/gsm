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

// contactDelegatesCmd represents the contactDelegates command
var contactDelegatesCmd = &cobra.Command{
	Use:               "contactDelegates",
	Short:             "Manage users' contact contact delegations (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/contact-delegation/guides",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var contactDelegateFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"create", "delete", "list"},
		Type:         "string",
		Description:  `The parent resource that will own the created delegate.`,
		Required:     []string{"create", "delete", "list"},
	},
	"email": {
		AvailableFor: []string{"create", "delete"},
		Type:         "string",
		Description:  `Email of the delegate.`,
		Required:     []string{"create", "delete"},
	},
}

func init() {
	rootCmd.AddCommand(contactDelegatesCmd)
}
