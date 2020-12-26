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

// sendAsCreateCmd represents the create command
var sendAsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: `Creates a custom "from" send-as alias`,
	Long: `If an SMTP MSA is specified, Gmail will attempt to connect to the SMTP service to validate the configuration before creating the alias.
If ownership verification is required for the alias, a message will be sent to the email address and the resource's verification status will be set to pending;
otherwise, the resource will be created with verification status set to accepted.
If a signature is provided, Gmail will sanitize the HTML before saving it with the alias.
https://developers.google.com/gmail/api/reference/rest/v1/users.settings.sendAs/create`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		s, err := mapToSendAs(flags)
		if err != nil {
			log.Fatalf("Error building send-as object: %v", err)
		}
		result, err := gsmgmail.CreateSendAs(flags["userId"].GetString(), flags["fields"].GetString(), s)
		if err != nil {
			log.Fatalf("Error creating send-as: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(sendAsCmd, sendAsCreateCmd, sendAsFlags)
}
