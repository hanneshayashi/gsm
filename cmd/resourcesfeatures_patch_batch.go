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

// resourcesFeaturesPatchBatchCmd represents the batch command
var resourcesFeaturesPatchBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch patches feature resources using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/resources/features/patch",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		results := []*admin.Feature{}
		for _, line := range csv {
			time.Sleep(300 * time.Millisecond)
			m := gsmhelpers.BatchFlagsToMap(flags, resourcesBuildingFlags, line, "patch")
			f, err := mapToFeature(m)
			if err != nil {
				log.Printf("Error building resourceFeature object: %v", err)
				continue
			}
			result, err := gsmadmin.PatchResourcesFeature(m["customer"].GetString(), m["featureKey"].GetString(), m["fields"].GetString(), f)
			if err != nil {
				log.Printf("Error getting building %s: %v\n", m["buildingId"].GetString(), err)
			}
			results = append(results, result)
		}
		fmt.Println(gsmhelpers.PrettyPrint(results, "json"))
	},
}

func init() {
	resourcesFeaturesPatchCmd.AddCommand(resourcesFeaturesPatchBatchCmd)
	flags := resourcesFeaturesPatchBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(resourcesFeatureFlags, flags, "patch")
	markFlagsRequired(resourcesFeaturesPatchBatchCmd, resourcesFeatureFlags, "patch")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(resourcesFeaturesPatchBatchCmd, batchFlags, "")
}
