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

	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// calendarsInsertCmd represents the insert command
var calendarsInsertCmd = &cobra.Command{
	Use:               "insert",
	Short:             `Creates a secondary calendar.`,
	Long:              "Implements the API documented at https://developers.google.com/calendar/api/v3/reference/calendars/insert",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToCalendar(flags)
		if err != nil {
			log.Fatalf("Error building calendar object: %v", err)
		}
		result, err := gsmcalendar.InsertCalendar(c, flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error inserting calendar: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(calendarsCmd, calendarsInsertCmd, calendarFlags)
}
