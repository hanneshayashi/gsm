/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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
	"google.golang.org/api/drive/v3"
)

// changesListCmd represents the list command
var changesListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists the changes for a user or shared drive.",
	Long:              "https://developers.google.com/drive/api/v3/reference/changes/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, nextStartPageToken, err := gsmdrive.ListChanges(flags["pageToken"].GetString(), flags["driveId"].GetString(), flags["spaces"].GetString(), flags["fields"].GetString(), flags["includePermissionsForView"].GetString(), flags["includeCorpusRemovals"].GetBool(), flags["includeItemsFromAllDrives"].GetBool(), flags["includeRemoved"].GetBool(), flags["restrictToMyDrive"].GetBool())
		if err != nil {
			log.Fatalf("Error listing changes: %v", err)
		}
		type resultStruct struct {
			NextStartPageToken string          `json:"nextStartPageToken"`
			Changes            []*drive.Change `json:"changes,omitempty"`
		}
		result := resultStruct{
			Changes:            r,
			NextStartPageToken: nextStartPageToken,
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(changesCmd, changesListCmd, changeFlags)
}
