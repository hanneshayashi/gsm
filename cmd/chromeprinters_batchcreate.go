/*
Copyright © 2020-2022 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// chromePrintersBatchCreateCmd represents the batchCreate command
var chromePrintersBatchCreateCmd = &cobra.Command{
	Use:               "batchCreate",
	Short:             "Creates printers under given Organization Unit.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/chrome-printer/reference/rest/v1/admin.directory.v1.customers.chrome.printers/batchCreatePrinters",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		parent := flags["parent"].GetString()
		if parent == "" {
			customerID, err := gsmadmin.GetOwnCustomerID()
			if err != nil {
				log.Printf("Error determining customer ID: %v\n", err)
			}
			parent = "customers/" + customerID
		}
		batchCreatePrintersRequest, err := mapToBatchCreatePrintersRequest(flags)
		if err != nil {
			log.Fatalf("Error building batch create printer request object: %v", err)
		}
		result, err := gsmadmin.BatchCreatePrinters(parent, flags["fields"].GetString(), batchCreatePrintersRequest)
		if err != nil {
			log.Fatalf("Error creating Chrome printer: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(chromePrintersCmd, chromePrintersBatchCreateCmd, chromePrinterFlags)
}
