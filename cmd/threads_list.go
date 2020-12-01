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

// threadsListCmd represents the list command
var threadsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the threads in the user's mailbox.",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.threads/list",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmail.ListThreads(flags["userId"].GetString(), flags["q"].GetString(), flags["fields"].GetString(), flags["labelIds"].GetStringSlice(), flags["includeSpamTrash"].GetBool())
		if err != nil {
			log.Fatalf("Error listing thread: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(threadsCmd, threadsListCmd, threadFlags)
}
