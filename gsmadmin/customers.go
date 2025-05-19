/*
Copyright © 2020-2025 Hannes Hayashi

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
	"fmt"
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// GetCustomer retrieves a customer.
func GetCustomer(id, fields string) (*admin.Customer, error) {
	srv := getCustomersService()
	c := srv.Get(id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(id), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Customer)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// GetCustomerID returns either your own customer ID or the provided one
func GetCustomerID(customerID string) string {
	if customerID == "" {
		var err error
		customerID, err = GetOwnCustomerID()
		if err != nil {
			log.Printf("Error determining customer ID: %v\n", err)
		}
	}
	return customerID
}

// GetOwnCustomerID returns your own customer ID
func GetOwnCustomerID() (string, error) {
	r, err := GetCustomer("my_customer", "id")
	if err != nil {
		return "", err
	}
	return r.Id, nil
}

// PatchCustomer updates a customer. This method supports patch semantics.
func PatchCustomer(customerKey, fields string, customer *admin.Customer) (*admin.Customer, error) {
	srv := getCustomersService()
	c := srv.Patch(customerKey, customer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerKey), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*admin.Customer)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
