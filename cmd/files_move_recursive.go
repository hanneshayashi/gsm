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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// filesMoveRecursiveCmd represents the movefoldertoshareddrive command
var filesMoveRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Moves a folder to a Shared Drive",
	Long: `WARNING: This command can "move" a folder to a Shared Drive, by creating a COPY(!) of its folder structure and MOVING(!)
all files to the new folders. For each source folder, a new folder will be created at the destination.
Files will be moved (not copied!!) to the new folders.
The original folders will be preserved at the source!`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		folderID := flags["folderId"].GetString()
		folderMap, files, err := gsmdrive.GetFilesAndFolders(folderID, threads)
		if err != nil {
			log.Fatalf("Error getting files and folders: %v", err)
		}
		filesChan := make(chan *drive.File, threads)
		results := make(chan *drive.File, threads)
		var wg sync.WaitGroup
		folderMap[folderID].NewParent = flags["parent"].GetString()
		go func() {
			for _, f := range files {
				filesChan <- f
			}
			close(filesChan)
		}()
		err = gsmdrive.CopyFolders(folderMap, "")
		if err != nil {
			log.Fatalf("Error creating new folder structure: %v", err)
		}
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for f := range filesChan {
						folder := folderMap[f.Parents[0]]
						u, err := gsmdrive.UpdateFile(f.Id, folder.NewID, folder.OldParent, "", "", "id", nil, nil, false, false)
						if err != nil {
							log.Println(err)
						} else {
							results <- u
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
			final := []*drive.File{}
			for r := range results {
				final = append(final, r)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(filesMoveCmd, filesMoveRecursiveCmd, fileFlags, recursiveFileFlags)
}
