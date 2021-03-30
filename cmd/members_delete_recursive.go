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

// membersDeleteRecursiveCmd represents the recursive command
var membersDeleteRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Removes users from a group by referencing one or more organizational units and/or groups.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/members/insert",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		type resultStruct struct {
			MemberKey string `json:"memberKey,omitempty"`
			Result    bool   `json:"result"`
		}
		results := make(chan resultStruct, threads)
		var wg sync.WaitGroup
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		groupKey := flags["groupKey"].GetString()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for uk := range userKeysUnique {
						result, err := gsmadmin.DeleteMember(groupKey, uk)
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{MemberKey: uk, Result: result}
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
			final := []resultStruct{}
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
	gsmhelpers.InitRecursiveCommand(membersDeleteCmd, membersDeleteRecursiveCmd, memberFlags, recursiveUserFlags)
}
