/*
Package gsmgmail implements the Gmail APIs
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

package gsmgmail

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var (
	client                                  *http.Client
	gmailService                            *gmail.Service
	usersMessagesService                    *gmail.UsersMessagesService
	usersSettingsDelegatesService           *gmail.UsersSettingsDelegatesService
	usersService                            *gmail.UsersService
	usersDraftsService                      *gmail.UsersDraftsService
	usersHistoryService                     *gmail.UsersHistoryService
	usersLabelsService                      *gmail.UsersLabelsService
	usersMessagesAttachmentsService         *gmail.UsersMessagesAttachmentsService
	usersSettingsService                    *gmail.UsersSettingsService
	usersSettingsFiltersService             *gmail.UsersSettingsFiltersService
	usersSettingsForwardingAddressesService *gmail.UsersSettingsForwardingAddressesService
	usersSettingsSendAsService              *gmail.UsersSettingsSendAsService
	usersSettingsSendAsSmimeInfoService     *gmail.UsersSettingsSendAsSmimeInfoService
	usersThreadsService                     *gmail.UsersThreadsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getGmailService() *gmail.Service {
	if client == nil {
		log.Fatalf("gsmgmail.client is not set. Set with gsmgmail.SetClient(client)")
	}
	if gmailService == nil {
		var err error
		gmailService, err = gmail.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating gmail service: %v", err)
		}
	}
	return gmailService
}

func getUsersService() *gmail.UsersService {
	if usersService == nil {
		usersService = gmail.NewUsersService(getGmailService())
	}
	return usersService
}

func getUsersDraftsService() *gmail.UsersDraftsService {
	if usersDraftsService == nil {
		usersDraftsService = gmail.NewUsersDraftsService(getGmailService())
	}
	return usersDraftsService
}

func getUsersMessagesService() *gmail.UsersMessagesService {
	if usersMessagesService == nil {
		usersMessagesService = gmail.NewUsersMessagesService(getGmailService())
	}
	return usersMessagesService
}

func getUsersSettingsDelegatesService() *gmail.UsersSettingsDelegatesService {
	if usersSettingsDelegatesService == nil {
		usersSettingsDelegatesService = gmail.NewUsersSettingsDelegatesService(getGmailService())
	}
	return usersSettingsDelegatesService
}

func getUsersHistoryService() *gmail.UsersHistoryService {
	if usersHistoryService == nil {
		usersHistoryService = gmail.NewUsersHistoryService(getGmailService())
	}
	return usersHistoryService
}

func getUsersLabelsService() *gmail.UsersLabelsService {
	if usersLabelsService == nil {
		usersLabelsService = gmail.NewUsersLabelsService(getGmailService())
	}
	return usersLabelsService
}

func getUsersMessagesAttachmentsService() *gmail.UsersMessagesAttachmentsService {
	if usersMessagesAttachmentsService == nil {
		usersMessagesAttachmentsService = gmail.NewUsersMessagesAttachmentsService(getGmailService())
	}
	return usersMessagesAttachmentsService
}

func getUsersSettingsService() *gmail.UsersSettingsService {
	if usersSettingsService == nil {
		usersSettingsService = gmail.NewUsersSettingsService(getGmailService())
	}
	return usersSettingsService
}

func getUsersSettingsFiltersService() *gmail.UsersSettingsFiltersService {
	if usersSettingsFiltersService == nil {
		usersSettingsFiltersService = gmail.NewUsersSettingsFiltersService(getGmailService())
	}
	return usersSettingsFiltersService
}

func getUsersSettingsForwardingAddressesService() *gmail.UsersSettingsForwardingAddressesService {
	if usersSettingsForwardingAddressesService == nil {
		usersSettingsForwardingAddressesService = gmail.NewUsersSettingsForwardingAddressesService(getGmailService())
	}
	return usersSettingsForwardingAddressesService
}

func getUsersSettingsSendAsService() *gmail.UsersSettingsSendAsService {
	if usersSettingsSendAsService == nil {
		usersSettingsSendAsService = gmail.NewUsersSettingsSendAsService(getGmailService())
	}
	return usersSettingsSendAsService
}

func getUsersSettingsSendAsSmimeInfoService() *gmail.UsersSettingsSendAsSmimeInfoService {
	if usersSettingsSendAsSmimeInfoService == nil {
		usersSettingsSendAsSmimeInfoService = gmail.NewUsersSettingsSendAsSmimeInfoService(getGmailService())
	}
	return usersSettingsSendAsSmimeInfoService
}

func getUsersThreadsService() *gmail.UsersThreadsService {
	if usersThreadsService == nil {
		usersThreadsService = gmail.NewUsersThreadsService(getGmailService())
	}
	return usersThreadsService
}
