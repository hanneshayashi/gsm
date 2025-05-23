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

package gsmdrive

import (
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// GetApp gets a specific app. For more information, see https://developers.google.com/workspace/drive/api/guides/user-info.
func GetApp(appId, fields string) (*drive.App, error) {
	srv := getAppsService()
	c := srv.Get(appId)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey("App"), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.App)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListApps lists a user's installed apps, see https://developers.google.com/workspace/drive/api/guides/user-info.
func ListApps(appFilterExtensions, appFilterMimeTypes, languageCode, fields string) (*drive.AppList, error) {
	srv := getAppsService()
	c := srv.List()
	if appFilterExtensions != "" {
		c.AppFilterExtensions(appFilterExtensions)
	}
	if appFilterMimeTypes != "" {
		c.AppFilterMimeTypes(appFilterMimeTypes)
	}
	if languageCode != "" {
		c.LanguageCode(languageCode)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey("ListApps"), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.AppList)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
