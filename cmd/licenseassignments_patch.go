/*
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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlicensing"

	"github.com/spf13/cobra"
)

// licenseAssignmentsPatchCmd represents the patch command
var licenseAssignmentsPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Patch a specific user's license by product SKU.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/licensing/v1/reference/licenseAssignments/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		l, err := mapToLicenseAssignment(flags)
		if err != nil {
			log.Fatalf("Error building licenseAssignmentPatch object: %v", err)
		}
		result, err := gsmlicensing.PatchLicenseAssignment(flags["productId"].GetString(), flags["skuId"].GetString(), flags["userId"].GetString(), flags["fields"].GetString(), l)
		if err != nil {
			log.Fatalf("Error patching licenseAssignment: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(licenseAssignmentsCmd, licenseAssignmentsPatchCmd, licenseAssignmentFlags)
}
