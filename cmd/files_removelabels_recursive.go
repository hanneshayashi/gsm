/*
Copyright © 2020-2025 Hannes Hayashi

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

// filesRemoveLabelsRecursiveCmd represents the recursive command
var filesRemoveLabelsRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively removes the specified labels on a folder and all of its children.",
	Long:  "Implements the API documented at https://developers.google.com/drive/api/v3/reference/files/modifyLabels",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), "files(id,mimeType),nextPageToken", flags["excludeFolders"].GetStringSlice(), flags["includeRoot"].GetBool(), threads)
		type resultStruct struct {
			FileID string         `json:"fileId,omitempty"`
			Labels []*drive.Label `json:"labels,omitempty"`
		}
		results := make(chan resultStruct, threads)
		wg := &sync.WaitGroup{}
		fields := flags["fields"].GetString()
		req, err := mapToRemoveLabelsRequest(flags)
		if err != nil {
			log.Fatalf("Error building remove labels request: %v", err)
		}
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for file := range files {
						r := resultStruct{FileID: file.Id}
						result, err := gsmdrive.ModifyLabels(file.Id, fields, req)
						if err != nil {
							log.Println(err)
						} else {
							r.Labels = result
						}
						results <- r
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
	gsmhelpers.InitRecursiveCommand(filesRemoveLabelsCmd, filesRemoveLabelsRecursiveCmd, fileFlags, recursiveFileFlags)
}
