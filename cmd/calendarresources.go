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

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// calendarResourcesCmd represents the calendarResources command
var calendarResourcesCmd = &cobra.Command{
	Use:               "calendarResources",
	Short:             "Manage resource calendars (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/resources/calendars",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var calendarResourceFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"calendarResourceId": {
		AvailableFor:   []string{"delete", "get", "patch"},
		Type:           "string",
		Description:    `The unique ID of the calendar resource`,
		Required:       []string{"delete", "get", "patch"},
		ExcludeFromAll: true,
	},
	"customer": {
		AvailableFor: []string{"delete", "get", "insert", "list", "patch"},
		Type:         "string",
		Description: `The unique ID for the customer's Workspace account.
As an account administrator, you can also use the my_customer alias to represent your account's customer ID.`,
		Defaults: map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer", "patch": "my_customer"},
	},
	"resourceId": {
		AvailableFor:   []string{"insert"},
		Type:           "string",
		Description:    `The unique ID of the calendar resource`,
		Required:       []string{"insert"},
		ExcludeFromAll: true,
	},
	"resourceName": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `The name of the calendar resource. For example, "Training Room 1A".`,
		Required:     []string{"insert"},
	},
	"buildingId": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Unique ID for the building a resource is located in.`,
	},
	"capacity": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "int64",
		Description:  `Capacity of a resource, number of seats in a room.`,
	},
	"featureInstances": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Instances of features for the calendar resource.`,
	},
	"floorName": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Name of the floor a resource is located on.`,
	},
	"floorSection": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Name of the section within a floor a resource is located in.`,
	},
	"resourceCategory": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The category of the calendar resource. Either CONFERENCE_ROOM or OTHER. Legacy data is set to CATEGORY_UNKNOWN.

Acceptable values are:
CATEGORY_UNKNOWN
CONFERENCE_ROOM
OTHER`,
	},
	"resourceDescription": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Description of the resource, visible only to admins.`,
	},
	"resourceType": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `The type of the calendar resource, intended for non-room resources.`,
	},
	"userVisibleDescription": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Description of the resource, visible to users and admins.`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Field(s) to sort results by in either ascending or descending order.
Supported fields include resourceId, resourceName, capacity, buildingId, and floorName.
If no order is specified, defaults to ascending.
Should be of the form "field [asc|desc], field [asc|desc], ...".
For example buildingId, capacity desc would return results sorted first by buildingId in ascending order then by capacity in descending order.`,
	},
	"query": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `	String query used to filter results.
Should be of the form "field operator value" where field can be any of supported fields and operators can be any of supported operations.
Operators include '=' for exact match and ':' for prefix match or HAS match, depending on type of field.
For ':', when the field supports a scalar value, such as a String, and the value is followed by an asterisk (*), the query is considered a prefix match.
In a prefix match, the value must be at the start of a string to be a match.
For example, resourceName:Conference* returns all strings whose resourceName starts with "Conference," such as "Conference-Room-1."
For ':', when the field supports repeated values, such as featureInstances[].features, use a colon (:) without an asterisk (*) to indicate a HAS match.
For example, featureInstances.feature.name:Phone would return any calendar resource that has a feature instance whose name is "Phone" (all rooms with phones).
An asterisk (*) is only valid at end of value, it cannot be used at start or middle of value. For example, resourceName:*Room* doesn't work.
Query strings are case sensitive.
Supported fields include generatedResourceName, resourceName, name, buildingId, featureInstances.feature.name.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var calendarResourceFlagsALL = gsmhelpers.GetAllFlags(calendarResourceFlags)

func init() {
	rootCmd.AddCommand(calendarResourcesCmd)
}

func mapToCalendarResource(flags map[string]*gsmhelpers.Value) (*admin.CalendarResource, error) {
	calendarResource := &admin.CalendarResource{}
	if flags["resourceId"].IsSet() {
		calendarResource.ResourceId = flags["resourceId"].GetString()
		if calendarResource.ResourceId == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "ResourceId")
		}
	}
	if flags["resourceName"].IsSet() {
		calendarResource.ResourceName = flags["resourceName"].GetString()
		if calendarResource.ResourceName == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "ResourceName")
		}
	}
	if flags["buildingId"].IsSet() {
		calendarResource.BuildingId = flags["buildingId"].GetString()
		if calendarResource.BuildingId == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "BuildingId")
		}
	}
	if flags["capacity"].IsSet() {
		calendarResource.Capacity = flags["capacity"].GetInt64()
		if calendarResource.Capacity == 0 {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "Capacity")
		}
	}
	if flags["featureInstances"].IsSet() {
		calendarResource.FeatureInstances = flags["featureInstances"].GetString()
		if calendarResource.FeatureInstances == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "FeatureInstances")
		}
	}
	if flags["floorName"].IsSet() {
		calendarResource.FloorName = flags["floorName"].GetString()
		if calendarResource.FloorName == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "FloorName")
		}
	}
	if flags["floorSection"].IsSet() {
		calendarResource.FloorSection = flags["floorSection"].GetString()
		if calendarResource.FloorSection == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "FloorSection")
		}
	}
	if flags["resourceCategory"].IsSet() {
		calendarResource.ResourceCategory = flags["resourceCategory"].GetString()
		if calendarResource.ResourceCategory == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "ResourceCategory")
		}
	}
	if flags["resourceType"].IsSet() {
		calendarResource.ResourceType = flags["resourceType"].GetString()
		if calendarResource.ResourceType == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "ResourceType")
		}
	}
	if flags["userVisibleDescription"].IsSet() {
		calendarResource.UserVisibleDescription = flags["userVisibleDescription"].GetString()
		if calendarResource.UserVisibleDescription == "" {
			calendarResource.ForceSendFields = append(calendarResource.ForceSendFields, "UserVisibleDescription")
		}
	}
	return calendarResource, nil
}
