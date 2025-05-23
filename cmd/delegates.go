/*
Copyright © 2020-2023 Hannes Hayashi

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
	"google.golang.org/api/gmail/v1"
)

// delegatesCmd represents the delegates command
var delegatesCmd = &cobra.Command{
	Use:               "delegates",
	Short:             "Manage Gmail Delegates (Part of Gmail API)",
	Long:              "Implements the API documented at https://developers.google.com/workspace/gmail/api/reference/rest/v1/users.settings.delegates",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var delegateFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"create", "delete", "get", "list"},
		Type:         "string",
		Description:  "The user's email address. The special value me can be used to indicate the authenticated user.",
		Defaults:     map[string]any{"create": "me", "delete": "me", "get": "me", "list": "me"},
	},
	"delegateEmail": {
		AvailableFor:   []string{"create", "delete", "get"},
		Type:           "string",
		Description:    "The email address of the delegate.",
		Required:       []string{"create", "delete", "get"},
		ExcludeFromAll: true,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var delegateFlagsALL = gsmhelpers.GetAllFlags(delegateFlags)

func init() {
	rootCmd.AddCommand(delegatesCmd)
}

func mapToDelegate(flags map[string]*gsmhelpers.Value) (*gmail.Delegate, error) {
	delegate := &gmail.Delegate{}
	delegate.DelegateEmail = flags["delegateEmail"].GetString()
	return delegate, nil
}
