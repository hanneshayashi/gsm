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

// filesCopyRecursiveCmd represents the recursive command
var filesCopyRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively copies a folder to a new destination.",
	Long:  "https://developers.google.com/drive/api/v3/reference/files/copy",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		folderID := flags["folderId"].GetString()
		results := make(chan *drive.File, threads)
		files, err := gsmdrive.CopyFoldersAndReturnFilesWithNewParents(folderID, flags["parent"].GetString(), results, flags["excludeFolders"].GetStringSlice(), threads)
		if err != nil {
			log.Fatalf("Error getting files and folders: %v", err)
		}
		var wg sync.WaitGroup
		if err != nil {
			log.Fatalf("Error creating new folder structure: %v", err)
		}
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for f := range files {
						c, err := gsmdrive.CopyFile(f.Id, "", "", "id,name,mimeType,parents", &drive.File{Parents: []string{f.Parents[1]}, Name: f.Name}, false, false)
						if err != nil {
							log.Println(err)
						} else {
							results <- c
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
	gsmhelpers.InitRecursiveCommand(filesCopyCmd, filesCopyRecursiveCmd, fileFlags, recursiveFileFlags)
}
