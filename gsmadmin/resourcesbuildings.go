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
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteBuilding deletes a building.
func DeleteBuilding(customer, buildingID string) (bool, error) {
	srv := getResourcesBuildingsService()
	c := srv.Delete(customer, buildingID)
	err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(customer, buildingID), err)
	}
	return true, nil
}

// GetBuilding retrieves a building.
func GetBuilding(customer, buildingID, fields string) (*admin.Building, error) {
	srv := getResourcesBuildingsService()
	c := srv.Get(customer, buildingID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(customer, buildingID), err)
	}
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(customer, building.BuildingName), err)
	}
	return r, nil
}

// ListBuildings retrieves a list of buildings for an account.
func ListBuildings(customer, fields string) iter.Seq2[*admin.Building, error] {
	srv := getResourcesBuildingsService()
	c := srv.List(customer).MaxResults(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*admin.Building, error) bool) {
		e := c.Pages(context.Background(), func(response *admin.Buildings) error {
			for i := range response.Buildings {
				if !yield(response.Buildings[i], nil) {
					return errIterStopped
				}
			}
			return nil
		})
		if e != nil && !errors.Is(e, errIterStopped) {
			yield(nil, e)
		}
	}
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(customer, buildingID), err)
	}
	return r, nil
}
