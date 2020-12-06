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
	"gsm/gsmdrive"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// permissionsDeleteCmd represents the delete command
var permissionsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a permission.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/delete",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrive.DeletePermission(flags["fileId"].GetString(), flags["permissionId"].GetString(), flags["useDomainAdminAccess"].GetBool())
		if err != nil {
			log.Fatalf("Error deleting permission %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(permissionsCmd, permissionsDeleteCmd, permissionFlags)
}
