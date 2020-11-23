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
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/gmail/v1"
)

// delegatesCreateBatchCmd represents the batch command
var delegatesCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch Adds a delegate with its verification status set directly to accepted, without sending any verification email using a CSV file as input.",
	Long: `The delegate user must be a member of the same Workspace organization as the delegator user.

Gmail imposes limitations on the number of delegates and delegators each user in a Workspace organization can have. These limits depend on your organization, but in general each user can have up to 25 delegates and up to 10 delegators.

Note that a delegate user must be referred to by their primary email address, and not an email alias.

Also note that when a new delegate is created, there may be up to a one minute delay before the new delegate is available for use.

https://developers.google.com/gmail/api/reference/rest/v1/users.settings.delegates/create`,
	Run: func(cmd *cobra.Command, args []string) {
		retrier := gsmhelpers.NewStandardRetrier()
		var wg sync.WaitGroup
		maps, err := gsmhelpers.GetBatchMaps(cmd, delegateFlags, viper.GetInt("threads"))
		cap := cap(maps)
		if err != nil {
			log.Fatalln(err)
		}
		results := make(chan *gmail.Delegate, cap)
		final := []*gmail.Delegate{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						var err error
						d, err := mapToDelegate(m)
						if err != nil {
							log.Printf("Error building delegate object: %v\n", err)
							continue
						}
						errKey := fmt.Sprintf("%s - %s:", m["userId"].GetString(), d.DelegateEmail)
						operation := func() error {
							result, err := gsmgmail.CreateDelegate(m["userId"].GetString(), m["fields"].GetString(), d)
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
	gsmhelpers.InitBatchCommand(delegatesCreateCmd, delegatesCreateBatchCmd, delegateFlags, delegateFlagsALL, batchFlags)
}
