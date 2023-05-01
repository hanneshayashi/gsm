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

	"github.com/hanneshayashi/gsm/gsmdrivelabels"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// driveLabelsPublishCmd represents the publish command
var driveLabelsPublishCmd = &cobra.Command{
	Use:   "publish",
	Short: `Publishes a Label`,
	Long: `Publish all draft changes to the Label.
Once published, the Label may not return to its draft state.
See google.apps.drive.labels.v2.Lifecycle for more information.

Publishing a Label will result in a new published revision.
All previous draft revisions will be deleted.
Previous published revisions will be kept but are subject to automated deletion as needed.

Once published, some changes are no longer permitted.
Generally, any change that would invalidate or cause new restrictions on existing metadata related to the Label will be rejected.
For example, the following changes to a Label will be rejected after the Label is published:
* The label cannot be directly deleted. It must be disabled first, then deleted.
* Field.FieldType cannot be changed.
* Changes to Field validation options cannot reject something that was previously accepted.
* Reducing the max entries.

Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels/publish`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToPublishDriveLabelRequest(flags)
		if err != nil {
			log.Fatalf("Error building Drive Label publish request object: %v\n", err)
		}
		result, err := gsmdrivelabels.Publish(gsmhelpers.EnsurePrefix(flags["name"].GetString(), "labels/"), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error publishing Drive Label: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelsCmd, driveLabelsPublishCmd, driveLabelFlags)
}
