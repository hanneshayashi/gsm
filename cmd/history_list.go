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
	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// historyListCmd represents the list command
var historyListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the history of all changes to the given mailbox. History results are returned in chronological order (increasing historyId).",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.history/list",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		historyTypes := flags["historyTypes"].GetStringSlice()
		for i := range historyTypes {
			historyTypes[i] = strings.ToUpper(historyTypes[i])
			if !gsmgmail.HistoryTypeIsValid(historyTypes[i]) {
				log.Fatalf("%s is not a valid history type", historyTypes[i])
			}
		}
		result, err := gsmgmail.ListHistory(flags["userId"].GetString(), flags["labelId"].GetString(), flags["fields"].GetString(), flags["startHistoryId"].GetUint64(), historyTypes...)
		if err != nil {
			log.Fatalf("Error listing history: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(historyCmd, historyListCmd, historyFlags)
}
