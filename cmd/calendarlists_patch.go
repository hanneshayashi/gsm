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

// calendarListsPatchCmd represents the patch command
var calendarListsPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates an existing calendar on the user's calendar list. This method supports patch semantics.",
	Long:              "Implements the API documented at https://developers.google.com/calendar/v3/reference/calendarList/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		calendarListEntry, err := mapToCalendarListEntry(flags)
		if err != nil {
			log.Fatalf("Error building calendarListEntry object: %v", err)
		}
		result, err := gsmcalendar.PatchCalendarListEntry(flags["calendarId"].GetString(), flags["fields"].GetString(), calendarListEntry, flags["colorRgbFormat"].GetBool())
		if err != nil {
			log.Fatalf("Error patching calendarListEntry: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(calendarListsCmd, calendarListsPatchCmd, calendarListFlags)
}
