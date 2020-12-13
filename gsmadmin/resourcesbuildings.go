/*
Package gsmadmin implements the Admin SDK APIs
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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteBuilding deletes a building.
func DeleteBuilding(customer, buildingID string) (bool, error) {
	srv := getResourcesBuildingsService()
	c := srv.Delete(customer, buildingID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customer, buildingID), func() error {
		return c.Do()
	})
	return result, err
}

// GetBuilding retrieves a building.
func GetBuilding(customer, buildingID, fields string) (*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.Get(customer, buildingID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, buildingID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Building)
	return r, nil
}

// InsertBuilding inserts a building.
func InsertBuilding(customer, coordinatesSource, fields string, building *admin.Building) (*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.Insert(customer, building)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if coordinatesSource != "" {
		c = c.CoordinatesSource(coordinatesSource)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, building.BuildingName), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Building)
	return r, nil
}

func makeListBuildingsCallAndAppend(c *admin.ResourcesBuildingsListCall, buildings []*admin.Building, errKey string) ([]*admin.Building, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Buildings)
	buildings = append(buildings, r.Buildings...)
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		buildings, err = makeListBuildingsCallAndAppend(c, buildings, errKey)
	}
	return buildings, err
}

// ListBuildings retrieves a list of buildings for an account.
func ListBuildings(customer, fields string) ([]*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.List(customer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var buildings []*admin.Building
	buildings, err := makeListBuildingsCallAndAppend(c, buildings, gsmhelpers.FormatErrorKey(customer))
	return buildings, err
}

// PatchBuilding updates a building. This method supports patch semantics.
func PatchBuilding(customer, buildingID, coordinatesSource, fields string, building *admin.Building) (*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.Patch(customer, buildingID, building)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if coordinatesSource != "" {
		c = c.CoordinatesSource(coordinatesSource)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, buildingID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Building)
	return r, nil
}
