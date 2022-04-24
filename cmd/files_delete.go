/*
Copyright © 2020-2022 Hannes Hayashi

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

// filesDeleteCmd represents the delete command
var filesDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: `Permanently deletes a file owned by the user without moving it to the trash.
If the file belongs to a shared drive the user must be an organizer on the parent.
If the target is a folder, all descendants owned by the user are also deleted.`,
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/files/delete",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrive.DeleteFile(flags["fileId"].GetString())
		if err != nil {
			log.Fatalf("Error deleting file: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesDeleteCmd, fileFlags)
}
