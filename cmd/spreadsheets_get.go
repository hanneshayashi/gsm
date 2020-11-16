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
	"fmt"
	"gsm/gsmhelpers"
	"gsm/gsmsheets"
	"log"

	"github.com/spf13/cobra"
)

// spreadsheetsGetCmd represents the get command
var spreadsheetsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a spreadsheet, returning the newly getd spreadsheet.",
	Long:  "https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets/get",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmsheets.GetSpreadsheet(flags["spreadsheetId"].GetString(), flags["fields"].GetString(), flags["ranges"].GetStringSlice(), flags["includeGridData"].GetBool())
		if err != nil {
			log.Fatalf("Error getting spreadsheet: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	spreadsheetsCmd.AddCommand(spreadsheetsGetCmd)
	gsmhelpers.AddFlags(spreadsheetFlags, spreadsheetsGetCmd.Flags(), spreadsheetsGetCmd.Use)
	markFlagsRequired(spreadsheetsGetCmd, spreadsheetFlags, "")
}
