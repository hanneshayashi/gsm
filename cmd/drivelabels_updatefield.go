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

	"github.com/hanneshayashi/gsm/gsmdrivelabels"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// driveLabelsUpdateFieldCmd represents the updateField command
var driveLabelsUpdateFieldCmd = &cobra.Command{
	Use:               "updateField",
	Short:             `Updates basic properties of a Field.`,
	Long:              `Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels/delta#UpdateFieldPropertiesRequest`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToUpdateDriveLabelFieldRequest(flags)
		if err != nil {
			log.Fatalf("Error building Drive Label UpdateFieldPropertiesRequest object: %v\n", err)
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
	gsmhelpers.InitCommand(driveLabelsCmd, driveLabelsUpdateFieldCmd, driveLabelFlags)
}
