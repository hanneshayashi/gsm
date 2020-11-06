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

// membersCmd represents the members command
var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "Manage group members (Part of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/members",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var memberFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"groupKey": {
		AvailableFor: []string{"delete", "get", "hasMember", "insert", "list", "patch"},
		Type:         "string",
		Description: `Identifies the group in the API request.
The value can be the group's email address, group alias, or the unique group ID.`,
		Required: []string{"delete", "get", "hasMember", "insert", "list", "patch"},
	},
	"memberKey": {
		AvailableFor: []string{"delete", "get", "hasMember"},
		Type:         "string",
		Description: `Identifies the group member in the API request.
A group member can be a user or another group.
The value can be the member's (group or user) primary email address, alias, or unique ID.`,
		Required: []string{"delete", "get", "hasMember", "insert"},
	},
	"delivery_settings": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Defines mail delivery preferences of member.
Acceptable values are:
"ALL_MAIL": All messages, delivered as soon as they arrive.
"DAILY": No more than one message a day.
"DIGEST": Up to 25 messages bundled into a single message.
"DISABLED": Remove subscription.
"NONE": No messages.`,
	},
	"role": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a member of group2, group2 cannot be a member of group1. For more information about a member's role, see the administration help center.

Acceptable values are:
"MANAGER": This role is only available if the Google Groups for Business is enabled using the Admin console. A MANAGER role can do everything done by an OWNER role except make a member an OWNER or delete the group. A group can have multiple MANAGER members.
"MEMBER": This role can subscribe to a group, view discussion archives, and view the group's membership list. For more information about member roles, see the administration help center.
"OWNER": This role can send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group. A group can have more than one OWNER.`,
		Defaults: map[string]interface{}{"insert": "MEMBER"},
	},
	"includeDerivedMembership": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether to list indirect memberships.`,
	},
	"roles": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The roles query parameter allows you to retrieve group members by role.
Allowed values are OWNER, MANAGER, and MEMBER.`,
	},
	"email": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description: `The member's email address. A member can be a user or another group.
This property is required when adding a member to a group.
The email must be unique and cannot be an alias of another group.
If the email address is changed, the API automatically reflects the email address changes.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(membersCmd)
}

func mapToMember(flags map[string]*gsmhelpers.Value) (*admin.Member, error) {
	member := &admin.Member{}
	if flags["email"].IsSet() {
		member.Email = flags["email"].GetString()
		if member.Email == "" {
			member.ForceSendFields = append(member.ForceSendFields, "Email")
		}
	}
	if flags["delivery_settings"].IsSet() {
		member.DeliverySettings = flags["delivery_settings"].GetString()
		if member.DeliverySettings == "" {
			member.ForceSendFields = append(member.ForceSendFields, "DeliverySettings")
		}
	}
	if flags["role"].IsSet() {
		member.Role = flags["role"].GetString()
		if member.Role == "" {
			member.ForceSendFields = append(member.ForceSendFields, "Role")
		}
	}
	member.Kind = "admin#directory#member"
	return member, nil
}
