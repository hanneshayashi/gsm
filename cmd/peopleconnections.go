/*
Copyright © 2020-2022 Hannes Hayashi

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
)

// peopleConnectionsCmd represents the peopleConnections command
var peopleConnectionsCmd = &cobra.Command{
	Use:               "peopleConnections",
	Short:             "Information about a person merged from various data sources such as the authenticated user's contacts and profile data. (Part of People API)",
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/people.connections",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var peopleConnectionFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"resourceName": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `The resource name to return connections for. Only people/me is valid.`,
		Defaults:     map[string]any{"list": "people/me"},
	},
	"personFields": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A field mask to restrict which fields on each person are returned.
Multiple fields can be specified by separating them with commas.
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
	"sortOrder": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Optional. The order in which the connections should be sorted.
Defaults to LAST_MODIFIED_ASCENDING.
Valid values are:
LAST_MODIFIED_ASCENDING   - Sort people by when they were changed; older entries first.
LAST_MODIFIED_DESCENDING  - Sort people by when they were changed; newer entries first.
FIRST_NAME_ASCENDING      - Sort people by first name.
LAST_NAME_ASCENDING       - Sort people by last name.`,
	},
	"sources": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A mask of what source types to return.
READ_SOURCE_TYPE_PROFILE         - Returns SourceType.ACCOUNT, SourceType.DOMAIN_PROFILE, and SourceType.PROFILE.
READ_SOURCE_TYPE_CONTACT         - Returns SourceType.CONTACT.
READ_SOURCE_TYPE_DOMAIN_CONTACT  - Returns SourceType.DOMAIN_CONTACT.`,
	},
	"fields": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(peopleConnectionsCmd)
}
