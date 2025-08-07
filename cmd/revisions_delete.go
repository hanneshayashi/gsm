/*
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

// revisionsDeleteCmd represents the delete command
var revisionsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently deletes a file version.",
	Long: `You can only delete revisions for files with binary content in Google Drive, like images or videos.
Revisions for other files, like Google Docs or Sheets, and the last remaining file version can't be deleted.
Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/revisions/delete`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrive.DeleteRevision(flags["fileId"].GetString(), flags["revisionId"].GetString())
		if err != nil {
			log.Fatalf("Error deleting revision: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(revisionsCmd, revisionsDeleteCmd, revisionFlags)
}
