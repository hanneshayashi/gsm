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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// sendAsDeleteCmd represents the delete command
var sendAsDeleteCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Deletes the specified send-as alias. Revokes any verification that may have been required for using it.",
	Long:              "Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users.settings.sendAs/delete",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmail.DeleteSendAs(flags["userId"].GetString(), flags["sendAsEmail"].GetString())
		if err != nil {
			log.Fatalf("Error deleting send-as: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(sendAsCmd, sendAsDeleteCmd, sendAsFlags)
}
