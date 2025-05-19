/*
Copyright © 2020-2025 Hannes Hayashi

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

// driveLabelPermissionsCreateCmd represents the create command
var driveLabelPermissionsCreateCmd = &cobra.Command{
	Use: "create",
	Short: `Updates a Label's permissions.
If a permission for the indicated principal doesn't exist, a new Label Permission is created, otherwise the existing permission is updated.
Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.`,
	Long:              "Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels.permissions/create",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToDriveLabelCreatePermission(flags)
		if err != nil {
			log.Fatalf("Error building Drive Label permission object: %v\n", err)
		}
		result, err := gsmdrivelabels.CreateLabelPermission(gsmhelpers.EnsurePrefix(flags["parent"].GetString(), "labels/"), flags["fields"].GetString(), flags["useAdminAccess"].GetBool(), p)
		if err != nil {
			log.Fatalf("Error creation Drive Label permission: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelPermissionsCmd, driveLabelPermissionsCreateCmd, driveLabelPermissionFlags)
}
