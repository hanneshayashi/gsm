/*
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
	"log"

	"github.com/hanneshayashi/gsm/gsmdrivelabels"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// driveLabelsDeleteSelectionChoiceCmd represents the deleteSelectionChoice command
var driveLabelsDeleteSelectionChoiceCmd = &cobra.Command{
	Use:               "deleteSelectionChoice",
	Short:             `Deletes a choice for an existing label selection field on a label`,
	Long:              `Implements the API documented at https://developers.google.com/workspace/drive/labels/reference/rest/v2/labels/delta#DeleteSelectionChoiceRequest`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToDeleteDriveLabelFieldSelectionChoiceRequest(flags)
		if err != nil {
			log.Fatalf("Error building Drive Label DeleteSelectionChoiceRequest object: %v\n", err)
		}
		result, err := gsmdrivelabels.Delta(gsmhelpers.EnsurePrefix(flags["name"].GetString(), "labels/"), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error updating Drive Label: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelsCmd, driveLabelsDeleteSelectionChoiceCmd, driveLabelFlags)
}
