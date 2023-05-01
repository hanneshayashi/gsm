/*
Copyright © 2020-2023 Hannes Hayashi

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

// drivesGetSizeCmd represents the getSize command
var drivesGetSizeCmd = &cobra.Command{
	Use:   "getSize",
	Short: "Counts the files in a Shared Drive and returns their number and total size",
	Long: `If you need to know the size of a folder, use
"gsm files count recursive"!`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		var q string
		if !flags["includeTrash"].GetBool() {
			q = "trashed = false"
		}
		files, err := gsmdrive.ListFiles(q, flags["driveId"].GetString(), "drive", "", "", "drive", "files(mimeType,size),nextPageToken", true, gsmhelpers.MaxThreads(0))
		result := gsmdrive.CountFilesAndFolders(files)
		e := <-err
		if e != nil {
			log.Fatalf("Error counting files: %v", e)
		}
		er := gsmhelpers.Output(result, "json", compressOutput)
		if er != nil {
			log.Fatalln(er)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(drivesCmd, drivesGetSizeCmd, driveFlags)
}
