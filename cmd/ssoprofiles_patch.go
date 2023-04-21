/*
Copyright © 2020-2023 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// ssoProfilesPatchCmd represents the patch command
var ssoProfilesPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates an InboundSamlSsoProfile.",
	Long:              `Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/inboundSamlSsoProfiles/patch`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, updateMask, err := mapToSsoProfile(flags)
		if err != nil {
			log.Fatalf("Error building InboundSamlSsoProfile object: %v", err)
		}
		result, err := gsmci.PatchSsoProfile(flags["name"].GetString(), updateMask, flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error patching inbound SAML SSO profile: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(ssoProfilesCmd, ssoProfilesPatchCmd, ssoProfileFlags)
}
