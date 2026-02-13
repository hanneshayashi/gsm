/*
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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// permissionsListRecursiveCmd represents the recursive command
var permissionsListRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively lists permissions on a folder and all of its children.",
	Long: `IMPORTANT:
If you are not specifying a folder in a Shared Drive, you can simply use "gsm files list recursive" with "permissions" in the fields parameter like so:
"gsm files list recursive --folderId <folderId> --fields "files(id,name,mimeType,permissions),nextPageToken"`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), "files(id,mimeType),nextPageToken", flags["excludeFolders"].GetStringSlice(), flags["includeRoot"].GetBool(), threads)
		type resultStruct struct {
			FileID      string              `json:"fileId,omitempty"`
			Permissions []*drive.Permission `json:"permissions,omitempty"`
		}
		results := make(chan resultStruct, threads)
		wg := &sync.WaitGroup{}
		fields := flags["fields"].GetString()
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		go func() {
			for range threads {
				wg.Go(func() {
					for file := range files {
						var iterErr error
						r := resultStruct{FileID: file.Id}
						for i, err := range gsmdrive.ListPermissions(file.Id, "", fields, useDomainAdminAccess) {
							if err != nil {
								iterErr = err
								break
							}
							r.Permissions = append(r.Permissions, i)
						}
						if iterErr != nil {
							log.Println(iterErr)
						} else {
							results <- r
						}
					}
				})
			}
			wg.Wait()
			close(results)
		}()
		gsmhelpers.StreamOrCollect(gsmhelpers.ChanToIter(results, nil), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(permissionsListCmd, permissionsListRecursiveCmd, permissionFlags, recursiveFileFlags)
}
