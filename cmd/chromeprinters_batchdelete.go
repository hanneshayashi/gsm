/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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

// chromePrintersBatchDeleteCmd represents the batchdelete command
var chromePrintersBatchDeleteCmd = &cobra.Command{
	Use:               "batchDelete",
	Short:             "Deletes printers in batch.",
	Long:              "https://developers.google.com/admin-sdk/chrome-printer/reference/rest/v1/admin.directory.v1.customers.chrome.printers/batchDeletePrinters",
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
		batchDeletePrintersRequest, err := mapToBatchDeletePrintersRequest(flags)
		if err != nil {
			log.Fatalf("Error building batch delete printer request object: %v", err)
		}
		result, err := gsmadmin.BatchDeletePrinters(parent, batchDeletePrintersRequest)
		if err != nil {
			log.Fatalf("Error deleting Chrome printers: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(chromePrintersCmd, chromePrintersBatchDeleteCmd, chromePrinterFlags)
}
