/*
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

// freeBusyQueryCmd represents the query command
var freeBusyQueryCmd = &cobra.Command{
	Use:               "query",
	Short:             "Returns free/busy information for a set of calendars.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/calendar/api/v3/reference/freebusy/query",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		f, err := mapToFreeBusyRequest(flags)
		if err != nil {
			log.Fatalf("Error building freeBusyRequest object: %v", err)
		}
		result, err := gsmcalendar.QueryFreeBusy(f, flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting IMAP settings for user %s: %v", flags["userId"].GetString(), err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(freeBusyCmd, freeBusyQueryCmd, freeBusyFlags)
}
