/*
Copyright © 2020-2023 Hannes Hayashi

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
	"github.com/hanneshayashi/gsm/gsmsheets"

	"github.com/spf13/cobra"
)

// spreadsheetsGetCmd represents the get command
var spreadsheetsGetCmd = &cobra.Command{
	Use:               "get",
	Short:             "Returns the spreadsheet at the given ID.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/sheets/api/reference/rest/v4/spreadsheets/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmsheets.GetSpreadsheet(flags["spreadsheetId"].GetString(), flags["fields"].GetString(), flags["ranges"].GetStringSlice(), flags["includeGridData"].GetBool())
		if err != nil {
			log.Fatalf("Error getting spreadsheet: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(spreadsheetsCmd, spreadsheetsGetCmd, spreadsheetFlags)
}
