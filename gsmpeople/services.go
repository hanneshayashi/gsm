/*
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

// Package gsmpeople implements the People API
package gsmpeople

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

var (
	client                      *http.Client
	peopleService               *people.Service
	pService                    *people.PeopleService
	contactGroupsService        *people.ContactGroupsService
	contactGroupsMembersService *people.ContactGroupsMembersService
	otherContactsService        *people.OtherContactsService
	peopleConnectionsService    *people.PeopleConnectionsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getPeopleService() *people.Service {
	if client == nil {
		log.Fatalf("gsmpeople.client is not set. Set with gsmpeople.SetClient(client)")
	}
	if peopleService == nil {
		var err error
		peopleService, err = people.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating people service: %v", err)
		}
	}
	return peopleService
}

func getpService() *people.PeopleService {
	if pService == nil {
		pService = people.NewPeopleService(getPeopleService())
	}
	return pService
}

func getContactGroupsService() *people.ContactGroupsService {
	if contactGroupsService == nil {
		contactGroupsService = people.NewContactGroupsService(getPeopleService())
	}
	return contactGroupsService
}

func getContactGroupsMembersService() *people.ContactGroupsMembersService {
	if contactGroupsMembersService == nil {
		contactGroupsMembersService = people.NewContactGroupsMembersService(getPeopleService())
	}
	return contactGroupsMembersService
}

func getOtherContactsService() *people.OtherContactsService {
	if otherContactsService == nil {
		otherContactsService = people.NewOtherContactsService(getPeopleService())
	}
	return otherContactsService
}

func getPeopleConnectionsService() *people.PeopleConnectionsService {
	if peopleConnectionsService == nil {
		peopleConnectionsService = people.NewPeopleConnectionsService(getPeopleService())
	}
	return peopleConnectionsService
}
