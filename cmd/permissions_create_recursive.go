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

// permissionsCreateRecursiveCmd represents the recursive command
var permissionsCreateRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively grant a permissions to a folder and all of its children.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/create",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToPermission(flags)
		if err != nil {
			log.Fatalf("Error building permission object: %v", err)
		}
		files, _ := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), "files(id,mimeType),nextPageToken", 10)
		var ids []string
		for _, f := range files {
			ids = append(ids, f.Id)
		}
		result, errs := gsmdrive.CreatePermissionRecursive(ids, flags["emailMessage"].GetString(), flags["fields"].GetString(), flags["useDomainAdminAccess"].GetBool(), flags["sendNotificationEmail"].GetBool(), flags["transferOwnership"].GetBool(), false, p, 10)
		if len(errs) > 0 {
			log.Fatalf("Error creating permissions %v", errs)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(permissionsCreateCmd, permissionsCreateRecursiveCmd, permissionFlags, recursiveFlags)
}
