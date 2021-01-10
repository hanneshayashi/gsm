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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// draftsGetCmd represents the get command
var draftsGetCmd = &cobra.Command{
	Use:               "get",
	Short:             "Gets the specified draft.",
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.drafts/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		if !gsmgmail.FormatIsValid(flags["format"].GetString()) {
			log.Fatalf("%s is not a format", flags["format"].GetString())
		}
		result, err := gsmgmail.GetDraft(flags["userId"].GetString(), flags["id"].GetString(), flags["format"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting draft: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(draftsCmd, draftsGetCmd, draftFlags)
}
