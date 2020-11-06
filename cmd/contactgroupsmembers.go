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
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/people/v1"
)

// contactGroupsMembersCmd represents the contactGroupsMembers command
var contactGroupsMembersCmd = &cobra.Command{
	Use:   "contactGroupsMembers",
	Short: "Modify members of contact groups (Part of People API)",
	Long:  "https://developers.google.com/people/api/rest/v1/contactGroups.members",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var contactGroupMemberFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"resourceName": {
		AvailableFor: []string{"modify"},
		Type:         "string",
		Description:  `The resource name of the contact group to modify.`,
		Required:     []string{"modify"},
	},
	"resourceNamesToAdd": {
		AvailableFor: []string{"modify"},
		Type:         "stringSlice",
		Description:  `The resource names of the contact people to add in the form of people/{person_id}.`,
	},
	"resourceNamesToRemove": {
		AvailableFor: []string{"modify"},
		Type:         "stringSlice",
		Description:  `The resource names of the contact people to remove in the form of people/{person_id}.`,
	},
	"fields": {
		AvailableFor: []string{"modify"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(contactGroupsMembersCmd)
}

func mapToModifyContactGroupMembersRequest(flags map[string]*gsmhelpers.Value) (*people.ModifyContactGroupMembersRequest, error) {
	modifyContactGroupMembersRequest := &people.ModifyContactGroupMembersRequest{}
	if flags["resourceNamesToAdd"].IsSet() {
		modifyContactGroupMembersRequest.ResourceNamesToAdd = flags["resourceNamesToAdd"].GetStringSlice()
	}
	if flags["resourceNamesToRemove"].IsSet() {
		modifyContactGroupMembersRequest.ResourceNamesToRemove = flags["resourceNamesToRemove"].GetStringSlice()
	}
	return modifyContactGroupMembersRequest, nil
}
