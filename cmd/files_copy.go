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

// filesCopyCmd represents the copy command
var filesCopyCmd = &cobra.Command{
	Use: "copy",
	Short: `Creates a copy of a file and applies any requested updates with patch semantics.
Use "files copy recursive" to copy folders.`,
	Long: "https://developers.google.com/drive/api/v3/reference/files/copy",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		f, err := mapToFile(flags)
		if err != nil {
			log.Fatalf("Error building file object: %v\n", err)
		}
		result, err := gsmdrive.CopyFile(flags["fileId"].GetString(), flags["includePermissionsForView"].GetString(), flags["ocrLanguage"].GetString(), flags["fields"].GetString(), f, flags["ignoreDefaultVisibility"].GetBool(), flags["keepRevisionForever"].GetBool())
		if err != nil {
			log.Fatalf("Error creating file %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesCopyCmd, fileFlags)
}
