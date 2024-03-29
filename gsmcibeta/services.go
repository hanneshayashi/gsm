/*
Copyright © 2020-2023 Hannes Hayashi

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

// Package gsmcibeta implements the Cloud Identity Beta API
package gsmcibeta

import (
	"context"
	"log"
	"net/http"

	cibeta "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/option"
)

var (
	client                     *http.Client
	ciBetaService              *cibeta.Service
	orgUnitsMembershipsService *cibeta.OrgUnitsMembershipsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getCiBetaService() *cibeta.Service {
	if client == nil {
		log.Fatalf("gsmcibeta.client is not set. Set with gsmcibeta.SetClient(client)")
	}
	if ciBetaService == nil {
		var err error
		ciBetaService, err = cibeta.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating ci beta service: %v", err)
		}
	}
	return ciBetaService
}

func getOrgUnitsMembershipsService() *cibeta.OrgUnitsMembershipsService {
	if orgUnitsMembershipsService == nil {
		orgUnitsMembershipsService = cibeta.NewOrgUnitsMembershipsService(getCiBetaService())
	}
	return orgUnitsMembershipsService
}
