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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// orgUnitsPatchCmd represents the patch command
var orgUnitsPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates an organizational unit. This method supports patch semantics.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/orgunits/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		o, err := mapToOrgUnit(flags)
		if err != nil {
			log.Fatalf("Error building org unit object: %v", err)
		}
		result, err := gsmadmin.PatchOrgUnit(flags["customerId"].GetString(), flags["orgUnitPath"].GetString(), flags["fields"].GetString(), o)
		if err != nil {
			log.Fatalf("Error patching org unit: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(orgUnitsCmd, orgUnitsPatchCmd, orgUnitFlags)
}
