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
	"google.golang.org/api/people/v1"
)

// peopleUpdateContactBatchCmd represents the batch command
var peopleUpdateContactBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch update contacts using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/people/updateContact",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		err = gsmhelpers.CheckBatchFlags(flags, peopleFlags, int64(len(csv[0])))
		if err != nil {
			log.Fatalf("Error with batch flag index: %v\n", err)
		}
		l := len(csv)
		results := make(chan *people.Person, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []*people.Person{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, peopleFlags, line, "updateContact")
				maps <- m
			}
			close(maps)
			wg1.Done()
		}()
		wg2.Add(1)
		retrier := gsmhelpers.NewStandardRetrier()
		for i := 0; i < gsmhelpers.MaxThreads(l); i++ {
			wg2.Add(1)
			go func() {
				for m := range maps {
					var err error
					errKey := fmt.Sprintf("%s:", m["resourceName"].GetString())
					operation := func() error {
						p, err := gsmpeople.GetContact(m["resourceName"].GetString(), m["personFields"].GetString(), m["sources"].GetString(), "*")
						if err != nil {
							retryable := gsmhelpers.ErrorIsRetryable(err)
							if retryable {
								log.Println(errKey, "Retrying after", err)
								return err
							}
							log.Println(errKey, "Giving up after", err)
							return nil
						}
						p, err = mapToPerson(m, p)
						if err != nil {
							log.Println(err)
							return nil
						}
						result, err := gsmpeople.UpdateContact(m["resourceName"].GetString(), m["updatePersonFields"].GetString(), m["personFields"].GetString(), m["sources"].GetString(), m["fields"].GetString(), p)
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
				wg2.Done()
			}()
		}
		wg3.Add(1)
		go func() {
			for res := range results {
				final = append(final, res)
			}
			wg3.Done()
		}()
		wg2.Done()
		wg1.Wait()
		wg2.Wait()
		close(results)
		wg3.Wait()
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json"))
	},
}

func init() {
	peopleUpdateContactCmd.AddCommand(peopleUpdateContactBatchCmd)
	flags := peopleUpdateContactBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(peopleFlags, flags, "updateContact")
	markFlagsRequired(peopleUpdateContactBatchCmd, peopleFlags, "updateContact")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(peopleUpdateContactBatchCmd, batchFlags, "")
}
