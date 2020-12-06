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
	"gsm/gsmcalendar"
	"gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// calendarACLInsertBatchCmd represents the batch command
var calendarACLInsertBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch inserts ACL rules using a CSV file as input.",
	Long:  `https://developers.google.com/calendar/v3/reference/acl/insert`,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, calendarACLFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *calendar.AclRule, cap)
		final := []*calendar.AclRule{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						a, err := mapToCalendarACLRule(m)
						if err != nil {
							log.Printf("Error building acl rule object: %v\n", err)
							continue
						}
						result, err := gsmcalendar.InsertACL(m["calendarId"].GetString(), m["fields"].GetString(), a, m["sendNotifications"].GetBool())
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
		for res := range results {
			final = append(final, res)
		}
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(calendarACLInsertCmd, calendarACLInsertBatchCmd, calendarACLFlags, calendarACLFlagsALL, batchFlags)
}
