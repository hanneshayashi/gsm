/*
Copyright © 2020-2021 Hannes Hayashi

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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// roleAssignmentsListRecursiveCmd represents the recursive command
var roleAssignmentsListRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "List users' role assignments by referencing one or more organizational units and/or groups.",
	Long:  "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/roleAssignments/get",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		type resultStruct struct {
			UserKey         string                  `json:"userKey,omitempty"`
			RoleAssignments []*admin.RoleAssignment `json:"roleAssignments,omitempty"`
		}
		results := make(chan resultStruct, threads)
		var wg sync.WaitGroup
		userKeysUnique, err := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		customer := flags["customer"].GetString()
		fields := flags["fields"].GetString()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for uk := range userKeysUnique {
						result, er := gsmadmin.ListRoleAssignments(customer, "", uk, fields, threads)
						r := resultStruct{UserKey: uk}
						for i := range result {
							r.RoleAssignments = append(r.RoleAssignments, i)
						}
						e := <-er
						if e != nil {
							log.Println(e)
						} else {
							results <- r
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
		e := <-err
		if e != nil {
			log.Fatalf("Error listing role assignments: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(roleAssignmentsListCmd, roleAssignmentsListRecursiveCmd, roleAssignmentFlags, recursiveUserFlags)
}
