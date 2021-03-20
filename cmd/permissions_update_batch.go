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

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// permissionsUpdateBatchCmd represents the batch command
var permissionsUpdateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch updates permissions for a file or shared drive using a CSV file as input.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/update",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, permissionFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			Permission *drive.Permission `json:"permission,omitempty"`
			FileID     string            `json:"fileId,omitempty"`
		}
		results := make(chan resultStruct, cap)
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
						permissionID, err := gsmdrive.GetPermissionID(m)
						if err != nil {
							log.Printf("Unable to determine permissionId: %v", err)
							continue
						}
						fileID := m["fileId"].GetString()
						result, err := gsmdrive.UpdatePermission(fileID, permissionID, m["fields"].GetString(), m["useDomainAdminAccess"].GetBool(), m["removeExpiration"].GetBool(), p)
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{FileID: fileID, Permission: result}
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
	gsmhelpers.InitBatchCommand(permissionsUpdateCmd, permissionsUpdateBatchCmd, permissionFlags, permissionFlagsALL, batchFlags)
}
