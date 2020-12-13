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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// drivesCmd represents the drives command
var drivesCmd = &cobra.Command{
	Use:   "drives",
	Short: "Manage Shared Drives (Part of Drive API)",
	Long:  "https://developers.google.com/drive/api/v3/reference/drives",	
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var driveFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"driveId": {
		AvailableFor:   []string{"delete", "get", "hide", "unhide", "update"},
		Type:           "string",
		Description:    "The ID of the shared drive",
		Required:       []string{"delete", "get", "hide", "unhide", "update"},
		ExcludeFromAll: true,
	},
	"themeId": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The ID of the theme from which the background image and color will be set.
The set of possible driveThemes can be retrieved from a drive.about.get response.
When not specified on a drive.drives.create request, a random theme is chosen from which the background image and color are set.
This is a write-only field; it can only be set on requests that don't set colorRgb or backgroundImageFile.`,
	},
	"name": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description:  "The name of this shared drive",
	},
	"colorRgb": {
		AvailableFor: []string{"update"},
		Type:         "string",
		Description: `The color of this shared drive as an RGB hex string.
It can only be set on a drive.drives.update request that does not set themeId.	`,
	},
	"useDomainAdminAccess": {
		AvailableFor: []string{"create", "get", "hide", "list", "unhide", "update"},
		Type:         "bool",
		Description:  "Issue the request as a domain administrator",
		Defaults:     map[string]interface{}{"create": true, "get": true, "hide": true, "unhide": true, "update": true},
	},
	"adminManagedRestrictions": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description:  "Whether administrative privileges on this shared drive are required to modify restrictions",
	},
	"copyRequiresWriterPermission": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description:  "Whether the options to copy, print, or download files inside this shared drive, should be disabled for readers and commenters. When this restriction is set to true, it will override the similarly named field to true for any file inside this shared drive",
	},
	"domainUsersOnly": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description:  "Whether access to this shared drive and items inside this shared drive is restricted to users of the domain to which this shared drive belongs. This restriction may be overridden by other sharing policies controlled outside of this shared drive",
	},
	"driveMembersOnly": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description:  "Whether access to items inside this shared drive is restricted to its members",
	},
	"q": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Query string for searching shared drives.
See the https://developers.google.com/drive/api/v3/search-shareddrives for supported syntax.`,
	},
	"fields": {
		AvailableFor: []string{"create", "get", "hide", "list", "unhide", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var driveFlagsALL = gsmhelpers.GetAllFlags(driveFlags)

func init() {
	rootCmd.AddCommand(drivesCmd)
}

func mapToDrive(flags map[string]*gsmhelpers.Value) (*drive.Drive, error) {
	drive := &drive.Drive{}
	drive.Name = flags["name"].GetString()

	return drive, nil
}
