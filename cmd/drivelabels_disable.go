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

// driveLabelsDisableCmd represents the disable command
var driveLabelsDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: `Disables a Label`,
	Long: `Disable a published Label.
Disabling a Label will result in a new disabled published revision based on the current published revision.
If there is a draft revision, a new disabled draft revision will be created based on the latest draft revision.
Older draft revisions will be deleted.

Once disabled, a label may be deleted with "gsm driveLabels delete".
Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels/disable`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToDisableDriveLabelRequest(flags)
		if err != nil {
			log.Fatalf("Error building Drive Label disable request object: %v\n", err)
		}
		result, err := gsmdrivelabels.Disable(gsmhelpers.EnsurePrefix(flags["name"].GetString(), "labels/"), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error disabling Drive Label: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelsCmd, driveLabelsDisableCmd, driveLabelFlags)
}
