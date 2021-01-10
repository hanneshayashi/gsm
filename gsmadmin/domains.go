/*
Package gsmadmin implements the Admin SDK APIs
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

// DeleteDomain deletes a domain of the customer.
func DeleteDomain(customerID, domainName string) (bool, error) {
	srv := getDomainsService()
	c := srv.Delete(customerID, domainName)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID, domainName), func() error {
		return c.Do()
	})
	return result, err
}

// GetDomain retrieves a domain of the customer.
func GetDomain(customerID, domainName, fields string) (*admin.Domains, error) {
	srv := getDomainsService()
	c := srv.Get(customerID, domainName)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, domainName), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Domains)
	return r, nil
}

// InsertDomain inserts a domain of the customer.
func InsertDomain(customerID, fields string, domain *admin.Domains) (*admin.Domains, error) {
	srv := getDomainsService()
	c := srv.Insert(customerID, domain)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, domain.DomainName), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Domains)
	return r, nil
}

// ListDomains lists the domains of the customer.
func ListDomains(customerID, fields string) ([]*admin.Domains, error) {
	srv := getDomainsService()
	c := srv.List(customerID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Domains2)
	return r.Domains, nil
}
