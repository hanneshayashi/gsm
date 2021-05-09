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
	"google.golang.org/api/calendar/v3"

	"github.com/spf13/cobra"
)

// eventsListCmd represents the list command
var eventsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Returns events on the specified calendar.",
	Long:              "Implements the API documented at https://developers.google.com/calendar/v3/reference/events/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmcalendar.ListEvents(flags["calendarId"].GetString(), flags["iCalUID"].GetString(), flags["orderBy"].GetString(), flags["q"].GetString(), flags["timeZone"].GetString(), flags["timeMax"].GetString(), flags["timeMin"].GetString(), flags["updatedMin"].GetString(), flags["fields"].GetString(), flags["privateExtendedProperty"].GetStringSlice(), flags["sharedExtendedProperty"].GetStringSlice(), flags["maxAttendees"].GetInt64(), flags["showDeleted"].GetBool(), flags["showHiddenInvitations"].GetBool(), flags["singleEvents"].GetBool(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*calendar.Event{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing events: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(eventsCmd, eventsListCmd, eventFlags)
}
