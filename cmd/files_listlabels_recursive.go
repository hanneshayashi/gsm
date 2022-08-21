/*
Copyright © 2020-2022 Hannes Hayashi

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

// filesListlabelsRecursiveCmd represents the recursive command
var filesListlabelsRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively lists labels on a folder and all of its children.",
	Long:  "Implements the API documented at https://developers.google.com/drive/api/v3/reference/files/listLabels",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		files := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), "files(id,mimeType),nextPageToken", flags["excludeFolders"].GetStringSlice(), threads)
		type resultStruct struct {
			FileID string         `json:"fileId,omitempty"`
			Labels []*drive.Label `json:"labels,omitempty"`
		}
		results := make(chan resultStruct, threads)
		wg := &sync.WaitGroup{}
		fields := flags["fields"].GetString()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for file := range files {
						result, err := gsmdrive.ListLabels(file.Id, fields, threads)
						r := resultStruct{FileID: file.Id}
						for i := range result {
							r.Labels = append(r.Labels, i)
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
	gsmhelpers.InitRecursiveCommand(filesListLabelsCmd, filesListlabelsRecursiveCmd, fileFlags, recursiveFileFlags)
}
