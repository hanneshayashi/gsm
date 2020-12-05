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

// permissionsCreateRecursiveCmd represents the recursive command
var permissionsCreateRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively grant a permissions to a folder and all of its children.",
	Long:  "https://developers.google.com/drive/api/v3/reference/permissions/create",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToPermission(flags)
		if err != nil {
			log.Fatalf("Error building permission object: %v", err)
		}
		folderID := flags["folderId"].GetString()
		folder, err := gsmdrive.GetFile(folderID, "id,mimeType", "")
		if err != nil {
			log.Fatalf("Error getting folder: %v", err)
		}
		if !gsmdrive.IsFolder(folder) {
			log.Fatalf("%s is not a folder", folderID)
		}
		threads := gsmhelpers.MaxThreads(viper.GetInt("threads"))
		files, err := gsmdrive.ListFilesRecursive(folderID, "files(id,mimeType),nextPageToken", threads)
		if err != nil {
			log.Fatalf("Error listing files: %v", err)
		}
		type resultStruct struct {
			FileID      string            `json:"fileId,omitempty"`
			Permissions *drive.Permission `json:"permissions,omitempty"`
		}
		resultsChan := make(chan resultStruct, threads)
		final := []resultStruct{}
		wgPermissions := &sync.WaitGroup{}
		wgFinal := &sync.WaitGroup{}
		idChan := make(chan string, threads)
		fields := flags["fields"].GetString()
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		emailMessage := flags["emailMessage"].GetString()
		sendNotificationEmail := flags["sendNotificationEmail"].GetBool()
		transferOwnership := flags["transferOwnership"].GetBool()
		moveToNewOwnersRoot := flags["moveToNewOwnersRoot"].GetBool()
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
					var move bool
					if moveToNewOwnersRoot && id == folderID {
						move = true
					} else {
						move = false
					}
					r, err := gsmdrive.CreatePermission(id, emailMessage, fields, useDomainAdminAccess, sendNotificationEmail, transferOwnership, move, p)
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
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(permissionsCreateCmd, permissionsCreateRecursiveCmd, permissionFlags, recursiveFileFlags)
}
