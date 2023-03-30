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

// Package gsmdrivelabels implements the Drive Labels API
package gsmdrivelabels

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/option"
)

var (
	client             *http.Client
	driveLabelsService *drivelabels.Service
	labelsService      *drivelabels.LabelsService
	labelsLocksService *drivelabels.LabelsLocksService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getDriveService() *drivelabels.Service {
	if client == nil {
		log.Fatalf("gsmdrivelabels.client is not set. Set with gsmdrivelabels.SetClient(client)")
	}
	if driveLabelsService == nil {
		var err error
		driveLabelsService, err = drivelabels.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating drive labels service: %v", err)
		}
	}
	return driveLabelsService
}

func getLabelsService() *drivelabels.LabelsService {
	if labelsService == nil {
		labelsService = drivelabels.NewLabelsService(getDriveService())
	}
	return labelsService
}

func getLabelsLocksService() *drivelabels.LabelsLocksService {
	if labelsLocksService == nil {
		labelsLocksService = drivelabels.NewLabelsLocksService(getDriveService())
	}
	return labelsLocksService
}
