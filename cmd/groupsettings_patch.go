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

	"github.com/hanneshayashi/gsm/gsmgroupssettings"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// groupSettingsPatchCmd represents the patch command
var groupSettingsPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates an existing resource. This method supports patch semantics.",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/groups-settings/v1/reference/groups/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
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
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupSettingsCmd, groupSettingsPatchCmd, groupSettingFlags)
}
