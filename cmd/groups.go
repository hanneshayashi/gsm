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
	admin "google.golang.org/api/admin/directory/v1"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Implements the groups API (Part of Admin SDK).",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/groups",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var groupFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"groupKey": {
		AvailableFor: []string{"delete", "get", "patch"},
		Type:         "string",
		Description: `Identifies the group in the API request.
The value can be the group's email address, group alias, or the unique group ID.`,
		Required:       []string{"delete", "get", "patch"},
		ExcludeFromAll: true,
	},
	"email": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Identifies the group in the API request.
The value can be the group's email address, group alias, or the unique group ID.`,
		Required:       []string{"insert"},
		ExcludeFromAll: true,
	},
	"description": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `An extended description to help users determine the purpose of a group.
For example, you can include information about who should join the group, the types of messages to send to the group, links to FAQs about the group, or related groups.
Maximum length is 4,096 characters.`,
	},
	"name": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `The group's display name.`,
	},
	"customer": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The unique ID for the customer's G Suite account.
In case of a multi-domain account, to fetch all groups for a customer, fill this field instead of domain.
As an account administrator, you can also use the my_customer alias to represent your account's customerId.
The customerId is also returned as part of the Users resource.`,
		Defaults: map[string]interface{}{"list": "my_customer"},
	},
	"domain": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The domain name.
Use this field to get fields from only one domain.
To return all domains for a customer account, use the customer query parameter instead.`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Column to use for sorting results
Acceptable values are:
"email": Email of the group.`,
	},
	"query": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Query string search. Should be of the form "".
Complete documentation is at https://developers.google.com/admin-sdk/directory/v1/guides/search-groups`,
	},
	"sortOrder": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Whether to return results in ascending or descending order. Only of use when orderBy is also used
Acceptable values are:
"ASCENDING": Ascending order.
"DESCENDING": Descending order.`,
	},
	"userKey": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Email or immutable ID of the user if only those groups are to be listed, the given user is a member of.
If it's an ID, it should match with the ID of the user object.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var groupFlagsALL = gsmhelpers.GetAllFlags(groupFlags)

func init() {
	rootCmd.AddCommand(groupsCmd)
}

func mapToGroup(flags map[string]*gsmhelpers.Value) (*admin.Group, error) {
	group := &admin.Group{}
	if flags["email"].IsSet() {
		group.Email = flags["email"].GetString()
		if group.Email == "" {
			group.ForceSendFields = append(group.ForceSendFields, "Email")
		}
	}
	if flags["description"].IsSet() {
		group.Description = flags["description"].GetString()
		if group.Description == "" {
			group.ForceSendFields = append(group.ForceSendFields, "Description")
		}
	}
	if flags["name"].IsSet() {
		group.Name = flags["name"].GetString()
		if group.Name == "" {
			group.ForceSendFields = append(group.ForceSendFields, "Name")
		}
	}
	return group, nil
}
