/*
Copyright © 2020-2024 Hannes Hayashi

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
	"github.com/hanneshayashi/gsm/gsmlicensing"
	"google.golang.org/api/licensing/v1"

	"github.com/spf13/cobra"
)

// licenseAssignmentsListForProductAndSkuCmd represents the listForProductAndSku command
var licenseAssignmentsListForProductAndSkuCmd = &cobra.Command{
	Use:               "listForProductAndSku",
	Short:             "List all users assigned licenses for a specific product SKU.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/licensing/reference/rest/v1/licenseAssignments/listForProductAndSku",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		customerID := gsmadmin.GetCustomerID(flags["customerId"].GetString())
		result, err := gsmlicensing.ListLicenseAssignmentsForProductAndSku(flags["productId"].GetString(), flags["skuId"].GetString(), customerID, flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*licensing.LicenseAssignment{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing license assignments for product and sku: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(licenseAssignmentsCmd, licenseAssignmentsListForProductAndSkuCmd, licenseAssignmentFlags)
}
