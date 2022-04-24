/*
Copyright © 2020-2022 Hannes Hayashi

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

// eventsInsertCmd represents the insert command
var eventsInsertCmd = &cobra.Command{
	Use:               "insert",
	Short:             "Creates an event.",
	Long:              "Implements the API documented at https://developers.google.com/calendar/v3/reference/events/insert",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		event, err := mapToEvent(flags)
		if err != nil {
			log.Fatalf("Error building event object: %v", err)
		}
		result, err := gsmcalendar.InsertEvent(flags["calendarId"].GetString(), flags["sendUpdates"].GetString(), flags["fields"].GetString(), event, flags["conferenceDataVersion"].GetInt64(), flags["maxAttendees"].GetInt64(), flags["supportsAttachments"].GetBool())
		if err != nil {
			log.Fatalf("Error inserting calendar event: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(eventsCmd, eventsInsertCmd, eventFlags)
}
