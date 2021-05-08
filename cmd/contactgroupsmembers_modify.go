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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmpeople"

	"github.com/spf13/cobra"
)

// contactGroupsMembersModifyCmd represents the modify command
var contactGroupsMembersModifyCmd = &cobra.Command{
	Use: "modify",
	Short: `Modify the members of a contact group owned by the authenticated user.
The only system contact groups that can have members added are contactGroups/myContacts and contactGroups/starred.
Other system contact groups are deprecated and can only have contacts removed.`,
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/contactGroups.members/modify",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		m, err := mapToModifyContactGroupMembersRequest(flags)
		if err != nil {
			log.Fatalf("Error building ModifyContactGroupMembersRequest object: %v", err)
		}
		result, err := gsmpeople.ModifyContactGroupMembers(flags["resourceName"].GetString(), flags["fields"].GetString(), m)
		if err != nil {
			log.Fatalf("Error creating contact group: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(contactGroupsMembersCmd, contactGroupsMembersModifyCmd, contactGroupMemberFlags)
}
