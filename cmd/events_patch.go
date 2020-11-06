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

// eventsPatchCmd represents the patch command
var eventsPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Updates an event. This method supports patch semantics.",
	Long:  "https://developers.google.com/calendar/v3/reference/events/patch",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		event, err := mapToEvent(flags)
		if err != nil {
			log.Fatalf("Error building event object: %v", err)
		}
		result, err := gsmcalendar.PatchEvent(flags["calendarId"].GetString(), flags["eventId"].GetString(), flags["sendUpdates"].GetString(), flags["fields"].GetString(), event, flags["conferenceDataVersion"].GetInt64(), flags["maxAttendees"].GetInt64(), flags["supportsAttachments"].GetBool())
		if err != nil {
			log.Fatalf("Error patching calendar event: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	eventsCmd.AddCommand(eventsPatchCmd)
	gsmhelpers.AddFlags(eventFlags, eventsPatchCmd.Flags(), eventsPatchCmd.Use)
	markFlagsRequired(eventsPatchCmd, eventFlags, "")
}
