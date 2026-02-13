/*
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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// calendarACLDeleteBatchCmd represents the batch command
var calendarACLDeleteBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch deletes ACL rules using a CSV file as input.",
	Long:  `Implements the API documented at https://developers.google.com/workspace/calendar/api/v3/reference/acl/delete`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, calendarACLFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			CalendarID string `json:"calendarId,omitempty"`
			RuleID     string `json:"ruleId,omitempty"`
			Result     bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for range cap {
				wg.Go(func() {
					for m := range maps {
						calendarID := m["calendarId"].GetString()
						ruleID := m["ruleId"].GetString()
						result, err := gsmcalendar.DeleteACL(calendarID, ruleID)
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{CalendarID: calendarID, RuleID: ruleID, Result: result}
					}
				})
			}
			wg.Wait()
			close(results)
		}()
		gsmhelpers.StreamOrCollect(gsmhelpers.ChanToIter(results, nil), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(calendarACLDeleteCmd, calendarACLDeleteBatchCmd, calendarACLFlags, calendarACLFlagsALL, batchFlags)
}
