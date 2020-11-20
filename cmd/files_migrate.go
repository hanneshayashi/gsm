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

// filesMigrateCmd represents the migrate command
var filesMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate a folder to a Shared Drive",
	Long: `Example:
migrate --folderId <folderId> --driveId <driveId>
For each source folder, a new folder will be created at the destination.
Files will be moved (not copied) to the new folders.
The original folders will be preserved at the source!`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		file, err := gsmdrive.GetFile(flags["folderId"].GetString(), "id, name, mimeType, parents", "")
		if err != nil {
			log.Fatalf("%v", err)
		}
		gsmdrive.Migrate(file, flags["driveId"].GetString(), flags["driveId"].GetString())
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesMigrateCmd, fileFlags)
}
