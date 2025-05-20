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
	"path/filepath"

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// filesCreateCmd represents the create command
var filesCreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Creates a new file or folder. Can also be used to upload files.",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/files/create",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		f, err := mapToFile(flags)
		if err != nil {
			log.Fatalf("Error building file object: %v\n", err)
		}
		var content *os.File
		if flags["localFilePath"].IsSet() {
			content, err = os.Open(flags["localFilePath"].GetString())
			if err != nil {
				log.Fatalf("Error opening file %s: %v", flags["localFilePath"].GetString(), err)
			}
			defer gsmhelpers.CloseLog(content, "fileContent")
			if f.Name == "" {
				f.Name = filepath.Base(content.Name())
			}
		}
		result, err := gsmdrive.CreateFile(f, content, flags["ignoreDefaultVisibility"].GetBool(), flags["keepRevisionForever"].GetBool(), flags["useContentAsIndexableText"].GetBool(), flags["includePermissionsForView"].GetString(), flags["ocrLanguage"].GetString(), flags["sourceMimeType"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error creating file: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesCreateCmd, fileFlags)
}
