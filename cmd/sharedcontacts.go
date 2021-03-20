/*
Package cmd contains the commands available to the end user
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
package cmd

import (
	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// sharedContactsCmd represents the sharedContacts command
var sharedContactsCmd = &cobra.Command{
	Use:               "sharedContacts",
	Short:             "Manage Domain Shared Contacts (Part of Shared Contacts API - not Admin SDK!)",
	Long:              "https://developers.google.com/admin-sdk/domain-shared-contacts",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var sharedContactFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"domain": {
		AvailableFor: []string{"create", "delete", "get", "list"},
		Type:         "string",
		Description:  "DNS domain the contact should be created in",
		Required:     []string{"create", "delete", "get", "list"},
	},
	"givenName": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Person's given name.",
	},
	"additionalName": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Additional name of the person, eg. middle name.",
	},
	"familyName": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Person's family name.",
	},
	"namePrefix": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Honorific prefix, eg. 'Mr' or 'Mrs'.",
	},
	"nameSuffix": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Honorific suffix, eg. 'san' or 'III'.",
	},
	"fullName": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "Unstructured representation of the name.",
	},
	"email": {
		AvailableFor: []string{"create", "update"},
		Type:         "stringSlice",
		Description: `Email.
Must be in the form of "address=user@domain.com;displayName=Some Name;primary=[true|false];label=[Work|Home]".
Can be used multiple time (although "primary" may only be used once`,
	},
	"phoneNumber": {
		AvailableFor: []string{"create", "update"},
		Type:         "stringSlice",
		Description: `Phone number.
Must be in the form of "phoneNumber=+1 212 213181;primary=[true|false]label=[Work|Home|Mobile]".
Can be used multiple time (although "primary" may only be used once`,
	},
	"im": {
		AvailableFor: []string{"create", "update"},
		Type:         "stringSlice",
		Description: `IM addresses.
Must be in the form of "protocol=http://schemas.google.com/g/2005#GOOGLE_TALK;address=some@address.com;primary=[true|false]label=[Work|Home|Mobile]".
Can be used multiple time (although "primary" may only be used once`,
	},
	"organization": {
		AvailableFor: []string{"create", "update"},
		Type:         "stringSlice",
		Description: `Organization of the contact.
Must be in the form of "orgName=Some Company;orgDepartment=Some Department;orgTitle=Some Title;orgJobDescription=Some Description;orgSymbol=Some Symbol"`,
	},
	"extendedProperty": {
		AvailableFor: []string{"create", "update"},
		Type:         "stringSlice",
		Description: `Extended Properties
Must be in the form of "name=Some Name;Value=Some Value;Realm=Some Realm"`,
	},
	"structuredPostalAddress": {
		AvailableFor: []string{"create", "update"},
		Type:         "stringSlice",
		Description: `Structed Postal Address
Must be in the form of "mailClass=...;label=...;usage=...;primary=[true|false];agent=...;housename=...;street=...;pobox=...neighborhood=...;city=...;subregion=...;region=...;postcode=...;country=...;formattedAddress=..."`,
	},
	"url": {
		AvailableFor: []string{"delete", "get", "update"},
		Type:         "string",
		Description: `URL of the Shared Contact (Retrieve with "list" and look for "id").
MUST BE https://!`,
		Required:       []string{"delete", "get", "update"},
		ExcludeFromAll: true,
	},
	"json": {
		AvailableFor: []string{"create", "get", "list", "update"},
		Type:         "bool",
		Description:  `Output as JSON"`,
	},
}
var sharedContactFlagsALL = gsmhelpers.GetAllFlags(sharedContactFlags)

func init() {
	rootCmd.AddCommand(sharedContactsCmd)
}

func mapToSharedContact(flags map[string]*gsmhelpers.Value, sharedContact *gsmadmin.Entry) (*gsmadmin.Entry, error) {
	if sharedContact == nil {
		sharedContact = &gsmadmin.Entry{}
	}
	if flags["givenName"].IsSet() || flags["additionalName"].IsSet() || flags["familyName"].IsSet() || flags["namePrefix"].IsSet() || flags["nameSuffix"].IsSet() || flags["fullName"].IsSet() {
		sharedContact.Name = gsmadmin.Name{}
		if flags["givenName"].IsSet() {
			sharedContact.Name.GivenName = flags["givenName"].GetString()
		}
		if flags["additionalName"].IsSet() {
			sharedContact.Name.AdditionalName = flags["additionalName"].GetString()
		}
		if flags["familyName"].IsSet() {
			sharedContact.Name.FamilyName = flags["familyName"].GetString()
		}
		if flags["namePrefix"].IsSet() {
			sharedContact.Name.NamePrefix = flags["namePrefix"].GetString()
		}
		if flags["nameSuffix"].IsSet() {
			sharedContact.Name.NameSuffix = flags["nameSuffix"].GetString()
		}
		if flags["fullName"].IsSet() {
			sharedContact.Name.FullName = flags["fullName"].GetString()
		}
	}
	if flags["email"].IsSet() {
		emails := flags["email"].GetStringSlice()
		sharedContact.Email = []gsmadmin.Email{}
		for i := range emails {
			m := gsmhelpers.FlagToMap(emails[i])
			sharedContact.Email = append(sharedContact.Email, gsmadmin.Email{Address: m["address"], DisplayName: m["displayName"], Primary: m["primary"], Label: m["label"]})
		}
	}
	if flags["phoneNumber"].IsSet() {
		phoneNumbers := flags["phoneNumber"].GetStringSlice()
		sharedContact.PhoneNumber = []gsmadmin.PhoneNumber{}
		for i := range phoneNumbers {
			m := gsmhelpers.FlagToMap(phoneNumbers[i])
			sharedContact.PhoneNumber = append(sharedContact.PhoneNumber, gsmadmin.PhoneNumber{PhoneNumber: m["phoneNumber"], Primary: m["primary"], Label: m["label"]})
		}
	}
	if flags["structuredPostalAddress"].IsSet() {
		structuredPostalAddresss := flags["structuredPostalAddress"].GetStringSlice()
		sharedContact.StructuredPostalAddress = []gsmadmin.StructuredPostalAddress{}
		for i := range structuredPostalAddresss {
			m := gsmhelpers.FlagToMap(structuredPostalAddresss[i])
			sharedContact.StructuredPostalAddress = append(sharedContact.StructuredPostalAddress, gsmadmin.StructuredPostalAddress{MailClass: m["mailClass"], Agent: m["agent"], City: m["city"], Country: m["country"], FormattedAddress: m["formattedAddress"], Housename: m["housename"], Label: m["label"], Neighborhood: m["neighborhood"], Pobox: m["pobox"], Postcode: m["postcode"], Primary: m["primary"], Region: m["region"], Street: m["street"], Subregion: m["subregion"], Usage: m["usage"]})
		}
	}
	if flags["organization"].IsSet() {
		organizations := flags["organization"].GetStringSlice()
		sharedContact.Organization = []gsmadmin.Organization{}
		for i := range organizations {
			m := gsmhelpers.FlagToMap(organizations[i])
			sharedContact.Organization = append(sharedContact.Organization, gsmadmin.Organization{Label: m["label"], OrgDepartment: m["orgDepartment"], OrgJobDescription: m["orgJobDescription"], OrgName: m["orgName"], OrgSymbol: m["orgSymbol"], OrgTitle: m["orgTitle"], Primary: m["primary"]})
		}
	}
	if flags["im"].IsSet() {
		ims := flags["im"].GetStringSlice()
		sharedContact.Im = []gsmadmin.Im{}
		for i := range ims {
			m := gsmhelpers.FlagToMap(ims[i])
			sharedContact.Im = append(sharedContact.Im, gsmadmin.Im{Label: m["label"], Address: m["address"], Primary: m["primary"], Protocol: m["protocol"]})
		}
	}
	if flags["extendedProperty"].IsSet() {
		extendedpropertys := flags["extendedProperty"].GetStringSlice()
		sharedContact.ExtendedProperty = []gsmadmin.ExtendedProperty{}
		for i := range extendedpropertys {
			m := gsmhelpers.FlagToMap(extendedpropertys[i])
			sharedContact.ExtendedProperty = append(sharedContact.ExtendedProperty, gsmadmin.ExtendedProperty{Name: m["name"], Value: m["value"], Realm: m["realm"]})
		}
	}
	return sharedContact, nil
}
