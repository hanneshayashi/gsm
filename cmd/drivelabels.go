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
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/drivelabels/v2"

	"github.com/spf13/cobra"
)

// driveLabelsCmd represents the files command
var driveLabelsCmd = &cobra.Command{
	Use:               "driveLabels",
	Short:             "Managed driveLabels (Part of Drive Labels API)",
	Long:              "Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var driveLabelFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Label resource name.
May be any of:
  - labels/{id} (equivalent to labels/{id}@latest)
  - labels/{id}@latest
  - labels/{id}@published
  - labels/{id}@{revisionId}
If you don't specify the "labels/" prefix, GSM will automatically prepend it to the request.`,
		Required:       []string{"get"},
		ExcludeFromAll: true,
	},
	"useAdminAccess": {
		AvailableFor: []string{"get", "list", "create"},
		Type:         "bool",
		Description: `Set to true in order to use the user's admin credentials.
The server verifies that the user is an admin for the label before allowing access.`,
	},
	"languageCode": {
		AvailableFor: []string{"get", "list", "create"},
		Type:         "string",
		Description: `The BCP-47 language code to use for evaluating localized field labels.
When not specified, values in the default configured language are used.`,
	},
	"view": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `When specified, only certain fields belonging to the indicated view are returned.
[LABEL_VIEW_BASIC|LABEL_VIEW_FULL]
LABEL_VIEW_BASIC - Implies the field mask: name,id,revisionId,labelType,properties.*
LABEL_VIEW_FULL  - All possible fields.`,
	},
	"publishedOnly": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description: `Whether to include only published labels in the results.

When true, only the current published label revisions are returned.
Disabled labels are included.
Returned label resource names reference the published revision (labels/{id}/{revisionId}).

When false, the current label revisions are returned, which might not be published.
Returned label resource names don't reference a specific revision (labels/{id}).`,
	},
	"minimumRole": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Specifies the level of access the user must have on the returned Labels.
The minimum role a user must have on a label.
Defaults to READER.
[READER|APPLIER|ORGANIZER|EDITOR]
READER     - A reader can read the label and associated metadata applied to Drive items.
APPLIER    - An applier can write associated metadata on Drive items in which they also have write access to. Implies READER.
ORGANIZER  - An organizer can pin this label in shared drives they manage and add new appliers to the label.
EDITOR     - Editors can make any update including deleting the label which also deletes the associated Drive item metadata. Implies APPLIER.`,
	},
	"labelType": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `The type of this label.
Defaults to SHARED.
[SHARED|ADMIN]
SHARED  - Shared labels may be shared with users to apply to Drive items.
ADMIN   - Admin-owned label. Only creatable and editable by admins. Supports some additional admin-only features.`,
		Defaults: map[string]any{"create": "SHARED"},
	},
	"title": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `Title of the label.`,
		Required:     []string{"create"},
	},
	"description": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `The description of the label.`,
		Required:     []string{"create"},
	},
	"field": {
		AvailableFor: []string{"create"},
		Type:         "stringSlice",
		Description: `Defines a field that has a display name, data type, and other configuration options.
This field defines the kind of metadata that may be set on a Drive item.

