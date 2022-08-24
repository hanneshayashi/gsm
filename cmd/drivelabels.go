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
		AvailableFor: []string{"get", "list"},
		Type:         "bool",
		Description: `Set to true in order to use the user's admin credentials.
The server verifies that the user is an admin for the label before allowing access.`,
	},
	"languageCode": {
		AvailableFor: []string{"get", "list"},
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
	"fields": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(driveLabelsCmd)
}
