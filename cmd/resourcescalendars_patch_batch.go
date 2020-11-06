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
	"fmt"
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"
	"time"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// resourcesCalendarsPatchBatchCmd represents the batch command
var resourcesCalendarsPatchBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch patches calendar resources using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/resources/calendars/patch",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		results := []*admin.CalendarResource{}
		for _, line := range csv {
			time.Sleep(300 * time.Millisecond)
			m := gsmhelpers.BatchFlagsToMap(flags, resourcesCalendarFlags, line, "patch")
			c, err := mapToResourceCalendar(m)
			if err != nil {
				log.Printf("Error building resourceCalendar object: %v", err)
				continue
			}
			result, err := gsmadmin.PatchResourcesCalendar(m["customer"].GetString(), m["calendarResourceId"].GetString(), m["fields"].GetString(), c)
			if err != nil {
				log.Printf("Error getting building %s: %v\n", m["buildingId"].GetString(), err)
			}
			results = append(results, result)
		}
		fmt.Println(gsmhelpers.PrettyPrint(results, "json"))
	},
}

func init() {
	resourcesCalendarsPatchCmd.AddCommand(resourcesCalendarsPatchBatchCmd)
	flags := resourcesCalendarsPatchBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(resourcesCalendarFlags, flags, "patch")
	markFlagsRequired(resourcesCalendarsPatchBatchCmd, resourcesCalendarFlags, "patch")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(resourcesCalendarsPatchBatchCmd, batchFlags, "")
}
