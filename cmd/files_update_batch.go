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
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// filesUpdateBatchCmd represents the batch command
var filesUpdateBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch update files using a CSV file as input.",
	Long:  "https://developers.google.com/drive/api/v3/reference/files/update",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)
		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
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
				m := gsmhelpers.BatchFlagsToMap(flags, fileFlags, line, "update")
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
						log.Printf("Error building file object: %v\n", err)
						continue
					}
					var removeParents string
					if m["parent"].IsSet() {
						fOld, err := gsmdrive.GetFile(m["fileId"].GetString(), m["fields"].GetString(), "")
						if err != nil {
							log.Printf("Error getting existing file %s: %v\n", m["fileId"].GetString(), err)
							continue
						}
						removeParents = strings.Join(fOld.Parents, ",")
					}
					var content *os.File
					if m["localFilePath"].IsSet() {
						content, err = os.Open(m["localFilePath"].GetString())
						if err != nil {
							log.Printf("Error opening file %s: %v", m["localFilePath"].GetString(), err)
							continue
						}
						defer content.Close()
					}
					errKey := fmt.Sprintf("%s:", m["fileId"].GetString())
					operation := func() error {
						result, err := gsmdrive.UpdateFile(m["fileId"].GetString(), m["parent"].GetString(), removeParents, m["includePermissionsForView"].GetString(), m["ocrLanguage"].GetString(), m["fields"].GetString(), f, content, m["keepRevisionForever"].GetBool(), m["useContentAsIndexableText"].GetBool())
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
	filesUpdateCmd.AddCommand(filesUpdateBatchCmd)
	flags := filesUpdateBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(fileFlags, flags, "update")
	markFlagsRequired(filesUpdateBatchCmd, fileFlags, "update")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(filesUpdateBatchCmd, batchFlags, "")
}
