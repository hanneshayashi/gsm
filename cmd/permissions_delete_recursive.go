/*
Copyright © 2020-2023 Hannes Hayashi

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
)

// permissionsDeleteRecursiveCmd represents the recursive command
var permissionsDeleteRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively deletes a permission from a folder and all of its children.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/permissions/delete",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		folderID := flags["folderId"].GetString()
		_, err := gsmdrive.GetFolder(folderID)
		if err != nil {
			log.Fatalf("Error getting folder: %v", err)
		}
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files := gsmdrive.ListFilesRecursive(folderID, "files(id,mimeType),nextPageToken", flags["excludeFolders"].GetStringSlice(), flags["includeRoot"].GetBool(), threads)
		type resultStruct struct {
			FileID string `json:"fileId,omitempty"`
			Result bool   `json:"result,omitempty"`
		}
		results := make(chan resultStruct, threads)
		var wg sync.WaitGroup
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		enforceExpansiveAccess := flags["enforceExpansiveAccess"].GetBool()
		permissionID, err := gsmdrive.GetPermissionID(flags)
		if err != nil {
			log.Fatalf("Unable to determine permissionId: %v", err)
		}
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for file := range files {
						r, err := gsmdrive.DeletePermission(file.Id, permissionID, useDomainAdminAccess, enforceExpansiveAccess)
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{FileID: file.Id, Result: r}
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
				err := enc.Encode(r)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []resultStruct{}
			for r := range results {
				final = append(final, r)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(permissionsDeleteCmd, permissionsDeleteRecursiveCmd, permissionFlags, recursiveFileFlags)
}
