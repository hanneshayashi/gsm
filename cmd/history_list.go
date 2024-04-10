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
	"strings"

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/gmail/v1"

	"github.com/spf13/cobra"
)

// historyListCmd represents the list command
var historyListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists the history of all changes to the given mailbox. History results are returned in chronological order (increasing historyId).",
	Long:              "Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users.history/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		historyTypes := flags["historyTypes"].GetStringSlice()
		for i := range historyTypes {
			historyTypes[i] = strings.ToUpper(historyTypes[i])
			if !gsmgmail.HistoryTypeIsValid(historyTypes[i]) {
				log.Fatalf("%s is not a valid history type", historyTypes[i])
			}
		}
		result, err := gsmgmail.ListHistory(flags["userId"].GetString(), flags["labelId"].GetString(), flags["fields"].GetString(), flags["startHistoryId"].GetUint64(), historyTypes, gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*gmail.History{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing history: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(historyCmd, historyListCmd, historyFlags)
}
