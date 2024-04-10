/*
Copyright © 2020-2024 Hannes Hayashi

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

// driveLabelPermissionsUpdateCmd represents the update command
var driveLabelPermissionsUpdateCmd = &cobra.Command{
	Use: "update",
	Short: `Updates a Label's permissions.
The permission must exist and be referenced with the "name" parameter.
Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.`,
	Long:              "Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels/updatePermissions",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToDriveLabelUpdatePermission(flags)
		if err != nil {
			log.Fatalf("Error building Drive Label permission object: %v\n", err)
		}
		result, err := gsmdrivelabels.UpdatePermissions(gsmhelpers.EnsurePrefix(flags["parent"].GetString(), "labels/"), flags["fields"].GetString(), flags["useAdminAccess"].GetBool(), p)
		if err != nil {
			log.Fatalf("Error updating Drive Label permission: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(driveLabelPermissionsCmd, driveLabelPermissionsUpdateCmd, driveLabelPermissionFlags)
}
