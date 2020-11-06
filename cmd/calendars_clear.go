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
	"gsm/gsmcalendar"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// calendarsClearCmd represents the clear command
var calendarsClearCmd = &cobra.Command{
	Use: "clear",
	Short: `Clears a primary calendar.
This operation deletes all events associated with the primary calendar of an account.`,
	Long: "https://developers.google.com/calendar/v3/reference/calendars/clear",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmcalendar.ClearCalendar(flags["calendarId"].GetString())
		if err != nil {
			log.Fatalf("Error clearing calendar: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	calendarsCmd.AddCommand(calendarsClearCmd)
	gsmhelpers.AddFlags(calendarFlags, calendarsClearCmd.Flags(), calendarsClearCmd.Use)
	markFlagsRequired(calendarsClearCmd, calendarFlags, "")
}
