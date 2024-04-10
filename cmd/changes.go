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
)

// changesCmd represents the changes command
var changesCmd = &cobra.Command{
	Use:               "changes",
	Short:             "View changes to user's or Shared Drive (Part of Drive API)",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/changes",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var changeFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"driveId": {
		AvailableFor: []string{"getStartPageToken", "list"},
		Type:         "string",
		Description:  `The ID of the shared drive`,
	},
	"pageToken": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The token for continuing a previous list request on the next page.
This should be set to the value of 'nextPageToken' from the previous response or to the response from the getStartPageToken method.`,
		Required: []string{"list"},
	},
	"includeCorpusRemovals": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether changes should include the file resource if the file is still accessible by the user at the time of the request, even when a file was removed from the list of changes and there will be no further change entries for this file.`,
	},
	"includeItemsFromAllDrives": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether both My Drive and shared drive items should be included in results.`,
	},
	"includePermissionsForView": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Specifies which additional view's permissions to include in the response.
Only 'published' is supported.`,
	},
	"includeRemoved": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether to include changes indicating that items have been removed from the list of changes, for example by deletion or loss of access.`,
	},
	"restrictToMyDrive": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description: `Whether to restrict the results to changes inside the My Drive hierarchy.
This omits changes to files such as those in the Application Data folder or shared files which have not been added to My Drive.`,
	},
	"spaces": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `A comma-separated list of spaces to query within the user corpus. Supported values are 'drive', 'appDataFolder' and 'photos'.`,
	},
	"fields": {
		AvailableFor: []string{"getStartPageToken", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(changesCmd)
}
