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

// resourcesCalendarsCmd represents the resourcesCalendars command
var resourcesCalendarsCmd = &cobra.Command{
	Use:   "resourcesCalendars",
	Short: "Manage resource calendars (Part of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/resources/calendars",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var resourcesCalendarFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
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
"CATEGORY_UNKNOWN"
"CONFERENCE_ROOM"
"OTHER"`,
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
var resourcesCalendarFlagsALL = gsmhelpers.GetAllFlags(resourcesCalendarFlags)

func init() {
	rootCmd.AddCommand(resourcesCalendarsCmd)
}

func mapToResourceCalendar(flags map[string]*gsmhelpers.Value) (*admin.CalendarResource, error) {
	resourceCalendar := &admin.CalendarResource{}
	if flags["resourceId"].IsSet() {
		resourceCalendar.ResourceId = flags["resourceId"].GetString()
		if resourceCalendar.ResourceId == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "ResourceId")
		}
	}
	if flags["resourceName"].IsSet() {
		resourceCalendar.ResourceName = flags["resourceName"].GetString()
		if resourceCalendar.ResourceName == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "ResourceName")
		}
	}
	if flags["buildingId"].IsSet() {
		resourceCalendar.BuildingId = flags["buildingId"].GetString()
		if resourceCalendar.BuildingId == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "BuildingId")
		}
	}
	if flags["capacity"].IsSet() {
		resourceCalendar.Capacity = flags["capacity"].GetInt64()
		if resourceCalendar.Capacity == 0 {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "Capacity")
		}
	}
	if flags["featureInstances"].IsSet() {
		resourceCalendar.FeatureInstances = flags["featureInstances"].GetString()
		if resourceCalendar.FeatureInstances == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "FeatureInstances")
		}
	}
	if flags["floorName"].IsSet() {
		resourceCalendar.FloorName = flags["floorName"].GetString()
		if resourceCalendar.FloorName == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "FloorName")
		}
	}
	if flags["floorSection"].IsSet() {
		resourceCalendar.FloorSection = flags["floorSection"].GetString()
		if resourceCalendar.FloorSection == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "FloorSection")
		}
	}
	if flags["resourceCategory"].IsSet() {
		resourceCalendar.ResourceCategory = flags["resourceCategory"].GetString()
		if resourceCalendar.ResourceCategory == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "ResourceCategory")
		}
	}
	if flags["resourceType"].IsSet() {
		resourceCalendar.ResourceType = flags["resourceType"].GetString()
		if resourceCalendar.ResourceType == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "ResourceType")
		}
	}
	if flags["userVisibleDescription"].IsSet() {
		resourceCalendar.UserVisibleDescription = flags["userVisibleDescription"].GetString()
		if resourceCalendar.UserVisibleDescription == "" {
			resourceCalendar.ForceSendFields = append(resourceCalendar.ForceSendFields, "UserVisibleDescription")
		}
	}
	return resourceCalendar, nil
}
