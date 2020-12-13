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

// messagesInsertCmd represents the insert command
var messagesInsertCmd = &cobra.Command{
	Use: "insert",
	Short: `Directly inserts a message into only this user's mailbox similar to IMAP APPEND, bypassing most scanning and classification.
Does not send a message.`,
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.messages/insert",
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
		result, err := gsmgmail.InsertMessage(flags["userId"].GetString(), internalDateSource, flags["fields"].GetString(), message, flags["deleted"].GetBool())
		if err != nil {
			log.Fatalf("Error inserting message: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(messagesCmd, messagesInsertCmd, messageFlags)
}
