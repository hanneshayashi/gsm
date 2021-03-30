/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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
	"sync"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// sharedContactsUpdateBatchCmd represents the batch command
var sharedContactsUpdateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch updates Domain Shared Contacts using a CSV file as input",
	Long:  `https://developers.google.com/admin-sdk/domain-shared-contacts`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, sharedContactFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *gsmadmin.Entry, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						url := m["url"].GetString()
						s, err := gsmadmin.GetSharedContact(url)
						if err != nil {
							log.Printf("Error getting shared contact: %v\n", err)
							continue
						}
						s, err = mapToSharedContact(m, s)
						if err != nil {
							log.Printf("Error building shared contact object: %v\n", err)
							continue
						}
						result, err := gsmadmin.UpdateSharedContact(url, s)
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
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for r := range results {
				err := enc.Encode(r)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*gsmadmin.Entry{}
			for res := range results {
				final = append(final, res)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	gsmhelpers.InitBatchCommand(sharedContactsUpdateCmd, sharedContactsUpdateBatchCmd, sharedContactFlags, sharedContactFlagsALL, batchFlags)
}
