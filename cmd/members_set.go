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
	admin "google.golang.org/api/admin/directory/v1"

	"github.com/spf13/cobra"
)

// membersSetCmd represents the set command
var membersSetCmd = &cobra.Command{
	Use:               "set",
	Short:             "Sets the members of a group to match the specified email addresses with the given role",
	Long:              "",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		groupKey := flags["groupKey"].GetString()
		membersToAdd, membersToRemove, err := gsmadmin.GetMembersToSet(groupKey, 4, flags["emails"].GetStringSlice()...)
		if err != nil {
			log.Fatalf("Error getting members to set: %v", err)
		}
		var addedMembers []*admin.Member
		type removed struct {
			Email  string
			Result bool
		}
		var removedMembers []removed
		type resultStruct struct {
			Added   []*admin.Member `json:"added,omitempty"`
			Removed []removed       `json:"removed,omitempty"`
		}
		fields := flags["fields"].GetString()
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for uk := range membersToAdd {
				m, er := mapToMember(flags)
				if er != nil {
					log.Printf("Error building member object: %v\n", er)
					continue
				}
				m.Email = uk
				result, er := gsmadmin.InsertMember(groupKey, fields, m)
				if err != nil {
					log.Println(er)
				} else {
					addedMembers = append(addedMembers, result)
				}
			}
			wg.Done()
		}()
		go func() {
			for uk := range membersToRemove {
				result, er := gsmadmin.DeleteMember(groupKey, uk)
				if er != nil {
					log.Println(er)
				}
				removedMembers = append(removedMembers, removed{Email: uk, Result: result})
			}
			wg.Done()
		}()
		wg.Wait()
		result := resultStruct{
			Added:   addedMembers,
			Removed: removedMembers,
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(membersCmd, membersSetCmd, memberFlags)
}
