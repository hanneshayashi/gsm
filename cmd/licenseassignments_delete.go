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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlicensing"

	"github.com/spf13/cobra"
)

// licenseAssignmentsDeleteCmd represents the delete command
var licenseAssignmentsDeleteCmd = &cobra.Command{
	Use:               "delete",
	Short:             "Delete a specific user's license by product SKU.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/licensing/reference/rest/v1/licenseAssignments/delete",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmlicensing.DeleteLicenseAssignment(flags["productId"].GetString(), flags["skuId"].GetString(), flags["userId"].GetString())
		if err != nil {
			log.Fatalf("Error deleting licenseAssignment: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(licenseAssignmentsCmd, licenseAssignmentsDeleteCmd, licenseAssignmentFlags)
}
