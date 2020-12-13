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
)

// permissionsDeleteRecursiveCmd represents the recursive command
var permissionsDeleteRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively grant a permissions to a folder and all of its children.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/delete",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		folderID := flags["folderId"].GetString()
		folder, err := gsmdrive.GetFile(folderID, "id,mimeType", "")
		if err != nil {
			log.Fatalf("Error getting folder: %v", err)
		}
		if !gsmdrive.IsFolder(folder) {
			log.Fatalf("%s is not a folder", folderID)
		}
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files, err := gsmdrive.ListFilesRecursive(folderID, "files(id,mimeType),nextPageToken", threads)
		if err != nil {
			log.Fatalf("Error listing files: %v", err)
		}
		type resultStruct struct {
			FileID string `json:"fileId,omitempty"`
			Result bool   `json:"result,omitempty"`
		}
		resultsChan := make(chan resultStruct, threads)
		final := []resultStruct{}
		wgPermissions := &sync.WaitGroup{}
		wgFinal := &sync.WaitGroup{}
		idChan := make(chan string, threads)
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		permissionID := flags["permissionId"].GetString()
		wgPermissions.Add(1)
		go func() {
			idChan <- folderID
			for _, f := range files {
				idChan <- f.Id
			}
			close(idChan)
			wgPermissions.Done()
		}()
		for i := 0; i < threads; i++ {
			wgPermissions.Add(1)
			go func() {
				for id := range idChan {
					r, err := gsmdrive.DeletePermission(id, permissionID, useDomainAdminAccess)
					if err != nil {
						log.Println(err)
					} else {
						resultsChan <- resultStruct{FileID: id, Result: r}
					}
				}
				wgPermissions.Done()
			}()
		}
		wgFinal.Add(1)
		go func() {
			for r := range resultsChan {
				final = append(final, r)
			}
			wgFinal.Done()
		}()
		wgPermissions.Wait()
		close(resultsChan)
		wgFinal.Wait()
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(permissionsDeleteCmd, permissionsDeleteRecursiveCmd, permissionFlags, recursiveFileFlags)
}
