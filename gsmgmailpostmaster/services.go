/*
Package gsmgmailpostmaster implements the Gmail Postmaster APIs
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
package gsmgmailpostmaster

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/gmailpostmastertools/v1"
	"google.golang.org/api/option"
)

var (
	client                     *http.Client
	gmailPostmasterService     *gmailpostmastertools.Service
	domainsService             *gmailpostmastertools.DomainsService
	domainsTrafficStatsService *gmailpostmastertools.DomainsTrafficStatsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getGmailPostmasterService() *gmailpostmastertools.Service {
	if client == nil {
		log.Fatalf("gsmgmailpostmastertools.client is not set. Set with gsmgmailpostmastertools.SetClient(client)")
	}
	if gmailPostmasterService == nil {
		var err error
		gmailPostmasterService, err = gmailpostmastertools.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating gmail postmaster service: %v", err)
		}
	}
	return gmailPostmasterService
}

func getDomainsService() *gmailpostmastertools.DomainsService {
	if domainsService == nil {
		domainsService = gmailpostmastertools.NewDomainsService(getGmailPostmasterService())
	}
	return domainsService
}

func getDomainsTrafficStatsService() *gmailpostmastertools.DomainsTrafficStatsService {
	if domainsTrafficStatsService == nil {
		domainsTrafficStatsService = gmailpostmastertools.NewDomainsTrafficStatsService(getGmailPostmasterService())
	}
	return domainsTrafficStatsService
}
