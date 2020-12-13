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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlicensing"

	"github.com/spf13/cobra"
)

// licenseAssignmentsGetCmd represents the get command
var licenseAssignmentsGetCmd = &cobra.Command{
	Use:               "get",
	Short:             "Get a specific user's license by product SKU.",
	Long:              "https://developers.google.com/admin-sdk/licensing/v1/reference/licenseAssignments/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmlicensing.GetLicenseAssignment(flags["productId"].GetString(), flags["skuId"].GetString(), flags["userId"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting licenseAssignment: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(licenseAssignmentsCmd, licenseAssignmentsGetCmd, licenseAssignmentFlags)
}
