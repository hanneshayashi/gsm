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
	"log"
	"sync"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// verificationcodesListRecursiveCmd represents the recursive command
var verificationcodesListRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: `Returns the current set of valid backup verification codes for users by referencing one or more organizational units and/or groups.`,
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/verificationcodes/list",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		threads := gsmhelpers.MaxThreads(flags["batchThreads"].GetInt())
		type resultStruct struct {
			UserKey           string                    `json:"userKey,omitempty"`
			VerificationCodes []*admin.VerificationCode `json:"verification_codes,omitempty"`
		}
		results := make(chan resultStruct, threads)
		var wg sync.WaitGroup
		userKeysUnique, _ := gsmadmin.GetUniqueUsersChannelRecursive(flags["orgUnit"].GetStringSlice(), flags["groupEmail"].GetStringSlice(), threads)
		fields := flags["fields"].GetString()
		go func() {
			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {
					for uk := range userKeysUnique {
						result, err := gsmadmin.ListVerificationCodes(uk, fields)
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{UserKey: uk, VerificationCodes: result}
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
			final := []resultStruct{}
			for r := range results {
				final = append(final, r)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(verificationCodesListCmd, verificationcodesListRecursiveCmd, verificationCodeFlags, recursiveUserFlags)
}
