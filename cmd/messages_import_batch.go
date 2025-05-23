/*
Copyright © 2020-2023 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// messagesImportBatchCmd represents the batch command
var messagesImportBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch imports messages using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/gmail/api/reference/rest/v1/users.messages/import",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, messageFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *gmail.Message, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						message, err := emlToMessage(m["eml"].GetString())
						if err != nil {
							log.Printf("Error with eml file: %v\n", err)
							continue
						}
						result, err := gsmgmail.ImportMessage(m["userId"].GetString(), m["internalDateSource"].GetString(), m["fields"].GetString(), message, m["deleted"].GetBool(), m["neverMarkSpam"].GetBool(), m["processForCalendar"].GetBool())
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
			final := []*gmail.Message{}
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
	gsmhelpers.InitBatchCommand(messagesImportCmd, messagesImportBatchCmd, messageFlags, messageFlagsALL, batchFlags)
}
