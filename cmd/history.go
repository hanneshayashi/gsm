/*
Copyright © 2020-2022 Hannes Hayashi

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

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:               "history",
	Short:             "Manage (list..) user's mailbox History (Part of Gmail API)",
	Long:              "Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users.history",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var historyFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The user's email address.
The special value me can be used to indicate the authenticated user.`,
		Defaults: map[string]any{"list": "me"},
	},
	"startHistoryId": {
		AvailableFor: []string{"list"},
		Type:         "uint64",
		Description: `Required. Returns history records after the specified startHistoryId.
The supplied startHistoryId should be obtained from the historyId of a message, thread, or previous list response.
History IDs increase chronologically but are not contiguous with random gaps in between valid IDs.`,
		Required: []string{"list"},
	},
	"labelId": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `Only return messages with a label matching the ID.`,
	},
	"historyTypes": {
		AvailableFor: []string{"list"},
		Type:         "stringSlice",
		Description: `History types to be returned by the function.
[MESSAGE_ADDED|MESSAGE_DELETED|LABEL_ADDED|LABEL_REMOVED]`,
	},
	"fields": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
