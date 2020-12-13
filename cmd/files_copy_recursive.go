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
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		folderID := flags["folderId"].GetString()
		folderMap, files, err := gsmdrive.GetFilesAndFolders(folderID, threads)
		if err != nil {
			log.Fatalf("Error getting files and folders: %v", err)
		}
		filesChan := make(chan *drive.File, threads)
		finalChan := make(chan *drive.File, threads)
		final := []*drive.File{}
		wgFiles := &sync.WaitGroup{}
		wgFinal := &sync.WaitGroup{}
		folderMap[folderID].NewParent = flags["parent"].GetString()
		wgFiles.Add(1)
		go func() {
			for _, f := range files {
				filesChan <- f
			}
			close(filesChan)
			wgFiles.Done()
		}()
		err = gsmdrive.CopyFolders(folderMap, "")
		if err != nil {
			log.Fatalf("Error creating new folder structure: %v", err)
		}
		for i := 0; i < threads; i++ {
			wgFiles.Add(1)
			go func() {
				for f := range filesChan {
					folder := folderMap[f.Parents[0]]
					c, err := gsmdrive.CopyFile(f.Id, "", "", "id,name,mimeType,parents", &drive.File{Parents: []string{folder.NewID}, Name: f.Name}, false, false)
					if err != nil {
						log.Println(err)
					} else {
						finalChan <- c
					}
				}
				wgFiles.Done()
			}()
		}
		wgFinal.Add(1)
		go func() {
			for r := range finalChan {
				final = append(final, r)
			}
			wgFinal.Done()
		}()
		wgFiles.Wait()
		close(finalChan)
		wgFinal.Wait()
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(filesCopyCmd, filesCopyRecursiveCmd, fileFlags, recursiveFileFlags)
}
