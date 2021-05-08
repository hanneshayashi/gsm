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

	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// eventsImportCmd represents the import command
var eventsImportCmd = &cobra.Command{
	Use: "import",
	Short: `Imports an event.
This operation is used to add a private copy of an existing event to a calendar.`,
	Long:              "https://developers.google.com/calendar/v3/reference/events/import",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		e, err := gsmcalendar.GetEvent(flags["calendarI"].GetString(), flags["eventId"].GetString(), "", "*", 0)
		if err != nil {
			log.Fatalf("Error getting source event: %v", err)
		}
		result, err := gsmcalendar.ImportEvent(flags["destination"].GetString(), flags["fields"].GetString(), e, flags["conferenceDataVersion"].GetInt64(), flags["supportsAttachments"].GetBool())
		if err != nil {
			log.Fatalf("Error importing event: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(eventsCmd, eventsImportCmd, eventFlags)
}
