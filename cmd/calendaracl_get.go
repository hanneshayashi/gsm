/*
Copyright © 2020-2024 Hannes Hayashi

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

// calendarACLGetCmd represents the get command
var calendarACLGetCmd = &cobra.Command{
	Use:               "get",
	Short:             "Returns an access control rule.",
	Long:              `Implements the API documented at https://developers.google.com/calendar/api/v3/reference/acl/get`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmcalendar.GetACL(flags["calendarId"].GetString(), flags["ruleId"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error deleting calendar acl rule: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(calendarACLCmd, calendarACLGetCmd, calendarACLFlags)
}
