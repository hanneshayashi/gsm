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
	"gsm/gsmgroupssettings"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// groupSettingsGetCmd represents the get command
var groupSettingsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a group's settings identified by the group email address.",
	Long:  "https://developers.google.com/admin-sdk/groups-settings/v1/reference/groups/get",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgroupssettings.GetGroupSettings(flags["groupUniqueId"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting group settings: %v", err)
		}
		if flags["ignoreDeprecated"].GetBool() {
			result = ignoreDeprecatedGroupSettings(result)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(groupSettingsCmd, groupSettingsGetCmd, groupSettingFlags)
}
