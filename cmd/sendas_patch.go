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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// sendAsPatchCmd represents the patch command
var sendAsPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Patch the specified send-as alias.",
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.settings.sendAs/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		s, err := mapToSendAs(flags)
		if err != nil {
			log.Fatalf("Error building send-as object: %v", err)
		}
		result, err := gsmgmail.PatchSendAs(flags["userId"].GetString(), flags["sendAsEmail"].GetString(), flags["fields"].GetString(), s)
		if err != nil {
			log.Fatalf("Error patching send-as: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(sendAsCmd, sendAsPatchCmd, sendAsFlags)
}
