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

// DeleteDomainAlias deletes a Domain Alias of the customer.
func DeleteDomainAlias(customerID, domainAliasName string) (bool, error) {
	srv := getDomainAliasesService()
	err := srv.Delete(customerID, domainAliasName).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetDomainAlias retrieves a domain alias of the customer.
func GetDomainAlias(customerID, domainAliasName, fields string) (*admin.DomainAlias, error) {
	srv := getDomainAliasesService()
	c := srv.Get(customerID, domainAliasName)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertDomainAlias inserts a Domain alias of the customer.
func InsertDomainAlias(customerID, fields string, domainAlias *admin.DomainAlias) (*admin.DomainAlias, error) {
	srv := getDomainAliasesService()
	c := srv.Insert(customerID, domainAlias)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// ListDomainAliases lists the domain aliases of the customer.
func ListDomainAliases(customerID, parentDomainName, fields string) ([]*admin.DomainAlias, error) {
	srv := getDomainAliasesService()
	c := srv.List(customerID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if parentDomainName != "" {
		c = c.ParentDomainName(parentDomainName)
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.DomainAliases, nil
}
