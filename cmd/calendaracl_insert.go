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

// calendarACLInsertCmd represents the insert command
var calendarACLInsertCmd = &cobra.Command{
	Use:               "insert",
	Short:             "Creates an access control rule.",
	Long:              `https://developers.google.com/calendar/v3/reference/acl/insert`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		a, err := mapToCalendarACLRule(flags)
		if err != nil {
			log.Fatalf("Error building acl rule object: %v", err)
		}
		result, err := gsmcalendar.InsertACL(flags["calendarId"].GetString(), flags["fields"].GetString(), a, flags["sendNotifications"].GetBool())
		if err != nil {
			log.Fatalf("Error inserting calendar acl rule: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(calendarACLCmd, calendarACLInsertCmd, calendarACLFlags)
}
