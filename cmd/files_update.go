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
	"os"
	"strings"

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/drive/v3"

	"github.com/spf13/cobra"
)

// filesUpdateCmd represents the update command
var filesUpdateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Updates a file's metadata and/or content. This method supports patch semantics.",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/files/update",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		f, err := mapToFile(flags)
		if err != nil {
			log.Fatalf("Error building file object: %v\n", err)
		}
		var removeParents string
		if flags["parent"].IsSet() {
			f.Parents = nil
			var fOld *drive.File
			fOld, err = gsmdrive.GetFile(flags["fileId"].GetString(), "parents", "")
			if err != nil {
				log.Fatalf("Error getting existing file: %v\n", err)
			}
			removeParents = strings.Join(fOld.Parents, ",")
		}
		var content *os.File
		if flags["localFilePath"].IsSet() {
			content, err = os.Open(flags["localFilePath"].GetString())
			if err != nil {
				log.Fatalf("Error opening file %s: %v", flags["localFilePath"].GetString(), err)
			}
			defer gsmhelpers.CloseLog(content, "fileContent")
		}
		result, err := gsmdrive.UpdateFile(flags["fileId"].GetString(), flags["parent"].GetString(), removeParents, flags["includePermissionsForView"].GetString(), flags["ocrLanguage"].GetString(), flags["fields"].GetString(), f, content, flags["keepRevisionForever"].GetBool(), flags["useContentAsIndexableText"].GetBool())
		if err != nil {
			log.Fatalf("Error updating file: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesUpdateCmd, fileFlags)
}
