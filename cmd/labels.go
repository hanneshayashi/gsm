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
	"google.golang.org/api/gmail/v1"
)

// labelsCmd represents the labels command
var labelsCmd = &cobra.Command{
	Use:               "labels",
	Short:             "Manage users' mailbox labels (Part of Gmail API)",
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.labels",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var labelFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"create", "delete", "get", "list", "patch"},
		Type:         "string",
		Description:  "The user's email address. The special value \"me\" can be used to indicate the authenticated user.",
		Defaults:     map[string]interface{}{"create": "me", "delete": "me", "get": "me", "list": "me", "patch": "me"},
	},
	"id": {
		AvailableFor:   []string{"delete", "get", "patch"},
		Type:           "string",
		Description:    "The ID of the label",
		Required:       []string{"delete", "get", "patch"},
		ExcludeFromAll: true,
	},
	"name": {
		AvailableFor:   []string{"create", "patch"},
		Type:           "string",
		Description:    "The display name of the label.",
		Required:       []string{"create"},
		ExcludeFromAll: true,
	},
	"messageListVisibility": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "The visibility of messages with this label in the message list in the Gmail web interface. [SHOW|HIDE]",
		Defaults:     map[string]interface{}{"create": "SHOW"},
	},
	"labelListVisibility": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "The visibility of the label in the label list in the Gmail web interface. [LABEL_SHOW|LABEL_SHOW_IF_UNREAD|LABEL_HIDE]",
		Defaults:     map[string]interface{}{"create": "LABEL_SHOW"},
	},
	"textColor": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "Text color",
	},
	"backgroundColor": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  "Background color",
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var labelFlagsALL = gsmhelpers.GetAllFlags(labelFlags)

func init() {
	rootCmd.AddCommand(labelsCmd)
}

func mapToLabel(flags map[string]*gsmhelpers.Value) (*gmail.Label, error) {
	label := &gmail.Label{}
	if flags["name"].IsSet() {
		label.Name = flags["name"].GetString()
		if label.Name == "" {
			label.ForceSendFields = append(label.ForceSendFields, "Name")
		}
	}
	if flags["messageListVisibility"].IsSet() {
		label.MessageListVisibility = flags["messageListVisibility"].GetString()
		if label.MessageListVisibility == "" {
			label.ForceSendFields = append(label.ForceSendFields, "MessageListVisibility")
		}
	}
	if flags["labelListVisibility"].IsSet() {
		label.LabelListVisibility = flags["labelListVisibility"].GetString()
		if label.LabelListVisibility == "" {
			label.ForceSendFields = append(label.ForceSendFields, "LabelListVisibility")
		}
	}
	if flags["textColor"].IsSet() || flags["backgroundColor"].IsSet() {
		label.Color = &gmail.LabelColor{}
		if flags["textColor"].IsSet() {
			label.Color.TextColor = flags["textColor"].GetString()
			if label.Color.TextColor == "" {
				label.Color.ForceSendFields = append(label.Color.ForceSendFields, "TextColor")
			}
		}
		if flags["backgroundColor"].IsSet() {
			label.Color.BackgroundColor = flags["backgroundColor"].GetString()
			if label.Color.BackgroundColor == "" {
				label.Color.ForceSendFields = append(label.Color.ForceSendFields, "BackgroundColor")
			}
		}
	}
	return label, nil
}
