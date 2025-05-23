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

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// permissionsUpdateCmd represents the update command
var permissionsUpdateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Updates a permission with patch semantics.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/permissions/update",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToPermission(flags)
		if err != nil {
			log.Fatalf("Error building permission object: %v", err)
		}
		permissionID, err := gsmdrive.GetPermissionID(flags)
		if err != nil {
			log.Fatalf("Unable to determine permissionId: %v", err)
		}
		result, err := gsmdrive.UpdatePermission(flags["fileId"].GetString(), permissionID, flags["fields"].GetString(), flags["useDomainAdminAccess"].GetBool(), flags["removeExpiration"].GetBool(), flags["enforceExpansiveAccess"].GetBool(), p)
		if err != nil {
			log.Fatalf("Error updating permission: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(permissionsCmd, permissionsUpdateCmd, permissionFlags)
}
