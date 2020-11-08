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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// sharedContactsCreateBatchCmd represents the batch command
var sharedContactsCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch Create a Domain Shared Contact",
	Long: `https://developers.google.com/admin-sdk/domain-shared-contacts
Example: gsm sharedContacts create --domain "example.org" --givenName "Jack" --familyName "Bauer" --email "displayName=Jack Bauer;address=jack@ctu.gov;primary=false" --email "displayName=Jack bauer;address=jack.bauer@ctu.gov;primary=true" --phoneNumber "phoneNumber=+49 127 12381;primary=true;label=Work" --phoneNumber "phoneNumber=+49 21891238;primary=false;label=Home" --organization "orgName=Counter Terrorist Unit;orgDepartment=Field Agents;orgTitle=Special Agent"`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		l := len(csv)
		results := make(chan *gsmadmin.Entry, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []*gsmadmin.Entry{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, sharedContactFlags, line, "create")
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
					s, err := mapToSharedContact(m, nil)
					if err != nil {
						log.Printf("Error building shared contact object: %v\n", err)
						continue
					}
					errKey := fmt.Sprintf("%s:", m["domain"].GetString())
					operation := func() error {
						result, statusCode, err := gsmadmin.CreateSharedContact(m["domain"].GetString(), s)
						if err != nil {
							if statusCode == 403 {
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
	sharedContactsCreateCmd.AddCommand(sharedContactsCreateBatchCmd)
	flags := sharedContactsCreateBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(sharedContactFlags, flags, "create")
	markFlagsRequired(sharedContactsCreateBatchCmd, sharedContactFlags, "create")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(sharedContactsCreateBatchCmd, batchFlags, "")
}
