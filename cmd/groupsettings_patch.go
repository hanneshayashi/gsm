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
	"gsm/gsmgroupssettings"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// groupSettingsPatchCmd represents the patch command
var groupSettingsPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Updates an existing resource. This method supports patch semantics.",
	Long:  "https://developers.google.com/admin-sdk/groups-settings/v1/reference/groups/patch",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		g, err := mapToGroupSettings(flags)
		if err != nil {
			log.Fatalf("Error building group settings object: %v", err)
		}
		result, err := gsmgroupssettings.PatchGroupSettings(flags["groupUniqueId"].GetString(), flags["fields"].GetString(), g)
		if err != nil {
			log.Fatalf("Error patching group settings: %v", err)
		}
		if flags["ignoreDeprecated"].GetBool() {
			result = ignoreDeprecatedGroupSettings(result)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(groupSettingsCmd, groupSettingsPatchCmd, groupSettingFlags)
}
