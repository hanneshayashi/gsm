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

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"google.golang.org/api/googleapi"

	"github.com/spf13/cobra"
)

// userInvitationsSendBatchCmd represents the batch command
var userInvitationsSendBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch sends user invitations using a CSV file as input.",
	Long:  "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/customers.userinvitations/send",
	Annotations: map[string]string{
		"crescendoAttachToParent": "true",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		maps, err := gsmhelpers.GetBatchMaps(cmd, userInvitationFlags)
		if err != nil {
			log.Fatalln(err)
		}
		var wg sync.WaitGroup
		cap := cap(maps)
		type resultStruct struct {
			Result *googleapi.RawMessage `json:"result"`
			Name   string                `json:"name,omitempty"`
		}
		results := make(chan resultStruct, cap)
		go func() {
			for i := 0; i < cap; i++ {
				wg.Add(1)
				go func() {
					for m := range maps {
						name := m["name"].GetString()
						sendUserInvitationRequest, err := mapToSendUserInvitationRequest(m)
						if err != nil {
							log.Fatalf("Error building sendUserInvitationRequest object: %v", err)
						}
						result, err := gsmci.SendInvitation(name, m["fields"].GetString(), sendUserInvitationRequest)
						if err != nil {
							log.Println(err)
						}
						results <- resultStruct{Name: name, Result: result}
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
			for res := range results {
				final = append(final, res)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	gsmhelpers.InitBatchCommand(userInvitationsSendCmd, userInvitationsSendBatchCmd, userInvitationFlags, userInvitationFlagsALL, batchFlags)
}
