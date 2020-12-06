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
	"gsm/gsmhelpers"
	"gsm/gsmpeople"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/api/people/v1"
)

// peopleUpdateContactBatchCmd represents the batch command
var peopleUpdateContactBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch update contacts using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/people/updateContact",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, peopleFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *people.Person, cap)
		final := []*people.Person{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						p, err := gsmpeople.GetContact(m["resourceName"].GetString(), m["personFields"].GetString(), m["sources"].GetString(), "*")
						if err != nil {
							log.Printf("Error getting contact: %v\n", err)
							continue
						}
						p, err = mapToPerson(m, p)
						if err != nil {
							log.Printf("Error building person object: %v\n", err)
							continue
						}
						result, err := gsmpeople.UpdateContact(m["resourceName"].GetString(), m["updatePersonFields"].GetString(), m["personFields"].GetString(), m["sources"].GetString(), m["fields"].GetString(), p)
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
	gsmhelpers.InitBatchCommand(peopleUpdateContactCmd, peopleUpdateContactBatchCmd, peopleFlags, peopleFlagsALL, batchFlags)
}
