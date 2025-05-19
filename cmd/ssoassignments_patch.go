/*
Copyright © 2020-2025 Hannes Hayashi

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

// ssoAssignmentsPatchCmd represents the patch command
var ssoAssignmentsPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates an InboundSsoAssignment.",
	Long:              `Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/inboundSsoAssignments/patch`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, updateMask, err := mapToSsoAssignment(flags)
		if err != nil {
			log.Fatalf("Error building InboundSamlSsoAssignment object: %v", err)
		}
		result, err := gsmci.PatchSsoAssignment(flags["name"].GetString(), updateMask, flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error patching inbound SAML SSO assignment: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(ssoAssignmentsCmd, ssoAssignmentsPatchCmd, ssoAssignmentFlags)
}
