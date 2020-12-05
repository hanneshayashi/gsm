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
	"gsm/gsmdrive"
	"gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/drive/v3"
)

// drivesUpdateBatchCmd represents the batch command
var drivesUpdateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch updates drives using a CSV file as input",
	Long:  "https://developers.google.com/drive/api/v3/reference/drives/update",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, driveFlags, viper.GetInt("threads"))
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *drive.Drive, cap)
		final := []*drive.Drive{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						d, err := mapToDrive(m)
						if err != nil {
							log.Printf("Error building drive object: %v\n", err)
							continue
						}
						result, err := gsmdrive.UpdateDrive(m["driveId"].GetString(), m["fields"].GetString(), m["useDomainAdminAccess"].GetBool(), d)
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
	gsmhelpers.InitBatchCommand(drivesUpdateCmd, drivesUpdateBatchCmd, driveFlags, driveFlagsALL, batchFlags)
}
