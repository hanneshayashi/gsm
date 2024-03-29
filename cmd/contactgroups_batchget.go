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

// contactGroupsBatchGetCmd represents the batchGet command
var contactGroupsBatchGetCmd = &cobra.Command{
	Use:               "batchGet",
	Short:             "Get a list of contact groups owned by the authenticated user by specifying a list of contact group resource names.",
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/contactGroups/batchGet",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmpeople.BatchGetContactGroups(flags["resourceNames"].GetStringSlice(), flags["maxMembers"].GetInt64(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error getting contact groups: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(contactGroupsCmd, contactGroupsBatchGetCmd, contactGroupFlags)
}
