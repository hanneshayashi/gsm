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
	"google.golang.org/api/drive/v3"

	"github.com/spf13/cobra"
)

// filesListCmd represents the list command
var filesListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists or searches files.",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/files/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrive.ListFiles(flags["q"].GetString(), flags["driveId"].GetString(), flags["corpora"].GetString(), flags["includePermissionsForView"].GetString(), flags["orderBy"].GetString(), flags["spaces"].GetString(), flags["fields"].GetString(), flags["includeItemsFromAllDrives"].GetBool(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*drive.File{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing files: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesListCmd, fileFlags)
}
