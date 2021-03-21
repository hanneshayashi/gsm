/*
Package cmd contains the commands available to the end user
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

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// devicesWipeBatchCmd represents the batch command
var devicesWipeBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch wipes devices using a CSV file as input.",
	Long:  `https://cloud.google.com/identity/docs/reference/rest/v1/devices/wipe`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, deviceFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			Result   map[string]interface{} `json:"result"`
			Name     string                 `json:"name,omitempty"`
			Customer string                 `json:"customer,omitempty"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						name := m["name"].GetString()
						wipeRequest, err := mapToWipeDeviceRequest(m)
						if err != nil {
							log.Fatalf("Error building wipeRequest object: %v", err)
						}
						result, err := gsmci.WipeDevice(name, m["fields"].GetString(), wipeRequest)
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{Name: name, Customer: wipeRequest.Customer, Result: result}
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
	gsmhelpers.InitBatchCommand(devicesWipeCmd, devicesWipeBatchCmd, deviceFlags, deviceFlagsALL, batchFlags)
}
