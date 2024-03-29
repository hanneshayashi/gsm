/*
Copyright © 2020-2023 Hannes Hayashi

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

package gsmlicensing

import (
	"context"
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/licensing/v1"
)

// DeleteLicenseAssignment revoke a license.
func DeleteLicenseAssignment(productID, skuID, userID string) (bool, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Delete(productID, skuID, userID)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(productID, skuID, userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetLicenseAssignment get a specific user's license by product SKU.
func GetLicenseAssignment(productID, skuID, userID, fields string) (*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Get(productID, skuID, userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(productID, skuID, userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*licensing.LicenseAssignment)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// InsertLicenseAssignment assign a license.
func InsertLicenseAssignment(productID, skuID, fields string, licenseAssignmentInsert *licensing.LicenseAssignmentInsert) (*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Insert(productID, skuID, licenseAssignmentInsert)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(productID, skuID, licenseAssignmentInsert.UserId), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*licensing.LicenseAssignment)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListLicenseAssignmentsForProduct list all users assigned licenses for a specific product SKU.
func ListLicenseAssignmentsForProduct(productID, customerID, fields string, cap int) (<-chan *licensing.LicenseAssignment, <-chan error) {
	srv := getLicenseAssignmentsService()
	c := srv.ListForProduct(productID, customerID).MaxResults(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *licensing.LicenseAssignment, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *licensing.LicenseAssignmentList) error {
			for i := range response.Items {
				ch <- response.Items[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// ListLicenseAssignmentsForProductAndSku list all users assigned licenses for a specific product SKU.
func ListLicenseAssignmentsForProductAndSku(productID, skuID, customerID, fields string, cap int) (<-chan *licensing.LicenseAssignment, <-chan error) {
	srv := getLicenseAssignmentsService()
	c := srv.ListForProductAndSku(productID, skuID, customerID).MaxResults(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *licensing.LicenseAssignment, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *licensing.LicenseAssignmentList) error {
			for i := range response.Items {
				ch <- response.Items[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// PatchLicenseAssignment reassign a user's product SKU with a different SKU in the same product.
func PatchLicenseAssignment(productID, skuID, userID, fields string, licenseAssignment *licensing.LicenseAssignment) (*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Patch(productID, skuID, userID, licenseAssignment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(productID, skuID, userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*licensing.LicenseAssignment)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
