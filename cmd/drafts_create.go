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
	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// draftsCreateCmd represents the create command
var draftsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new draft with the DRAFT label.",
	Long:  "https://developers.google.com/gmail/api/reference/rest/v1/users.drafts/create",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		d, err := mapToDraft(flags)
		if err != nil {
			log.Fatalf("Error building draft object: %v\n", err)
		}
		result, err := gsmgmail.CreateDraft(flags["userId"].GetString(), flags["fields"].GetString(), d)
		if err != nil {
			log.Fatalf("Error creating draft: %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(draftsCmd, draftsCreateCmd, draftFlags)
}
