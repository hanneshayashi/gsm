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

package cmd

import (
	"bufio"
	"encoding/base64"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/people/v1"
)

// peopleCmd represents the people command
var peopleCmd = &cobra.Command{
	Use:               "people",
	Short:             "Manage people's contacts (Part of People API)",
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/people",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var peopleFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"personFields": {
		AvailableFor: []string{"createContact", "deleteContactPhoto", "get", "getBatchGet", "updateContact", "updateContactPhoto"},
		Type:         "string",
		Description: `A field mask to restrict which fields on each person are returned.
Multiple fields can be specified by separating them with commas.
Defaults to all fields if not set.
Valid values are:
addresses
ageRanges
biographies
birthdays
calendarUrls
clientData
coverPhotos
emailAddresses
events
externalIds
genders
imClients
interests
locales
locations
memberships
metadata
miscKeywords
names
nicknames
occupations
organizations
phoneNumbers
photos
relations
sipAddresses
skills
urls
userDefined`,
		Required: []string{"get", "getBatchGet", "updateContact"},
	},
	"sources": {
		AvailableFor: []string{"createContact", "deleteContactPhoto", "get", "getBatchGet", "listDirectoryPeople", "searchDirectoryPeople", "updateContact", "updateContactPhoto"},
		Type:         "string",
		Description: `A mask of what source types to return.
DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact.
DIRECTORY_SOURCE_TYPE_DOMAIN_PROFILE  - Workspace domain profile.
READ_SOURCE_TYPE_PROFILE              - Returns SourceType.ACCOUNT, SourceType.DOMAIN_PROFILE, and SourceType.PROFILE.
READ_SOURCE_TYPE_CONTACT              - Returns SourceType.CONTACT.
READ_SOURCE_TYPE_DOMAIN_CONTACT       - Returns SourceType.DOMAIN_CONTACT.`,
		Required: []string{"listDirectoryPeople", "searchDirectoryPeople", "updateContact"},
	},
	"resourceName": {
		AvailableFor:   []string{"deleteContact", "deleteContactPhoto", "get", "updateContact", "updateContactPhoto"},
		Type:           "string",
		Description:    `The resource name of the contact-`,
		Required:       []string{"deleteContact", "deleteContactPhoto", "get", "updateContact", "updateContactPhoto"},
		ExcludeFromAll: true,
	},
	"resourceNames": {
		AvailableFor: []string{"getBatchGet"},
		Type:         "stringSlice",
		Description: `The resource names of the people to provide information about.
It's repeatable. The URL query parameter should be

resourceNames=<name1>&resourceNames=<name2>&...

To get information about the authenticated user, specify people/me.
To get information about a google account, specify people/{account_id}.
To get information about a contact, specify the resource name that identifies the contact as returned by people.connections.list.
You can include up to 50 resource names in one request.`,
		Required:       []string{"getBatchGet"},
		ExcludeFromAll: true,
	},
	"readMask": {
		AvailableFor: []string{"listDirectoryPeople", "searchDirectoryPeople"},
		Type:         "string",
		Description: `A field mask to restrict which fields on each person are returned.
Multiple fields can be specified by separating them with commas.
Valid values are:
addresses
ageRanges
biographies
birthdays
calendarUrls
clientData
coverPhotos
emailAddresses
events
externalIds
genders
imClients
interests
locales
locations
memberships
metadata
miscKeywords
names
nicknames
occupations
organizations
phoneNumbers
photos
relations
sipAddresses
skills
urls
userDefined`,
		Required: []string{"listDirectoryPeople", "searchDirectoryPeople"},
	},
	"mergeSources": {
		AvailableFor: []string{"listDirectoryPeople", "searchDirectoryPeople"},
		Type:         "stringSlice",
		Description:  `Additional data to merge into the directory sources if they are connected through verified join keys such as email addresses or phone numbers.`,
	},
	"query": {
		AvailableFor: []string{"searchDirectoryPeople"},
		Type:         "string",
		Description: `Prefix query that matches fields in the person.
Does NOT use the readMask for determining what fields to match.`,
		Required: []string{"searchDirectoryPeople"},
	},
	"updatePersonFields": {
		AvailableFor: []string{"updateContact"},
		Type:         "string",
		Description: `A field mask to restrict which fields on the person are updated.
Multiple fields can be specified by separating them with commas.
All updated fields will be replaced.
Valid values are:
addresses
biographies
birthdays
calendarUrls
clientData
emailAddresses
events
externalIds
genders
imClients
interests
locales
locations
memberships
miscKeywords
names
nicknames
occupations
organizations
phoneNumbers
relations
sipAddresses
urls
userDefined`,
		Required: []string{"updateContact"},
	},
	"photo": {
		AvailableFor: []string{"updateContactPhoto"},
		Type:         "string",
		Description:  `Path to a photo file.`,
		Required:     []string{"updateContactPhoto"},
	},
	"addresses": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `A person's physical address.
May be a P.O. box or street address.
All fields are optional.
May be used multiple times in the for of "formattedValue=...;type=...;poBox...", etc.
You may use the following fields:
primary          - True if the field is the primary field; false if the field is a secondary field.
formattedValue   - The unstructured value of the address.
				   If this is not set by the user it will be automatically constructed from structured values.
type             - The type of the address.
                   The type can be custom or one of these predefined values:
                     - home
                     - work
                     - other
poBox            - The P.O. box of the address.
streetAddress    - The street address.
extendedAddress	 - The extended address of the address; for example, the apartment number.
city             - The city of the address.
region           - The region of the address; for example, the state or province.
postalCode	     - The postal code of the address.
country          - The country of the address.
countryCode	     - The ISO 3166-1 alpha-2 country code of the address.`,
	},
	"biographyValue": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The short biography.`,
	},
	"biographyContentType": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description: `The content type of the biography.
