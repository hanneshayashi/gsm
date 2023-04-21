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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmpeople"

	"github.com/spf13/cobra"
)

// peopleGetBatchGetCmd represents the getBatchGet command
var peopleGetBatchGetCmd = &cobra.Command{
	Use:               "getBatchGet",
	Short:             "Provides information about a list of specific people by specifying a list of requested resource names.",
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/people/getBatchGet",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmpeople.GetContactsBatch(flags["resourceNames"].GetStringSlice(), flags["personFields"].GetString(), flags["sources"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting contacts: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(peopleCmd, peopleGetBatchGetCmd, peopleFlags)
}
