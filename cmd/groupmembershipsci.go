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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	ci "google.golang.org/api/cloudidentity/v1beta1"
)

// groupMembershipsCiCmd represents the groupMembershipsCi command
var groupMembershipsCiCmd = &cobra.Command{
	Use:   "groupMembershipsCi",
	Short: "Manage group memberships (Part of Cloud Identity Beta API)",
	Long:  "https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups.memberships",	
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var groupMembershipCiFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"checkTransitiveMembership", "create", "getMembershipGraph", "list", "lookup", "searchTransitiveGroups", "searchTransitiveMemberships"},
		Type:         "string",
		Description: `Resource name of the group.
Format: groups/{group_id}, where group_id is the unique id assigned to the Group to which the Membership belongs to.`,
	},
	"query": {
		AvailableFor: []string{"checkTransitiveMembership", "getMembershipGraph", "searchTransitiveGroups"},
		Type:         "string",
		Description: `A CEL expression that MUST include:
getMembershipGraph      - member specification AND label(s)
                          (Example query: member_key_id == 'member_key_id_value' && <label_value> in labels)
list                    - member specification
                          (Example query: member_key_id == 'member_key_id_value')
searchTransitiveGroups  - member specification AND label(s)
                          Users can search on label attributes of groups. CONTAINS match ('in') is supported on labels.
                          (Example query: member_key_id == 'member_key_id_value' && <label_value> in labels)
Certain groups are uniquely identified by both a 'member_key_id' and a 'member_key_namespace', which requires an additional query input: 'member_key_namespace'.`,
		Required: []string{"checkTransitiveMembership", "getMembershipGraph", "searchTransitiveGroups"},
	},
	"memberKeyId": {
		AvailableFor: []string{"create", "lookup"},
		Type:         "string",
		Description: `The ID of the entity.

For Google-managed entities, the id must be the email address of an existing group or user.

For external-identity-mapped entities, the id must be a string conforming to the Identity Source's requirements.

Must be unique within a namespace.`,
		Required: []string{"create", "lookup"},
	},
	"memberKeyNamespace": {
		AvailableFor: []string{"create", "lookup"},
		Type:         "string",
		Description: `The namespace in which the entity exists.

If not specified, the EntityKey represents a Google-managed entity such as a Google user or a Google Group.

If specified, the EntityKey represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of identitysources/{identity_source_id}.`,
	},
	"roles": {
		AvailableFor: []string{"create"},
		Type:         "stringSlice",
		Description: `The MembershipRoles that apply to the Membership.

If unspecified, defaults to a single MembershipRole with name MEMBER.

Must not contain duplicate MembershipRoles with the same name.

Can be used multiple times in the form of "--roles name=...;expiryDate...
You may use the following properties:
name  - The name of the MembershipRole.
		Must be one of OWNER, MANAGER, MEMBER.
expireTime    - The time at which the MembershipRole will expire.
				A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
				Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
                Expiry details are only supported for MEMBER MembershipRoles.`,
		Required: []string{"create"},
	},
	"name": {
		AvailableFor: []string{"delete", "get", "modifyMembershipRoles"},
		Type:         "string",
		Description: `The resource name of the Membership.
Must be of the form groups/{group_id}/memberships/{membership_id}.`,
		ExcludeFromAll: true,
	},
	"email": {
		AvailableFor: []string{"checkTransitiveMembership", "create", "delete", "get", "getMembershipGraph", "list", "lookup", "searchTransitiveGroups", "searchTransitiveMemberships"},
		Type:         "string",
		Description: `Email address of the group.
This may be used instead of the name to do a lookup of the group resource name.
Note that this will cause an additional API call.`,
	},
	"view": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The level of detail to be returned.
BASIC  - Default. Only basic resource information is returned.
FULL   - All resource information is returned.`,
	},
	"addRoles": {
		AvailableFor: []string{"modifyMembershipRoles"},
		Type:         "stringSlice",
		Description: `The MembershipRoles to be added.

Adding or removing roles in the same request as updating roles is not supported.

Must not be set if updateRolesParams is set.
Can be used multiple times in the form of "--addRoles name=...;expiryDate...
You may use the following properties:
name        - The name of the MembershipRole.
              Must be one of OWNER, MANAGER, MEMBER.
expireTime  - The time at which the MembershipRole will expire.
			  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
			  Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
              Expiry details are only supported for MEMBER MembershipRoles.`,
	},
	"removeRoles": {
		AvailableFor: []string{"modifyMembershipRoles"},
		Type:         "stringSlice",
		Description: `The names of the MembershipRoles to be removed.

Adding or removing roles in the same request as updating roles is not supported.

It is not possible to remove the MEMBER MembershipRole. If you wish to delete a Membership, call MembershipsService.DeleteMembership instead.

Must not contain MEMBER. Must not be set if updateRolesParams is set.`,
	},
	"updateRolesParams": {
		AvailableFor: []string{"modifyMembershipRoles"},
		Type:         "stringSlice",
		Description: `The MembershipRoles to be updated.

