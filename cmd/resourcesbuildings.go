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
	admin "google.golang.org/api/admin/directory/v1"
)

// resourcesBuildingsCmd represents the resourcesBuildings command
var resourcesBuildingsCmd = &cobra.Command{
	Use:   "resourcesBuildings",
	Short: "Manage Buildings (Resources) (Part of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/resources/buildings",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var resourcesBuildingFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"buildingId": {
		AvailableFor:   []string{"delete", "get", "insert", "patch"},
		Type:           "string",
		Description:    `The ID of the file.`,
		Required:       []string{"delete", "get", "insert", "patch"},
		ExcludeFromAll: true,
	},
	"customer": {
		AvailableFor: []string{"delete", "get", "insert", "list", "patch"},
		Type:         "string",
		Description:  `The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer alias to represent your account's customer ID.`,
		Defaults:     map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer", "patch": "my_customer"},
	},
	"coordinatesSource": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Source from which Building.coordinates are derived.

Acceptable values are:
CLIENT_SPECIFIED       - Building.coordinates are set to the coordinates included in the request.
RESOLVED_FROM_ADDRESS  - Building.coordinates are automatically populated based on the postal address.
SOURCE_UNSPECIFIED     - Defaults to RESOLVED_FROM_ADDRESS if postal address is provided. Otherwise, defaults to CLIENT_SPECIFIED if coordinates are provided. (default)`,
	},
	"addressLines": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description:  `Unstructured address lines describing the lower levels of an address.`,
	},
	"administrativeArea": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Optional. Highest administrative subdivision which is used for postal addresses of a country or region.`,
	},
	"languageCode": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Optional. BCP-47 language code of the contents of this address (if known).`,
	},
	"locality": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Optional. Generally refers to the city/town portion of the address.
Examples: US city, IT comune, UK post town.
In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.`,
	},
	"postalCode": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Optional. Postal code of the address.`,
	},
	"regionCode": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Required. CLDR region code of the country/region of the address.`,
		Required:     []string{"insert"},
	},
	"sublocality": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Optional. Sublocality of the address.`,
	},
	"buildingName": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The building name as seen by users in Calendar.
Must be unique for the customer. For example, "NYC-CHEL".
The maximum length is 100 characters.`,
	},
	"latitude": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "float64",
		Description:  `Latitude in decimal degrees.`,
	},
	"longitude": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "float64",
		Description:  `Longitude in decimal degrees.`,
	},
	"description": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `A brief description of the building. For example, "Chelsea Market".`,
	},
	"floorNames": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The display names for all floors in this building.
The floors are expected to be sorted in ascending order, from lowest floor to highest floor.
For example, ["B2", "B1", "L", "1", "2", "2M", "3", "PH"] Must contain at least one entry.`,
		Required: []string{"insert"},
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var resourcesBuildingFlagsALL = gsmhelpers.GetAllFlags(resourcesBuildingFlags)

func init() {
	rootCmd.AddCommand(resourcesBuildingsCmd)
}

func mapToBuilding(flags map[string]*gsmhelpers.Value) (*admin.Building, error) {
	building := &admin.Building{}
	if flags["buildingId"].IsSet() {
		building.BuildingId = flags["buildingId"].GetString()
		if building.BuildingId == "" {
			building.ForceSendFields = append(building.ForceSendFields, "BuildingId")
		}
	}
	if flags["addressLines"].IsSet() || flags["administrativeArea"].IsSet() || flags["languageCode"].IsSet() || flags["locality"].IsSet() || flags["postalCode"].IsSet() || flags["regionCode"].IsSet() || flags["sublocality"].IsSet() {
		building.Address = &admin.BuildingAddress{}
		if flags["addressLines"].IsSet() {
			building.Address.AddressLines = flags["addressLines"].GetStringSlice()
			if len(building.Address.AddressLines) == 0 {
				building.Address.ForceSendFields = append(building.Address.ForceSendFields, "AddressLines")
			}
		}
		if flags["administrativeArea"].IsSet() {
			building.Address.AdministrativeArea = flags["administrativeArea"].GetString()
			if building.Address.AdministrativeArea == "" {
				building.Address.ForceSendFields = append(building.Address.ForceSendFields, "AdministrativeArea")
			}
		}
		if flags["languageCode"].IsSet() {
			building.Address.LanguageCode = flags["languageCode"].GetString()
			if building.Address.LanguageCode == "" {
				building.Address.ForceSendFields = append(building.Address.ForceSendFields, "LanguageCode")
			}
		}
		if flags["locality"].IsSet() {
			building.Address.Locality = flags["locality"].GetString()
			if building.Address.Locality == "" {
				building.Address.ForceSendFields = append(building.Address.ForceSendFields, "Locality")
			}
		}
		if flags["postalCode"].IsSet() {
			building.Address.PostalCode = flags["postalCode"].GetString()
			if building.Address.PostalCode == "" {
				building.Address.ForceSendFields = append(building.Address.ForceSendFields, "PostalCode")
			}
		}
		if flags["regionCode"].IsSet() {
			building.Address.RegionCode = flags["regionCode"].GetString()
			if building.Address.RegionCode == "" {
				building.Address.ForceSendFields = append(building.Address.ForceSendFields, "RegionCode")
			}
		}
		if flags["sublocality"].IsSet() {
			building.Address.Sublocality = flags["sublocality"].GetString()
			if building.Address.Sublocality == "" {
				building.Address.ForceSendFields = append(building.Address.ForceSendFields, "Sublocality")
			}
		}
	}
	if flags["buildingName"].IsSet() {
		building.BuildingName = flags["buildingName"].GetString()
		if building.BuildingName == "" {
			building.ForceSendFields = append(building.ForceSendFields, "BuildingName")
		}
	}
	if flags["latitude"].IsSet() || flags["longitude"].IsSet() {
		building.Coordinates = &admin.BuildingCoordinates{}
		if flags["latitude"].IsSet() {
			building.Coordinates.Latitude = flags["latitude"].GetFloat64()
			if building.Coordinates.Latitude == 0.0 {
				building.Coordinates.ForceSendFields = append(building.Coordinates.ForceSendFields, "Latitude")
			}
		}
		if flags["longitude"].IsSet() {
			building.Coordinates.Longitude = flags["longitude"].GetFloat64()
			if building.Coordinates.Longitude == 0.0 {
				building.Coordinates.ForceSendFields = append(building.Coordinates.ForceSendFields, "Longitude")
			}
		}
	}
	if flags["description"].IsSet() {
		building.Description = flags["description"].GetString()
		if building.Description == "" {
			building.ForceSendFields = append(building.ForceSendFields, "Description")
		}
	}
	if flags["floorNames"].IsSet() {
		building.FloorNames = flags["floorNames"].GetStringSlice()
		if len(building.FloorNames) == 0 {
			building.ForceSendFields = append(building.ForceSendFields, "FloorNames")
		}
	}
	return building, nil
}
