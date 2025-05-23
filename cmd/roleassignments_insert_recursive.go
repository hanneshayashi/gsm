/*
Copyright © 2020-2023 Hannes Hayashi

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

// roleAssignmentsInsertRecursiveCmd represents the recursive command
var roleAssignmentsInsertRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Creates role assignments for users by referencing one or more organizational units and/or groups.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/admin/directory/reference/rest/v1/roleAssignments/insert",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		type resultStruct struct {
			RoleAssignment *admin.RoleAssignment `json:"roleAssignment,omitempty"`
			UserKey        string                `json:"userKey,omitempty"`
		}
		results := make(chan resultStruct, threads)
		userIdsUnique := make(chan string, threads)
		var wgUserIds sync.WaitGroup
		var wg sync.WaitGroup
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		customer := flags["customer"].GetString()
		fields := flags["fields"].GetString()
		go func() {
			for i := 0; i < threads; i++ {
				wgUserIds.Add(1)
				go func() {
					for uk := range userKeysUnique {
						u, err := gsmadmin.GetUser(uk, "id", "", "", "")
						if err != nil {
							log.Println(err)
						} else {
							userIdsUnique <- u.Id
						}
					}
					wgUserIds.Done()
				}()
			}
			wgUserIds.Wait()
			close(userIdsUnique)
		}()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
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
							results <- resultStruct{UserKey: uid, RoleAssignment: result}
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
	gsmhelpers.InitRecursiveCommand(roleAssignmentsInsertCmd, roleAssignmentsInsertRecursiveCmd, roleAssignmentFlags, recursiveUserFlags)
}
