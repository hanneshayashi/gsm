/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmcibeta"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// userInvitationsIsInvitableUserCmd represents the isinvitableuser command
var userInvitationsIsInvitableUserCmd = &cobra.Command{
	Use:   "isInvitableUser",
	Short: "Retrieves a UserInvitation resource.",
	Long: `Eligibility is based on the following criteria:
 - the email address is a consumer account and it's the primary email address of the account, and
 - the domain of the email address matches an existing verified Google Workspace or Cloud Identity domain
If both conditions are met, the user is eligible.
https://cloud.google.com/identity/docs/reference/rest/v1beta1/customers.userinvitations/isInvitableUser`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmcibeta.IsInvitableUser(flags["name"].GetString())
		if err != nil {
			log.Fatalf("Error checking if user is invitable: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(userInvitationsCmd, userInvitationsIsInvitableUserCmd, userInvitationFlags)
}
