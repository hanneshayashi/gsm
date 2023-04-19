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
		AvailableFor: []string{"get", "delete", "updateLabel", "createField", "deleteField", "disableField", "updateField", "updateFieldType", "enableField", "createSelectionChoice", "updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice", "disable"},
		Type:         "string",
		Description: `Label resource name.
May be any of:
  - labels/{id} (equivalent to labels/{id}@latest)
  - labels/{id}@latest
  - labels/{id}@published
  - labels/{id}@{revisionId}
If you don't specify the "labels/" prefix, GSM will automatically prepend it to the request.`,
		Required:       []string{"get", "delete", "updateLabel", "createField", "deleteField", "disableField", "updateField", "updateFieldType", "enableField", "createSelectionChoice", "updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice", "disable"},
		ExcludeFromAll: true,
	},
	"useAdminAccess": {
		AvailableFor: []string{"get", "list", "create", "delete", "updateLabel", "createField", "deleteField", "disableField", "updateField", "updateFieldType", "enableField", "createSelectionChoice", "updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice", "disable"},
		Type:         "bool",
		Description: `Set to true in order to use the user's admin credentials.
The server verifies that the user is an admin for the label before allowing access.`,
	},
	"requiredRevisionId": {
		AvailableFor: []string{"delete"},
		Type:         "string",
		Description: `The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied to.
If this is not the latest revision of the label, the request will not be processed and will return a 400 Bad Request error.`,
	},
	"languageCode": {
		AvailableFor: []string{"get", "list", "create", "disable"},
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
		AvailableFor: []string{"create", "updateLabel"},
		Type:         "string",
		Description:  `Title of the label.`,
		Required:     []string{"create"},
	},
	"description": {
		AvailableFor: []string{"create", "updateLabel"},
		Type:         "string",
		Description:  `The description of the label.`,
		Required:     []string{"create"},
	},
	"fieldId": {
		AvailableFor: []string{"updateField", "updateFieldType", "disableField", "enableField", "deleteField", "createSelectionChoice", "updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice", "disable"},
		Type:         "string",
		Description:  `The ID of the field.`,
		Required:     []string{"updateField", "updateFieldType", "disableField", "enableField", "deleteField", "createSelectionChoice", "updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice", "disable"},
	},
	"choiceId": {
		AvailableFor: []string{"updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice"},
		Type:         "string",
		Description:  `The ID of the choice.`,
		Required:     []string{"updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice"},
	},
	"hideInSearch": {
		AvailableFor: []string{"disableField", "disableSelectionChoice", "disable"},
		Type:         "bool",
		Description: `Whether to hide this disabled object in the search menu for Drive items.
When false, the object is generally shown in the UI as disabled but it appears in the search results when searching for Drive items.
When true, the object is generally hidden in the UI when searching for Drive items.`,
	},
	"showInApply": {
		AvailableFor: []string{"disableField", "disableSelectionChoice", "disable"},
		Type:         "bool",
		Description: `Whether to show this disabled object in the apply menu on Drive items.
When true, the object is generally shown in the UI as disabled and is unselectable.
When false, the object is generally hidden in the UI.`,
	},
	"displayName": {
		AvailableFor: []string{"createField", "updateField", "createSelectionChoice", "updateSelectionChoiceProperties"},
		Type:         "string",
		Description:  `The display text to show in the UI identifying this item.`,
		Required:     []string{"createField", "createSelectionChoice"},
	},
	"insertBeforeField": {
		AvailableFor: []string{"createField", "updateField"},
		Type:         "string",
		Description: `Input only.
Insert or move this field before the indicated field.
If empty, the field is placed at the end of the list.`,
	},
	"insertBeforeChoice": {
		AvailableFor: []string{"createSelectionChoice", "updateSelectionChoiceProperties"},
		Type:         "string",
		Description: `Input only.
Insert or move this choice before the indicated choice.
If empty, the choice is placed at the end of the list.`,
	},
	"choice": {
		AvailableFor: []string{"createField", "updateFieldType"},
		Type:         "stringSlice",
		Description: `A choice for a selection field.
Can be used multiple times to create multiple choices that will be set in the order specified.`,
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
                         - textString        - A text field
                         - integer           - An integer field
                         - user              - A field that allows the selection of one or more user(s)
                           - maxEntries      - The maximum number of entries for the field as a whole number
                         - selection         - A field that allows the user to select on or more choices from a predefined set
                           - maxEntries      - The maximum number of entries for the field as a whole number
                           - choices         - A list of choices in the format "choices=choice1|choice2|choice3"
                         - dateString        - A date field
                           - dateFormatType  - Localized date format options. Possible values are:
						     - LONG_DATE     - Includes full month name. For example, January 12, 1999 (MMMM d, y)
                             - SHORT_DATE    - Short, numeric, representation. For example, 12/13/99 (M/d/yy)`,
	},
	"learnMoreUri": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `Custom URL to present to users to allow them to learn more about this label and how it should be used.`,
	},
	"dateFormatType": {
		AvailableFor: []string{"updateFieldType"},
		Type:         "string",
		Description: `Localized date format options.
May be one of the following:
LONG_DATE   - Includes full month name. For example, January 12, 1999 (MMMM d, y)
SHORT_DATE  - Short, numeric, representation. For example, 12/13/99 (M/d/yy)`,
	},
	"required": {
		AvailableFor: []string{"createField", "updateField"},
		Type:         "bool",
		Description:  `Whether the field should be marked as required.`,
	},
	"valueType": {
		AvailableFor: []string{"createField", "updateFieldType"},
		Type:         "string",
		Description: `The type of the field
May be one of the following:
- dateString  - A date field
- integer     - An integer field
- selection   - A field that allows the user to select on or more choices from a predefined set
- textString  - A text field
- user        - A field that allows the selection of one or more user(s)`,
	},
	"maxEntries": {
		AvailableFor: []string{"updateFieldType"},
		Type:         "int64",
		Description: `The maximum number of entries for the field as a whole number.
Can be used with "user" or "selection type fields`,
	},
	"fields": {
		AvailableFor: []string{"get", "list", "create", "updateLabel", "createField", "deleteField", "disableField", "updateField", "updateFieldType", "enableField", "createSelectionChoice", "updateSelectionChoiceProperties", "disableSelectionChoice", "enableSelectionChoice", "deleteSelectionChoice", "disable"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(driveLabelsCmd)
}

func mapToDriveLabelField(m map[string]string) (*drivelabels.GoogleAppsDriveLabelsV2Field, error) {
	var err error
	field := &drivelabels.GoogleAppsDriveLabelsV2Field{}
	field.Properties = &drivelabels.GoogleAppsDriveLabelsV2FieldProperties{
		DisplayName:       m["displayName"],
		InsertBeforeField: m["insertBeforeField"],
	}
	if m["required"] != "" {
		field.Properties.Required, err = strconv.ParseBool(m["required"])
		if err != nil {
			return nil, err
		}
		if !field.Properties.Required {
			field.Properties.ForceSendFields = append(field.Properties.ForceSendFields, "Required")
		}
	}
	switch m["valueType"] {
	case "textString":
		field.TextOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldTextOptions{}
	case "integer":
		field.IntegerOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldIntegerOptions{}
	case "dateString":
		field.DateOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldDateOptions{}
		dateFormatType, ok := m["dateFormatType"]
		if ok {
			field.DateOptions.DateFormatType = dateFormatType
			if field.DateOptions.DateFormatType == "" {
				field.DateOptions.ForceSendFields = append(field.DateOptions.ForceSendFields, "DateFormatType")
			}
		}
	case "user":
		field.UserOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldUserOptions{}
		maxEntries, ok := m["maxEntries"]
		if ok {
			field.UserOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{}
			field.UserOptions.ListOptions.MaxEntries, err = strconv.ParseInt(maxEntries, 10, 64)
			if err != nil {
				return nil, err
			}
			if field.UserOptions.ListOptions.MaxEntries == 0 {
				field.UserOptions.ListOptions.ForceSendFields = append(field.UserOptions.ListOptions.ForceSendFields, "MaxEntries")
			}
		}
	case "selection":
		field.SelectionOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptions{}
		maxEntries, ok := m["maxEntries"]
		if ok {
			field.SelectionOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{}
			field.SelectionOptions.ListOptions.MaxEntries, err = strconv.ParseInt(maxEntries, 10, 64)
			if err != nil {
				return nil, err
			}
			if field.SelectionOptions.ListOptions.MaxEntries == 0 {
				field.SelectionOptions.ListOptions.ForceSendFields = append(field.SelectionOptions.ListOptions.ForceSendFields, "MaxEntries")
			}
		}
		_, ok = m["choices"]
		if ok {
			choices := strings.Split(m["choices"], "|")
			field.SelectionOptions.Choices = make([]*drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice, len(choices))
			for i := range choices {
				field.SelectionOptions.Choices[i] = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice{
					Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties{
						DisplayName: choices[i],
					},
				}
			}
		}
	default:
		return nil, fmt.Errorf("Unknown valueType: %s", m["valueType"])
	}
	return field, nil
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
				label.Fields[i], err = mapToDriveLabelField(gsmhelpers.FlagToMap(fields[i]))
				if err != nil {
					return nil, err
				}
			}
		} else {
			label.ForceSendFields = append(label.ForceSendFields, "Fields")
		}
	}
	return label, err
}

func mapToUpdateDriveLabelRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				UpdateLabel: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateLabelPropertiesRequest{
					Properties: &drivelabels.GoogleAppsDriveLabelsV2LabelProperties{},
				},
			},
		},
	}
	if flags["title"].IsSet() {
		request.Requests[0].UpdateLabel.Properties.Title = flags["title"].GetString()
		if request.Requests[0].UpdateLabel.Properties.Title == "" {
			request.Requests[0].UpdateLabel.Properties.ForceSendFields = append(request.Requests[0].UpdateLabel.Properties.ForceSendFields, "Title")
		}
	}
	if flags["description"].IsSet() {
		request.Requests[0].UpdateLabel.Properties.Description = flags["description"].GetString()
		if request.Requests[0].UpdateLabel.Properties.Description == "" {
			request.Requests[0].UpdateLabel.Properties.ForceSendFields = append(request.Requests[0].UpdateLabel.Properties.ForceSendFields, "Description")
		}
	}
	return request, nil
}

func mapToUpdateDriveLabelFieldTypeRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				UpdateFieldType: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateFieldTypeRequest{
					Id: flags["fieldId"].GetString(),
				},
			},
		},
	}
	valueType := flags["valueType"].GetString()
	switch valueType {
	case "textString":
		request.Requests[0].UpdateFieldType.TextOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldTextOptions{}
	case "integer":
		request.Requests[0].UpdateFieldType.IntegerOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldIntegerOptions{}
	case "dateString":
		request.Requests[0].UpdateFieldType.DateOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldDateOptions{}
		if flags["dateFormatType"].IsSet() {
			request.Requests[0].UpdateFieldType.DateOptions.DateFormatType = flags["dateFormatType"].GetString()
			if request.Requests[0].UpdateFieldType.DateOptions.DateFormatType == "" {
				request.Requests[0].UpdateFieldType.DateOptions.ForceSendFields = append(request.Requests[0].UpdateFieldType.DateOptions.ForceSendFields, "DateFormatType")
			}
		}
	case "user":
		request.Requests[0].UpdateFieldType.UserOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldUserOptions{}
		if flags["maxEntries"].IsSet() {
			request.Requests[0].UpdateFieldType.UserOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{
				MaxEntries: flags["maxEntries"].GetInt64(),
			}
			if request.Requests[0].UpdateFieldType.UserOptions.ListOptions.MaxEntries == 0 {
				request.Requests[0].UpdateFieldType.UserOptions.ListOptions.ForceSendFields = append(request.Requests[0].UpdateFieldType.UserOptions.ListOptions.ForceSendFields, "MaxEntries")
			}
		}
	case "selection":
		request.Requests[0].UpdateFieldType.SelectionOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptions{}
		if flags["maxEntries"].IsSet() {
			request.Requests[0].UpdateFieldType.SelectionOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{
				MaxEntries: flags["maxEntries"].GetInt64(),
			}
			if request.Requests[0].UpdateFieldType.SelectionOptions.ListOptions.MaxEntries == 0 {
				request.Requests[0].UpdateFieldType.SelectionOptions.ListOptions.ForceSendFields = append(request.Requests[0].UpdateFieldType.SelectionOptions.ListOptions.ForceSendFields, "MaxEntries")
			}
		}
		if flags["choice"].IsSet() {
			choices := flags["choice"].GetStringSlice()
			request.Requests[0].UpdateFieldType.SelectionOptions.Choices = make([]*drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice, len(choices))
			for i := range choices {
				request.Requests[0].UpdateFieldType.SelectionOptions.Choices[i] = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice{
					Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties{
						DisplayName: choices[i],
					},
				}
			}
		}
	default:
		return nil, fmt.Errorf("Unknown valueType: %s", valueType)
	}
	return request, nil
}

func mapToUpdateDriveLabelFieldRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				UpdateField: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateFieldPropertiesRequest{
					Id:         flags["fieldId"].GetString(),
					Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldProperties{},
				},
			},
		},
	}
	if flags["displayName"].IsSet() {
		request.Requests[0].UpdateField.Properties.DisplayName = flags["displayName"].GetString()
		if request.Requests[0].UpdateField.Properties.DisplayName == "" {
			request.Requests[0].UpdateField.Properties.ForceSendFields = append(request.Requests[0].UpdateField.Properties.ForceSendFields, "DisplayName")
		}
	}
	if flags["insertBeforeField"].IsSet() {
		request.Requests[0].UpdateField.Properties.InsertBeforeField = flags["insertBeforeField"].GetString()
		if request.Requests[0].UpdateField.Properties.InsertBeforeField == "" {
			request.Requests[0].UpdateField.Properties.ForceSendFields = append(request.Requests[0].UpdateField.Properties.ForceSendFields, "InsertBeforeField")
		}
	}
	if flags["required"].IsSet() {
		request.Requests[0].UpdateField.Properties.Required = flags["required"].GetBool()
		if !request.Requests[0].UpdateField.Properties.Required {
			request.Requests[0].UpdateField.Properties.ForceSendFields = append(request.Requests[0].UpdateField.Properties.ForceSendFields, "Required")
		}
	}
	return request, nil
}

func mapToCreateDriveLabelFieldRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				CreateField: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestCreateFieldRequest{
					Field: &drivelabels.GoogleAppsDriveLabelsV2Field{
						Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldProperties{
							DisplayName: flags["displayName"].GetString(),
						},
					},
				},
			},
		},
	}
	if flags["insertBeforeField"].IsSet() {
		request.Requests[0].UpdateField.Properties.InsertBeforeField = flags["insertBeforeField"].GetString()
		if request.Requests[0].UpdateField.Properties.InsertBeforeField == "" {
			request.Requests[0].UpdateField.Properties.ForceSendFields = append(request.Requests[0].UpdateField.Properties.ForceSendFields, "InsertBeforeField")
		}
	}
	if flags["required"].IsSet() {
		request.Requests[0].UpdateField.Properties.Required = flags["required"].GetBool()
		if !request.Requests[0].UpdateField.Properties.Required {
			request.Requests[0].UpdateField.Properties.ForceSendFields = append(request.Requests[0].UpdateField.Properties.ForceSendFields, "Required")
		}
	}
	valueType := flags["valueType"].GetString()
	switch valueType {
	case "textString":
		request.Requests[0].CreateField.Field.TextOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldTextOptions{}
	case "integer":
		request.Requests[0].CreateField.Field.IntegerOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldIntegerOptions{}
	case "dateString":
		request.Requests[0].CreateField.Field.DateOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldDateOptions{}
		if flags["dateFormatType"].IsSet() {
			request.Requests[0].CreateField.Field.DateOptions.DateFormatType = flags["dateFormattype"].GetString()
			if request.Requests[0].CreateField.Field.DateOptions.DateFormatType == "" {
				request.Requests[0].CreateField.Field.DateOptions.ForceSendFields = append(request.Requests[0].CreateField.Field.DateOptions.ForceSendFields, "DateFormatType")
			}
		}
	case "user":
		request.Requests[0].CreateField.Field.UserOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldUserOptions{}
		if flags["maxEntries"].IsSet() {
			request.Requests[0].CreateField.Field.UserOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{
				MaxEntries: flags["maxEntries"].GetInt64(),
			}
			if request.Requests[0].CreateField.Field.UserOptions.ListOptions.MaxEntries == 0 {
				request.Requests[0].CreateField.Field.UserOptions.ListOptions.ForceSendFields = append(request.Requests[0].CreateField.Field.UserOptions.ListOptions.ForceSendFields, "MaxEntries")
			}
		}
	case "selection":
		request.Requests[0].CreateField.Field.SelectionOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptions{}
		if flags["maxEntries"].IsSet() {
			request.Requests[0].CreateField.Field.SelectionOptions.ListOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldListOptions{
				MaxEntries: flags["maxEntries"].GetInt64(),
			}
			if request.Requests[0].CreateField.Field.SelectionOptions.ListOptions.MaxEntries == 0 {
				request.Requests[0].CreateField.Field.SelectionOptions.ListOptions.ForceSendFields = append(request.Requests[0].CreateField.Field.SelectionOptions.ListOptions.ForceSendFields, "MaxEntries")
			}
		}
		if flags["choice"].IsSet() {
			choices := flags["choice"].GetStringSlice()
			request.Requests[0].CreateField.Field.SelectionOptions.Choices = make([]*drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice, len(choices))
			for i := range choices {
				request.Requests[0].CreateField.Field.SelectionOptions.Choices[i] = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice{
					Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties{
						DisplayName: choices[i],
					},
				}
			}
		}
	default:
		return nil, fmt.Errorf("Unknown valueType: %s", valueType)
	}
	return request, nil
}

func mapToDeleteDriveLabelFieldRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				DeleteField: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDeleteFieldRequest{
					Id: flags["fieldId"].GetString(),
				},
			},
		},
	}
	return request, nil
}

func mapToDisableDriveLabelFieldRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				DisableField: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDisableFieldRequest{
					Id:             flags["fieldId"].GetString(),
					DisabledPolicy: &drivelabels.GoogleAppsDriveLabelsV2LifecycleDisabledPolicy{},
				},
			},
		},
	}
	if flags["hideInSearch"].IsSet() {
		request.Requests[0].DisableField.DisabledPolicy.HideInSearch = flags["hideInSearch"].GetBool()
		if !request.Requests[0].DisableField.DisabledPolicy.HideInSearch {
			request.Requests[0].DisableField.DisabledPolicy.ForceSendFields = append(request.Requests[0].DisableField.DisabledPolicy.ForceSendFields, "HideInSearch")
		}
	}
	if flags["showInApply"].IsSet() {
		request.Requests[0].DisableField.DisabledPolicy.ShowInApply = flags["showInApply"].GetBool()
		if !request.Requests[0].DisableField.DisabledPolicy.ShowInApply {
			request.Requests[0].DisableField.DisabledPolicy.ForceSendFields = append(request.Requests[0].DisableField.DisabledPolicy.ForceSendFields, "ShowInApply")
		}
	}
	return request, nil
}

func mapToEnableDriveLabelFieldRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				EnableField: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestEnableFieldRequest{
					Id: flags["fieldId"].GetString(),
				},
			},
		},
	}
	return request, nil
}

func mapToCreateDriveLabelFieldSelectionChoiceRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				CreateSelectionChoice: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestCreateSelectionChoiceRequest{
					FieldId: flags["fieldId"].GetString(),
					Choice: &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice{
						Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties{
							DisplayName: flags["displayName"].GetString(),
						},
					},
				},
			},
		},
	}
	if request.Requests[0].CreateSelectionChoice.Choice.Properties.DisplayName == "" {
		request.Requests[0].CreateSelectionChoice.Choice.Properties.ForceSendFields = append(request.Requests[0].CreateSelectionChoice.Choice.Properties.ForceSendFields, "DisplayName")
	}
	if flags["insertBeforeChoice"].IsSet() {
		request.Requests[0].CreateSelectionChoice.Choice.Properties.InsertBeforeChoice = flags["insertBeforeChoice"].GetString()
		if request.Requests[0].CreateSelectionChoice.Choice.Properties.InsertBeforeChoice == "" {
			request.Requests[0].CreateSelectionChoice.Choice.Properties.ForceSendFields = append(request.Requests[0].CreateSelectionChoice.Choice.Properties.ForceSendFields, "InsertBeforeChoice")
		}
	}
	return request, nil
}

