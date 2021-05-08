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
	"google.golang.org/api/drive/v3"
)

// repliesCmd represents the replies command
var repliesCmd = &cobra.Command{
	Use:               "replies",
	Short:             "Manage replies to comments (Part of Drive API)",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/replies",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var replyFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fileId": {
		AvailableFor: []string{"create", "delete", "get", "list", "update"},
		Type:         "string",
		Description:  "The ID of the file.",
		Required:     []string{"create", "delete", "get", "list", "update"},
	},
	"commentId": {
		AvailableFor: []string{"create", "delete", "get", "list", "update"},
		Type:         "string",
		Description:  "The ID of the comment.",
		Required:     []string{"create", "delete", "get", "list", "update"},
	},
	"content": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The plain text content of the comment.
This field is used for setting the content, while htmlContent should be displayed.`,
		Required: []string{"update"},
	},
	"action": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The action the reply performed to the parent comment.
[resolve|reopen]`,
	},
	"replyId": {
		AvailableFor:   []string{"delete", "get", "update"},
		Type:           "string",
		Description:    `The ID of the reply.`,
		Required:       []string{"delete", "get", "update"},
		ExcludeFromAll: true,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var replyFlagsALL = gsmhelpers.GetAllFlags(replyFlags)

func init() {
	rootCmd.AddCommand(repliesCmd)
}

func mapToReply(flags map[string]*gsmhelpers.Value) (*drive.Reply, error) {
	reply := &drive.Reply{}
	if flags["content"].IsSet() {
		reply.Content = flags["content"].GetString()
		if reply.Content == "" {
			reply.ForceSendFields = append(reply.ForceSendFields, "Content")
		}
	}
	if flags["action"].IsSet() {
		reply.Action = flags["action"].GetString()
		if reply.Action == "" {
			reply.ForceSendFields = append(reply.ForceSendFields, "Action")
		}
	}
	return reply, nil
}
