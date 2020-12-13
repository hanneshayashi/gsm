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
	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// permissionsCreateCmd represents the create command
var permissionsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a permission for a file or shared drive.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/create",	
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToPermission(flags)
		if err != nil {
			log.Fatalf("Error building permission object: %v", err)
		}
		result, err := gsmdrive.CreatePermission(flags["fileId"].GetString(), flags["emailMessage"].GetString(), flags["fields"].GetString(), flags["useDomainAdminAccess"].GetBool(), flags["sendNotificationEmail"].GetBool(), flags["transferOwnership"].GetBool(), flags["moveToNewOwnersRoot"].GetBool(), p)
		if err != nil {
			log.Fatalf("Error creating permission %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(permissionsCmd, permissionsCreateCmd, permissionFlags)
}
