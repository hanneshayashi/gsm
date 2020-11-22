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
	"gsm/gsmhelpers"
	"gsm/gsmpeople"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/people/v1"
)

// contactGroupsCreateBatchCmd represents the batch command
var contactGroupsCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch creates contact groups using a CSV file as input.",
	Long:  "https://developers.google.com/people/api/rest/v1/contactGroups/create",
	Run: func(cmd *cobra.Command, args []string) {
		retrier := gsmhelpers.NewStandardRetrier()
		var wg sync.WaitGroup
		maps, err := gsmhelpers.GetBatchMaps(cmd, contactGroupFlags, viper.GetInt("threads"))
		cap := cap(maps)
		if err != nil {
			log.Fatalln(err)
		}
		results := make(chan *people.ContactGroup, cap)
		final := []*people.ContactGroup{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						var err error
						c, err := mapToCreateContactGroupRequest(m)
						if err != nil {
							log.Printf("Error building createContactGroupRequest object: %v\n", err)
							continue
						}
						errKey := fmt.Sprintf("%s:", c.ContactGroup.Name)
						operation := func() error {
							result, err := gsmpeople.CreateContactGroup(c, m["fields"].GetString())
							if err != nil {
								retryable := gsmhelpers.ErrorIsRetryable(err)
								if retryable {
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
	gsmhelpers.InitBatchCommand(contactGroupsCreateCmd, contactGroupsCreateBatchCmd, contactGroupFlags, contactGroupFlagsALL, batchFlags)
}
