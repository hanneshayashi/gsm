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
	"gsm/gsmgmail"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// gmailSettingsUpdateAutoForwardingCmd represents the updateAutoForwarding command
var gmailSettingsUpdateAutoForwardingCmd = &cobra.Command{
	Use: "updateAutoForwarding",
	Short: `Updates the auto-forwarding setting for the specified account.
A verified forwarding address must be specified when auto-forwarding is enabled.`,
	Long: "https://developers.google.com/gmail/api/reference/rest/v1/users.settings/updateAutoForwarding",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		a, err := mapToAutoforwarding(flags)
		if err != nil {
			log.Fatalf("Error building auto-forwarding object: %v", err)
		}
		result, err := gsmgmail.UpdateAutoForwardingSettings(flags["userId"].GetString(), flags["fields"].GetString(), a)
		if err != nil {
			log.Fatalf("Error updating auto-forwarding settings for user %s: %v", flags["userId"].GetString(), err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(gmailSettingsCmd, gmailSettingsUpdateAutoForwardingCmd, gmailSettingFlags)
}
