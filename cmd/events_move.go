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
	"log"

	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// eventsMoveCmd represents the move command
var eventsMoveCmd = &cobra.Command{
	Use:               "move",
	Short:             "Moves an event to another calendar, i.e. changes an event's organizer.",
	Long:              "https://developers.google.com/calendar/v3/reference/events/move",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmcalendar.MoveEvent(flags["calendarId"].GetString(), flags["eventId"].GetString(), flags["destination"].GetString(), flags["sendUpdates"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error moving event: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(eventsCmd, eventsMoveCmd, eventFlags)
}
