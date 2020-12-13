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
	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// userPhotosGetBatchCmd represents the batch command
var userPhotosGetBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch get user photos using a CSV file as input.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/users/photos/get",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},	
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, userPhotoFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			UserKey   string           `json:"userKey,omitempty"`
			UserPhoto *admin.UserPhoto `json:"userPhoto,omitempty"`
		}
		results := make(chan resultStruct, cap)
		final := []resultStruct{}
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						result, err := gsmadmin.GetUserPhoto(m["userKey"].GetString(), m["fields"].GetString())
						if err != nil {
							log.Println(err)
						} else {
							results <- resultStruct{UserKey: m["userKey"].GetString(), UserPhoto: result}
						}
					}
					wg.Done()
				}()
			}
			wg.Wait()
			close(results)
		}()
		for res := range results {
			final = append(final, res)
		}
		gsmhelpers.StreamOutput(final, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitBatchCommand(userPhotosGetCmd, userPhotosGetBatchCmd, userPhotoFlags, userPhotoFlagsALL, batchFlags)
}
