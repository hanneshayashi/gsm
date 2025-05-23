/*
Copyright © 2020-2023 Hannes Hayashi

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

// revisionsCmd represents the revisions command
var revisionsCmd = &cobra.Command{
	Use:               "revisions",
	Short:             "Manage revisions of non-Google files (Part of Drive API)",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/revisions",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var revisionFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fileId": {
		AvailableFor: []string{"delete", "get", "list", "update"},
		Type:         "string",
		Description:  `The ID of the file.`,
		Required:     []string{"delete", "get", "list", "update"},
	},
	"revisionId": {
		AvailableFor:   []string{"delete", "get", "update"},
		Type:           "string",
		Description:    `The ID of the revision.`,
		Required:       []string{"delete", "get", "update"},
		ExcludeFromAll: true,
	},
	"acknowledgeAbuse": {
		AvailableFor: []string{"delete", "get", "update"},
		Type:         "bool",
		Description: `Whether the user is acknowledging the risk of downloading known malware or other abusive files.
This is only applicable when alt=media.`,
	},
	"keepForever": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description: `Whether to keep this revision forever, even if it is no longer the head revision.
If not set, the revision will be automatically purged 30 days after newer content is uploaded.
This can be set on a maximum of 200 revisions for a file.
This field is only applicable to files with binary content in Drive.`,
	},
	"publishAuto": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description: `Whether subsequent revisions will be automatically republished.
This is only applicable to Google Docs.`,
	},
	"published": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description: `Whether this revision is published.
This is only applicable to Google Docs.`,
	},
	"publishedOutsideDomain": {
		AvailableFor: []string{"update"},
		Type:         "bool",
		Description: `Whether this revision is published outside the domain.
This is only applicable to Google Docs.`,
	},
	"fields": {
		AvailableFor: []string{"get", "list", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var revisionFlagsALL = gsmhelpers.GetAllFlags(revisionFlags)

func init() {
	rootCmd.AddCommand(revisionsCmd)
}

func mapToRevision(flags map[string]*gsmhelpers.Value) (*drive.Revision, error) {
	revision := &drive.Revision{}
	if flags["keepForever"].IsSet() {
		revision.KeepForever = flags["keepForever"].GetBool()
		if !revision.KeepForever {
			revision.ForceSendFields = append(revision.ForceSendFields, "KeepForever")
		}
	}
	if flags["publishAuto"].IsSet() {
		revision.PublishAuto = flags["publishAuto"].GetBool()
		if !revision.PublishAuto {
			revision.ForceSendFields = append(revision.ForceSendFields, "PublishAuto")
		}
	}
	if flags["published"].IsSet() {
		revision.Published = flags["published"].GetBool()
		if !revision.Published {
			revision.ForceSendFields = append(revision.ForceSendFields, "Published")
		}
	}
	if flags["publishedOutsideDomain"].IsSet() {
		revision.PublishedOutsideDomain = flags["publishedOutsideDomain"].GetBool()
		if !revision.PublishedOutsideDomain {
			revision.ForceSendFields = append(revision.ForceSendFields, "PublishedOutsideDomain")
		}
	}
	return revision, nil
}
