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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// resourcesCalendarsPatchCmd represents the patch command
var resourcesCalendarsPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patches a calendar resource.",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/resources/calendars/patch",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToResourceCalendar(flags)
		if err != nil {
			log.Fatalf("Error building resourceCalendar object: %v", err)

		}
		result, err := gsmadmin.PatchResourcesCalendar(flags["customer"].GetString(), flags["calendarResourceId"].GetString(), flags["fields"].GetString(), c)
		if err != nil {
			log.Fatalf("Error patching calendar resource %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(resourcesCalendarsCmd, resourcesCalendarsPatchCmd, resourcesCalendarFlags)
}
