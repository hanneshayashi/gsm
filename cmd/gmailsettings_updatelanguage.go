/*
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

// gmailSettingsUpdateLanguageCmd represents the updateLanguage command
var gmailSettingsUpdateLanguageCmd = &cobra.Command{
	Use:   "updateLanguage",
	Short: "Updates language settings.",
	Long: `If successful, the return object contains the displayLanguage that was saved for the user, which may differ from the value passed into the request.
This is because the requested displayLanguage may not be directly supported by Gmail but have a close variant that is, and so the variant may be chosen and saved instead.
Implements the API documented at https://developers.google.com/workspace/gmail/api/reference/rest/v1/users.settings/updateLanguage`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		l, err := mapToLanguageSettings(flags)
		if err != nil {
			log.Fatalf("Error building language settings object: %v", err)
		}
		result, err := gsmgmail.UpdateLanguageSettings(flags["userId"].GetString(), flags["fields"].GetString(), l)
		if err != nil {
			log.Fatalf("Error updating language settings for user %s: %v", flags["userId"].GetString(), err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(gmailSettingsCmd, gmailSettingsUpdateLanguageCmd, gmailSettingFlags)
}
