/*
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

// userInvitationsSendCmd represents the send command
var userInvitationsSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends a UserInvitation to email.",
	Long: `If the UserInvitation does not exist for this request and it is a valid request, the request creates a UserInvitation.
Note: The get and list methods have a 48-hour delay where newly-created consumer accounts will not appear in the results.
You can still send a UserInvitation to those accounts if you know the unmanaged email address and userinvitations.isInvitableUser==True.
Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1beta1/customers.userinvitations/send`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		sendUserInvitationRequest, err := mapToSendUserInvitationRequest(flags)
		if err != nil {
			log.Fatalf("Error building sendUserInvitationRequest object: %v", err)
		}
		result, err := gsmcibeta.SendInvitation(flags["name"].GetString(), flags["fields"].GetString(), sendUserInvitationRequest)
		if err != nil {
			log.Fatalf("Error sending user invitation: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(userInvitationsCmd, userInvitationsSendCmd, userInvitationFlags)
}