CONTENT_TYPE_UNSPECIFIED  - Unspecified.
TEXT_PLAIN                - Plain text.
TEXT_HTML                 - HTML text.`,
	},
	"birthdayYear": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "int64",
		Description: `Year of date.
Must be from 1 to 9999, or 0 if specifying a date without a year.`,
	},
	"birthdayMonth": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "int64",
		Description: `Month of year.
Must be from 1 to 12, or 0 if specifying a year without a month and day.`,
	},
	"birthdayDay": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "int64",
		Description: `Day of month.
Must be from 1 to 31 and valid for the year and month, or 0 if specifying a year by itself or a year and month where the day is not significant.`,
	},
	"birthdayText": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `A free-form string representing the user's birthday.`,
	},
	"calendarUrls": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's calendar URLs.
Can be used multiple times in the form of "url=...,type=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
url      - The calendar URL.
type     - The type of the calendar URL.
           The type can be custom or one of these predefined values:
             - home
             - work
             - other`,
	},
	"clientData": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's client data.
Arbitrary client data that is populated by clients. Duplicate keys and values are allowed.
Can be used multiple times in the form of "key=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
key      - The client specified key of the client data.
value    - The client specified value of the client data.`,
	},
	"emailAddresses": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's email addresses.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary    - True if the field is the primary field; false if the field is a secondary field.
value      - The email address.
type       - The type of the email address.
             The type can be custom or one of these predefined values:
               - home
               - work
               - other
displayName  - The display name of the email.`,
	},
	"events": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's events.
Can be used multiple times in the form of "year=...,month=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
year     - Year of date.
           Must be from 1 to 9999, or 0 if specifying a date without a year.
month    - Month of year.
           Must be from 1 to 12, or 0 if specifying a year without a month and day.
day      - Day of month.
           Must be from 1 to 31 and valid for the year and month, or 0 if specifying a year by itself or a year and month where the day is not significant.
type     - The type of the event.
           The type can be custom or one of these predefined values:
             - anniversary
             - other`,
	},
	"externalIds": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's external IDs.
Can be used multiple times in the form of "value=...,type=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The value of the external ID.
type     - The type of the external ID.
           The type can be custom or one of these predefined values:
             - account
             - customer
             - loginId
             - network
             - organization`,
	},
	"fileAses": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's file-ases.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The file-as value`,
	},
	"genderValue": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description: `The gender for the person.
The gender can be custom or one of these predefined values:
  - male
  - female
  - unspecified`,
	},
	"addressMeAs": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description: `The type of pronouns that should be used to address the person.
The value can be custom or one of these predefined values:
  - male
  - female
  - other`,
	},
	"imClients": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's instant messaging clients.
Can be used multiple times in the form of "primary=...,username=...", etc.
You may use the following fields:
primary   - True if the field is the primary field; false if the field is a secondary field.
username  - The user name used in the IM client.
type      - The type of the IM client.
            The type can be custom or one of these predefined values:
              - home
              - work
			  - other
protocol  - The protocol of the IM client.
            The protocol can be custom or one of these predefined values:
              - aim
              - msn
              - yahoo
              - skype
              - qq
              - googleTalk
              - icq
              - jabber
              - netMeeting`,
	},
	"interests": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's interests.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The interest; for example, stargazing.`,
	},
	"locales": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's locale preferences.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The well-formed IETF BCP 47 language tag representing the locale.`,
	},
	"locations": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's locations.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary       - True if the field is the primary field; false if the field is a secondary field.
