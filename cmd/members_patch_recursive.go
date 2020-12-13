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
	admin "google.golang.org/api/admin/directory/v1"
)

// membersPatchRecursiveCmd represents the recursive command
var membersPatchRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Updates the membership properties of users by referencing one or more organizational units and/or groups.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/members/patch",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		finalChan := make(chan *admin.Member, threads)
		final := []*admin.Member{}
		wgOps := &sync.WaitGroup{}
		wgFinal := &sync.WaitGroup{}
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		groupKey := flags["groupKey"].GetString()
		fields := flags["fields"].GetString()
		for i := 0; i < threads; i++ {
			wgOps.Add(1)
			go func() {
				for uk := range userKeysUnique {
					m, err := mapToMember(flags)
					if err != nil {
						log.Printf("Error building user object: %v\n", err)
						continue
					}
					m.Email = uk
					result, err := gsmadmin.PatchMember(groupKey, uk, fields, m)
					if err != nil {
						log.Println(err)
					} else {
						finalChan <- result
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
	gsmhelpers.InitRecursiveCommand(membersPatchCmd, membersPatchRecursiveCmd, memberFlags, recursiveUserFlags)
}
