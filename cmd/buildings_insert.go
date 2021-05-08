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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// buildingsInsertCmd represents the insert command
var buildingsInsertCmd = &cobra.Command{
	Use:               "insert",
	Short:             "Inserts a building.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/resources/buildings/insert",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		b, err := mapToBuilding(flags)
		if err != nil {
			log.Fatalf("Error building building object: %v", err)
		}
		result, err := gsmadmin.InsertBuilding(flags["customer"].GetString(), flags["coordinatesSource"].GetString(), flags["fields"].GetString(), b)
		if err != nil {
			log.Fatalf("Error inserting building: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(buildingsCmd, buildingsInsertCmd, buildingFlags)
}