Updating roles in the same request as adding or removing roles is not supported.

Must not be set if either addRoles or removeRoles is set.

Can be used multiple times in the form of "--updateRolesParams fieldMask=...;membershipRole=..."
You can use the following properties:
name        - The name of the MembershipRole.
		      Must be one of OWNER, MANAGER, MEMBER.
expireTime  - The time at which the MembershipRole will expire.
			  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
              Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
              Expiry details are only supported for MEMBER MembershipRoles.`,
	},
	"fields": {
		AvailableFor: []string{"create", "list", "get", "getMembershipGraph", "modifyMembershipRoles", "searchTransitiveGroups", "searchTransitiveMemberships"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(groupMembershipsCiCmd)
}

func mapToGroupMemberShipCi(flags map[string]*gsmhelpers.Value) (*ci.Membership, error) {
	membership := &ci.Membership{}
	if flags["memberKeyId"].IsSet() || flags["memberKeyNamespace"].IsSet() {
		membership.MemberKey = &ci.EntityKey{}
		if flags["memberKeyId"].IsSet() {
			membership.MemberKey.Id = flags["memberKeyId"].GetString()
			if membership.MemberKey.Id == "" {
				membership.MemberKey.ForceSendFields = append(membership.MemberKey.ForceSendFields, "Id")
			}
		}
		if flags["memberKeyNamespace"].IsSet() {
			membership.MemberKey.Namespace = flags["memberKeyNamespace"].GetString()
			if membership.MemberKey.Namespace == "" {
				membership.MemberKey.ForceSendFields = append(membership.MemberKey.ForceSendFields, "Namespace")
			}
		}
	}
	if flags["roles"].IsSet() {
		membership.Roles = []*ci.MembershipRole{}
		roles := flags["roles"].GetStringSlice()
		if len(roles) > 0 {
			for _, r := range roles {
				m := gsmhelpers.FlagToMap(r)
				role := &ci.MembershipRole{
					Name: m["name"],
				}
				if m["expireTime"] != "" {
					role.ExpiryDetail = &ci.ExpiryDetail{
						ExpireTime: m["expireTime"],
					}
				}
				membership.Roles = append(membership.Roles, role)
			}
		} else {
			membership.ForceSendFields = append(membership.ForceSendFields, "Roles")
		}
	}
	return membership, nil
}

func mapToModifyMembershipRolesRequestCi(flags map[string]*gsmhelpers.Value) (*ci.ModifyMembershipRolesRequest, error) {
	modifyMembershipRolesRequest := &ci.ModifyMembershipRolesRequest{}
	if flags["addRoles"].IsSet() {
		modifyMembershipRolesRequest.AddRoles = []*ci.MembershipRole{}
		addRoles := flags["addRoles"].GetStringSlice()
		if len(addRoles) > 0 {
			for _, r := range addRoles {
				m := gsmhelpers.FlagToMap(r)
				addRole := &ci.MembershipRole{
					Name: m["name"],
				}
				if m["expireTime"] != "" {
					addRole.ExpiryDetail = &ci.ExpiryDetail{
						ExpireTime: m["expireTime"],
					}
				}
				modifyMembershipRolesRequest.AddRoles = append(modifyMembershipRolesRequest.AddRoles, addRole)
			}
		} else {
			modifyMembershipRolesRequest.ForceSendFields = append(modifyMembershipRolesRequest.ForceSendFields, "AddRoles")
		}
	}
	if flags["removeRoles"].IsSet() {
		modifyMembershipRolesRequest.RemoveRoles = flags["removeRoles"].GetStringSlice()
		if len(modifyMembershipRolesRequest.RemoveRoles) > 0 {
			modifyMembershipRolesRequest.ForceSendFields = append(modifyMembershipRolesRequest.ForceSendFields, "RemoveRoles")
		}
	}
	if flags["updateRolesParams"].IsSet() {
		modifyMembershipRolesRequest.UpdateRolesParams = []*ci.UpdateMembershipRolesParams{}
		updateRolesParams := flags["updateRolesParams"].GetStringSlice()
		if len(updateRolesParams) > 0 {
			for _, u := range updateRolesParams {
				m := gsmhelpers.FlagToMap(u)
				updateRolesParam := &ci.UpdateMembershipRolesParams{
					FieldMask: m["fieldMask"],
					MembershipRole: &ci.MembershipRole{
						Name: m["name"],
					},
				}
				if m["expireTime"] != "" {
					updateRolesParam.MembershipRole.ExpiryDetail = &ci.ExpiryDetail{
						ExpireTime: m["expireTime"],
					}
					updateRolesParam.FieldMask = "expiryDetail"
				}
				modifyMembershipRolesRequest.UpdateRolesParams = append(modifyMembershipRolesRequest.UpdateRolesParams, updateRolesParam)
			}
		}
	}
	return modifyMembershipRolesRequest, nil
}
