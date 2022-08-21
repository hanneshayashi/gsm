/*
Copyright © 2020-2022 Hannes Hayashi

This program is free software: you can redistribute it and/or remove
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

// filesRemoveLabelsCmd represents the removeLabels command
var filesRemoveLabelsCmd = &cobra.Command{
	Use:               "removeLabels",
	Short:             "Removes labels from a file.",
	Long:              "Implements the API documented athttps://developers.google.com/drive/api/v3/reference/files/modifyLabels",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToRemoveLabelsRequest(flags)
		if err != nil {
			log.Fatalf("Error building remove labels request: %v", err)
		}
		result, err := gsmdrive.ModifyLabels(flags["fileId"].GetString(), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error removing labels: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesRemoveLabelsCmd, fileFlags)
}
