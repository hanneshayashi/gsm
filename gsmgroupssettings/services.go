/*
Copyright © 2020-2025 Hannes Hayashi

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

// Package gsmgroupssettings implements the Group Settings API
package gsmgroupssettings

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/groupssettings/v1"
	"google.golang.org/api/option"
)

var (
	client                *http.Client
	groupssettingsService *groupssettings.Service
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getGroupssettingsService() *groupssettings.Service {
	if client == nil {
		log.Fatalf("gsmgroupssettings.client is not set. Set with gsmgroupssettings.SetClient(client)")
	}
	if groupssettingsService == nil {
		var err error
		groupssettingsService, err = groupssettings.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating groupssettings service: %v", err)
		}
	}
	return groupssettingsService
}
