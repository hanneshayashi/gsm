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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// messagesDeleteCmd represents the messagesDelete command
var messagesDeleteCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Immediately and permanently deletes the specified message. This operation cannot be undone. Prefer messages.trash instead.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/gmail/api/reference/rest/v1/users.messages/delete",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmail.DeleteMessage(flags["userId"].GetString(), flags["id"].GetString())
		if err != nil {
			log.Fatalf("Error deleting message: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(messagesCmd, messagesDeleteCmd, messageFlags)
}
