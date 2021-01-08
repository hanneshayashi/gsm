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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/gmail/v1"

	"github.com/spf13/cobra"
)

// messagesListCmd represents the list command
var messagesListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists the messages in the user's mailbox.",
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.messages/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmail.ListMessages(flags["userId"].GetString(), flags["q"].GetString(), flags["fields"].GetString(), flags["labelIds"].GetStringSlice(), flags["includeSpamTrash"].GetBool(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*gmail.Message{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing messages: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(messagesCmd, messagesListCmd, messageFlags)
}
