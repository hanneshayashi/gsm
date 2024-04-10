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
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// drivesCmd represents the drives command
var drivesCmd = &cobra.Command{
	Use:               "drives",
	Short:             "Manage Shared Drives (Part of Drive API)",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/drives",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var driveFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"driveId": {
		AvailableFor:   []string{"delete", "get", "hide", "unhide", "update", "getSize"},
		Type:           "string",
		Description:    "The ID of the shared drive",
		Required:       []string{"delete", "get", "hide", "unhide", "update", "getSize"},
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
It can only be set on a drive.drives.update request that does not set themeId.`,
	},
	"backgroundImageFile": {
		AvailableFor: []string{"update"},
		Type:         "string",
		Description: `An image file and cropping parameters from which a background image for this shared drive is set.
This is a write only field; it can only be set on drive.drives.update requests that don't set themeId.
When specified, all fields of the backgroundImageFile must be set.
Specify in the following format:
'--backgroundImageFile "id=...;width=...;xCoordinate=...;yCoordinate=..."'
Use ALL of the following parameters:
id:           The ID of an image file in Google Drive to use for the background image.
width:        The width of the cropped image in the closed range of 0 to 1.
              This value represents the width of the cropped image divided by the width of the entire image.
              The height is computed by applying a width to height aspect ratio of 80 to 9.
              The resulting image must be at least 1280 pixels wide and 144 pixels high.
xCoordinate:  The X coordinate of the upper left corner of the cropping area in the background image.
              This is a value in the closed range of 0 to 1.
              This value represents the horizontal distance from the left side of the entire image to the left side of the cropping area divided by the width of the entire image.
yCoordinate:  The Y coordinate of the upper left corner of the cropping area in the background image.
              This is a value in the closed range of 0 to 1.
              This value represents the vertical distance from the top side of the entire image to the top side of the cropping area divided by the height of the entire image.`,
	},
	"useDomainAdminAccess": {
		AvailableFor: []string{"create", "get", "hide", "list", "unhide", "update", "delete"},
		Type:         "bool",
		Description:  "Issue the request as a domain administrator",
	},
	"adminManagedRestrictions": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description:  "Whether administrative privileges on this shared drive are required to modify restrictions",
	},
	"copyRequiresWriterPermission": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description: `Whether the options to copy, print, or download files inside this shared drive, should be disabled for readers and commenters.
When this restriction is set to true, it will override the similarly named field to true for any file inside this shared drive`,
	},
	"domainUsersOnly": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description: `Whether access to this shared drive and items inside this shared drive is restricted to users of the domain to which this shared drive belongs.
This restriction may be overridden by other sharing policies controlled outside of this shared drive`,
	},
	"driveMembersOnly": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description:  "Whether access to items inside this shared drive is restricted to its members",
	},
	"q": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Query string for searching shared drives.
See the https://developers.google.com/drive/api/v3/search-shareddrives for supported syntax.`,
	},
	"includeTrash": {
		AvailableFor: []string{"getSize"},
		Type:         "bool",
		Description:  `Whether to include trashed items.`,
	},
	"returnWhenReady": {
		AvailableFor: []string{"create"},
		Type:         "bool",
		Description: `The Google Drive API returns the drive after creation immediately, but usually before it can be used in subsequent requests.
Setting this flag will cause GSM to try and do the follwing on the newly created drive to make sure that it is available before returning it:
1. Get the Drive by its driveId
2. Get the user's permissionId from the about method
3. Get the user's permission on the newly created Drive using the Drive's driveId and the user's permissionId
4. Update the drive's DomainUsersOnly restriction to 'true'
5. Update the drive's DomainUsersOnly restriction to 'false' (default)
Note that this will cause additional API requests that may be subject to your quota.
The API requests are made with useDomainAdminAccess set to 'false'`,
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
	var err error
	d := &drive.Drive{}
	if flags["backgroundImageFile"].IsSet() {
		backgroundImageFile := flags["backgroundImageFile"].GetString()
		if backgroundImageFile != "" {
			m := gsmhelpers.FlagToMap(backgroundImageFile)
			d.BackgroundImageFile = &drive.DriveBackgroundImageFile{
				Id: m["id"],
			}
			d.BackgroundImageFile.Width, err = strconv.ParseFloat(m["width"], 64)
			if err != nil {
				return nil, err
			}
			d.BackgroundImageFile.XCoordinate, err = strconv.ParseFloat(m["xCoordinate"], 64)
			if err != nil {
				return nil, err
			}
			d.BackgroundImageFile.YCoordinate, err = strconv.ParseFloat(m["yCoordinate"], 64)
			if err != nil {
				return nil, err
			}
		} else {
			d.ForceSendFields = append(d.ForceSendFields, "BackgroundImageFile")
		}
	}
	if flags["themeId"].IsSet() {
		d.ThemeId = flags["themeId"].GetString()
		if d.ThemeId == "" {
			d.ForceSendFields = append(d.ForceSendFields, "ThemeId")
		}
	}
	if flags["colorRgb"].IsSet() {
		d.ColorRgb = flags["colorRgb"].GetString()
		if d.ColorRgb == "" {
			d.ForceSendFields = append(d.ForceSendFields, "ColorRgb")
		}
	}
	if flags["name"].IsSet() {
		d.Name = flags["name"].GetString()
		if d.Name == "" {
			d.ForceSendFields = append(d.ForceSendFields, "Name")
		}
	}
	adminManagedRestrictionsSet := flags["adminManagedRestrictions"].IsSet()
	copyRequiresWriterPermissionSet := flags["copyRequiresWriterPermission"].IsSet()
	domainUsersOnlySet := flags["domainUsersOnly"].IsSet()
	driveMembersOnlySet := flags["driveMembersOnly"].IsSet()
	if adminManagedRestrictionsSet || copyRequiresWriterPermissionSet || domainUsersOnlySet || driveMembersOnlySet {
		d.Restrictions = &drive.DriveRestrictions{}
		if adminManagedRestrictionsSet {
			d.Restrictions.AdminManagedRestrictions = flags["adminManagedRestrictions"].GetBool()
			if !d.Restrictions.AdminManagedRestrictions {
				d.Restrictions.ForceSendFields = append(d.Restrictions.ForceSendFields, "AdminManagedRestrictions")
			}
		}
		if copyRequiresWriterPermissionSet {
			d.Restrictions.CopyRequiresWriterPermission = flags["copyRequiresWriterPermission"].GetBool()
			if !d.Restrictions.CopyRequiresWriterPermission {
				d.Restrictions.ForceSendFields = append(d.Restrictions.ForceSendFields, "CopyRequiresWriterPermission")
			}
		}
		if domainUsersOnlySet {
			d.Restrictions.DomainUsersOnly = flags["domainUsersOnly"].GetBool()
			if !d.Restrictions.DomainUsersOnly {
				d.Restrictions.ForceSendFields = append(d.Restrictions.ForceSendFields, "DomainUsersOnly")
			}
		}
		if driveMembersOnlySet {
			d.Restrictions.DriveMembersOnly = flags["driveMembersOnly"].GetBool()
			if !d.Restrictions.DriveMembersOnly {
				d.Restrictions.ForceSendFields = append(d.Restrictions.ForceSendFields, "DriveMembersOnly")
			}
		}
	}
	return d, nil
}
