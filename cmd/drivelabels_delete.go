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

// driveLabelsDeleteCmd represents the delete command
var driveLabelsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: `Deletes a Label`,
	Long: `Permanently deletes a Label and related metadata on Drive Items.
Once deleted, the Label and related Drive item metadata will be deleted.
Only draft Labels, and disabled Labels may be deleted.

Implements the API documented at https://developers.google.com/workspace/drive/labels/reference/rest/v2/labels/delete`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrivelabels.DeleteLabel(gsmhelpers.EnsurePrefix(flags["name"].GetString(), "labels/"), flags["requiredRevisionId"].GetString(), flags["useAdminAccess"].GetBool())
		if err != nil {
			log.Fatalf("Error deleting Drive Label: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelsCmd, driveLabelsDeleteCmd, driveLabelFlags)
}
