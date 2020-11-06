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

// DeleteOrgUnit removes an organizational unit.
func DeleteOrgUnit(customerID, orgUnitPath string) (bool, error) {
	srv := getOrgunitsService()
	err := srv.Delete(customerID, orgUnitPath).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetOrgUnit retrieves an organizational unit.
func GetOrgUnit(customerID, orgUnitPath, fields string) (*admin.OrgUnit, error) {
	srv := getOrgunitsService()
	c := srv.Get(customerID, orgUnitPath)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertOrgUnit adds an organizational unit.
func InsertOrgUnit(customerID, fields string, OrgUnit *admin.OrgUnit) (*admin.OrgUnit, error) {
	srv := getOrgunitsService()
	c := srv.Insert(customerID, OrgUnit)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
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
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.OrganizationUnits, nil
}

// PatchOrgUnit updates an organizational unit. This method supports patch semantics.
func PatchOrgUnit(customerID, orgUnitPath, fields string, OrgUnit *admin.OrgUnit) (*admin.OrgUnit, error) {
	srv := getOrgunitsService()
	c := srv.Patch(customerID, orgUnitPath, OrgUnit)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
