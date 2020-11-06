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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/go-test/deep"
	"google.golang.org/api/drive/v3"
)

func TestAboutGet(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    *drive.About
		wantErr bool
	}{
		{name: "all fields", args: []string{"about", "get", "--fields", "*"}, want: &drive.About{User: &drive.User{DisplayName: "Hannes Hayashi", EmailAddress: "hannes.siefert@gmail.com", Me: true, Kind: "drive#user", PermissionId: "06181217780026370843", PhotoLink: "https://lh3.googleusercontent.com/a-/AOh14Girr9JGV8udLqPRxQjH8XeWEaVLNlfLKOf965pkIMs=s64"}}, wantErr: false},
		{name: "diplayName", args: []string{"about", "get", "--fields", "user(displayName)"}, want: &drive.About{User: &drive.User{DisplayName: "Hannes Hayashi"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := bytes.NewBufferString("")
			rootCmd.SetOut(out)
			rootCmd.SetArgs(tt.args)
			rootCmd.Execute()
			got, err := ioutil.ReadAll(out)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAbout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotObj := &drive.About{}
			json.Unmarshal(got, gotObj)
			diff := deep.Equal(tt.want.User, gotObj.User)
			if len(diff) > 0 {
				t.Error(diff)
				return
			}
		})
	}
}
