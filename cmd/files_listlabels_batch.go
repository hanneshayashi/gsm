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

// filesListLabelsBatchCmd represents the batch command
var filesListLabelsBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch list the labels on files using a CSV file as input.",
	Long:  "Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/files/listLabels",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, fileFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			FileID string         `json:"fileId,omitempty"`
			Labels []*drive.Label `json:"labels,omitempty"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						fileID := m["fileId"].GetString()
						var iterErr error
						r := resultStruct{FileID: fileID}
						for i, err := range gsmdrive.ListLabels(fileID, m["fields"].GetString()) {
							if err != nil {
								iterErr = err
								break
							}
							r.Labels = append(r.Labels, i)
						}
						if iterErr != nil {
							log.Println(iterErr)
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
		gsmhelpers.StreamOrCollect(gsmhelpers.ChanToIter(results, nil), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(filesListLabelsCmd, filesListLabelsBatchCmd, fileFlags, fileFlagsALL, batchFlags)
}
