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
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// GetStartPageToken gets the starting pageToken for listing future changes.
func GetStartPageToken(driveID, fields string) (*drive.StartPageToken, error) {
	srv := getChangesService()
	c := srv.GetStartPageToken().SupportsAllDrives(true)
	if driveID != "" {
		c = c.DriveId(driveID)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

func listChanges(pageToken, driveID, spaces, fields, includePermissionsForView string, includeCorpusRemovals, includeItemsFromAllDrives, includeRemoved, restrictToMyDrive bool, changes []*drive.Change) (*drive.ChangeList, error) {
	srv := getChangesService()
	c := srv.List(pageToken).IncludeCorpusRemovals(includeCorpusRemovals).IncludeItemsFromAllDrives(includeItemsFromAllDrives).IncludeRemoved(includeRemoved).RestrictToMyDrive(restrictToMyDrive).SupportsAllDrives(true)
	if driveID != "" {
		c = c.DriveId(driveID)
	}
	if spaces != "" {
		c = c.Spaces(spaces)
	}
	if includePermissionsForView != "" {
		c = c.IncludePermissionsForView(includePermissionsForView)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	changes = append(changes, r.Changes...)
	if r.NextPageToken != "" {
		r, err = listChanges(r.NextPageToken, driveID, spaces, fields, includePermissionsForView, includeCorpusRemovals, includeItemsFromAllDrives, includeRemoved, restrictToMyDrive, changes)
		if err != nil {
			return nil, err
		}
		changes = append(changes, r.Changes...)
	}
	return &drive.ChangeList{Changes: changes}, err
}

// ListChanges lists the changes for a user or shared drive.
func ListChanges(pageToken, driveID, spaces, fields, includePermissionsForView string, includeCorpusRemovals, includeItemsFromAllDrives, includeRemoved, restrictToMyDrive bool) ([]*drive.Change, string, error) {
	var changes []*drive.Change
	r, err := listChanges(pageToken, driveID, spaces, fields, includePermissionsForView, includeCorpusRemovals, includeItemsFromAllDrives, includeRemoved, restrictToMyDrive, changes)
	if err != nil {
		return nil, "", err
	}
	return r.Changes, r.NewStartPageToken, err
}
