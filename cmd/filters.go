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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// filtersCmd represents the filters command
var filtersCmd = &cobra.Command{
	Use:               "filters",
	Short:             "Manage users' Gmail message filters (Part of Gmail API)",
	Long:              "Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users.settings.filters",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var filterFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userId": {
		AvailableFor: []string{"create", "delete", "get", "list"},
		Type:         "string",
		Description:  "The user's email address. The special value \"me\" can be used to indicate the authenticated user.",
		Defaults:     map[string]any{"create": "me", "delete": "me", "get": "me", "list": "me"},
	},
	"addLabelIds": {
		AvailableFor: []string{"create"},
		Type:         "stringSlice",
		Description:  "A list of IDs of labels to add to this message. Can be used multiple times.",
	},
	"removeLabelIds": {
		AvailableFor: []string{"create"},
		Type:         "stringSlice",
		Description:  "A list of IDs of labels to remove from this message. Can be used multiple times.",
	},
	"forward": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  "Email address that the message should be forwarded to.",
	},
	"from": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  "The sender's display name or email address.",
	},
	"to": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields.
You can use simply the local part of the email address
For example, "example" and "example@" both match "example@gmail.com".
This field is case-insensitive.`,
	},
	"subject": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  "Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent spaces are collapsed.",
	},
	"query": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `Only return messages matching the specified query.
Supports the same query format as the Gmail search box.
For example, "from:someuser@example.com rfc822msgid:<somemsgid@example.com> is:unread".`,
	},
	"negatedQuery": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `Only return messages not matching the specified query.
Supports the same query format as the Gmail search box.
For example, "from:someuser@example.com rfc822msgid:<somemsgid@example.com> is:unread".`,
	},
	"hasAttachment": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description:  "Whether the message has any attachment.",
	},
	"excludeChats": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description:  "Whether the response should exclude chats.",
	},
	"size": {
		AvailableFor: []string{"create"},
		Type:         "int64",
		Description:  "The size of the entire RFC822 message in bytes, including all headers and attachments.",
	},
	"sizeComparison": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `How the message size in bytes should be in relation to the size field.
"[SMALLER|LARGER]
SMALLER  - Find messages smaller than the given size.
LARGER   - Find messages larger than the given size.`,
	},
	"id": {
		AvailableFor:   []string{"delete", "get"},
		Type:           "string",
		Description:    `The ID of the filter.`,
		ExcludeFromAll: true,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var filterFlagsALL = gsmhelpers.GetAllFlags(filterFlags)

func init() {
	rootCmd.AddCommand(filtersCmd)
}

func mapToFilter(flags map[string]*gsmhelpers.Value) (*gmail.Filter, error) {
	filter := &gmail.Filter{}
	if flags["addLabelIds"].IsSet() || flags["removeLabelIds"].IsSet() || flags["forward"].IsSet() {
		filter.Action = &gmail.FilterAction{}
		if flags["addLabelIds"].IsSet() {
			filter.Action.AddLabelIds = flags["addLabelIds"].GetStringSlice()
			if len(filter.Action.AddLabelIds) == 0 {
				filter.Action.ForceSendFields = append(filter.Action.ForceSendFields, "AddLabelIds")
			}
		}
		if flags["removeLabelIds"].IsSet() {
			filter.Action.RemoveLabelIds = flags["removeLabelIds"].GetStringSlice()
			if len(filter.Action.RemoveLabelIds) == 0 {
				filter.Action.ForceSendFields = append(filter.Action.ForceSendFields, "RemoveLabelIds")
			}
		}
		if flags["forward"].IsSet() {
			filter.Action.Forward = flags["forward"].GetString()
			if filter.Action.Forward == "" {
				filter.Action.ForceSendFields = append(filter.Action.ForceSendFields, "Forward")
			}
		}
	}
	if flags["excludeChats"].IsSet() || flags["from"].IsSet() || flags["hasAttachment"].IsSet() || flags["negatedQuery"].IsSet() || flags["query"].IsSet() || flags["size"].IsSet() || flags["sizeComparison"].IsSet() || flags["subject"].IsSet() || flags["to"].IsSet() {
		filter.Criteria = &gmail.FilterCriteria{}
		if flags["excludeChats"].IsSet() {
			filter.Criteria.ExcludeChats = flags["excludeChats"].GetBool()
			if !filter.Criteria.ExcludeChats {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "ExcludeChats")
			}
		}
		if flags["from"].IsSet() {
			filter.Criteria.From = flags["from"].GetString()
			if filter.Criteria.From == "" {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "From")
			}
		}
		if flags["hasAttachment"].IsSet() {
			filter.Criteria.HasAttachment = flags["hasAttachment"].GetBool()
			if !filter.Criteria.HasAttachment {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "HasAttachment")
			}
		}
		if flags["negatedQuery"].IsSet() {
			filter.Criteria.NegatedQuery = flags["negatedQuery"].GetString()
			if filter.Criteria.NegatedQuery == "" {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "NegatedQuery")
			}
		}
		if flags["query"].IsSet() {
			filter.Criteria.Query = flags["query"].GetString()
			if filter.Criteria.Query == "" {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "Query")
			}
		}
		if flags["size"].IsSet() {
			filter.Criteria.Size = flags["size"].GetInt64()
			if filter.Criteria.Size == 0 {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "Size")
			}
		}
		if flags["sizeComparison"].IsSet() {
			sizeComparison := flags["sizeComparison"].GetString()
			if gsmgmail.SizeComparisonIsValid(sizeComparison) {
				filter.Criteria.SizeComparison = sizeComparison
				if filter.Criteria.SizeComparison == "" {
					filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "SizeComparison")
				}
			} else {
				log.Printf("%s is not a valid value for sizeComparison", sizeComparison)
			}
		}
		if flags["subject"].IsSet() {
			filter.Criteria.Subject = flags["subject"].GetString()
			if filter.Criteria.Subject == "" {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "Subject")
			}
		}
		if flags["to"].IsSet() {
			filter.Criteria.To = flags["to"].GetString()
			if filter.Criteria.To == "" {
				filter.Criteria.ForceSendFields = append(filter.Criteria.ForceSendFields, "To")
			}
		}
	}
	return filter, nil
}
