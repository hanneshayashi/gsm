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
	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
)

// sharedContactsCreateBatchCmd represents the batch command
var sharedContactsCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch Create a Domain Shared Contact",
	Long:  `https://developers.google.com/admin-sdk/domain-shared-contacts`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, sharedContactFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *gsmadmin.Entry, cap)
		final := []*gsmadmin.Entry{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						s, err := mapToSharedContact(m, nil)
						if err != nil {
							log.Printf("Error building shared contact object: %v\n", err)
							continue
						}
						result, err := gsmadmin.CreateSharedContact(m["domain"].GetString(), s)
						if err != nil {
							log.Println(err)
						} else {
							results <- result
						}
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
	gsmhelpers.InitBatchCommand(sharedContactsCreateCmd, sharedContactsCreateBatchCmd, sharedContactFlags, sharedContactFlagsALL, batchFlags)
}
