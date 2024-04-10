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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// userPhotosUpdateCmd represents the update command
var userPhotosUpdateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Adds a photo for the user.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/users.photos/update",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToUserPhoto(flags)
		if err != nil {
			log.Fatalf("Error building userPhoto object: %v", err)
		}
		result, err := gsmadmin.UpdateUserPhoto(flags["userKey"].GetString(), flags["fields"].GetString(), p)
		if err != nil {
			log.Fatalf("Error updating user photo: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(userPhotosCmd, userPhotosUpdateCmd, userPhotoFlags)
}