func mapToUpdateDriveLabelFieldSelectionChoiceRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				UpdateSelectionChoiceProperties: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateSelectionChoicePropertiesRequest{
					FieldId:    flags["fieldId"].GetString(),
					Id:         flags["choiceId"].GetString(),
					Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties{},
				},
			},
		},
	}
	if flags["displayName"].IsSet() {
		request.Requests[0].UpdateSelectionChoiceProperties.Properties.DisplayName = flags["displayName"].GetString()
		if request.Requests[0].UpdateSelectionChoiceProperties.Properties.DisplayName == "" {
			request.Requests[0].UpdateSelectionChoiceProperties.Properties.ForceSendFields = append(request.Requests[0].UpdateSelectionChoiceProperties.Properties.ForceSendFields, "DisplayName")
		}
	}
	if flags["insertBeforeChoice"].IsSet() {
		request.Requests[0].UpdateSelectionChoiceProperties.Properties.InsertBeforeChoice = flags["insertBeforeChoice"].GetString()
		if request.Requests[0].UpdateSelectionChoiceProperties.Properties.InsertBeforeChoice == "" {
			request.Requests[0].UpdateSelectionChoiceProperties.Properties.ForceSendFields = append(request.Requests[0].UpdateSelectionChoiceProperties.Properties.ForceSendFields, "InsertBeforeChoice")
		}
	}
	return request, nil
}

func mapToDisableDriveLabelFieldSelectionChoiceRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				DisableSelectionChoice: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDisableSelectionChoiceRequest{
					FieldId:        flags["fieldId"].GetString(),
					Id:             flags["choiceId"].GetString(),
					DisabledPolicy: &drivelabels.GoogleAppsDriveLabelsV2LifecycleDisabledPolicy{},
				},
			},
		},
	}
	if flags["hideInSearch"].IsSet() {
		request.Requests[0].DisableSelectionChoice.DisabledPolicy.HideInSearch = flags["hideInSearch"].GetBool()
		if !request.Requests[0].DisableSelectionChoice.DisabledPolicy.HideInSearch {
			request.Requests[0].DisableSelectionChoice.DisabledPolicy.ForceSendFields = append(request.Requests[0].DisableSelectionChoice.DisabledPolicy.ForceSendFields, "HideInSearch")
		}
	}
	if flags["showInApply"].IsSet() {
		request.Requests[0].DisableSelectionChoice.DisabledPolicy.ShowInApply = flags["showInApply"].GetBool()
		if !request.Requests[0].DisableSelectionChoice.DisabledPolicy.ShowInApply {
			request.Requests[0].DisableSelectionChoice.DisabledPolicy.ForceSendFields = append(request.Requests[0].DisableSelectionChoice.DisabledPolicy.ForceSendFields, "ShowInApply")
		}
	}
	return request, nil
}

func mapToEnableDriveLabelFieldSelectionChoiceRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				EnableSelectionChoice: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestEnableSelectionChoiceRequest{
					FieldId: flags["fieldId"].GetString(),
					Id:      flags["choiceId"].GetString(),
				},
			},
		},
	}
	return request, nil
}

func mapToDeleteDriveLabelFieldSelectionChoiceRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		Requests: []*drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest{
			{
				DeleteSelectionChoice: &drivelabels.GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDeleteSelectionChoiceRequest{
					FieldId: flags["fieldId"].GetString(),
					Id:      flags["choiceId"].GetString(),
				},
			},
		},
	}
	return request, nil
}

func mapToDisableDriveLabelRequest(flags map[string]*gsmhelpers.Value) (*drivelabels.GoogleAppsDriveLabelsV2DisableLabelRequest, error) {
	request := &drivelabels.GoogleAppsDriveLabelsV2DisableLabelRequest{
		UseAdminAccess: flags["useAdminAccess"].GetBool(),
		DisabledPolicy: &drivelabels.GoogleAppsDriveLabelsV2LifecycleDisabledPolicy{},
	}
	if flags["hideInSearch"].IsSet() {
		request.DisabledPolicy.HideInSearch = flags["hideInSearch"].GetBool()
		if !request.DisabledPolicy.HideInSearch {
			request.DisabledPolicy.ForceSendFields = append(request.DisabledPolicy.ForceSendFields, "HideInSearch")
		}
	}
	if flags["showInApply"].IsSet() {
		request.DisabledPolicy.ShowInApply = flags["showInApply"].GetBool()
		if !request.DisabledPolicy.ShowInApply {
			request.DisabledPolicy.ForceSendFields = append(request.DisabledPolicy.ForceSendFields, "ShowInApply")
		}
	}
	if flags["languageCode"].IsSet() {
		request.LanguageCode = flags["languageCode"].GetString()
		if request.LanguageCode == "" {
			request.ForceSendFields = append(request.ForceSendFields, "LanguageCode")
		}
	}
	return request, nil
}
