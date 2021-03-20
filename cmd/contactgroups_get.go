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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmpeople"

	"github.com/spf13/cobra"
)

// contactGroupsGetCmd represents the get command
var contactGroupsGetCmd = &cobra.Command{
	Use:               "get",
	Short:             "Get a specific contact group owned by the authenticated user by specifying a contact group resource name.",
	Long:              "https://developers.google.com/people/api/rest/v1/contactGroups/get",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmpeople.GetContactGroup(flags["resourceName"].GetString(), flags["fields"].GetString(), flags["maxMembers"].GetInt64())
		if err != nil {
			log.Fatalf("Error getting contact group: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(contactGroupsCmd, contactGroupsGetCmd, contactGroupFlags)
}
