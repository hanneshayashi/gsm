/*
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
	admin "google.golang.org/api/admin/directory/v1"
)

// membersPatchRecursiveCmd represents the recursive command
var membersPatchRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Updates the membership properties of users by referencing one or more organizational units and/or groups.",
	Long:  "Implements the API documented at https://developers.google.com/admin-sdk/directory/v1/reference/members/patch",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		results := make(chan *admin.Member, threads)
		var wg sync.WaitGroup
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		groupKey := flags["groupKey"].GetString()
		fields := flags["fields"].GetString()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for uk := range userKeysUnique {
						m, err := mapToMember(flags)
						if err != nil {
							log.Printf("Error building member object: %v\n", err)
							continue
						}
						m.Email = uk
						result, err := gsmadmin.PatchMember(groupKey, uk, fields, m)
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
			final := []*admin.Member{}
			for r := range results {
				final = append(final, r)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(membersPatchCmd, membersPatchRecursiveCmd, memberFlags, recursiveUserFlags)
}