Can be used multiple times in the format: "displayName=...;valueType=...;[required=[true|false];insertBeforeField=...;minLength=]"
The following options are available:
- displayName        - (string) Required. The display text to show in the UI identifying this field.
- required           - (bool) Whether the field should be marked as required.
- insertBeforeField  - (string) Input only. Insert or move this field before the indicated field. If empty, the field is placed at the end of the list.
- valueType          - The type of the field
                       Where valueType may be one of the following, with different options available for each type:
                         - dateString    - A date field
                           - maxValue    - The maximum value for the field in the format "YYYY/MM/DD"
                           - minValue    - The minimum value for the field in the format "YYYY/MM/DD"
                         - integer       - An integer field
                           - maxValue    - The maximum value for the field as a whole number
                           - minValue    - The minimum value for the field as a whole number
                         - selection     - A field that allows the user to select on or more choices from a predefined set
                           - maxEntries  - The maximum number of entries for the field as a whole number
                           - choices     - A list of choices in the format "choices=choice1|choice2|choice3"
                         - text          - A text field
                           - maxLength   - The maximum length of the text as a whole number
                           - minLength   - The minimum length of the text as a whole number
                         - user          - A field that allows the selection of one or more user(s)
                           - maxEntries  - The maximum number of entries for the field as a whole number`,
	},
	"learnMoreUri": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `Custom URL to present to users to allow them to learn more about this label and how it should be used.`,
	},
	"fields": {
		AvailableFor: []string{"get", "list", "create"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(driveLabelsCmd)
}

func stringToLabelDate(s string) (*drivelabels.GoogleTypeDate, error) {
	date, errDate := time.Parse("2006/01/02", s)
	if errDate != nil {
		return nil, errDate
	}
	gDate := &drivelabels.GoogleTypeDate{
		Year:  int64(date.Year()),
		Month: int64(date.Month()),
		Day:   int64(date.Day()),
	}
	return gDate, nil
}

func mapToDriveLabel(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
	label := &drivelabels.GoogleAppsDriveLabelsV2Label{}
	var err error
	label.LabelType = flags["labelType"].GetString()
	if label.LabelType == "" {
		label.ForceSendFields = append(label.ForceSendFields, "LabelType")
	}
	titleSet := flags["title"].IsSet()
	descriptionSet := flags["description"].IsSet()
	if titleSet || descriptionSet {
		label.Properties = &drivelabels.GoogleAppsDriveLabelsV2LabelProperties{}
		if titleSet {
			label.Properties.Title = flags["title"].GetString()
			if label.Properties.Title == "" {
				label.Properties.ForceSendFields = append(label.Properties.ForceSendFields, "Title")
			}
		}
		if descriptionSet {
			label.Properties.Description = flags["description"].GetString()
			if label.Properties.Description == "" {
				label.Properties.ForceSendFields = append(label.Properties.ForceSendFields, "Description")
			}
		}
	}
	if flags["learnMoreUri"].IsSet() {
		label.LearnMoreUri = flags["learnMoreUri"].GetString()
		if label.LearnMoreUri == "" {
			label.ForceSendFields = append(label.ForceSendFields, "LearnMoreUri")
		}
	}
	if flags["field"].IsSet() {
		fields := flags["field"].GetStringSlice()
		if len(fields) > 0 {
			label.Fields = make([]*drivelabels.GoogleAppsDriveLabelsV2Field, len(fields))
			for i := range fields {
				m := gsmhelpers.FlagToMap(fields[i])
				label.Fields[i] = &drivelabels.GoogleAppsDriveLabelsV2Field{}
				label.Fields[i].Properties = &drivelabels.GoogleAppsDriveLabelsV2FieldProperties{
					DisplayName:       m["displayName"],
					InsertBeforeField: m["insertBeforeField"],
				}
				if m["required"] != "" {
					label.Fields[i].Properties.Required, err = strconv.ParseBool(m["required"])
					if err != nil {
						return nil, err
					}
				}
				switch m["valueType"] {
				case "textString":
					label.Fields[i].TextOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldTextOptions{}
					if m["maxLength"] != "" {
						label.Fields[i].TextOptions.MaxLength, err = strconv.ParseInt(m["maxLength"], 10, 64)
						if err != nil {
							return nil, err
						}
					}
					if m["minLength"] != "" {
						label.Fields[i].TextOptions.MinLength, err = strconv.ParseInt(m["minLength"], 10, 64)
						if err != nil {
							return nil, err
						}
					}
				case "integer":
					label.Fields[i].IntegerOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldIntegerOptions{}
					if m["maxValue"] != "" {
						label.Fields[i].IntegerOptions.MaxValue, err = strconv.ParseInt(m["maxValue"], 10, 64)
						if err != nil {
							return nil, err
						}
					}
					if m["minValue"] != "" {
						label.Fields[i].IntegerOptions.MinValue, err = strconv.ParseInt(m["minValue"], 10, 64)
						if err != nil {
							return nil, err
						}
					}
				case "dateString":
					label.Fields[i].DateOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldDateOptions{
						DateFormatType: m["dateFormatType"],
					}
					if m["maxValue"] != "" {
						label.Fields[i].DateOptions.MaxValue, err = stringToLabelDate(m["maxValue"])
						if err != nil {
							return nil, err
						}
					}
					if m["minValue"] != "" {
						label.Fields[i].DateOptions.MinValue, err = stringToLabelDate(m["minValue"])
						if err != nil {
							return nil, err
						}
					}
				case "user":
					label.Fields[i].UserOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldUserOptions{}
					if m["maxEntries"] != "" {
						label.Fields[i].UserOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{}
						label.Fields[i].UserOptions.ListOptions.MaxEntries, err = strconv.ParseInt(m["maxEntries"], 10, 64)
						if err != nil {
							return nil, err
						}
					}
				case "selection":
					label.Fields[i].SelectionOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptions{}
					if m["maxEntries"] != "" {
						label.Fields[i].SelectionOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{}
						label.Fields[i].SelectionOptions.ListOptions.MaxEntries, err = strconv.ParseInt(m["maxEntries"], 10, 64)
						if err != nil {
							return nil, err
						}
					}
					if m["choices"] != "" {
						choices := strings.Split(m["choices"], "|")
						label.Fields[i].SelectionOptions.Choices = make([]*drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice, len(choices))
						for c := range choices {
							label.Fields[i].SelectionOptions.Choices[c] = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice{
								Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties{
									DisplayName: choices[c],
								},
							}
						}
					}
				default:
					return nil, fmt.Errorf("Unknown valueType: %s", m["valueType"])
				}
			}
		} else {
			label.ForceSendFields = append(label.ForceSendFields, "Fields")
		}
	}
	return label, err
}
