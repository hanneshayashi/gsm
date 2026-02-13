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

package gsmlicensing

import (
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/licensing/v1"
)

// DeleteLicenseAssignment revoke a license.
func DeleteLicenseAssignment(productID, skuID, userID string) (bool, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Delete(productID, skuID, userID)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(productID, skuID, userID), err)
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(productID, skuID, userID), err)
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
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(productID, skuID, licenseAssignmentInsert.UserId), err)
	}
	return r, nil
}

// ListLicenseAssignmentsForProduct list all users assigned licenses for a specific product SKU.
func ListLicenseAssignmentsForProduct(productID, customerID, fields string) iter.Seq2[*licensing.LicenseAssignment, error] {
	srv := getLicenseAssignmentsService()
	c := srv.ListForProduct(productID, customerID).MaxResults(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*licensing.LicenseAssignment, error) bool) {
		e := c.Pages(context.Background(), func(response *licensing.LicenseAssignmentList) error {
			for i := range response.Items {
				if !yield(response.Items[i], nil) {
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

// ListLicenseAssignmentsForProductAndSku list all users assigned licenses for a specific product SKU.
func ListLicenseAssignmentsForProductAndSku(productID, skuID, customerID, fields string) iter.Seq2[*licensing.LicenseAssignment, error] {
	srv := getLicenseAssignmentsService()
	c := srv.ListForProductAndSku(productID, skuID, customerID).MaxResults(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*licensing.LicenseAssignment, error) bool) {
		e := c.Pages(context.Background(), func(response *licensing.LicenseAssignmentList) error {
			for i := range response.Items {
				if !yield(response.Items[i], nil) {
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

// PatchLicenseAssignment reassign a user's product SKU with a different SKU in the same product.
func PatchLicenseAssignment(productID, skuID, userID, fields string, licenseAssignment *licensing.LicenseAssignment) (*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Patch(productID, skuID, userID, licenseAssignment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(productID, skuID, userID), err)
	}
	return r, nil
}
