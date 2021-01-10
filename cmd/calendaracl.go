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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// calendarACLCmd represents the calendarAcl command
var calendarACLCmd = &cobra.Command{
	Use:               "calendarAcl",
	Short:             "Manage entries in users' calendar acl (Part of Calendar API)",
	Long:              "https://developers.google.com/calendar/v3/reference/calendarAcl",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var calendarACLFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"calendarId": {
		AvailableFor: []string{"delete", "get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Calendar identifier. To retrieve calendar IDs call the calendarAcl.acl method.
If you want to access the primary calendar of the currently logged in user, use the "primary" keyword.`,
		Defaults: map[string]interface{}{"delete": "primary", "get": "primary", "insert": "primary", "list": "primary", "patch": "primary"},
	},
	"ruleId": {
		AvailableFor:   []string{"delete", "get", "patch"},
		Type:           "string",
		Description:    `ACL rule identifier.`,
		Required:       []string{"delete", "get", "patch"},
		ExcludeFromAll: true,
	},
	"sendNotifications": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description: `Whether to send notifications about the calendar sharing change.
Optional. The default is True.`,
		Defaults: map[string]interface{}{"insert": true, "patch": true},
	},
	"role": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The role assigned to the scope. Possible values are:
"none" - Provides no access.
"freeBusyReader" - Provides read access to free/busy information.
"reader" - Provides read access to the calendar. Private events will appear to users with reader access, but event details will be hidden.
"writer" - Provides read and write access to the calendar. Private events will appear to users with writer access, and event details will be visible.
"owner" - Provides ownership of the calendar. This role has all of the permissions of the writer role with the additional ability to see and manipulate ACLs.`,
		Required: []string{"insert", "patch"},
	},
	"scopeType": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The type of the scope. Possible values are:
"default" - The public scope. This is the default value.
"user" - Limits the scope to a single user.
"group" - Limits the scope to a group.
"domain" - Limits the scope to a domain.
Note: The permissions granted to the "default", or public, scope apply to any user, authenticated or not.`,
		Defaults: map[string]interface{}{"insert": "default"},
	},
	"scopeValue": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The email address of a user or group, or the name of a domain, depending on the scope type.
Omitted for type "default".`,
	},
	"showDeleted": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description: `Whether to include deleted ACLs in the result.
Deleted ACLs are represented by role equal to "none".
Deleted ACLs will always be included if syncToken is provided.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var calendarACLFlagsALL = gsmhelpers.GetAllFlags(calendarACLFlags)

func init() {
	rootCmd.AddCommand(calendarACLCmd)
}
func mapToCalendarACLRule(flags map[string]*gsmhelpers.Value) (*calendar.AclRule, error) {
	aclRule := &calendar.AclRule{}
	if flags["role"].IsSet() {
		aclRule.Role = flags["role"].GetString()
		if aclRule.Role == "" {
			aclRule.ForceSendFields = append(aclRule.ForceSendFields, "Role")
		}
	}
	if flags["scopeType"].IsSet() || flags["scopeValue"].IsSet() {
		aclRule.Scope = &calendar.AclRuleScope{}
		if flags["scopeType"].IsSet() {
			aclRule.Scope.Type = flags["scopeType"].GetString()
			if aclRule.Scope.Type == "" {
				aclRule.Scope.ForceSendFields = append(aclRule.Scope.ForceSendFields, "Type")
			}
		}
		if flags["scopeValue"].IsSet() {
			aclRule.Scope.Value = flags["scopeValue"].GetString()
			if aclRule.Scope.Value == "" {
				aclRule.Scope.ForceSendFields = append(aclRule.Scope.ForceSendFields, "Value")
			}
		}
	}
	return aclRule, nil
}
