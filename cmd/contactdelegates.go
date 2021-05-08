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

// contactDelegatesCmd represents the contactDelegates command
var contactDelegatesCmd = &cobra.Command{
	Use:               "contactDelegates",
	Short:             "Manage users' contact contact delegations (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/contact-delegation/guides",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var contactDelegateFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"create", "delete", "list"},
		Type:         "string",
		Description:  `The email address of the user whose contacts should be delegated.`,
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
