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
	"google.golang.org/api/gmail/v1"
)

// delegatesCmd represents the delegates command
var delegatesCmd = &cobra.Command{
	Use:   "delegates",
	Short: "Manage Gmail Delegates (Part of Gmail API)",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.settings.delegates",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var delegateFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"create", "delete", "get", "list"},
		Type:         "string",
		Description:  "The user's email address. The special value me can be used to indicate the authenticated user.",
		Defaults:     map[string]interface{}{"create": "me", "delete": "me", "get": "me", "list": "me"},
	},
	"delegateEmail": {
		AvailableFor: []string{"create", "delete", "get"},
		Type:         "string",
		Description:  "The email address of the delegate.",
		Required:     []string{"create", "delete", "get"},
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
