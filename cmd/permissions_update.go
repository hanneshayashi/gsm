/*
Package cmd contains the commands available to the end user
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
	"fmt"
	"gsm/gsmdrive"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// permissionsUpdateCmd represents the update command
var permissionsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/update",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToPermission(flags)
		if err != nil {
			log.Fatalf("Error building permission object: %v", err)
		}
		result, err := gsmdrive.UpdatePermission(flags["fileId"].GetString(), flags["permissionId"].GetString(), flags["fields"].GetString(), flags["useDomainAdminAccess"].GetBool(), flags["removeExpiration"].GetBool(), p)
		if err != nil {
			log.Fatalf("Error updating permission %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(permissionsCmd, permissionsUpdateCmd, permissionFlags)
}
