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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
)

// membersDeleteRecursiveCmd represents the recursive command
var membersDeleteRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Removes users from a group by referencing one or more organizational units and/or groups.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/members/insert",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		type resultStruct struct {
			MemberKey string `json:"memberKey,omitempty"`
			Result    bool   `json:"result"`
		}
		final := []resultStruct{}
		finalChan := make(chan resultStruct, threads)
		wgOps := &sync.WaitGroup{}
		wgFinal := &sync.WaitGroup{}
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		groupKey := flags["groupKey"].GetString()
		for i := 0; i < threads; i++ {
			wgOps.Add(1)
			go func() {
				for uk := range userKeysUnique {
					result, err := gsmadmin.DeleteMember(groupKey, uk)
					if err != nil {
						log.Println(err)
					} else {
						finalChan <- resultStruct{MemberKey: uk, Result: result}
					}
				}
				wgOps.Done()
			}()
		}
		wgFinal.Add(1)
		go func() {
			for r := range finalChan {
				final = append(final, r)
			}
			wgFinal.Done()
		}()
		wgOps.Wait()
		close(finalChan)
		wgFinal.Wait()
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(membersDeleteCmd, membersDeleteRecursiveCmd, memberFlags, recursiveUserFlags)
}