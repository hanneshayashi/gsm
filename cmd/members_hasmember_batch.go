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

// membersHasMemberBatchCmd represents the batch command
var membersHasMemberBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Checks whether the given user is a member of the group. Membership can be direct or nested.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/members/hasMember",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, memberFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			GroupKey  string `json:"groupKey,omitempty"`
			MemberKey string `json:"memberKey,omitempty"`
			Result    bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		final := []resultStruct{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmadmin.HasMember(m["groupKey"].GetString(), m["memberKey"].GetString())
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{GroupKey: m["groupKey"].GetString(), MemberKey: m["memberKey"].GetString(), Result: result}
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
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(membersHasMemberCmd, membersHasMemberBatchCmd, memberFlags, memberFlagsALL, batchFlags)
}
