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

// sharedContactsUpdateBatchCmd represents the batch command
var sharedContactsUpdateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch updates Domain Shared Contacts using a CSV file as input",
	Long:  `https://developers.google.com/admin-sdk/domain-shared-contacts`,
	Run: func(cmd *cobra.Command, args []string) {
		retrier := gsmhelpers.NewStandardRetrier()
		var wg sync.WaitGroup
		maps, err := gsmhelpers.GetBatchMaps(cmd, sharedContactFlags, batchThreads)
		if err != nil {
			log.Fatalln(err)
		}
		results := make(chan *gsmadmin.Entry, batchThreads)
		final := []*gsmadmin.Entry{}
		go func() {
			for i := 0; i < batchThreads; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						var err error
						errKey := fmt.Sprintf("%s:", m["url"].GetString())
						operation := func() error {
							s, _, err := gsmadmin.GetSharedContact(m["url"].GetString())
							if err != nil {
								log.Printf("Error getting shared contact: %v\n", err)
								return nil
							}
							s, err = mapToSharedContact(m, s)
							if err != nil {
								log.Printf("Error building shared contact object: %v\n", err)
								return nil
							}
							result, statusCode, err := gsmadmin.UpdateSharedContact(m["url"].GetString(), s)
							if err != nil {
								if statusCode == 403 {
									log.Println(errKey, "Retrying after", err)
									return err
								}
								log.Println(errKey, "Giving up after", err)
								return nil
							}
							results <- result
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
	gsmhelpers.InitBatchCommand(sharedContactsUpdateCmd, sharedContactsUpdateBatchCmd, sharedContactFlags, sharedContactFlagsALL, batchFlags)
}
