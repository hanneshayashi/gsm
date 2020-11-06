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
	"gsm/gsmgmail"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/api/gmail/v1"
)

// forwardingAddressesCreateBatchCmd represents the batch command
var forwardingAddressesCreateBatchCmd = &cobra.Command{
	Use: "batch",
	Short: `Creates a forwarding address.
If ownership verification is required, a message will be sent to the recipient and the resource's verification status will be set to pending;
otherwise, the resource will be created with verification status set to accepted.`,
	Long: "https://developers.google.com/gmail/api/reference/rest/v1/users.settings.forwardingAddresses/create",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		cmd.Flags().VisitAll(gsmhelpers.CheckBatchFlags)
		csv, err := gsmhelpers.GetCSV(flags)

		if err != nil {
			log.Fatalf("Error with CSV file: %v\n", err)
		}
		results := []*gmail.ForwardingAddress{}
		for _, line := range csv {
			m := gsmhelpers.BatchFlagsToMap(flags, forwardingAddressFlags, line, "create")
			f, err := mapToForwardingAddress(m)
			if err != nil {
				log.Printf("Error building forwarding address object: %v\n", err)
				continue
			}
			result, err := gsmgmail.CreateForwardingAddress(m["userId"].GetString(), m["fields"].GetString(), f)
			if err != nil {
				log.Printf("Error creating forwarding address for user %s: %v", m["userId"].GetString(), err)
			}
			results = append(results, result)
		}
		fmt.Println(gsmhelpers.PrettyPrint(results, "json"))
	},
}

func init() {
	forwardingAddressesCreateCmd.AddCommand(forwardingAddressesCreateBatchCmd)
	flags := forwardingAddressesCreateBatchCmd.Flags()
	gsmhelpers.AddFlagsBatch(forwardingAddressFlags, flags, "create")
	markFlagsRequired(forwardingAddressesCreateBatchCmd, forwardingAddressFlags, "create")
	gsmhelpers.AddFlags(batchFlags, flags, "batch")
	markFlagsRequired(forwardingAddressesCreateBatchCmd, batchFlags, "")
}
