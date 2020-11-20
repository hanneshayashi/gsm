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
	"google.golang.org/api/drive/v3"
)

// changesListCmd represents the list command
var changesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the changes for a user or shared drive.",
	Long:  "https://developers.google.com/drive/api/v3/reference/changes/list",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, nextStartPageToken, err := gsmdrive.ListChanges(flags["pageToken"].GetString(), flags["driveId"].GetString(), flags["spaces"].GetString(), flags["fields"].GetString(), flags["includePermissionsForView"].GetString(), flags["includeCorpusRemovals"].GetBool(), flags["includeItemsFromAllDrives"].GetBool(), flags["includeRemoved"].GetBool(), flags["restrictToMyDrive"].GetBool())
		if err != nil {
			log.Fatalf("Error listing changes: %v", err)
		}
		type resultStruct struct {
			Changes            []*drive.Change `json:"changes,omitempty"`
			NextStartPageToken string          `json:"nextStartPageToken"`
		}
		r := resultStruct{
			Changes:            result,
			NextStartPageToken: nextStartPageToken,
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(r, "json"))
	},
}

func init() {
	gsmhelpers.InitCommand(changesCmd, changesListCmd, changeFlags)
}
