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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// chromeOsDevicesActionBatchCmd represents the batch command
var chromeOsDevicesActionBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch takes actions that affect Chrome OS devices using a CSV file as input",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/chromeosdevices/action",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, chromeOsDeviceFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			ResourceID string `json:"resourceId,omitempty"`
			Result     bool   `json:"result"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						a, err := mapToChromeOsDeviceAction(m)
						if err != nil {
							log.Printf("Error building chromeOsDeviceAction object: %v", err)
							continue
						}
						resourceID := m["resourceId"].GetString()
						result, err := gsmadmin.TakeActionOnChromeOsDevice(m["customerId"].GetString(),resourceID, a)
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{ResourceID:resourceID, Result: result}
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
				enc.Encode(r)
			}
		} else {
			final := []resultStruct{}
			for res := range results {
				final = append(final, res)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitBatchCommand(chromeOsDevicesActionCmd, chromeOsDevicesActionBatchCmd, chromeOsDeviceFlags, chromeOsDeviceFlagsALL, batchFlags)
}