value         - The free-form value of the location.
type          - The type of the location.
                The type can be custom or one of these predefined values:
                  - desk
			      - grewUp
current       - Whether the location is the current location.
buildingId    - The building identifier.
floor         - The floor name or number.
floorSection  - The floor section in floor_name.
deskCode      - The individual desk location.`,
	},
	"memberships": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's group memberships.
Can be used multiple times in the form of "primary=...,contactGroupResourceName=...", etc.
You may use the following fields:
primary                          - True if the field is the primary field; false if the field is a secondary field.
contactGroupResourceName         - The resource name for the contact group, assigned by the server.
								   An ASCII string, in the form of contactGroups/{contactGroupId}.
								   Only contactGroupResourceName can be used for modifying memberships.
								   Any contact group membership can be removed, but only user group or "myContacts" or "starred" system groups memberships can be added.
								   A contact must always have at least one contact group membership.`,
	},
	"miscKeywords": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's miscellaneous keywords.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The value of the miscellaneous keyword.
type     - The miscellaneous keyword type.
           Allowed values are:
             - TYPE_UNSPECIFIED             - Unspecified.
             - OUTLOOK_BILLING_INFORMATION  - Outlook field for billing information.
             - OUTLOOK_DIRECTORY_SERVER     - Outlook field for directory server.
             - OUTLOOK_KEYWORD              - Outlook field for keyword.
             - OUTLOOK_MILEAGE              - Outlook field for mileage.
             - OUTLOOK_PRIORITY             - Outlook field for priority.
             - OUTLOOK_SENSITIVITY          - Outlook field for sensitivity.
             - OUTLOOK_SUBJECT              - Outlook field for subject.
             - OUTLOOK_USER                 - Outlook field for user.
             - HOME                         - Home.
             - WORK                         - Work.
             - OTHER                        - Other.`,
	},
	"unstructuredName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The free form name value.`,
	},
	"familyName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The family name.`,
	},
	"givenName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The given name.`,
	},
	"middleName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The middle name(s).`,
	},
	"honorificPrefix": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The honorific prefixes, such as Mrs. or Dr.`,
	},
	"honorificSuffix": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The honorific suffixes, such as Jr.`,
	},
	"phoneticFullName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The full name spelled as it sounds.`,
	},
	"phoneticFamilyName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The family name spelled as it sounds.`,
	},
	"phoneticGivenName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The given name spelled as it sounds.`,
	},
	"phoneticMiddleName": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The middle name(s) spelled as they sound.`,
	},
	"phoneticHonorificPrefix": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The honorific prefixes spelled as they sound.`,
	},
	"phoneticHonorificSuffix": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "string",
		Description:  `The honorific suffixes spelled as they sound.`,
	},
	"nicknames": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's nicknames.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The nickname.
type     - The type of a nickname.
           Allowed values are:
			 - DEFAULT         - Generic nickname.
             - ALTERNATE_NAME  - Alternate name person is known by.`,
	},
	"occupations": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's occupations.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The occupation; for example, carpenter.`,
	},
	"organizations": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's past or current organizations.
Can be used multiple times in the form of "primary=...,type=...", etc.
You may use the following fields:
primary         - True if the field is the primary field; false if the field is a secondary field.
type            - The type of the organization.
                  The type can be custom or one of these predefined values:
                    - work
			        - school
startDateYear   - Year of date.
                  Must be from 1 to 9999, or 0 if specifying a date without a year.
startDateMonth  - Month of year.
                  Must be from 1 to 12, or 0 if specifying a year without a month and day.
startDateDay    - Day of month.
                  Must be from 1 to 31 and valid for the year and month, or 0 if specifying a year by itself or a year and month where the day is not significant.
endDateYear     - Year of date.
                  Must be from 1 to 9999, or 0 if specifying a date without a year.
endDateMonth    - Month of year.
                  Must be from 1 to 12, or 0 if specifying a year without a month and day.
endDateDay      - Day of month.
                  Must be from 1 to 31 and valid for the year and month, or 0 if specifying a year by itself or a year and month where the day is not significant.
current         - True if the organization is the person's current organization; false if the organization is a past organization.
name            - The name of the organization.
phoneticName    - The phonetic name of the organization.
department      - The person's department at the organization.
title           - The person's job title at the organization.
jobDescription  - The person's job description at the organization.
symbol          - The symbol associated with the organization; for example, a stock ticker symbol, abbreviation, or acronym.
domain          - The domain name associated with the organization; for example, google.com.
location        - The location of the organization office the person works at.`,
	},
	"phoneNumbers": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's phone numbers.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The phone number.
type     - The type of the phone number.
           The type can be custom or one of these predefined values:
             - home
             - work
             - mobile
             - homeFax
             - workFax
             - otherFax
             - pager
             - workMobile
             - workPager
             - main
             - googleVoice
             - other`,
	},
	"relations": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's relations.
