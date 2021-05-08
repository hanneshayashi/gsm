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
	"errors"
	"log"

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	ci "google.golang.org/api/cloudidentity/v1"
)

// groupsCiCmd represents the groupsCi command
var groupsCiCmd = &cobra.Command{
	Use:               "groupsCi",
	Short:             "Manage Google Groups with the Cloud Identity API",
	Long:              "https://cloud.google.com/identity/docs/reference/rest/v1/groups",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var groupCiFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"initialGroupConfig": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `Required. The initial configuration option for the Group.
WITH_INITIAL_OWNER  - The end user making the request will be added as the initial owner of the Group.
EMPTY               - An empty group is created without any initial owners.
                      This can only be used by admins of the domain.`,
		Defaults: map[string]interface{}{"create": "EMPTY"},
	},
	"labels": {
		AvailableFor: []string{"create", "patch"},
		Type:         "stringSlice",
		Description: ` One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.

Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value.

Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added.

Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic.

Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.

Examples: {"cloudidentity.googleapis.com/groups.discussion_forum": ""} or {"system/groups/external": ""}.

An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
	},
	"name": {
		AvailableFor: []string{"get", "delete", "patch"},
		Type:         "string",
		Description: `The resource name of the Group.

Must be of the form groups/{group_id}.`,
		ExcludeFromAll: true,
	},
	"email": {
		AvailableFor: []string{"get", "delete", "patch"},
		Type:         "string",
		Description: `Email address of the group.
This may be used instead of the name to do a lookup of the group resource name.
Note that this will cause an additional API call.`,
		ExcludeFromAll: true,
	},
	"parent": {
		AvailableFor: []string{"create", "list"},
		Type:         "string",
		Description:  `Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_id} for Google Groups.`,
	},
	"queries": {
		AvailableFor: []string{"create"},
		Type:         "stringArray",
		Description: `Memberships will be the union of all queries.
Only one entry with USER resource is currently supported.
Can be used multiple times in the form of "--queries query=...;resourceType=..."
You may use the following properties:
resourceType  - The following values are valid:
				  - USER - For queries on User
query         - Query that determines the memberships of the dynamic group.
				Examples:
			      - All users with at least one organizations.department of engineering:
                    user.organizations.exists(org, org.department=='engineering')
                  - All users with at least one location that has area of foo and building_id of bar:
                    user.locations.exists(loc, loc.area=='foo' && loc.building_id=='bar')`,
	},
	"id": {
		AvailableFor: []string{"create", "lookup"},
		Type:         "string",
		Description: `The ID of the entity.

For Google-managed entities, the id must be the email address.

For external-identity-mapped entities, the id must be a string conforming to the Identity Source's requirements.

Must be unique within a namespace.`,
		Required:       []string{"create", "lookup"},
		ExcludeFromAll: true,
	},
	"namespace": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `The namespace in which the entity exists.

If not specified, the EntityKey represents a Google-managed entity such as a Google user or a Google Group.

If specified, the EntityKey represents an external-identity-mapped group.
The namespace must correspond to an identity source created in Admin Console and must be in the form of identitysources/{identity_source_id}.`,
	},
	// 	"additionalGroupKeys": {
	// 		AvailableFor: []string{"create"},
	// 		Type:         "stringSlice",
	// 		Description: `Additional entity key aliases for a Group.
	// Can be used multiple times in the form of "id=...;namespace=..."`,
	// },
	"displayName": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `The display name of the Group.`,
	},
	"description": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description: `An extended description to help users determine the purpose of a Group.
Must not be longer than 4,096 characters.`,
	},
	"view": {
		AvailableFor: []string{"list", "search"},
		Type:         "string",
		Description: `The level of detail to be returned.
BASIC  - Default. Only basic resource information is returned.
FULL   - All resource information is returned.`,
	},
	"query": {
		AvailableFor: []string{"search"},
		Type:         "string",
		Description: `The search query.
Must be specified in Common Expression Language.
May only contain equality operators on the parent and inclusion operators on labels (e.g., parent == 'customers/{customer_id}' && 'cloudidentity.googleapis.com/groups.discussion_forum' in labels).`,
	},
	"updateMask": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The fully-qualified names of fields to update.

