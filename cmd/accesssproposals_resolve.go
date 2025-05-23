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

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// accessProposalsResolveCmd represents the resolve command
var accessProposalsResolveCmd = &cobra.Command{
	Use:               "resolve",
	Short:             "Used to approve or deny an Access Proposal.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/accessproposals/resolve",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		request, err := mapToResolveAccessProposalRequest(flags)
		if err != nil {
			log.Fatalf("Error building access proposal request object: %v", err)
		}
		result, err := gsmdrive.ResolveAccessProposal(flags["fileId"].GetString(), flags["proposalId"].GetString(), request)
		if err != nil {
			log.Fatalf("Error resolving access proposal: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(accessProposalsCmd, accessProposalsResolveCmd, accessProposalsFlags)
}
