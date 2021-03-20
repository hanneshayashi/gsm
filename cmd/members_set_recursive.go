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
	admin "google.golang.org/api/admin/directory/v1"
)

// membersSetRecursiveCmd represents the recursive command
var membersSetRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Sets the memberships of the group by referencing one or more organizational units and/or groups.",
	Long:  "",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		addedMembers := make(chan *admin.Member, threads)
		type removed struct {
			Email  string
			Result bool
		}
		removedMembers := make(chan removed, threads)
		type resultStruct struct {
			Added   []*admin.Member `json:"added,omitempty"`
			Removed []removed       `json:"removed,omitempty"`
		}
		var wgResults sync.WaitGroup
		var wgFinal sync.WaitGroup
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		var emails []string
		for uk := range userKeysUnique {
			emails = append(emails, uk)
		}
		groupKey := flags["groupKey"].GetString()
		membersToAdd, membersToRemove, err := gsmadmin.GetMembersToSet(groupKey, threads, emails...)
		if err != nil {
			log.Fatalf("Error getting members to set: %v", err)
		}
		fields := flags["fields"].GetString()
		go func() {
			for i := 0; i < threads/2; i++ {
				wgResults.Add(1)
				go func() {
					for uk := range membersToAdd {
						m, err := mapToMember(flags)
						if err != nil {
							log.Printf("Error building member object: %v\n", err)
							continue
						}
						m.Email = uk
						result, err := gsmadmin.InsertMember(groupKey, fields, m)
						if err != nil {
							log.Println(err)
						} else {
							addedMembers <- result
						}
					}
					wgResults.Done()
				}()
			}
			wgResults.Wait()
			close(addedMembers)
		}()
		go func() {
			for i := 0; i < threads/2; i++ {
				wgResults.Add(1)
				go func() {
					for uk := range membersToRemove {
						result, err := gsmadmin.DeleteMember(groupKey, uk)
						if err != nil {
							log.Println(err)
						}
						removedMembers <- removed{Email: uk, Result: result}
					}
					wgResults.Done()
				}()
			}
			wgResults.Wait()
			close(removedMembers)
		}()
		final := resultStruct{}
		wgFinal.Add(2)
		go func() {
			for a := range addedMembers {
				final.Added = append(final.Added, a)
			}
			wgFinal.Done()
		}()
		go func() {
			for r := range removedMembers {
				final.Removed = append(final.Removed, r)
			}
			wgFinal.Done()
		}()
		wgFinal.Wait()
		gsmhelpers.Output(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(membersSetCmd, membersSetRecursiveCmd, memberFlags, recursiveUserFlags)
}
