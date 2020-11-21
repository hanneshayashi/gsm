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
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// filesCreateBatchCmd represents the batch command
var filesCreateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch Creates a new file or folder. Can also be used to upload files using a CSV file as input.",
	Long:  "https://developers.google.com/drive/api/v3/reference/files/create",
	Run: func(cmd *cobra.Command, args []string) {
		flags, err := gsmhelpers.ConsolidateFlags(cmd, fileFlags)
		if err != nil {
			log.Fatalf("Error consolidating flags: %v", err)
		}
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		err = gsmhelpers.CheckBatchFlags(flags, fileFlags, int64(len(csv[0])))
		if err != nil {
			log.Fatalf("Error with batch flag index: %v\n", err)
		}
		l := len(csv)
		results := make(chan *drive.File, l)
		maps := make(chan map[string]*gsmhelpers.Value, l)
		final := []*drive.File{}
		var wg1 sync.WaitGroup
		var wg2 sync.WaitGroup
		var wg3 sync.WaitGroup
		wg1.Add(1)
		go func() {
			for _, line := range csv {
				m := gsmhelpers.BatchFlagsToMap(flags, fileFlags, line, "create")
				maps <- m
			}
			close(maps)
			wg1.Done()
		}()
		wg2.Add(1)
		retrier := gsmhelpers.NewStandardRetrier()
		for i := 0; i < gsmhelpers.MaxThreads(l); i++ {
			wg2.Add(1)
			go func() {
				for m := range maps {
					var err error
					f, err := mapToFile(m)
					if err != nil {
						log.Printf("Error building drive object: %v\n", err)
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
					errKey := fmt.Sprintf("%s:", f.Name)
					operation := func() error {
						result, err := gsmdrive.CreateFile(f, content, m["ignoreDefaultVisibility"].GetBool(), m["keepRevisionForever"].GetBool(), m["useContentAsIndexableText"].GetBool(), m["includePermissionsForView"].GetString(), m["ocrLanguage"].GetString(), m["fields"].GetString())
						if err != nil {
							retryable := gsmhelpers.ErrorIsRetryable(err)
							if retryable {
								log.Println(errKey, "Retrying after", err)
								return err
							}
							log.Println(errKey, "Giving up after", err)
							return nil
						}
						results <- result
						return nil
					}
					err = retrier.Run(operation)
					if err != nil {
						log.Println(errKey, "Max retries reached. Giving up after", err)
					}
					time.Sleep(200 * time.Millisecond)
				}
				wg2.Done()
			}()
		}
		wg3.Add(1)
		go func() {
			for res := range results {
				final = append(final, res)
			}
			wg3.Done()
		}()
		wg2.Done()
		wg1.Wait()
		wg2.Wait()
		close(results)
		wg3.Wait()
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(final, "json"))
	},
}

func init() {
	gsmhelpers.InitBatchCommand(filesCreateCmd, filesCreateBatchCmd, fileFlags, fileFlagsALL, batchFlags)
}
