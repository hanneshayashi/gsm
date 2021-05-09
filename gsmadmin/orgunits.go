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

package gsmadmin

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteOrgUnit removes an organizational unit.
func DeleteOrgUnit(customerID, orgUnitPath string) (bool, error) {
	srv := getOrgunitsService()
	c := srv.Delete(customerID, orgUnitPath)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID, orgUnitPath), func() error {
		return c.Do()
	})
	return result, err
}

// GetOrgUnit retrieves an organizational unit.
func GetOrgUnit(customerID, orgUnitPath, fields string) (*admin.OrgUnit, error) {
	srv := getOrgunitsService()
	c := srv.Get(customerID, orgUnitPath)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, orgUnitPath), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.OrgUnit)
	return r, nil
}

// InsertOrgUnit adds an organizational unit.
func InsertOrgUnit(customerID, fields string, OrgUnit *admin.OrgUnit) (*admin.OrgUnit, error) {
	srv := getOrgunitsService()
	c := srv.Insert(customerID, OrgUnit)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, OrgUnit.Name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.OrgUnit)
	return r, nil
}

// ListOrgUnits retrieves a list of all organizational units for an account.
func ListOrgUnits(customerID, t, orgUnitPath, fields string) ([]*admin.OrgUnit, error) {
	srv := getOrgunitsService()
	c := srv.List(customerID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if orgUnitPath != "" {
		c = c.OrgUnitPath(orgUnitPath)
	}
	if t != "" {
		c = c.Type(t)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.OrgUnits)
	return r.OrganizationUnits, nil
}

// PatchOrgUnit updates an organizational unit. This method supports patch semantics.
func PatchOrgUnit(customerID, orgUnitPath, fields string, OrgUnit *admin.OrgUnit) (*admin.OrgUnit, error) {
	srv := getOrgunitsService()
	c := srv.Patch(customerID, orgUnitPath, OrgUnit)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, orgUnitPath), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.OrgUnit)
	return r, nil
}
