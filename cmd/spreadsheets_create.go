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

// spreadsheetsCreateCmd represents the create command
var spreadsheetsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a spreadsheet, returning the newly created spreadsheet.",
	Long:  "https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets/create",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		s, err := mapToSpreadsheet(flags)
		if err != nil {
			log.Fatalf("Error building spreadsheet object: %v\n", err)
		}
		result, err := gsmsheets.CreateSpreadsheet(s, flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error creating spreadsheet: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	spreadsheetsCmd.AddCommand(spreadsheetsCreateCmd)
	gsmhelpers.AddFlags(spreadsheetFlags, spreadsheetsCreateCmd.Flags(), spreadsheetsCreateCmd.Use)
	markFlagsRequired(spreadsheetsCreateCmd, spreadsheetFlags, "")
}