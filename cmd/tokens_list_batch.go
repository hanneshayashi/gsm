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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	admin "google.golang.org/api/admin/directory/v1"
)

// tokensListBatchCmd represents the batch command
var tokensListBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch list access tokens issued by a user using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/tokens/list",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, tokenFlags, viper.GetInt("threads"))
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			UserKey string         `json:"userKey,omitempty"`
			Clients []*admin.Token `json:"clients,omitempty"`
		}
		results := make(chan resultStruct, cap)
		final := []resultStruct{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmadmin.ListTokens(m["userKey"].GetString(), m["fields"].GetString())
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{UserKey: m["userKey"].GetString(), Clients: result}
						}
						time.Sleep(200 * time.Millisecond)
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
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(tokensListCmd, tokensListBatchCmd, tokenFlags, tokenFlagsALL, batchFlags)
}
