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

// orgUnitsListCmd represents the list command
var orgUnitsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Retrieves a list of all organizational units for an account.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/orgunits/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.ListOrgUnits(flags["customerId"].GetString(), flags["type"].GetString(), flags["orgUnitPath"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error listing org units: %v", err)
		}
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(result[i])
			}
		} else {
			gsmhelpers.Output(result, "json", compressOutput)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(orgUnitsCmd, orgUnitsListCmd, orgUnitFlags)
}
