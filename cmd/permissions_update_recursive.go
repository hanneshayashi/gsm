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
	"google.golang.org/api/drive/v3"

	"github.com/spf13/cobra"
)

// permissionsUpdateRecursiveCmd represents the recursive command
var permissionsUpdateRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively updates a permission on a folder and all of its children.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/update",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		folderID := flags["folderId"].GetString()
		_, err := gsmdrive.GetFolder(folderID)
		if err != nil {
			log.Fatalf("Error getting folder: %v", err)
		}
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files := gsmdrive.ListFilesRecursive(folderID, "files(id,mimeType),nextPageToken", threads)
		type resultStruct struct {
			FileID     string            `json:"fileId,omitempty"`
			Permission *drive.Permission `json:"permission,omitempty"`
		}
		results := make(chan resultStruct, threads)
		var wg sync.WaitGroup
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		removeExpiration := flags["removeExpiration"].GetBool()
		fields := flags["fields"].GetString()
		p, err := mapToPermission(flags)
		if err != nil {
			log.Fatalf("Error building permission object: %v", err)
		}
		permissionID, err := gsmdrive.GetPermissionID(flags)
		if err != nil {
			log.Fatalf("Unable to determine permissionId: %v", err)
		}
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for file := range files {
						r, err := gsmdrive.UpdatePermission(file.Id, permissionID, fields, useDomainAdminAccess, removeExpiration, p)
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{FileID: file.Id, Permission: r}
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
			for r := range results {
				final = append(final, r)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(permissionsUpdateCmd, permissionsUpdateRecursiveCmd, permissionFlags, recursiveFileFlags)
}
