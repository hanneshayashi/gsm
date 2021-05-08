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

// schemasInsertCmd represents the insert command
var schemasInsertCmd = &cobra.Command{
	Use:               "insert",
	Short:             "Create a custom schema.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/v1/reference/schemas/insert",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		s, err := mapToSchema(flags)
		if err != nil {
			log.Fatalf("Error building schema object: %v", err)
		}
		result, err := gsmadmin.InsertSchema(flags["customerId"].GetString(), flags["fields"].GetString(), s)
		if err != nil {
			log.Fatalf("Error inserting schema: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(schemasCmd, schemasInsertCmd, schemaFlags)
}
