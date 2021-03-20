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

// contactGroupsUpdateCmd represents the update command
var contactGroupsUpdateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Update the name of an existing contact group owned by the authenticated user.",
	Long:              "https://developers.google.com/people/api/rest/v1/contactGroups/update",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := gsmpeople.GetContactGroup(flags["resourceName"].GetString(), "*", 0)
		if err != nil {
			log.Fatalf("Error getting contact group: %v", err)
		}
		u, err := mapToUpdateContactGroupRequest(flags, c)
		if err != nil {
			log.Fatalf("Error building updateContactGroupRequest object: %v", err)
		}
		result, err := gsmpeople.UpdateContactGroup(flags["resourceName"].GetString(), flags["fields"].GetString(), u)
		if err != nil {
			log.Fatalf("Error creating contact group: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(contactGroupsCmd, contactGroupsUpdateCmd, contactGroupFlags)
}
