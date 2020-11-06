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
	"github.com/google/uuid"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// CreateDrive creates a new shared drive.
func CreateDrive(drive *drive.Drive, fields string) (*drive.Drive, error) {
	srv := getDrivesService()
	u, _ := uuid.NewRandom()
	requestID := u.String()
	c := srv.Create(requestID, drive)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// DeleteDrive permanently deletes a shared drive for which the user is an organizer. The shared drive cannot contain any untrashed items.
func DeleteDrive(driveID string) (bool, error) {
	srv := getDrivesService()
	err := srv.Delete(driveID).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetDrive gets a shared drive's metadata by ID.
func GetDrive(driveID, fields string, useDomainAdminAccess bool) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Get(driveID).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// HideDrive hides a shared drive from the default view.
func HideDrive(driveID, fields string) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Hide(driveID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

func makeListDrivesCallAndAppend(c *drive.DrivesListCall, drives []*drive.Drive) ([]*drive.Drive, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, d := range r.Drives {
		drives = append(drives, d)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		drives, err = makeListDrivesCallAndAppend(c, drives)
	}
	return drives, err
}

// ListDrives lists the user's shared drives.
// This method accepts the q parameter, which is a search query combining one or more search terms.
// For more information, see https://developers.google.com/drive/api/v3/search-shareddrives.
func ListDrives(filter, fields string, useDomainAdminAccess bool) ([]*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.List().UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c = c.Q(filter)
	}
	var drives []*drive.Drive
	drives, err := makeListDrivesCallAndAppend(c, drives)
	return drives, err
}

// UnhideDrive restores a shared drive to the default view.
func UnhideDrive(driveID, fields string) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Unhide(driveID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// UpdateDrive updates the metadate for a shared drive.
func UpdateDrive(driveID, fields string, useDomainAdminAccess bool, drive *drive.Drive) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Update(driveID, drive).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
