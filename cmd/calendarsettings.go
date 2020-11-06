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
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// calendarSettingsCmd represents the calendarSettings command
var calendarSettingsCmd = &cobra.Command{
	Use:   "calendarSettings",
	Short: "See users' calendar settings (Part of Calendar API)",
	Long:  "https://developers.google.com/calendar/v3/reference/settings",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var calendarSettingFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"setting": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Calendar identifier. To retrieve calendar IDs call the calendarList.list method.
If you want to access the primary calendar of the currently logged in user, use the "primary" keyword.`,
		Required: []string{"get"},
	},
	"fields": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(calendarSettingsCmd)
}
