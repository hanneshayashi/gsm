/*
Copyright © 2020-2022 Hannes Hayashi

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
	"log"

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// groupsCiGetSecuritySettingsCmd represents the getSecuritySettings command
var groupsCiGetSecuritySettingsCmd = &cobra.Command{
	Use:               "getSecuritySettings",
	Short:             "Retrieves the security settings (member restrictions) of a group.",
	Long:              "Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/groups/getSecuritySettings",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		name, err := getGroupCiName(flags["name"].GetString(), flags["email"].GetString())
		if err != nil {
			log.Fatalf("%v", err)
		}
		result, err := gsmci.GetSecuritySettings(fmt.Sprintf("%s/securitySettings", name), flags["readMask"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting group's security settings: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupsCiCmd, groupsCiGetSecuritySettingsCmd, groupCiFlags)
}
