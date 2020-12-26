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
	"log"

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// filesExportCmd represents the export command
var filesExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports a Google Doc to the requested MIME type and returns the exported content.",
	Long: `Please note that the exported content is limited to 10MB.
https://developers.google.com/drive/api/v3/reference/files/export`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrive.ExportFile(flags["fileId"].GetString(), flags["mimeType"].GetString())
		if err != nil {
			log.Fatalf("Error downloading file: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesExportCmd, fileFlags)
}
