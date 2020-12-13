/*
Package cmd contains the commands available to the end user
Copyright © 2020 Hannes Hayashi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public Role as published by
the Free Software Foundation, either version 3 of the Role, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public Role for more details.

You should have received a copy of the GNU General Public Role
along with this program. If not, see <http://www.gnu.org/roles/>.
*/
package cmd

import (
	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// roleAssignmentsInsertRecursiveCmd represents the recursive command
var roleAssignmentsInsertRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Creates role assignments for users by referencing one or more organizational units and/or groups.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/roleassignments/get",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},	
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		type resultStruct struct {
			UserKey        string                `json:"userKey,omitempty"`
			RoleAssignment *admin.RoleAssignment `json:"roleAssignment,omitempty"`
		}
		finalChan := make(chan resultStruct, threads)
		final := []resultStruct{}
		userIdsUnique := make(chan string, threads)
		wgGets := &sync.WaitGroup{}
		wgOps := &sync.WaitGroup{}
		wgFinal := &sync.WaitGroup{}
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		customer := flags["customer"].GetString()
		fields := flags["fields"].GetString()
		for i := 0; i < threads; i++ {
			wgGets.Add(1)
			go func() {
				for uk := range userKeysUnique {
					u, err := gsmadmin.GetUser(uk, "id", "", "", "")
					if err != nil {
						log.Println(err)
					} else {
						userIdsUnique <- u.Id
					}
				}
				wgGets.Done()
			}()
		}
		for i := 0; i < threads; i++ {
			wgOps.Add(1)
			go func() {
				for uid := range userIdsUnique {
					r, err := mapToRoleAssignment(flags)
					if err != nil {
						log.Fatalf("Error building role assignment object: %v", err)
					}
					r.AssignedTo = uid
					result, err := gsmadmin.InsertRoleAssignment(customer, fields, r)
					if err != nil {
						log.Println(err)
					} else {
						finalChan <- resultStruct{UserKey: uid, RoleAssignment: result}
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
		wgGets.Wait()
		close(userIdsUnique)
		wgOps.Wait()
		close(finalChan)
		wgFinal.Wait()
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(roleAssignmentsInsertCmd, roleAssignmentsInsertRecursiveCmd, roleAssignmentFlags, recursiveUserFlags)
}
