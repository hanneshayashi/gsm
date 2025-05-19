/*
Copyright © 2020-2025 Hannes Hayashi

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

// orgUnitsInsertCmd represents the insert command
var orgUnitsInsertCmd = &cobra.Command{
	Use:               "insert",
	Short:             "Adds an organizational unit.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/orgunits/insert",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		o, err := mapToOrgUnit(flags)
		if err != nil {
			log.Fatalf("Error building org unit object: %v", err)
		}
		result, err := gsmadmin.InsertOrgUnit(flags["customerId"].GetString(), flags["fields"].GetString(), o)
		if err != nil {
			log.Fatalf("Error inserting org unit: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(orgUnitsCmd, orgUnitsInsertCmd, orgUnitFlags)
}
