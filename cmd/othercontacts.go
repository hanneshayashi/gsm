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
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/people/v1"
)

// otherContactsCmd represents the otherContacts command
var otherContactsCmd = &cobra.Command{
	Use:               "otherContacts",
	Short:             "Manage 'other' contacts (Part of People API)",
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/otherContacts",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var otherContactFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"resourceName": {
		AvailableFor: []string{"copyOtherContactToMyContactsGroup"},
		Type:         "string",
		Description:  `The resource name of the "Other contact".`,
		Required:     []string{"copyOtherContactToMyContactsGroup"},
	},
	"copyMask": {
		AvailableFor: []string{"copyOtherContactToMyContactsGroup"},
		Type:         "string",
		Description: `A field mask to restrict which fields are copied into the new contact.
Valid values are:
  - emailAddresses
  - names
  - phoneNumbers`,
		Required: []string{"copyOtherContactToMyContactsGroup"},
	},
	"readMask": {
		AvailableFor: []string{"copyOtherContactToMyContactsGroup", "list"},
		Type:         "string",
		Description: `A field mask to restrict which fields on the person are returned. Multiple fields can be specified by separating them with commas.
Defaults to the copy mask with metadata and membership fields if not set.
Valid values are:
  - addresses
  - ageRanges
  - biographies
  - birthdays
  - calendarUrls
  - clientData
  - coverPhotos
  - emailAddresses
  - events
  - externalIds
  - genders
  - imClients
  - interests
  - locales
  - locations
  - memberships
  - metadata
  - miscKeywords
  - names
  - nicknames
  - occupations
  - organizations
  - phoneNumbers
  - photos
  - relations
  - sipAddresses
  - skills
  - urls
  - userDefined`,
		Required: []string{"list"},
	},
	"sources": {
		AvailableFor: []string{"copyOtherContactToMyContactsGroup"},
		Type:         "stringSlice",
		Description: `A mask of what source types to return.
READ_SOURCE_TYPE_PROFILE         - Returns SourceType.ACCOUNT, SourceType.DOMAIN_PROFILE, and SourceType.PROFILE.
READ_SOURCE_TYPE_CONTACT         - Returns SourceType.CONTACT.
READ_SOURCE_TYPE_DOMAIN_CONTACT  - Returns SourceType.DOMAIN_CONTACT.`,
	},
	"fields": {
		AvailableFor: []string{"copyOtherContactToMyContactsGroup", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(otherContactsCmd)
}

func mapToCopyOtherContactToMyContactsGroupRequest(flags map[string]*gsmhelpers.Value) (*people.CopyOtherContactToMyContactsGroupRequest, error) {
	copyOtherContactToMyContactsGroupRequest := &people.CopyOtherContactToMyContactsGroupRequest{}
	if flags["copyMask"].IsSet() {
		copyOtherContactToMyContactsGroupRequest.CopyMask = flags["copyMask"].GetString()
		if copyOtherContactToMyContactsGroupRequest.CopyMask == "" {
			copyOtherContactToMyContactsGroupRequest.ForceSendFields = append(copyOtherContactToMyContactsGroupRequest.ForceSendFields, "CopyMask")
		}
	}
	if flags["readMask"].IsSet() {
		copyOtherContactToMyContactsGroupRequest.ReadMask = flags["readMask"].GetString()
		if copyOtherContactToMyContactsGroupRequest.ReadMask == "" {
			copyOtherContactToMyContactsGroupRequest.ForceSendFields = append(copyOtherContactToMyContactsGroupRequest.ForceSendFields, "ReadMask")
		}
	}
	if flags["sources"].IsSet() {
		copyOtherContactToMyContactsGroupRequest.Sources = flags["sources"].GetStringSlice()
		if len(copyOtherContactToMyContactsGroupRequest.Sources) == 0 {
			copyOtherContactToMyContactsGroupRequest.ForceSendFields = append(copyOtherContactToMyContactsGroupRequest.ForceSendFields, "Sources")
			copyOtherContactToMyContactsGroupRequest.NullFields = append(copyOtherContactToMyContactsGroupRequest.NullFields, "Sources")
		}
	}
	return copyOtherContactToMyContactsGroupRequest, nil
}
