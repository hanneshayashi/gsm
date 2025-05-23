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
	"google.golang.org/api/drivelabels/v2"

	"github.com/spf13/cobra"
)

// driveLabelPermissionsListCmd represents the list command
var driveLabelPermissionsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists the LabelPermissions on a Label.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/labels/reference/rest/v2/labels.permissions/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrivelabels.ListLabelPermissions(gsmhelpers.EnsurePrefix(flags["parent"].GetString(), "labels/"), flags["fields"].GetString(), flags["useAdminAccess"].GetBool(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*drivelabels.GoogleAppsDriveLabelsV2LabelPermission{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing Drive Label permissions: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelPermissionsCmd, driveLabelPermissionsListCmd, driveLabelPermissionFlags)
}