Can be used multiple times in the form of "primary=...,person=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
person   - The name of the other person this relation refers to.
type     - The person's relation to the other person.
           The type can be custom or one of these predefined values:
             - spouse
             - child
             - mother
             - father
             - parent
             - brother
             - sister
             - friend
             - relative
             - domesticPartner
             - manager
             - assistant
             - referredBy
             - partner`,
	},
	"sipAddresses": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's SIP addresses.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The SIP address in the RFC 3261 19.1 SIP URI format.
type     - The type of the SIP address.
           The type can be custom or or one of these predefined values:
             - home
             - work
             - mobile
             - other`,
	},
	"skills": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's skills.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The skill; for example, underwater basket weaving.`,
	},
	"urls": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's associated URLs.
Can be used multiple times in the form of "primary=...,value=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
value    - The URL.
type     - The type of the URL.
           The type can be custom or one of these predefined values:
             - home
             - work
             - blog
             - profile
             - homePage
             - ftp
             - reservations
             - appInstallPage: website for a Currents application.
             - other`,
	},
	"userDefined": {
		AvailableFor: []string{"createContact", "updateContact"},
		Type:         "stringSlice",
		Description: `The person's user defined data.
Can be used multiple times in the form of "primary=...,type=...", etc.
You may use the following fields:
primary  - True if the field is the primary field; false if the field is a secondary field.
key      - The end user specified key of the user defined data.
value    - The end user specified value of the user defined data.`,
	},
	"fields": {
		AvailableFor: []string{"createContact", "deleteContactPhoto", "get", "getBatchGet", "listDirectoryPeople", "searchDirectoryPeople", "updateContact", "updateContactPhoto"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var peopleFlagsALL = gsmhelpers.GetAllFlags(peopleFlags)

func init() {
	rootCmd.AddCommand(peopleCmd)
}

func stringsToDate(year, month, day string) *people.Date {
	date := &people.Date{}
	var err error
	date.Year, err = strconv.ParseInt(year, 10, 64)
	if err == nil && date.Year == 0 {
		date.ForceSendFields = append(date.ForceSendFields, "Year")
	}
	date.Month, err = strconv.ParseInt(month, 10, 64)
	if err == nil && date.Month == 0 {
		date.ForceSendFields = append(date.ForceSendFields, "Month")
	}
	date.Day, err = strconv.ParseInt(day, 10, 64)
	if err == nil && date.Day == 0 {
		date.ForceSendFields = append(date.ForceSendFields, "Day")
	}
	return date
}

func addMetaData(primary string) *people.FieldMetadata {
	if primary != "" {
		b, err := strconv.ParseBool(primary)
		if err != nil {
			log.Printf("Error parsing %s to bool: %v. Setting to false.", primary, err)
		}
		return &people.FieldMetadata{Primary: b}
	}
	return nil
}

func mapToPerson(flags map[string]*gsmhelpers.Value, person *people.Person) (*people.Person, error) {
	if person == nil {
		person = &people.Person{}
	}
	if flags["addresses"].IsSet() {
		person.Addresses = []*people.Address{}
		addresses := flags["addresses"].GetStringSlice()
		if len(addresses) > 0 {
			for i := range addresses {
				m := gsmhelpers.FlagToMap(addresses[i])
				address := &people.Address{
					City:            m["city"],
					Country:         m["country"],
					CountryCode:     m["countryCode"],
					ExtendedAddress: m["extendedAddress"],
					FormattedValue:  m["formattedValue"],
					PoBox:           m["poBox"],
					PostalCode:      m["postalCode"],
					Region:          m["region"],
					StreetAddress:   m["streetAddress"],
					Type:            m["type"],
				}
				address.Metadata = addMetaData(m["primary"])
				person.Addresses = append(person.Addresses, address)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Addresses")
		}
	}
	if flags["biographyValue"].IsSet() || flags["biographyContentType"].IsSet() {
		person.Biographies = make([]*people.Biography, 1)
		person.Biographies[0] = new(people.Biography)
		if flags["biographyValue"].IsSet() {
			person.Biographies[0].Value = flags["biographyValue"].GetString()
			if person.Biographies[0].Value == "" {
				person.Biographies[0].ForceSendFields = append(person.Biographies[0].ForceSendFields, "Value")
			}
		}
		if flags["biographyContentType"].IsSet() {
			person.Biographies[0].ContentType = flags["biographyContentType"].GetString()
			if person.Biographies[0].ContentType == "" {
				person.Biographies[0].ForceSendFields = append(person.Biographies[0].ForceSendFields, "ContentType")
			}
		}
	}
	if flags["birthdayYear"].IsSet() || flags["birthdayMonth"].IsSet() || flags["birthdayDay"].IsSet() || flags["birthdayText"].IsSet() {
		person.Birthdays = make([]*people.Birthday, 1)
		person.Birthdays[0] = new(people.Birthday)
		if flags["birthdayYear"].IsSet() || flags["birthdayMonth"].IsSet() || flags["birthdayDay"].IsSet() {
			person.Birthdays[0].Date = &people.Date{}
			if flags["birthdayYear"].IsSet() {
				person.Birthdays[0].Date.Year = flags["birthdayYear"].GetInt64()
				if person.Birthdays[0].Date.Year == 0 {
					person.Birthdays[0].Date.ForceSendFields = append(person.Birthdays[0].Date.ForceSendFields, "Year")
				}
			}
			if flags["birthdayMonth"].IsSet() {
				person.Birthdays[0].Date.Month = flags["birthdayMonth"].GetInt64()
				if person.Birthdays[0].Date.Month == 0 {
					person.Birthdays[0].Date.ForceSendFields = append(person.Birthdays[0].Date.ForceSendFields, "Month")
				}
			}
			if flags["birthdayDay"].IsSet() {
				person.Birthdays[0].Date.Day = flags["birthdayDay"].GetInt64()
				if person.Birthdays[0].Date.Day == 0 {
					person.Birthdays[0].Date.ForceSendFields = append(person.Birthdays[0].Date.ForceSendFields, "Day")
				}
			}
		}
		if flags["birthdayText"].IsSet() {
			person.Birthdays[0].Text = flags["birthdayText"].GetString()
			if person.Birthdays[0].Text == "" {
				person.Birthdays[0].ForceSendFields = append(person.Birthdays[0].ForceSendFields, "Text")
			}
		}
	}
	if flags["calendarUrls"].IsSet() {
		person.CalendarUrls = []*people.CalendarUrl{}
		calendarURL := flags["calendarUrls"].GetStringSlice()
		if len(calendarURL) > 0 {
			for i := range calendarURL {
				m := gsmhelpers.FlagToMap(calendarURL[i])
				calendarURL := &people.CalendarUrl{
					Type: m["type"],
					Url:  m["url"],
				}
				calendarURL.Metadata = addMetaData(m["primary"])
				person.CalendarUrls = append(person.CalendarUrls, calendarURL)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "CalendarUrls")
		}
	}
	if flags["clientData"].IsSet() {
		person.ClientData = []*people.ClientData{}
		clientData := flags["clientData"].GetStringSlice()
		if len(clientData) > 0 {
			for i := range clientData {
				m := gsmhelpers.FlagToMap(clientData[i])
				cData := &people.ClientData{
					Key:   m["key"],
					Value: m["value"],
				}
				cData.Metadata = addMetaData(m["primary"])
				person.ClientData = append(person.ClientData, cData)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "ClientData")
		}
	}
	if flags["emailAddresses"].IsSet() {
		person.EmailAddresses = []*people.EmailAddress{}
		emailAddresses := flags["emailAddresses"].GetStringSlice()
		if len(emailAddresses) > 0 {
			for i := range emailAddresses {
				m := gsmhelpers.FlagToMap(emailAddresses[i])
				emailAddress := &people.EmailAddress{
					Value:       m["value"],
					DisplayName: m["displayName"],
					Type:        m["type"],
				}
				emailAddress.Metadata = addMetaData(m["primary"])
				person.EmailAddresses = append(person.EmailAddresses, emailAddress)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "EmailAddresses")
		}
	}
	if flags["events"].IsSet() {
		person.Events = []*people.Event{}
		events := flags["events"].GetStringSlice()
		if len(events) > 0 {
			for i := range events {
				m := gsmhelpers.FlagToMap(events[i])
				event := &people.Event{
					Date: stringsToDate(m["year"], m["month"], m["day"]),
					Type: m["type"],
				}
				event.Metadata = addMetaData(m["primary"])
				person.Events = append(person.Events, event)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Events")
		}
	}
	if flags["externalIds"].IsSet() {
		person.ExternalIds = []*people.ExternalId{}
		externalIDs := flags["externalIds"].GetStringSlice()
		if len(externalIDs) > 0 {
			for i := range externalIDs {
				m := gsmhelpers.FlagToMap(externalIDs[i])
				externalID := &people.ExternalId{
					Value: m["value"],
					Type:  m["type"],
				}
				externalID.Metadata = addMetaData(m["primary"])
				person.ExternalIds = append(person.ExternalIds, externalID)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "ExternalIds")
		}
	}
	if flags["fileAses"].IsSet() {
		person.FileAses = []*people.FileAs{}
		fileAses := flags["fileAses"].GetStringSlice()
		if len(fileAses) > 0 {
			for i := range fileAses {
				m := gsmhelpers.FlagToMap(fileAses[i])
				FileAs := &people.FileAs{
					Value: m["value"],
				}
				FileAs.Metadata = addMetaData(m["primary"])
				person.FileAses = append(person.FileAses, FileAs)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "FileAses")
		}
	}
	if flags["genderValue"].IsSet() || flags["addressMeAs"].IsSet() {
		person.Genders = make([]*people.Gender, 1)
		person.Genders[0] = new(people.Gender)
		if flags["genderValue"].IsSet() {
			person.Genders[0].Value = flags["genderValue"].GetString()
			if person.Genders[0].Value == "" {
				person.Genders[0].ForceSendFields = append(person.Genders[0].ForceSendFields, "Value")
			}
		}
		if flags["addressMeAs"].IsSet() {
			person.Genders[0].Value = flags["addressMeAs"].GetString()
			if person.Genders[0].AddressMeAs == "" {
				person.Genders[0].ForceSendFields = append(person.Genders[0].ForceSendFields, "AddressMeAs")
			}
		}
	}
	if flags["imClients"].IsSet() {
		person.ImClients = []*people.ImClient{}
		imClients := flags["imClients"].GetStringSlice()
		if len(imClients) > 0 {
			for i := range imClients {
				m := gsmhelpers.FlagToMap(imClients[i])
				imClient := &people.ImClient{
					Type:     m["type"],
					Username: m["username"],
					Protocol: m["protocol"],
				}
				imClient.Metadata = addMetaData(m["primary"])
				person.ImClients = append(person.ImClients, imClient)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "ImClients")
		}
	}
	if flags["interests"].IsSet() {
		person.Interests = []*people.Interest{}
		interests := flags["interests"].GetStringSlice()
		if len(interests) > 0 {
			for i := range interests {
				m := gsmhelpers.FlagToMap(interests[i])
				interest := &people.Interest{
					Value: m["value"],
				}
				interest.Metadata = addMetaData(m["primary"])
				person.Interests = append(person.Interests, interest)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Interests")
		}
	}
	if flags["locales"].IsSet() {
		person.Locales = []*people.Locale{}
		locales := flags["locales"].GetStringSlice()
		if len(locales) > 0 {
			for i := range locales {
				m := gsmhelpers.FlagToMap(locales[i])
				locale := &people.Locale{
					Value: m["value"],
				}
				locale.Metadata = addMetaData(m["primary"])
				person.Locales = append(person.Locales, locale)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Locales")
		}
	}
	if flags["locations"].IsSet() {
		person.Locations = []*people.Location{}
		locations := flags["locations"].GetStringSlice()
		if len(locations) > 0 {
			for i := range locations {
				m := gsmhelpers.FlagToMap(locations[i])
				location := &people.Location{
					Value:        m["value"],
					BuildingId:   m["buildingId"],
					DeskCode:     m["deskCode"],
					Floor:        m["floor"],
					FloorSection: m["floorSection"],
					Type:         m["type"],
				}
				location.Current, _ = strconv.ParseBool(m["current"])
				location.Metadata = addMetaData(m["primary"])
				person.Locations = append(person.Locations, location)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Locations")
		}
	}
	if flags["memberships"].IsSet() {
		person.Memberships = []*people.Membership{}
		memberships := flags["memberships"].GetStringSlice()
		if len(memberships) > 0 {
			for i := range memberships {
				m := gsmhelpers.FlagToMap(memberships[i])
				membership := &people.Membership{
					ContactGroupMembership: &people.ContactGroupMembership{
						ContactGroupResourceName: m["contactGroupResourceName"],
					},
				}
				membership.Metadata = addMetaData(m["primary"])
				person.Memberships = append(person.Memberships, membership)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Memberships")
		}
	}
	if flags["miscKeywords"].IsSet() {
		person.MiscKeywords = []*people.MiscKeyword{}
		miscKeywords := flags["miscKeywords"].GetStringSlice()
		if len(miscKeywords) > 0 {
			for i := range miscKeywords {
				m := gsmhelpers.FlagToMap(miscKeywords[i])
				miscKeyword := &people.MiscKeyword{
					Type:  m["type"],
					Value: m["value"],
				}
				miscKeyword.Metadata = addMetaData(m["primary"])
				person.MiscKeywords = append(person.MiscKeywords, miscKeyword)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "MiscKeywords")
		}
	}
	if flags["unstructuredName"].IsSet() || flags["familyName"].IsSet() || flags["givenName"].IsSet() || flags["middleName"].IsSet() || flags["honorificPrefix"].IsSet() || flags["honorificSuffix"].IsSet() || flags["phoneticFullName"].IsSet() || flags["phoneticFamilyName"].IsSet() || flags["phoneticGivenName"].IsSet() || flags["phoneticMiddleName"].IsSet() || flags["phoneticHonorificPrefix"].IsSet() || flags["phoneticHonorificSuffix"].IsSet() {
		if person.Names == nil {
			person.Names = make([]*people.Name, 1)
			person.Names[0] = new(people.Name)
		}
		if flags["unstructuredName"].IsSet() {
			person.Names[0].UnstructuredName = flags["unstructuredName"].GetString()
			if person.Names[0].UnstructuredName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "UnstructuredName")
			}
		}
		if flags["familyName"].IsSet() {
			person.Names[0].FamilyName = flags["familyName"].GetString()
			if person.Names[0].FamilyName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "FamilyName")
			}
		}
		if flags["givenName"].IsSet() {
			person.Names[0].GivenName = flags["givenName"].GetString()
			if person.Names[0].GivenName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "GivenName")
			}
		}
		if flags["middleName"].IsSet() {
			person.Names[0].MiddleName = flags["middleName"].GetString()
			if person.Names[0].MiddleName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "MiddleName")
			}
		}
		if flags["honorificPrefix"].IsSet() {
			person.Names[0].HonorificPrefix = flags["honorificPrefix"].GetString()
			if person.Names[0].HonorificPrefix == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "HonorificPrefix")
			}
		}
		if flags["honorificSuffix"].IsSet() {
			person.Names[0].HonorificSuffix = flags["honorificSuffix"].GetString()
			if person.Names[0].HonorificSuffix == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "HonorificSuffix")
			}
		}
		if flags["phoneticFullName"].IsSet() {
			person.Names[0].PhoneticFullName = flags["phoneticFullName"].GetString()
			if person.Names[0].PhoneticFullName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "PhoneticFullName")
			}
		}
		if flags["phoneticFamilyName"].IsSet() {
			person.Names[0].PhoneticFamilyName = flags["phoneticFamilyName"].GetString()
			if person.Names[0].PhoneticFamilyName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "PhoneticFamilyName")
			}
		}
		if flags["phoneticGivenName"].IsSet() {
			person.Names[0].PhoneticGivenName = flags["phoneticGivenName"].GetString()
			if person.Names[0].PhoneticGivenName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "PhoneticGivenName")
			}
		}
		if flags["phoneticMiddleName"].IsSet() {
			person.Names[0].PhoneticMiddleName = flags["phoneticMiddleName"].GetString()
			if person.Names[0].PhoneticMiddleName == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "PhoneticMiddleName")
			}
		}
		if flags["phoneticHonorificPrefix"].IsSet() {
			person.Names[0].PhoneticHonorificPrefix = flags["phoneticHonorificPrefix"].GetString()
			if person.Names[0].PhoneticHonorificPrefix == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "PhoneticHonorificPrefix")
			}
		}
		if flags["phoneticHonorificSuffix"].IsSet() {
			person.Names[0].PhoneticHonorificSuffix = flags["phoneticHonorificSuffix"].GetString()
			if person.Names[0].PhoneticHonorificSuffix == "" {
				person.Names[0].ForceSendFields = append(person.Names[0].ForceSendFields, "PhoneticHonorificSuffix")
			}
		}
	}
	if flags["nicknames"].IsSet() {
		person.Nicknames = []*people.Nickname{}
		nicknames := flags["nicknames"].GetStringSlice()
		if len(nicknames) > 0 {
			for i := range nicknames {
				m := gsmhelpers.FlagToMap(nicknames[i])
				nickname := &people.Nickname{
					Type:  m["type"],
					Value: m["value"],
				}
				nickname.Metadata = addMetaData(m["primary"])
				person.Nicknames = append(person.Nicknames, nickname)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Nicknames")
		}
	}
	if flags["occupations"].IsSet() {
		person.Occupations = []*people.Occupation{}
		occupations := flags["occupations"].GetStringSlice()
		if len(occupations) > 0 {
			for i := range occupations {
				m := gsmhelpers.FlagToMap(occupations[i])
				occupation := &people.Occupation{
					Value: m["value"],
				}
				occupation.Metadata = addMetaData(m["primary"])
				person.Occupations = append(person.Occupations, occupation)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Occupations")
		}
	}
	if flags["organizations"].IsSet() {
		person.Organizations = []*people.Organization{}
		organizations := flags["organizations"].GetStringSlice()
		if len(organizations) > 0 {
			for i := range organizations {
				m := gsmhelpers.FlagToMap(organizations[i])
				organization := &people.Organization{
					Type:           m["type"],
					Department:     m["department"],
					Domain:         m["domain"],
					EndDate:        stringsToDate(m["endDateYear"], m["endDateMonth"], m["endDateDay"]),
					JobDescription: m["jobDescription"],
					Location:       m["location"],
					Name:           m["name"],
					PhoneticName:   m["phoneticName"],
					StartDate:      stringsToDate(m["startDateYear"], m["startDateMonth"], m["startDateDay"]),
					Symbol:         m["symbol"],
					Title:          m["title"],
				}
				organization.Current, _ = strconv.ParseBool(m["current"])
				organization.Metadata = addMetaData(m["primary"])
				person.Organizations = append(person.Organizations, organization)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Organizations")
		}
	}
	if flags["phoneNumbers"].IsSet() {
		person.PhoneNumbers = []*people.PhoneNumber{}
		phoneNumbers := flags["phoneNumbers"].GetStringSlice()
		if len(phoneNumbers) > 0 {
			for i := range phoneNumbers {
				m := gsmhelpers.FlagToMap(phoneNumbers[i])
				phoneNumber := &people.PhoneNumber{
					Type:  m["type"],
					Value: m["value"],
				}
				phoneNumber.Metadata = addMetaData(m["primary"])
				person.PhoneNumbers = append(person.PhoneNumbers, phoneNumber)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "PhoneNumbers")
		}
	}
	if flags["relations"].IsSet() {
		person.Relations = []*people.Relation{}
		relations := flags["relations"].GetStringSlice()
		if len(relations) > 0 {
			for i := range relations {
				m := gsmhelpers.FlagToMap(relations[i])
				relation := &people.Relation{
					Type:   m["type"],
					Person: m["person"],
				}
				relation.Metadata = addMetaData(m["primary"])
				person.Relations = append(person.Relations, relation)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Relations")
		}
	}
	if flags["sipAddresses"].IsSet() {
		person.SipAddresses = []*people.SipAddress{}
		sipAddresses := flags["sipAddresses"].GetStringSlice()
		if len(sipAddresses) > 0 {
			for i := range sipAddresses {
				m := gsmhelpers.FlagToMap(sipAddresses[i])
				sipaddress := &people.SipAddress{
					Type:  m["type"],
					Value: m["value"],
				}
				sipaddress.Metadata = addMetaData(m["primary"])
				person.SipAddresses = append(person.SipAddresses, sipaddress)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "SipAddresses")
		}
	}
	if flags["skills"].IsSet() {
		person.Skills = []*people.Skill{}
		skills := flags["skills"].GetStringSlice()
		if len(skills) > 0 {
			for i := range skills {
				m := gsmhelpers.FlagToMap(skills[i])
				skill := &people.Skill{
					Value: m["value"],
				}
				skill.Metadata = addMetaData(m["primary"])
				person.Skills = append(person.Skills, skill)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Skills")
		}
	}
	if flags["urls"].IsSet() {
		person.Urls = []*people.Url{}
		urls := flags["urls"].GetStringSlice()
		if len(urls) > 0 {
			for i := range urls {
				m := gsmhelpers.FlagToMap(urls[i])
				url := &people.Url{
					Value: m["value"],
				}
				url.Metadata = addMetaData(m["primary"])
				person.Urls = append(person.Urls, url)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "Urls")
		}
	}
	if flags["userDefined"].IsSet() {
		person.UserDefined = []*people.UserDefined{}
		userDefined := flags["userDefined"].GetStringSlice()
		if len(userDefined) > 0 {
			for i := range userDefined {
				m := gsmhelpers.FlagToMap(userDefined[i])
				uDef := &people.UserDefined{
					Value: m["value"],
					Key:   m["key"],
				}
				uDef.Metadata = addMetaData(m["primary"])
				person.UserDefined = append(person.UserDefined, uDef)
			}
		} else {
			person.ForceSendFields = append(person.ForceSendFields, "UserDefined")
		}
	}
	return person, nil
}

func mapToUpdateContactPhotoRequest(flags map[string]*gsmhelpers.Value) (*people.UpdateContactPhotoRequest, error) {
	updateContactPhotoRequest := &people.UpdateContactPhotoRequest{}
	f, err := os.Open(flags["photo"].GetString())
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(f)
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	updateContactPhotoRequest.PhotoBytes = base64.StdEncoding.EncodeToString(content)
	return updateContactPhotoRequest, nil
}
