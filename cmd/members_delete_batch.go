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
)

// membersDeleteBatchCmd represents the batch command
var membersDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch deletes group members using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/members/delete",
	Run: func(cmd *cobra.Command, args []string) {
		retrier := gsmhelpers.NewStandardRetrier()
		var wg sync.WaitGroup
		maps, err := gsmhelpers.GetBatchMaps(cmd, memberFlags, batchThreads)
		if err != nil {
			log.Fatalln(err)
		}
		type resultStruct struct {
			GroupKey  string `json:"groupKey,omitempty"`
			MemberKey string `json:"memberKey,omitempty"`
			Result    bool   `json:"result"`
		}
		results := make(chan resultStruct, batchThreads)
		final := []resultStruct{}
		go func() {
			for i := 0; i < batchThreads; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						var err error
						errKey := fmt.Sprintf("%s - %s:", m["groupKey"].GetString(), m["memberKey"].GetString())
						operation := func() error {
							result, err := gsmadmin.DeleteMember(m["groupKey"].GetString(), m["memberKey"].GetString())
							if err != nil {
								retryable := gsmhelpers.ErrorIsRetryable(err)
								if retryable {
									log.Println(errKey, "Retrying after", err)
									return err
								}
								log.Println(errKey, "Giving up after", err)
								return nil
							}
							results <- resultStruct{GroupKey: m["groupKey"].GetString(), MemberKey: m["memberKey"].GetString(), Result: result}
							return nil
						}
						err = retrier.Run(operation)
						if err != nil {
							log.Println(errKey, "Max retries reached. Giving up after", err)
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
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json"))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(membersDeleteCmd, membersDeleteBatchCmd, memberFlags, memberFlagsALL, batchFlags)
}
