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
	"fmt"
	"gsm/gsmdrive"
	"gsm/gsmhelpers"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// filesCreateBatchCmd represents the batch command
var filesCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch Creates a new file or folder. Can also be used to upload files using a CSV file as input.",
	Long:  "https://developers.google.com/drive/api/v3/reference/files/create",
	Run: func(cmd *cobra.Command, args []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, fileFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		results := make(chan *drive.File, cap)
		final := []*drive.File{}
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
						var content *os.File
						if m["localFilePath"].IsSet() {
							content, err = os.Open(m["localFilePath"].GetString())
							if err != nil {
								log.Printf("Error opening file %s: %v", m["localFilePath"].GetString(), err)
								continue
							}
							defer content.Close()
							if f.Name == "" {
								f.Name = filepath.Base(content.Name())
							}
						}
						result, err := gsmdrive.CreateFile(f, content, m["ignoreDefaultVisibility"].GetBool(), m["keepRevisionForever"].GetBool(), m["useContentAsIndexableText"].GetBool(), m["includePermissionsForView"].GetString(), m["ocrLanguage"].GetString(), m["fields"].GetString())
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
		for res := range results {
			final = append(final, res)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(filesCreateCmd, filesCreateBatchCmd, fileFlags, fileFlagsALL, batchFlags)
}
