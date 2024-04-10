/*
Copyright © 2020-2024 Hannes Hayashi

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

// Package gsmlicensing implements the Enterprise License Manager API
package gsmlicensing

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/licensing/v1"
	"google.golang.org/api/option"
)

var (
	client                    *http.Client
	licensingService          *licensing.Service
	licenseAssignmentsService *licensing.LicenseAssignmentsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getLicensingService() *licensing.Service {
	if client == nil {
		log.Fatalf("gsmlicensing.client is not set. Set with gsmlicensing.SetClient(client)")
	}
	if licensingService == nil {
		var err error
		licensingService, err = licensing.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating licensing service: %v", err)
		}
	}
	return licensingService
}

func getLicenseAssignmentsService() *licensing.LicenseAssignmentsService {
	if licenseAssignmentsService == nil {
		licenseAssignmentsService = licensing.NewLicenseAssignmentsService(getLicensingService())
	}
	return licenseAssignmentsService
}
