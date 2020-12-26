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
	"strings"

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// messagesImportCmd represents the import command
var messagesImportCmd = &cobra.Command{
	Use: "import",
	Short: `Imports a message into only this user's mailbox, with standard email delivery scanning and classification similar to receiving via SMTP.
Does not send a message.`,
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.messages/import",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		internalDateSource := flags["internalDateSource"].GetString()
		internalDateSource = strings.ToUpper(internalDateSource)
		if !gsmgmail.InternalDateSourceIsValid(internalDateSource) {
			log.Fatalf("%s is not a valid value for internalDateSource", internalDateSource)
		}
		message, err := emlToMessage(flags["eml"].GetString())
		if err != nil {
			log.Fatalf("Error with eml file: %v", err)
		}
		result, err := gsmgmail.ImportMessage(flags["userId"].GetString(), internalDateSource, flags["fields"].GetString(), message, flags["deleted"].GetBool(), flags["neverMarkSpam"].GetBool(), flags["processForCalendar"].GetBool())
		if err != nil {
			log.Fatalf("Error importing message: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(messagesCmd, messagesImportCmd, messageFlags)
}
