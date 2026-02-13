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
	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlicensing"

	"github.com/spf13/cobra"
)

// licenseAssignmentsListForProductCmd represents the listForProduct command
var licenseAssignmentsListForProductCmd = &cobra.Command{
	Use:               "listForProduct",
	Short:             "List all users assigned licenses for a specific product SKU.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/admin/licensing/reference/rest/v1/licenseAssignments/listForProduct",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		customerID := gsmadmin.GetCustomerID(flags["customerId"].GetString())
		gsmhelpers.StreamOrCollect(gsmlicensing.ListLicenseAssignmentsForProduct(flags["productId"].GetString(), customerID, flags["fields"].GetString()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(licenseAssignmentsCmd, licenseAssignmentsListForProductCmd, licenseAssignmentFlags)
}
