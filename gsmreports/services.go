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

// Package gsmreports implements the Reports API of Admin SDK
package gsmreports

import (
	"context"
	"log"
	"net/http"

	reports "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/option"
)

var (
	client                      *http.Client
	reportsService              *reports.Service
	activitiesService           *reports.ActivitiesService
	customerUsageReportsService *reports.CustomerUsageReportsService
	entityUsageReportsService   *reports.EntityUsageReportsService
	userUsageReportService      *reports.UserUsageReportService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getReportsService() *reports.Service {
	if client == nil {
		log.Fatalf("gsmreports.client is not set. Set with gsmreports.SetClient(client)")
	}
	if reportsService == nil {
		var err error
		reportsService, err = reports.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating reports service: %v", err)
		}
	}
	return reportsService
}

func getActivitiesService() *reports.ActivitiesService {
	if activitiesService == nil {
		activitiesService = reports.NewActivitiesService(getReportsService())
	}
	return activitiesService
}

func getCustomerUsageReportsService() *reports.CustomerUsageReportsService {
	if customerUsageReportsService == nil {
		customerUsageReportsService = reports.NewCustomerUsageReportsService(getReportsService())
	}
	return customerUsageReportsService
}

func getEntityUsageReportsService() *reports.EntityUsageReportsService {
	if entityUsageReportsService == nil {
		entityUsageReportsService = reports.NewEntityUsageReportsService(getReportsService())
	}
	return entityUsageReportsService
}

func getUserUsageReportsService() *reports.UserUsageReportService {
	if userUsageReportService == nil {
		userUsageReportService = reports.NewUserUsageReportService(getReportsService())
	}
	return userUsageReportService
}
