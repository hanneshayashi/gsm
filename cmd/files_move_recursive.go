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

// filesMoveRecursiveCmd represents the movefoldertoshareddrive command
var filesMoveRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Moves a folder to a Shared Drive",
	Long: `WARNING: This command can "move" a folder to a Shared Drive outside your organization, by creating a COPY(!) of its folder structure and MOVING(!)
all files to the new folders. For each source folder, a new folder will be created at the destination.
Files will be moved (not copied!!) to the new folders.
The original folders will be preserved at the source!`,
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		folderID := flags["folderId"].GetString()
		results := make(chan *drive.File, threads)
		files, err := gsmdrive.CopyFoldersAndReturnFilesWithNewParents(folderID, flags["parent"].GetString(), results, flags["excludeFolders"].GetStringSlice(), threads)
		if err != nil {
			log.Fatalf("Error getting files and folders: %v", err)
		}
		var wg sync.WaitGroup
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for f := range files {
						u, err := gsmdrive.UpdateFile(f.Id, f.Parents[1], f.Parents[0], "", "", "id", nil, nil, false, false)
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
				err := enc.Encode(r)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*drive.File{}
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
	gsmhelpers.InitRecursiveCommand(filesMoveCmd, filesMoveRecursiveCmd, fileFlags, recursiveFileFlags)
}
