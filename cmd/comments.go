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
	"google.golang.org/api/drive/v3"
)

// commentsCmd represents the comments command
var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "Manage comments in Google files (Part of Drive API)",
	Long:  "https://developers.google.com/drive/api/v3/reference/comments",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var commentFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fileId": {
		AvailableFor: []string{"create", "delete", "get", "list", "update"},
		Type:         "string",
		Description:  `The ID of the file.`,
		Required:     []string{"create", "delete", "get", "list", "update"},
	},
	"content": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The plain text content of the comment.
This field is used for setting the content, while htmlContent should be displayed.`,
		Required: []string{"create", "update"},
	},
	"anchor": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `A region of the document represented as a JSON string.
See anchor documentation for details on how to define and interpret anchor properties.`,
	},
	"quotedFileContentValue": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The quoted content itself.
This is interpreted as plain text if set through the API.`,
	},
	"commentId": {
		AvailableFor: []string{"delete", "get", "update"},
		Type:         "string",
		Description:  `The ID of the comment.`,
		Required:     []string{"delete", "get", "update"},
	},
	"includeDeleted": {
		AvailableFor: []string{"get", "list"},
		Type:         "bool",
		Description: `Whether to return deleted comments.
Deleted comments will not include their original content.`,
	},
	"startModifiedTime": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `The minimum value of 'modifiedTime' for the result comments (RFC 3339 date-time).`,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Required: []string{"create", "get", "list", "update"},
	},
}
var commentFlagsALL = gsmhelpers.GetAllFlags(commentFlags)

func init() {
	rootCmd.AddCommand(commentsCmd)
}

func mapToComment(flags map[string]*gsmhelpers.Value) (*drive.Comment, error) {
	comment := &drive.Comment{}
	if flags["content"].IsSet() {
		comment.Content = flags["content"].GetString()
		if comment.Content == "" {
			comment.ForceSendFields = append(comment.ForceSendFields, "content")
		}
	}
	if flags["anchor"].IsSet() {
		comment.Anchor = flags["anchor"].GetString()
		if comment.Anchor == "" {
			comment.ForceSendFields = append(comment.ForceSendFields, "anchor")
		}
	}
	if flags["quotedFileContentValue"].IsSet() {
		comment.QuotedFileContent = &drive.CommentQuotedFileContent{}
		comment.QuotedFileContent.Value = flags["quotedFileContentValue"].GetString()
		if comment.QuotedFileContent.Value == "" {
			comment.QuotedFileContent.ForceSendFields = append(comment.QuotedFileContent.ForceSendFields, "Value")
		}
	}
	return comment, nil
}
