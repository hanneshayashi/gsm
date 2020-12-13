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
	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// draftsUpdateBatchCmd represents the batch command
var draftsUpdateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch updates drafts using a CSV file as input.",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.drafts/update",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, draftFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *gmail.Draft, cap)
		final := []*gmail.Draft{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						d, err := mapToDraft(m)
						if err != nil {
							log.Printf("Error building draft object: %v\n", err)
							continue
						}
						result, err := gsmgmail.UpdateDraft(m["userId"].GetString(), m["id"].GetString(), m["fields"].GetString(), d)
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
	gsmhelpers.InitBatchCommand(draftsUpdateCmd, draftsUpdateBatchCmd, draftFlags, draftFlagsALL, batchFlags)
}
