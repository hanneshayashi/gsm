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
	"fmt"
	"log"

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// filesCountCmd represents the count command
var filesCountCmd = &cobra.Command{
	Use:               "count",
	Short:             "Counts files in a folder and returns their number and size.",
	Long:              "Use the recursive subcommand to also scan subfolders",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		filesCh, err := gsmdrive.ListFiles(fmt.Sprintf("'%s' in parents", flags["folderId"].GetString()), "", "allDrives", "", "", "", "files(mimeType,size),nextPageToken", true, gsmhelpers.MaxThreads(0))
		result := gsmdrive.CountFilesAndFolders(filesCh)
		e := <-err
		if e != nil {
			log.Fatalf("Error listing files: %v", e)
		}
		er := gsmhelpers.Output(result, "json", compressOutput)
		if er != nil {
			log.Fatalln(er)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesCountCmd, fileFlags)
}
