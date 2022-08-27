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

// filesCountRecursiveCmd represents the recursive command
var filesCountRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively count files in a folder",
	Long: `If you need to know the size of a Shared Drive, use
"gsm drives getSize", because it will be faster!`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		files := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), "files(id,size,mimeType),nextPageToken", flags["excludeFolders"].GetStringSlice(), flags["includeRoot"].GetBool(), gsmhelpers.MaxThreads(flags["batchThreads"].GetInt()))
		result := gsmdrive.CountFilesAndFolders(files)
		err := gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(filesCountCmd, filesCountRecursiveCmd, fileFlags, recursiveFileFlags)
}
