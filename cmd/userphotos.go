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
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// userPhotosCmd represents the userPhotos command
var userPhotosCmd = &cobra.Command{
	Use:               "userPhotos",
	Short:             "Manage user photos (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/users.photos",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var userPhotoFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userKey": {
		AvailableFor: []string{"delete", "get", "update"},
		Type:         "string",
		Description: `Identifies the user in the API request.
The value can be the user's primary email address, alias email address, or unique user ID.`,
		Required:       []string{"delete", "get", "update"},
		ExcludeFromAll: true,
	},
	"photo": {
		AvailableFor: []string{"update"},
		Type:         "string",
		Description: `Path to the photo file.
Allowed formats are: jpeg, png, gif, bmp and tiff.`,
		Required:  []string{"update"},
		Recursive: []string{"update"},
	},
	"fields": {
		AvailableFor: []string{"get", "update"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
		Recursive: []string{"get", "update"},
	},
}
var userPhotoFlagsALL = gsmhelpers.GetAllFlags(userPhotoFlags)

func init() {
	rootCmd.AddCommand(userPhotosCmd)
}

func mapToUserPhoto(flags map[string]*gsmhelpers.Value) (*admin.UserPhoto, error) {
	userPhoto := &admin.UserPhoto{}
	if flags["photo"].IsSet() {
		f, err := os.Open(flags["photo"].GetString())
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		userPhoto.PhotoData = base64.RawURLEncoding.EncodeToString(b)
		if userPhoto.PhotoData == "" {
			userPhoto.ForceSendFields = append(userPhoto.ForceSendFields, "PhotoData")
		}
	}
	return userPhoto, nil
}