May only contain the following fields: displayName, description.

A comma-separated list of fully qualified names of fields. Example: "user.displayName,photo".`,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "list", "lookup", "patch", "search"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var groupCiFlagsALL = gsmhelpers.GetAllFlags(groupCiFlags)

func init() {
	rootCmd.AddCommand(groupsCiCmd)
}

func getGroupCiName(name, email string) (string, error) {
	if name != "" {
		return name, nil
	}
	if email == "" {
		return "", errors.New("either name or email must be supplied")
	}
	name, err := gsmci.LookupGroup(email)
	if err != nil {
		return "", err
	}
	return name, nil
}

func mapToGroupCi(flags map[string]*gsmhelpers.Value) (*ci.Group, error) {
	group := &ci.Group{}
	if flags["labels"].IsSet() {
		group.Labels = make(map[string]string)
		labels := flags["labels"].GetStringSlice()
		if len(labels) > 0 {
			for i := range labels {
				group.Labels[labels[i]] = ""
			}
		} else {
			group.ForceSendFields = append(group.ForceSendFields, "Labels")
		}
	}
	if flags["parent"].IsSet() {
		group.Parent = flags["parent"].GetString()
		if group.Parent == "" {
			group.ForceSendFields = append(group.ForceSendFields, "Parent")
		}
	}
	if flags["id"].IsSet() || flags["namespace"].IsSet() {
		group.GroupKey = &ci.EntityKey{}
		if flags["id"].IsSet() {
			group.GroupKey.Id = flags["id"].GetString()
			if group.GroupKey.Id == "" {
				group.GroupKey.ForceSendFields = append(group.ForceSendFields, "Id")
			}
		}
		if flags["namespace"].IsSet() {
			group.GroupKey.Namespace = flags["namespace"].GetString()
			if group.GroupKey.Namespace == "" {
				group.GroupKey.ForceSendFields = append(group.ForceSendFields, "Namespace")
			}
		}
	}
	// if flags["additionalGroupKeys"].IsSet() {
	// 	group.AdditionalGroupKeys = []*ci.EntityKey{}
	// 	additionalGroupKeys := flags["additionalGroupKeys"].GetStringSlice()
	// 	if len(additionalGroupKeys) > 0 {
	// 		for i := range additionalGroupKeys {
	// 			m := gsmhelpers.FlagToMap(a)
	// 			group.AdditionalGroupKeys = append(group.AdditionalGroupKeys, &ci.EntityKey{Id: m["id"], Namespace: m["namespace"]})
	// 		}
	// 	} else {
	// 		group.ForceSendFields = append(group.ForceSendFields, "AdditionalGroupKeys")
	// 	}
	// }
	if flags["queries"].IsSet() {
		group.DynamicGroupMetadata = &ci.DynamicGroupMetadata{}
		queries := flags["queries"].GetStringSlice()
		if len(queries) > 0 {
			group.DynamicGroupMetadata.Queries = []*ci.DynamicGroupQuery{}
			for i := range queries {
				m := gsmhelpers.FlagToMap(queries[i])
				group.DynamicGroupMetadata.Queries = append(group.DynamicGroupMetadata.Queries, &ci.DynamicGroupQuery{ResourceType: m["resourceType"], Query: m["query"]})
			}

		} else {
			group.DynamicGroupMetadata.ForceSendFields = append(group.DynamicGroupMetadata.ForceSendFields, "Queries")
		}
	}
	if flags["displayName"].IsSet() {
		group.DisplayName = flags["displayName"].GetString()
		if group.DisplayName == "" {
			group.ForceSendFields = append(group.ForceSendFields, "DisplayName")
		}
	}
	return group, nil
}
