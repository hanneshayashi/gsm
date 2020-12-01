/*
Package gsmlicensing implements the Enterprise License Manager API
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
	"gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/licensing/v1"
)

// DeleteLicenseAssignment revoke a license.
func DeleteLicenseAssignment(productID, skuID, userID string) (bool, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Delete(productID, skuID, userID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(productID, skuID, userID), func() error {
		return c.Do()
	})
	return result, err
}

// GetLicenseAssignment get a specific user's license by product SKU.
func GetLicenseAssignment(productID, skuID, userID, fields string) (*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Get(productID, skuID, userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(productID, skuID, userID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*licensing.LicenseAssignment)
	return r, nil
}

// InsertLicenseAssignment assign a license.
func InsertLicenseAssignment(productID, skuID, fields string, licenseAssignmentInsert *licensing.LicenseAssignmentInsert) (*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Insert(productID, skuID, licenseAssignmentInsert)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(productID, skuID, licenseAssignmentInsert.UserId), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*licensing.LicenseAssignment)
	return r, nil
}

func makeListLicenseAssignmentsForProductCallAndAppend(c *licensing.LicenseAssignmentsListForProductCall, licenseAssignments []*licensing.LicenseAssignment, errKey string) ([]*licensing.LicenseAssignment, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*licensing.LicenseAssignmentList)
	licenseAssignments = append(licenseAssignments, r.Items...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		licenseAssignments, err = makeListLicenseAssignmentsForProductCallAndAppend(c, licenseAssignments, errKey)
	}
	return licenseAssignments, err
}

// ListLicenseAssignmentsForProduct list all users assigned licenses for a specific product SKU.
func ListLicenseAssignmentsForProduct(productID, customerID, fields string) ([]*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.ListForProduct(productID, customerID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var licenseAssignments []*licensing.LicenseAssignment
	licenseAssignments, err := makeListLicenseAssignmentsForProductCallAndAppend(c, licenseAssignments, gsmhelpers.FormatErrorKey(productID, customerID))
	return licenseAssignments, err
}

func makeListLicenseAssignmentsForProductAndSkuCallAndAppend(c *licensing.LicenseAssignmentsListForProductAndSkuCall, licenseAssignments []*licensing.LicenseAssignment, errKey string) ([]*licensing.LicenseAssignment, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*licensing.LicenseAssignmentList)
	licenseAssignments = append(licenseAssignments, r.Items...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		licenseAssignments, err = makeListLicenseAssignmentsForProductAndSkuCallAndAppend(c, licenseAssignments, errKey)
	}
	return licenseAssignments, err
}

// ListLicenseAssignmentsForProductAndSku list all users assigned licenses for a specific product SKU.
func ListLicenseAssignmentsForProductAndSku(productID, skuID, customerID, fields string) ([]*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.ListForProductAndSku(productID, skuID, customerID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var licenseAssignments []*licensing.LicenseAssignment
	licenseAssignments, err := makeListLicenseAssignmentsForProductAndSkuCallAndAppend(c, licenseAssignments, gsmhelpers.FormatErrorKey(productID, skuID, customerID))
	return licenseAssignments, err
}

// PatchLicenseAssignment reassign a user's product SKU with a different SKU in the same product.
func PatchLicenseAssignment(productID, skuID, userID, fields string, licenseAssignment *licensing.LicenseAssignment) (*licensing.LicenseAssignment, error) {
	srv := getLicenseAssignmentsService()
	c := srv.Patch(productID, skuID, userID, licenseAssignment)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(productID, skuID, userID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*licensing.LicenseAssignment)
	return r, nil
}
