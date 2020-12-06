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
	"gsm/gsmhelpers"
	"gsm/gsmsheets"
	"log"

	"github.com/spf13/cobra"
)

// spreadsheetsBatchUpdateCmd represents the batchupdate command
var spreadsheetsBatchUpdateCmd = &cobra.Command{
	Use:   "batchUpdate",
	Short: "Applies one or more updates to the spreadsheet.",
	Long:  "https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets/batchupdate",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		b, err := mapToBatchUpdateSpreadsheetRequest(flags)
		if err != nil {
			log.Fatalf("Error building spreadsheet object: %v\n", err)
		}
		result, err := gsmsheets.BatchUpdateSpreadsheet(flags["spreadsheetId"].GetString(), flags["fields"].GetString(), b)
		if err != nil {
			log.Fatalf("Error creating spreadsheet: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(spreadsheetsCmd, spreadsheetsBatchUpdateCmd, spreadsheetFlags)
}
