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

package cmd

import (
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Information apps a user has installed (Part of Drive API)",
	Long: `This API only works with the currently authenticated user!
Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/apps`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var appsFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fields": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
	"appId": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description:  `The ID of the app.`,
	},
	"appFilterExtensions": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A comma-separated list of file extensions to limit returned results.
All results within the given app query scope which can open any of the given file extensions are included in the response.
If appFilterMimeTypes are provided as well, the result is a union of the two resulting app lists.`,
	},
	"appFilterMimeTypes": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A comma-separated list of file extensions to limit returned results.
All results within the given app query scope which can open any of the given MIME types will be included in the response.
If appFilterExtensions are provided as well, the result is a union of the two resulting app lists.`,
	},
	"languageCode": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `A language or locale code, as defined by BCP 47, with some extensions from Unicode's LDML format (http://www.unicode.org/reports/tr35/).`,
	},
}

func init() {
	rootCmd.AddCommand(appsCmd)
}
