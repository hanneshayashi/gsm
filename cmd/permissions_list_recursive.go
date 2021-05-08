/*
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

// permissionsListRecursiveCmd represents the recursive command
var permissionsListRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively lists permissions to a folder and all of its children.",
	Long: `IMPORTANT:
If you are not specifying a folder in a Shared Drive, you can simply use "gsm files list recursive" with "permissions" in the fields parameter like so:
"gsm files list recursive --folderId <folderId> --fields "files(id,name,mimeType,permissions)"`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), "files(id,mimeType),nextPageToken", flags["excludeFolders"].GetStringSlice(), threads)
		type resultStruct struct {
			FileID      string              `json:"fileId,omitempty"`
			Permissions []*drive.Permission `json:"permissions,omitempty"`
		}
		results := make(chan resultStruct, threads)
		wg := &sync.WaitGroup{}
		fields := flags["fields"].GetString()
		useDomainAdminAccess := flags["useDomainAdminAccess"].GetBool()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for file := range files {
						result, err := gsmdrive.ListPermissions(file.Id, "", fields, useDomainAdminAccess, threads)
						r := resultStruct{FileID: file.Id}
						for i := range result {
							r.Permissions = append(r.Permissions, i)
						}
						e := <-err
						if e != nil {
							log.Println(e)
						} else {
							results <- r
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
	gsmhelpers.InitRecursiveCommand(permissionsListCmd, permissionsListRecursiveCmd, permissionFlags, recursiveFileFlags)
}
