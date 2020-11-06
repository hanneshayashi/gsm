/*
Package cmd contains the commands available to the end user
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
package cmd

import (
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// customersCmd represents the customers command
var customersCmd = &cobra.Command{
	Use:   "customers",
	Short: "Implements customers API (Part of Admin SDK).",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/customers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var customerFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customerKey": {
		AvailableFor: []string{"get", "patch"},
		Type:         "string",
		Description:  `Id of the customer.`,
		Defaults:     map[string]interface{}{"get": "my_customer", "patch": "my_customer"},
	},
	"alternateEmail": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The customer's secondary contact email address.
This email address cannot be on the same domain as the customerDomain`,
	},
	"customerDomain": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The customer's primary domain name string.
Do not include the www prefix when creating a new customer.`,
	},
	"language": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The customer's ISO 639-2 language code.
See the Language Codes page for the list of supported codes.
Valid language codes outside the supported set will be accepted by the API but may lead to unexpected behavior.
The default value is en.`,
	},
	"phoneNumber": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `The customer's contact phone number in E.164 format.`,
	},
	"addressLine1": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `A customer's physical address.
The address can be composed of one to three lines.`,
	},
	"addressLine2": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `Address line 2 of the address.`,
	},
	"addressLine3": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `Address line 3 of the address.`,
	},
	"contactName": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `The customer contact's name.`,
	},
	"countryCode": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `This is a required property.
For countryCode information see the ISO 3166 country code elements.(http://www.iso.org/iso/country_codes.htm)`,
	},
	"locality": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Name of the locality.
An example of a locality value is the city of San Francisco.`,
	},
	"organizationName": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `The company or company division name.`,
	},
	"postalCode": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The postal code. A postalCode example is a postal zip code such as 10009.
This is in accordance with - http://portablecontacts.net/draft-spec.html#address_element.`,
	},
	"region": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Name of the region.
An example of a region value is NY for the state of New York.`,
	},
	"fields": {
		AvailableFor: []string{"get", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(customersCmd)
}

func mapToCustomer(flags map[string]*gsmhelpers.Value) (*admin.Customer, error) {
	customer := &admin.Customer{}
	if flags["alternateEmail"].IsSet() {
		customer.AlternateEmail = flags["alternateEmail"].GetString()
		if customer.AlternateEmail == "" {
			customer.ForceSendFields = append(customer.ForceSendFields, "AlternateEmail")
		}
	}
	if flags["customerDomain"].IsSet() {
		customer.CustomerDomain = flags["customerDomain"].GetString()
		if customer.CustomerDomain == "" {
			customer.ForceSendFields = append(customer.ForceSendFields, "CustomerDomain")
		}
	}
	if flags["language"].IsSet() {
		customer.Language = flags["language"].GetString()
		if customer.Language == "" {
			customer.ForceSendFields = append(customer.ForceSendFields, "Language")
		}
	}
	if flags["phoneNumber"].IsSet() {
		customer.PhoneNumber = flags["phoneNumber"].GetString()
		if customer.PhoneNumber == "" {
			customer.ForceSendFields = append(customer.ForceSendFields, "PhoneNumber")
		}
	}
	if flags["addressLine1"].IsSet() || flags["addressLine2"].IsSet() || flags["addressLine3"].IsSet() || flags["contactName"].IsSet() || flags["countryCode"].IsSet() || flags["locality"].IsSet() || flags["organizationName"].IsSet() || flags["postalCode"].IsSet() || flags["region"].IsSet() {
		customer.PostalAddress = &admin.CustomerPostalAddress{}
		if flags["addressLine1"].IsSet() {
			customer.PostalAddress.AddressLine1 = flags["addressLine1"].GetString()
			if customer.PostalAddress.AddressLine1 == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "AddressLine1")
			}
		}
		if flags["addressLine2"].IsSet() {
			customer.PostalAddress.AddressLine2 = flags["addressLine2"].GetString()
			if customer.PostalAddress.AddressLine2 == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "AddressLine2")
			}
		}
		if flags["addressLine3"].IsSet() {
			customer.PostalAddress.AddressLine3 = flags["addressLine3"].GetString()
			if customer.PostalAddress.AddressLine3 == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "AddressLine3")
			}
		}
		if flags["contactName"].IsSet() {
			customer.PostalAddress.ContactName = flags["contactName"].GetString()
			if customer.PostalAddress.ContactName == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "ContactName")
			}
		}
		if flags["countryCode"].IsSet() {
			customer.PostalAddress.CountryCode = flags["countryCode"].GetString()
			if customer.PostalAddress.CountryCode == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "CountryCode")
			}
		}
		if flags["locality"].IsSet() {
			customer.PostalAddress.Locality = flags["locality"].GetString()
			if customer.PostalAddress.Locality == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "Locality")
			}
		}
		if flags["organizationName"].IsSet() {
			customer.PostalAddress.OrganizationName = flags["organizationName"].GetString()
			if customer.PostalAddress.OrganizationName == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "OrganizationName")
			}
		}
		if flags["postalCode"].IsSet() {
			customer.PostalAddress.PostalCode = flags["postalCode"].GetString()
			if customer.PostalAddress.PostalCode == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "PostalCode")
			}
		}
		if flags["region"].IsSet() {
			customer.PostalAddress.Region = flags["region"].GetString()
			if customer.PostalAddress.Region == "" {
				customer.PostalAddress.ForceSendFields = append(customer.PostalAddress.ForceSendFields, "Region")
			}
		}
	}
	return customer, nil
}
