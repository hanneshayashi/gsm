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
	"gsm/gsmdrive"
	"gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// permissionsListRecursiveCmd represents the recursive command
var permissionsListRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively lists permissions to a folder and all of its children.",
	Long: `IMPORTANT:
If you are not specifying a folder in a Shared Drive, you can simply use "files list recursive" with "permissions" in the fields parameter like so:
"files list recursive --folder <folderId> --fields "files(id,name,permissions)"`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files, err := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), "files(id,mimeType),nextPageToken", threads)
		if err != nil {
			log.Fatalf("Error listing files: %v", err)
		}
		type resultStruct struct {
			FileID      string              `json:"fileId,omitempty"`
			Permissions []*drive.Permission `json:"permissions,omitempty"`
		}
		resultsChan := make(chan resultStruct, threads)
		final := []resultStruct{}
		wgPermissions := &sync.WaitGroup{}
		wgFinal := &sync.WaitGroup{}
		idChan := make(chan string, threads)
		fields := flags["fields"].GetString()
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		wgPermissions.Add(1)
		go func() {
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
					r, err := gsmdrive.ListPermissions(id, "", fields, useDomainAdminAccess)
					if err != nil {
						log.Println(err)
					} else {
						resultsChan <- resultStruct{FileID: id, Permissions: r}
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
	gsmhelpers.InitRecursiveCommand(permissionsListCmd, permissionsListRecursiveCmd, permissionFlags, recursiveFileFlags)
}
