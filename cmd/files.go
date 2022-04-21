/*
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
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:               "files",
	Short:             "Managed files (Part of Drive API)",
	Long:              "Implements the API documented at https://developers.google.com/drive/api/v3/reference/files",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var fileFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fileId": {
		AvailableFor:   []string{"copy", "delete", "export", "get", "move", "update", "download"},
		Type:           "string",
		Description:    "The ID of the file",
		Required:       []string{"copy", "delete", "export", "get", "move", "update", "download"},
		ExcludeFromAll: true,
	},
	"ignoreDefaultVisibility": {
		AvailableFor: []string{"copy", "create"},
		Type:         "bool",
		Description: `Whether to ignore the domain's default visibility settings for the created file.
Domain administrators can choose to make all uploaded files visible to the domain by default;
this parameter bypasses that behavior for the request.
Permissions are still inherited from parent folders.`,
	},
	"includePermissionsForView": {
		AvailableFor: []string{"copy", "create", "get", "list", "update"},
		Type:         "string",
		Description: `Specifies which additional view's permissions to include in the response.
Only 'published' is supported.`,
	},
	"keepRevisionForever": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "bool",
		Description: `Whether to set the 'keepForever' field in the new head revision.
This is only applicable to files with binary content in Google Drive.
Only 200 revisions for the file can be kept forever.
If the limit is reached, try deleting pinned revisions.`,
	},
	"ocrLanguage": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description:  "A language hint for OCR processing during image import (ISO 639-1 code).",
	},
	"appProperties": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "stringSlice",
		Description: `A collection of arbitrary key-value pairs which are private to the requesting app.
Entries with null values are cleared in update and copy requests.`,
	},
	"thumbnailImage": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description:  `The thumbnail data encoded with URL-safe Base64 (RFC 4648 section 5).`,
	},
	"thumbnailMimeType": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description:  `The MIME type of the thumbnail.`,
	},
	"readOnly": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "bool",
		Description: `Whether the content of the file is read-only.
If a file is read-only, a new revision of the file may not be added, comments may not be added or modified, and the title of the file may not be modified.`,
	},
	"readOnlyReason": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description: `Reason for why the content of the file is restricted.
This is only mutable on requests that also set readOnly=true.`,
	},
	"copyRequiresWriterPermission": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "bool",
		Description:  `Whether the options to copy, print, or download this file, should be disabled for readers and commenters.`,
	},
	"description": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description:  `A short description of the file.`,
	},
	"mimeType": {
		AvailableFor: []string{"copy", "create", "export", "update"},
		Type:         "string",
		Description: `The MIME type of the file.
Google Drive will attempt to automatically detect an appropriate value from uploaded content if no value is provided.
The value cannot be changed unless a new revision is uploaded.

If a file is created with a Google Doc MIME type, the uploaded content will be imported if possible.
The supported import formats are published in the About resource.`,
		Required: []string{"export"},
	},
	"modifiedTime": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description: `The last time the file was modified by anyone (RFC 3339 date-time).
Note that setting modifiedTime will also update modifiedByMeTime for the user.`,
	},
	"name": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description: `The name of the file. This is not necessarily unique within a folder.
Note that for immutable items such as the top level folders of shared drives, My Drive root folder, and Application Data folder the name is constant.`,
	},
	"parent": {
		AvailableFor: []string{"copy", "create", "move", "update"},
		Type:         "string",
		Description:  `The single parent of the file.`,
		Required:     []string{"move"},
		Recursive:    []string{"copy", "move"},
	},
	"properties": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "stringSlice",
		Description: `A collection of arbitrary key-value pairs which are visible to all apps.
Entries with null values are cleared in update and copy requests.`,
	},
	"starred": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "bool",
		Description:  `Whether the user has starred the file.`,
	},
	"viewedByMeTime": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "string",
		Description:  `The last time the file was viewed by the user (RFC 3339 date-time).`,
	},
	"writersCanShare": {
		AvailableFor: []string{"copy", "create", "update"},
		Type:         "bool",
		Description: `Whether users with only writer permission can modify the file's permissions.
Not populated for items in shared drives.`,
	},
	"useContentAsIndexableText": {
		AvailableFor: []string{"create", "update"},
		Type:         "bool",
		Description: `Whether users with only writer permission can modify the file's permissions.
Not populated for items in shared drives.`,
	},
	"indexableText": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `Text to be indexed for the file to improve fullText queries.
This is limited to 128KB in length and may contain HTML elements.`,
	},
	"createdTime": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `The time at which the file was created (RFC 3339 date-time).`,
	},
	"folderColorRgb": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The color for a folder as an RGB hex string.
The supported colors are published in the folderColorPalette field of the About resource.
If an unsupported color is specified, the closest color in the palette will be used instead.`,
	},
	"id": {
		AvailableFor:   []string{"copy", "create"},
		Type:           "string",
		Description:    `The ID of the file.`,
		ExcludeFromAll: true,
	},
	"originalFilename": {
		AvailableFor: []string{"create", "update"},
		Type:         "string",
		Description: `The original filename of the uploaded content if available, or else the original value of the name field.
This is only available for files with binary content in Google Drive.`,
	},
	"targetId": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `The ID of the file that this shortcut points to.`,
	},
	"count": {
		AvailableFor: []string{"generateIds"},
		Type:         "int64",
		Description:  `The number of IDs to return. Acceptable values are 1 to 1000, inclusive. (Default: 10)`,
		Defaults:     map[string]any{"generateIds": int64(10)},
	},
	"space": {
		AvailableFor: []string{"generateIds"},
		Type:         "string",
		Description: `The space in which the IDs can be used to create new files.
Supported values are 'drive' and 'appDataFolder'.`,
	},
	"corpora": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Groupings of files to which the query applies.
Supported groupings are:
'user' (files created by, opened by, or shared directly with the user)
'drive' (files in the specified shared drive as indicated by the 'driveId')
'domain' (files shared to the user's domain)
'allDrives' (A combination of 'user' and 'drive' for all drives where the user is a member).
When able, use 'user' or 'drive', instead of 'allDrives', for efficiency.`,
	},
	"driveId": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `ID of the shared drive.`,
	},
	"includeItemsFromAllDrives": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `Whether both My Drive and shared drive items should be included in results.`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A comma-separated list of sort keys.
Valid keys are 'createdTime', 'folder', 'modifiedByMeTime', 'modifiedTime', 'name', 'name_natural', 'quotaBytesUsed', 'recency', 'sharedWithMeTime', 'starred', and 'viewedByMeTime'.
Each key sorts ascending by default, but may be reversed with the 'desc' modifier.
Example usage: ?orderBy=folder,modifiedTime desc,name.
Please note that there is a current limitation for users with approximately one million files in which the requested sort order is ignored.`,
	},
	"q": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A query for filtering the file results.
See the https://developers.google.com/drive/api/v3/search-files for the supported syntax.`,
	},
	"spaces": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `A comma-separated list of spaces to query within the corpus.
Supported values are 'drive', 'appDataFolder' and 'photos'.`,
	},
	"trashed": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description: `Whether the file has been trashed, either explicitly or from a trashed parent folder.
Only the owner may trash a file.
The trashed item is excluded from all files.list responses returned for any user who does not own the file.
However, all users with access to the file can see the trashed item metadata in an API response.
All users with access can copy, download, export, and share the file.`,
	},
	"localFilePath": {
		AvailableFor: []string{"create", "update", "download", "export"},
		Type:         "string",
		Description:  `Path to a file or folder on the local disk.`,
		Required:     []string{"download", "export"},
	},
	"acknowledgeAbuse": {
		AvailableFor: []string{"download"},
		Type:         "bool",
		Description:  `Whether the user is acknowledging the risk of downloading known malware or other abusive files.`,
	},
	"folderId": {
		AvailableFor: []string{"count"},
		Type:         "string",
		Description:  `Id of the folder.`,
	},
	"fields": {
		AvailableFor: []string{"copy", "create", "get", "list", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Recursive: []string{"copy", "create", "get", "list", "update"},
	},
}
var fileFlagsALL = gsmhelpers.GetAllFlags(fileFlags)

func init() {
	rootCmd.AddCommand(filesCmd)
}

func mapToFile(flags map[string]*gsmhelpers.Value) (*drive.File, error) {
	file := &drive.File{}
	if flags["appProperties"].IsSet() {
		file.AppProperties = make(map[string]string)
		ap := flags["appProperties"].GetString()
		if len(ap) > 0 {
			file.AppProperties = gsmhelpers.FlagToMap(ap)
		} else {
			file.ForceSendFields = append(file.ForceSendFields, "AppProperties")
		}
	}
	if flags["thumbnailImage"].IsSet() || flags["thumbnailMimeType"].IsSet() || flags["indexableText"].IsSet() {
		file.ContentHints = &drive.FileContentHints{}
		if flags["thumbnailImage"].IsSet() || flags["thumbnailMimeType"].IsSet() {
			file.ContentHints.Thumbnail = &drive.FileContentHintsThumbnail{}
			if flags["thumbnailImage"].IsSet() {
				file.ContentHints.Thumbnail.Image = flags["thumbnailImage"].GetString()
				if file.ContentHints.Thumbnail.Image == "" {
					file.ContentHints.Thumbnail.ForceSendFields = append(file.ContentHints.Thumbnail.ForceSendFields, "Image")
				}
			}
			if flags["thumbnailMimeType"].IsSet() {
				file.ContentHints.Thumbnail.MimeType = flags["thumbnailMimeType"].GetString()
				if file.ContentHints.Thumbnail.MimeType == "" {
					file.ContentHints.Thumbnail.ForceSendFields = append(file.ContentHints.Thumbnail.ForceSendFields, "MimeType")
				}
			}
		}
		if flags["indexableText"].IsSet() {
			file.ContentHints.IndexableText = flags["indexableText"].GetString()
			if file.ContentHints.IndexableText == "" {
				file.ContentHints.ForceSendFields = append(file.ContentHints.ForceSendFields, "IndexableText")
			}
		}
	}
	if flags["readOnly"].IsSet() || flags["readOnlyReason"].IsSet() {
		file.ContentRestrictions = []*drive.ContentRestriction{}
		file.ContentRestrictions = append(file.ContentRestrictions, &drive.ContentRestriction{})
		if flags["readOnly"].IsSet() {
			file.ContentRestrictions[0].ReadOnly = flags["readOnly"].GetBool()
			if !file.ContentRestrictions[0].ReadOnly {
				file.ContentRestrictions[0].ForceSendFields = append(file.ContentRestrictions[0].ForceSendFields, "ReadOnly")
			}
		}
		if flags["readOnlyReason"].IsSet() {
			file.ContentRestrictions[0].Reason = flags["readOnlyReason"].GetString()
			if file.ContentRestrictions[0].Reason == "" {
				file.ContentRestrictions[0].ForceSendFields = append(file.ContentRestrictions[0].ForceSendFields, "Reason")
			}
		}
	}
	if flags["copyRequiresWriterPermission"].IsSet() {
		file.CopyRequiresWriterPermission = flags["copyRequiresWriterPermission"].GetBool()
		if !file.CopyRequiresWriterPermission {
			file.ForceSendFields = append(file.ForceSendFields, "CopyRequiresWriterPermission")
		}
	}
	if flags["description"].IsSet() {
		file.Description = flags["description"].GetString()
		if file.Description == "" {
			file.ForceSendFields = append(file.ForceSendFields, "Description")
		}
	}
	if flags["mimeType"].IsSet() {
		file.MimeType = flags["mimeType"].GetString()
		if file.MimeType == "" {
			file.ForceSendFields = append(file.ForceSendFields, "MimeType")
		}
	}
	if flags["modifiedTime"].IsSet() {
		file.ModifiedTime = flags["modifiedTime"].GetString()
		if file.ModifiedTime == "" {
			file.ForceSendFields = append(file.ForceSendFields, "ModifiedTime")
		}
	}
	if flags["name"].IsSet() {
		file.Name = flags["name"].GetString()
		if file.Name == "" {
			file.ForceSendFields = append(file.ForceSendFields, "Name")
		}
	}
	if flags["properties"].IsSet() {
		file.Properties = make(map[string]string)
		p := flags["properties"].GetString()
		if len(p) > 0 {
			file.Properties = gsmhelpers.FlagToMap(p)
		} else {
			file.ForceSendFields = append(file.ForceSendFields, "Properties")
		}
	}
	if flags["starred"].IsSet() {
		file.Starred = flags["starred"].GetBool()
		if !file.Starred {
			file.ForceSendFields = append(file.ForceSendFields, "Starred")
		}
	}
	if flags["viewedByMeTime"].IsSet() {
		file.ViewedByMeTime = flags["viewedByMeTime"].GetString()
		if file.ViewedByMeTime == "" {
			file.ForceSendFields = append(file.ForceSendFields, "ViewedByMeTime")
		}
	}
	if flags["writersCanShare"].IsSet() {
		file.WritersCanShare = flags["writersCanShare"].GetBool()
		if !file.WritersCanShare {
			file.ForceSendFields = append(file.ForceSendFields, "WritersCanShare")
		}
	}
	if flags["createdTime"] != nil && flags["createdTime"].IsSet() {
		file.CreatedTime = flags["createdTime"].GetString()
		if file.CreatedTime == "" {
			file.ForceSendFields = append(file.ForceSendFields, "CreatedTime")
		}
	}
	if flags["folderColorRgb"].IsSet() {
		file.FolderColorRgb = flags["folderColorRgb"].GetString()
		if file.FolderColorRgb == "" {
			file.ForceSendFields = append(file.ForceSendFields, "FolderColorRgb")
		}
	}
	if flags["id"].IsSet() {
		file.Id = flags["id"].GetString()
		if file.Id == "" {
			file.ForceSendFields = append(file.ForceSendFields, "Id")
		}
	}
	if flags["originalFilename"].IsSet() {
		file.OriginalFilename = flags["originalFilename"].GetString()
		if file.OriginalFilename == "" {
			file.ForceSendFields = append(file.ForceSendFields, "OriginalFilename")
		}
	}
	if flags["targetId"].IsSet() {
		file.ShortcutDetails = &drive.FileShortcutDetails{}
		file.ShortcutDetails.TargetId = flags["targetId"].GetString()
		if file.ShortcutDetails.TargetId == "" {
			file.ShortcutDetails.ForceSendFields = append(file.ShortcutDetails.ForceSendFields, "TargetId")
		}
	}
	if flags["trashed"].IsSet() {
		file.Trashed = flags["trashed"].GetBool()
		if !file.Trashed {
			file.ForceSendFields = append(file.ForceSendFields, "Trashed")
		}
	}
	if flags["parent"].IsSet() {
		parent := flags["parent"].GetString()
		if parent == "" {
			file.ForceSendFields = append(file.ForceSendFields, "Parents")
		} else {
			file.Parents = append(file.Parents, parent)
		}
	}
	return file, nil
}
