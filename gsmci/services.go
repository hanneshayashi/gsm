/*
Package gsmci implements the Cloud Identity (Beta) API
Copyright © 2020-2021 Hannes Hayashi

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
package gsmci

import (
	"context"
	"log"
	"net/http"

	ci "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/option"
)

var (
	client                   *http.Client
	ciService                *ci.Service
	groupsService            *ci.GroupsService
	groupsMembershipsService *ci.GroupsMembershipsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getCiService() *ci.Service {
	if client == nil {
		log.Fatalf("gsmci.client is not set. Set with gsmci.SetClient(client)")
	}
	if ciService == nil {
		var err error
		ciService, err = ci.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating ci service: %v", err)
		}
	}
	return ciService
}

func getGroupsService() *ci.GroupsService {
	if groupsService == nil {
		groupsService = ci.NewGroupsService(getCiService())
	}
	return groupsService
}

func getGroupsMembershipsService() *ci.GroupsMembershipsService {
	if groupsMembershipsService == nil {
		groupsMembershipsService = ci.NewGroupsMembershipsService(getCiService())
	}
	return groupsMembershipsService
}
