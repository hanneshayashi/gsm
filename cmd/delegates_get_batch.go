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

// delegatesGetBatchCmd represents the batch command
var delegatesGetBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch gets the specified delegates using a CSV file as input.",
	Long: `Note that a delegate user must be referred to by their primary email address, and not an email alias.
	https://developers.google.com/gmail/api/reference/rest/v1/users.settings.delegates/get`,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, delegateFlags, viper.GetInt("threads"))
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *gmail.Delegate, cap)
		final := []*gmail.Delegate{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmgmail.GetDelegate(m["userId"].GetString(), m["delegateEmail"].GetString(), m["fields"].GetString())
						if err != nil {
							log.Println(err)
						} else {
							results <- result
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
	gsmhelpers.InitBatchCommand(delegatesGetCmd, delegatesGetBatchCmd, delegateFlags, delegateFlagsALL, batchFlags)
}
