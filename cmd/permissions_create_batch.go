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
	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// permissionsCreateBatchCmd represents the batch command
var permissionsCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch Creates a permission for a file or shared drive using a CSV file as input.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/create",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, permissionFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *drive.Permission, cap)
		final := []*drive.Permission{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						p, err := mapToPermission(m)
						if err != nil {
							log.Printf("Error building permission object: %v\n", err)
							continue
						}
						result, err := gsmdrive.CreatePermission(m["fileId"].GetString(), m["emailMessage"].GetString(), m["fields"].GetString(), m["useDomainAdminAccess"].GetBool(), m["sendNotificationEmail"].GetBool(), m["transferOwnership"].GetBool(), m["moveToNewOwnersRoot"].GetBool(), p)
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
	gsmhelpers.InitBatchCommand(permissionsCreateCmd, permissionsCreateBatchCmd, permissionFlags, permissionFlagsALL, batchFlags)
}
