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

package gsmdrive

import (
	"context"
	"fmt"
	"time"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/google/uuid"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// CreateDrive creates a new shared drive.
func CreateDrive(d *drive.Drive, fields string, returnWhenReady bool) (*drive.Drive, error) {
	srv := getDrivesService()
	u, _ := uuid.NewRandom()
	requestID := u.String()
	c := srv.Create(requestID, d)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(d.Name), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Drive)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	if returnWhenReady {
		about, err := GetAbout("user(permissionId)")
		if err != nil {
			return nil, err
		}
		_, err = GetPermission(r.Id, about.User.PermissionId, "", false)
		for err != nil {
			_, err = GetPermission(r.Id, about.User.PermissionId, "", false)
			time.Sleep(1 * time.Second)
		}
	}
	return r, nil
}

// DeleteDrive permanently deletes a shared drive for which the user is an organizer. The shared drive cannot contain any untrashed items.
func DeleteDrive(driveID string) (bool, error) {
	srv := getDrivesService()
	c := srv.Delete(driveID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(driveID), func() error {
		return c.Do()
	})
	return result, err
}

// GetDrive gets a shared drive's metadata by ID.
func GetDrive(driveID, fields string, useDomainAdminAccess bool) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Get(driveID).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(driveID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Drive)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// HideDrive hides a shared drive from the default view.
func HideDrive(driveID, fields string) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Hide(driveID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(driveID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Drive)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// ListDrives lists the user's shared drives.
// This method accepts the q parameter, which is a search query combining one or more search terms.
// For more information, see https://developers.google.com/drive/api/v3/search-shareddrives.
func ListDrives(filter, fields string, useDomainAdminAccess bool, cap int) (<-chan *drive.Drive, <-chan error) {
	srv := getDrivesService()
	c := srv.List().UseDomainAdminAccess(useDomainAdminAccess).PageSize(100)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c = c.Q(filter)
	}
	ch := make(chan *drive.Drive, cap)
	err := make(chan error, 1)
	go func() {
		e := c.Pages(context.Background(), func(response *drive.DriveList) error {
			for i := range response.Drives {
				ch <- response.Drives[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// UnhideDrive restores a shared drive to the default view.
func UnhideDrive(driveID, fields string) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Unhide(driveID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(driveID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Drive)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// UpdateDrive updates the metadate for a shared drive.
func UpdateDrive(driveID, fields string, useDomainAdminAccess bool, d *drive.Drive) (*drive.Drive, error) {
	srv := getDrivesService()
	c := srv.Update(driveID, d).UseDomainAdminAccess(useDomainAdminAccess)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(driveID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*drive.Drive)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}
