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

	"github.com/spf13/cobra"
	"google.golang.org/api/people/v1"
)

// contactGroupsCmd represents the contactGroups command
var contactGroupsCmd = &cobra.Command{
	Use:               "contactGroups",
	Short:             "Manage users' contact groups (Part of People API)",
	Long:              "https://developers.google.com/people/api/rest/v1/contactGroups",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var contactGroupFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"resourceNames": {
		AvailableFor: []string{"batchGet"},
		Type:         "stringSlice",
		Description:  `The resource names of the contact groups.`,
		Required:     []string{"batchGet"},
	},
	"maxMembers": {
		AvailableFor: []string{"batchGet", "get"},
		Type:         "int64",
		Description: `Specifies the maximum number of members to return for each group.
Defaults to 0 if not set, which will return zero members.`,
	},
	"name": {
		AvailableFor:   []string{"create", "update"},
		Type:           "string",
		Description:    `The contact group name set by the group owner or a system provided name for system groups.`,
		Required:       []string{"create"},
		ExcludeFromAll: true,
	},
	"resourceName": {
		AvailableFor:   []string{"delete", "get", "update"},
		Type:           "string",
		Description:    `The resource name of the contact group.`,
		Required:       []string{"delete", "get", "update"},
		ExcludeFromAll: true,
	},
	"deleteContacts": {
		AvailableFor: []string{"delete"},
		Type:         "bool",
		Description:  `Set to true to also delete the contacts in the specified group.`,
	},
	"fields": {
		AvailableFor: []string{"batchGet", "create", "get", "list", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var contactGroupFlagsALL = gsmhelpers.GetAllFlags(contactGroupFlags)

func init() {
	rootCmd.AddCommand(contactGroupsCmd)
}

func mapToCreateContactGroupRequest(flags map[string]*gsmhelpers.Value) (*people.CreateContactGroupRequest, error) {
	createContactGroupRequest := &people.CreateContactGroupRequest{}
	createContactGroupRequest.ContactGroup = &people.ContactGroup{}
	createContactGroupRequest.ContactGroup.Name = flags["name"].GetString()
	return createContactGroupRequest, nil
}

func mapToUpdateContactGroupRequest(flags map[string]*gsmhelpers.Value, contactGroup *people.ContactGroup) (*people.UpdateContactGroupRequest, error) {
	updateContactGroupRequest := &people.UpdateContactGroupRequest{}
	updateContactGroupRequest.ContactGroup = contactGroup
	if flags["name"].IsSet() {
		updateContactGroupRequest.ContactGroup.Name = flags["name"].GetString()
		if updateContactGroupRequest.ContactGroup.Name == "" {
			updateContactGroupRequest.ContactGroup.ForceSendFields = append(updateContactGroupRequest.ContactGroup.ForceSendFields, "Name")
		}
	}
	return updateContactGroupRequest, nil
}
