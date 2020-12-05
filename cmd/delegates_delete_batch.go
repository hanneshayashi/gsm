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

	"github.com/spf13/cobra"
)

// delegatesDeleteBatchCmd represents the batch command
var delegatesDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch deletes the specified delegates using a CSV file as input.",
	Long: `Note that a delegate user must be referred to by their primary email address, and not an email alias.
	https://developers.google.com/gmail/api/reference/rest/v1/users.settings.delegates/delete`,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, delegateFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			UserID        string `json:"userId,omitempty"`
			DelegateEmail string `json:"delegateEmail,omitempty"`
			Result        bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		final := []resultStruct{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmgmail.DeleteDelegate(m["userId"].GetString(), m["delegateEmail"].GetString())
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{UserID: m["userId"].GetString(), DelegateEmail: m["delegateEmail"].GetString(), Result: result}
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
	gsmhelpers.InitBatchCommand(delegatesDeleteCmd, delegatesDeleteBatchCmd, delegateFlags, delegateFlagsALL, batchFlags)
}
