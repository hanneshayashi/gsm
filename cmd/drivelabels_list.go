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
	"github.com/hanneshayashi/gsm/gsmdrivelabels"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// driveLabelsListCmd represents the list command
var driveLabelsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List labels.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/labels/reference/rest/v2/labels/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		gsmhelpers.StreamOrCollect(gsmdrivelabels.ListLabels(flags["languageCode"].GetString(), flags["view"].GetString(), flags["minimumRole"].GetString(), flags["fields"].GetString(), flags["useAdminAccess"].GetBool(), flags["publishedOnly"].GetBool()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelsCmd, driveLabelsListCmd, driveLabelFlags)
}
