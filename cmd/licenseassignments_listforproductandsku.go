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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"gsm/gsmlicensing"
	"log"

	"github.com/spf13/cobra"
)

// licenseAssignmentsListForProductAndSkuCmd represents the listForProductAndSku command
var licenseAssignmentsListForProductAndSkuCmd = &cobra.Command{
	Use:   "listForProductAndSku",
	Short: "List all users assigned licenses for a specific product SKU.",
	Long:  "https://developers.google.com/admin-sdk/licensing/v1/reference/licenseAssignments/listForProductAndSku",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		customerID := gsmadmin.GetCustomerID(flags["customerId"].GetString())
		result, err := gsmlicensing.ListLicenseAssignmentsForProductAndSku(flags["productId"].GetString(), flags["skuId"].GetString(), customerID, flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error listing license assignments for product: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	gsmhelpers.InitCommand(licenseAssignmentsCmd, licenseAssignmentsListForProductAndSkuCmd, licenseAssignmentFlags)
}
