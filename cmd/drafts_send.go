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
	"gsm/gsmgmail"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// draftsSendCmd represents the send command
var draftsSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends the specified, existing draft to the recipients in the To, Cc, and Bcc headers.",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.drafts/send",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		draft, err := gsmgmail.GetDraft(flags["userId"].GetString(), flags["id"].GetString(), "FULL", flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting draft: %v", err)
		}
		result, err := gsmgmail.SendDraft(flags["userId"].GetString(), draft)
		if err != nil {
			log.Fatalf("Error sending draft: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	draftsCmd.AddCommand(draftsSendCmd)
	gsmhelpers.AddFlags(draftFlags, draftsSendCmd.Flags(), draftsSendCmd.Use)
	markFlagsRequired(draftsSendCmd, draftFlags, "")
}
