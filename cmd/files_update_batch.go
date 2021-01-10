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
	"os"
	"strings"
	"sync"

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// filesUpdateBatchCmd represents the batch command
var filesUpdateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch update files using a CSV file as input.",
	Long:  "https://developers.google.com/drive/api/v3/reference/files/update",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, fileFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *drive.File, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						f, err := mapToFile(m)
						if err != nil {
							log.Printf("Error building file object: %v\n", err)
							continue
						}
						var removeParents string
						fileID := m["fileId"].GetString()
						fields := m["fields"].GetString()
						if m["parent"].IsSet() {
							f.Parents = nil
							fOld, err := gsmdrive.GetFile(fileID, fields, "")
							if err != nil {
								log.Printf("Error getting existing file %s: %v\n", fileID, err)
								continue
							}
							removeParents = strings.Join(fOld.Parents, ",")
						}
						var content *os.File
						if m["localFilePath"].IsSet() {
							localFilePath := m["localFilePath"].GetString()
							content, err = os.Open(localFilePath)
							if err != nil {
								log.Printf("Error opening file %s: %v", localFilePath, err)
								continue
							}
							defer content.Close()
						}
						result, err := gsmdrive.UpdateFile(fileID, m["parent"].GetString(), removeParents, m["includePermissionsForView"].GetString(), m["ocrLanguage"].GetString(), fields, f, content, m["keepRevisionForever"].GetBool(), m["useContentAsIndexableText"].GetBool())
						if err != nil {
							log.Println(err)
						} else {
							results <- result
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
			for res := range results {
				final = append(final, res)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitBatchCommand(filesUpdateCmd, filesUpdateBatchCmd, fileFlags, fileFlagsALL, batchFlags)
}
