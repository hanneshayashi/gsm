/*
Copyright © 2020-2024 Hannes Hayashi

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

// filesModifyLabelsCmd represents the modifyLabels command
var filesModifyLabelsCmd = &cobra.Command{
	Use:               "modifyLabels",
	Short:             "Modifies the set of labels on a file.",
	Long:              "Implements the API documented athttps://developers.google.com/drive/api/v3/reference/files/modifyLabels",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToModifyLabelsRequest(flags)
		if err != nil {
			log.Fatalf("Error building modify labels request: %v", err)
		}
		result, err := gsmdrive.ModifyLabels(flags["fileId"].GetString(), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error modifying labels: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesModifyLabelsCmd, fileFlags)
}
