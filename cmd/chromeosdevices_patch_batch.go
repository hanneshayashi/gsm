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
	"sync"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// chromeOsDevicesPatchBatchCmd represents the batch command
var chromeOsDevicesPatchBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch patch Chrome OS devices using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/chromeosdevices/patch",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, chromeOsDeviceFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *admin.ChromeOsDevice, cap)
		final := []*admin.ChromeOsDevice{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						c, err := mapToChromeOsDevice(m)
						if err != nil {
							log.Printf("Error building chromeOsDevice object: %v\n", err)
							continue
						}
						result, err := gsmadmin.PatchChromeOsDevice(m["customerId"].GetString(), m["deviceId"].GetString(), m["fields"].GetString(), m["projection"].GetString(), c)
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
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(chromeOsDevicesPatchCmd, chromeOsDevicesPatchBatchCmd, chromeOsDeviceFlags, chromeOsDeviceFlagsALL, batchFlags)
}
