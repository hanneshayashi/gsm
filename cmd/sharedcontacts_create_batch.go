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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sharedContactsCreateBatchCmd represents the batch command
var sharedContactsCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch Create a Domain Shared Contact",
	Long: `https://developers.google.com/admin-sdk/domain-shared-contacts
Example: gsm sharedContacts create --domain "example.org" --givenName "Jack" --familyName "Bauer" --email "displayName=Jack Bauer;address=jack@ctu.gov;primary=false" --email "displayName=Jack bauer;address=jack.bauer@ctu.gov;primary=true" --phoneNumber "phoneNumber=+49 127 12381;primary=true;label=Work" --phoneNumber "phoneNumber=+49 21891238;primary=false;label=Home" --organization "orgName=Counter Terrorist Unit;orgDepartment=Field Agents;orgTitle=Special Agent"`,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, sharedContactFlags, viper.GetInt("threads"))
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
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(sharedContactsCreateCmd, sharedContactsCreateBatchCmd, sharedContactFlags, sharedContactFlagsALL, batchFlags)
}
