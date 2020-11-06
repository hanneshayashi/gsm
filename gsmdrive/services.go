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
	"context"
	"log"
	"net/http"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var (
	client             *http.Client
	driveService       *drive.Service
	filesService       *drive.FilesService
	permissionsService *drive.PermissionsService
	drivesService      *drive.DrivesService
	aboutService       *drive.AboutService
	changesService     *drive.ChangesService
	commentsService    *drive.CommentsService
	repliesService     *drive.RepliesService
	revisionsService   *drive.RevisionsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getDriveService() (driveService *drive.Service) {
	if client == nil {
		log.Fatalf("gsmdrive.client is not set. Set with gsmdrive.SetClient(client)")
	}
	if driveService == nil {
		var err error
		driveService, err = drive.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating drive service: %v", err)
		}
	}
	return driveService
}

func getFilesService() (filesService *drive.FilesService) {
	if filesService == nil {
		filesService = drive.NewFilesService(getDriveService())
	}
	return
}

func getPermissionsService() *drive.PermissionsService {
	if permissionsService == nil {
		permissionsService = drive.NewPermissionsService(getDriveService())
	}
	return permissionsService
}

func getDrivesService() *drive.DrivesService {
	if drivesService == nil {
		drivesService = drive.NewDrivesService(getDriveService())
	}
	return drivesService
}

func getAboutService() *drive.AboutService {
	if aboutService == nil {
		aboutService = drive.NewAboutService(getDriveService())
	}
	return aboutService
}

func getChangesService() *drive.ChangesService {
	if changesService == nil {
		changesService = drive.NewChangesService(getDriveService())
	}
	return changesService
}

func getCommentsService() *drive.CommentsService {
	if commentsService == nil {
		commentsService = drive.NewCommentsService(getDriveService())
	}
	return commentsService
}

func getRepliesService() *drive.RepliesService {
	if repliesService == nil {
		repliesService = drive.NewRepliesService(getDriveService())
	}
	return repliesService
}

func getRevisionsService() *drive.RevisionsService {
	if revisionsService == nil {
		revisionsService = drive.NewRevisionsService(getDriveService())
	}
	return revisionsService
}
