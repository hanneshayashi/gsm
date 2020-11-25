/*
Package gsmgmail implements the Gmail APIs
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
package gsmgmail

import (
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// DeleteSmimeInfo deletes the specified S/MIME config for the specified send-as alias.
func DeleteSmimeInfo(userID, sendAsEmail, id string) (bool, error) {
	srv := getUsersSettingsSendAsSmimeInfoService()
	err := srv.Delete(userID, sendAsEmail, id).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetSmimeInfo gets the specified S/MIME config for the specified send-as alias.
func GetSmimeInfo(userID, sendAsEmail, id, fields string) (*gmail.SmimeInfo, error) {
	srv := getUsersSettingsSendAsSmimeInfoService()
	c := srv.Get(userID, sendAsEmail, id)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertSmimeInfo uploads the given S/MIME config for the specified send-as alias.
// Note that pkcs12 format is required for the key.
func InsertSmimeInfo(userID, sendAsEmail, fields string, smimeInfo *gmail.SmimeInfo) (*gmail.SmimeInfo, error) {
	srv := getUsersSettingsSendAsSmimeInfoService()
	c := srv.Insert(userID, sendAsEmail, smimeInfo)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// ListSmimeInfo lists S/MIME configs for the specified send-as alias.
func ListSmimeInfo(userID, sendAsEmail, fields string) ([]*gmail.SmimeInfo, error) {
	srv := getUsersSettingsSendAsSmimeInfoService()
	c := srv.List(userID, sendAsEmail)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	return r.SmimeInfo, err
}

// SetDefaultSmimeInfo sets the default S/MIME config for the specified send-as alias.
func SetDefaultSmimeInfo(userID, sendAsEmail, id string) (bool, error) {
	srv := getUsersSettingsSendAsSmimeInfoService()
	err := srv.SetDefault(userID, sendAsEmail, id).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}
