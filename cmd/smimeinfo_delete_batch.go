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
	"gsm/gsmgmail"
	"gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
)

// smimeInfoDeleteBatchCmd represents the batch command
var smimeInfoDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch deletes the specified S/MIME config for the specified send-as aliases using a CSV file as input.",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.settings.sendAs.smimeInfo/delete",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, smimeInfoFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			UserID      string `json:"userId,omitempty"`
			SendAsEmail string `json:"sendAsEmail,omitempty"`
			ID          string `json:"id,omitempty"`
			Result      bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		final := []resultStruct{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmgmail.DeleteSmimeInfo(m["userId"].GetString(), m["sendAsEmail"].GetString(), m["id"].GetString())
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{ID: m["id"].GetString(), SendAsEmail: m["sendAsEmail"].GetString(), UserID: m["userId"].GetString(), Result: result}
					}
					wg.Done()
				}()
			}
			wg.Wait()
			close(results)
		}()
		for res := range results {
			final = append(final, res)
		}
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(smimeInfoDeleteCmd, smimeInfoDeleteBatchCmd, smimeInfoFlags, smimeInfoFlagsALL, batchFlags)
}
