/*
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
package gsmadmin

import (
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteResourcesBuilding deletes a building.
func DeleteResourcesBuilding(customer, buildingID string) (bool, error) {
	srv := getResourcesBuildingsService()
	err := srv.Delete(customer, buildingID).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetResourcesBuilding retrieves a building.
func GetResourcesBuilding(customer, buildingID, fields string) (*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.Get(customer, buildingID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertResourcesBuilding inserts a building.
func InsertResourcesBuilding(customer, coordinatesSource, fields string, building *admin.Building) (*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.Insert(customer, building)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if coordinatesSource != "" {
		c = c.CoordinatesSource(coordinatesSource)
	}
	r, err := c.Do()
	return r, err
}

func makeListResourceBuildingsCallAndAppend(c *admin.ResourcesBuildingsListCall, buildings []*admin.Building) ([]*admin.Building, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, b := range r.Buildings {
		buildings = append(buildings, b)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		buildings, err = makeListResourceBuildingsCallAndAppend(c, buildings)
	}
	return buildings, err
}

// ListResourcesBuildings retrieves a list of buildings for an account.
func ListResourcesBuildings(customer, fields string) ([]*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.List(customer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var buildings []*admin.Building
	buildings, err := makeListResourceBuildingsCallAndAppend(c, buildings)
	return buildings, err
}

// PatchResourcesBuilding updates a building. This method supports patch semantics.
func PatchResourcesBuilding(customer, buildingID, coordinatesSource, fields string, building *admin.Building) (*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.Patch(customer, buildingID, building)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if coordinatesSource != "" {
		c = c.CoordinatesSource(coordinatesSource)
	}
	r, err := c.Do()
	return r, err
}
