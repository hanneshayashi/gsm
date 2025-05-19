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

// ssoAssignmentsCreateCmd represents the create command
var ssoAssignmentsCreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Creates an InboundSsoAssignment for users and devices in a Customer under a given Group or OrgUnit.",
	Long:              `Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/inboundSsoAssignments/create`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, _, err := mapToSsoAssignment(flags)
		if err != nil {
			log.Fatalf("Error building InboundSamlSsoAssignment object: %v", err)
		}
		result, err := gsmci.CreateSsoAssignment(flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error createing inbound SAML SSO assignment: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(ssoAssignmentsCmd, ssoAssignmentsCreateCmd, ssoAssignmentFlags)
}
