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
	"google.golang.org/api/drive/v3"
)

// permissionsCreateRecursiveCmd represents the recursive command
var permissionsCreateRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively grant a permissions to a folder and all of its children.",
	Long:  "Implements the API documented at https://developers.google.com/drive/api/v3/reference/permissions/create",
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
		p, err := mapToPermission(flags)
		if err != nil {
			log.Fatalf("Error building permission object: %v", err)
		}
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files := gsmdrive.ListFilesRecursive(folderID, "files(id,mimeType),nextPageToken", flags["excludeFolders"].GetStringSlice(), flags["includeRoot"].GetBool(), threads)
		type resultStruct struct {
			Permissions *drive.Permission `json:"permissions,omitempty"`
			FileID      string            `json:"fileId,omitempty"`
		}
		results := make(chan resultStruct, threads)
		var wg sync.WaitGroup
		fields := flags["fields"].GetString()
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		emailMessage := flags["emailMessage"].GetString()
		sendNotificationEmail := flags["sendNotificationEmail"].GetBool()
		transferOwnership := flags["transferOwnership"].GetBool()
		moveToNewOwnersRoot := flags["moveToNewOwnersRoot"].GetBool()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for file := range files {
						var move bool
						if moveToNewOwnersRoot && file.Id == folderID {
							move = true
						} else {
							move = false
						}
						r, err := gsmdrive.CreatePermission(file.Id, emailMessage, fields, useDomainAdminAccess, sendNotificationEmail, transferOwnership, move, p)
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{FileID: file.Id, Permissions: r}
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
	gsmhelpers.InitRecursiveCommand(permissionsCreateCmd, permissionsCreateRecursiveCmd, permissionFlags, recursiveFileFlags)
}
