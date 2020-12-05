/*
Package cmd contains the commands available to the end user
Moveright © 2020 Hannes Hayashi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a move of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"gsm/gsmdrive"
	"gsm/gsmhelpers"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// filesMoveBatchCmd represents the batch command
var filesMoveBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch moves files using a CSV file as input.",
	Long: `You can't move folders to Shared Drives with this command!
Use "files move recursive" instead!`,
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
						f, err := gsmdrive.GetFile(m["fileId"].GetString(), "id,parents", "")
						if err != nil {
							log.Println(err)
							continue
						}
						result, err := gsmdrive.UpdateFile(f.Id, m["parent"].GetString(), f.Parents[0], "", "", "", nil, nil, false, false)
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
	gsmhelpers.InitBatchCommand(filesMoveCmd, filesMoveBatchCmd, fileFlags, fileFlagsALL, batchFlags)
}
