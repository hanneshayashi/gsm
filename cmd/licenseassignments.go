/*
Copyright © 2020-2024 Hannes Hayashi

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

package cmd

import (
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/licensing/v1"
)

// licenseAssignmentsCmd represents the licenseAssignments command
var licenseAssignmentsCmd = &cobra.Command{
	Use:               "licenseAssignments",
	Short:             "Manage user license assignments (Part of Enterprise License Manager API)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/licensing/reference/rest/v1/licenseAssignments",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}
var licenseAssignmentFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"productId": {
		AvailableFor: []string{"delete", "get", "insert", "listForProduct", "listForProductAndSku", "patch"},
		Type:         "string",
		Description: `A product's unique identifier.
For more information about products in this version of the API, see https://developers.google.com/admin-sdk/licensing/v1/how-tos/products.`,
		Defaults:  map[string]any{"delete": "Google-Apps", "get": "Google-Apps", "insert": "Google-Apps", "listForProduct": "Google-Apps", "listForProductAndSku": "Google-Apps", "patch": "Google-Apps"},
		Recursive: []string{"delete", "get", "insert", "patch"},
	},
	"skuId": {
		AvailableFor: []string{"delete", "get", "insert", "listForProductAndSku", "patch"},
		Type:         "string",
		Description: `A product SKU's unique identifier.
For more information about available SKUs in this version of the API, see https://developers.google.com/admin-sdk/licensing/v1/how-tos/products.`,
		Required:  []string{"delete", "get", "insert", "listForProductAndSku", "patch"},
		Recursive: []string{"delete", "get", "insert", "patch"},
	},
	"skuIdNew": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The product's new unique identifier.
For more information about products in this version of the API, see https://developers.google.com/admin-sdk/licensing/v1/how-tos/products.`,
		Required:  []string{"patch"},
		Recursive: []string{"patch"},
	},
	"userId": {
		AvailableFor: []string{"delete", "get", "insert", "patch"},
		Type:         "string",
		Description: `The user's current primary email address.
If the user's email address changes, use the new email address in your API requests.
Since a userId is subject to change, do not use a userId value as a key for persistent data.
This key could break if the current user's email address changes.
If the userId is suspended, the license status changes.`,
		Required: []string{"delete", "get", "insert", "patch"},
	},
	"customerId": {
		AvailableFor: []string{"listForProduct", "listForProductAndSku"},
		Type:         "string",
		Description: `The user's current primary email address.
If the user's email address changes, use the new email address in your API requests.
Since a userId is subject to change, do not use a userId value as a key for persistent data.
This key could break if the current user's email address changes.
If the userId is suspended, the license status changes.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "listForProduct", "listForProductAndSku", "patch"},
		Type:         "string",
		Description: `The user's current primary email address.
If the user's email address changes, use the new email address in your API requests.
Since a userId is subject to change, do not use a userId value as a key for persistent data.
This key could break if the current user's email address changes.
If the userId is suspended, the license status changes.`,
		Recursive: []string{"get", "insert", "patch"},
	},
}
var licenseAssignmentFlagsALL = gsmhelpers.GetAllFlags(licenseAssignmentFlags)

func init() {
	rootCmd.AddCommand(licenseAssignmentsCmd)
}

func mapToLicenseAssignmentInsert(flags map[string]*gsmhelpers.Value) (*licensing.LicenseAssignmentInsert, error) {
	licenseAssignmentInsert := &licensing.LicenseAssignmentInsert{}
	if flags["userId"].IsSet() {
		licenseAssignmentInsert.UserId = flags["userId"].GetString()
		if licenseAssignmentInsert.UserId == "" {
			licenseAssignmentInsert.ForceSendFields = append(licenseAssignmentInsert.ForceSendFields, "UserId")
		}
	}
	return licenseAssignmentInsert, nil
}

func mapToLicenseAssignment(flags map[string]*gsmhelpers.Value) (*licensing.LicenseAssignment, error) {
	licenseAssignment := &licensing.LicenseAssignment{}
	if flags["skuIdNew"].IsSet() {
		licenseAssignment.SkuId = flags["skuIdNew"].GetString()
		if licenseAssignment.SkuId == "" {
			licenseAssignment.ForceSendFields = append(licenseAssignment.ForceSendFields, "SkuId")
		}
	}
	return licenseAssignment, nil
}
