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

// driveLabelsEnableCmd represents the enable command
var driveLabelsEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: `Enables a Label`,
	Long: `Enable a disabled Label and restore it to its published state.
This will result in a new published revision based on the current disabled published revision.
If there is an existing disabled draft revision, a new revision will be created based on that draft and will be enabled.

Implements the API documented at https://developers.google.com/workspace/drive/labels/reference/rest/v2/labels/enable`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToEnableDriveLabelRequest(flags)
		if err != nil {
			log.Fatalf("Error building Drive Label enable request object: %v\n", err)
		}
		result, err := gsmdrivelabels.Enable(gsmhelpers.EnsurePrefix(flags["name"].GetString(), "labels/"), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error enabling Drive Label: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelsCmd, driveLabelsEnableCmd, driveLabelFlags)
}
