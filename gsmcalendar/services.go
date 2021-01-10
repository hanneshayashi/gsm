/*
Package gsmcalendar implements the Calendar API
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
package gsmcalendar

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var (
	client              *http.Client
	calendarService     *calendar.Service
	calendarListService *calendar.CalendarListService
	eventsService       *calendar.EventsService
	aclService          *calendar.AclService
	calendarsService    *calendar.CalendarsService
	colorsService       *calendar.ColorsService
	freebusyService     *calendar.FreebusyService
	settingsService     *calendar.SettingsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getCalendarService() *calendar.Service {
	if client == nil {
		log.Fatalf("gsmcalendar.client is not set. Set with gsmcalendar.SetClient(client)")
	}
	if calendarService == nil {
		var err error
		calendarService, err = calendar.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating calendar service: %v", err)
		}
	}
	return calendarService
}

func getCalendarListService() *calendar.CalendarListService {
	if calendarListService == nil {
		calendarListService = calendar.NewCalendarListService(getCalendarService())
	}
	return calendarListService
}

func getEventsService() *calendar.EventsService {
	if eventsService == nil {
		eventsService = calendar.NewEventsService(getCalendarService())
	}
	return eventsService
}

func getACLService() *calendar.AclService {
	if aclService == nil {
		aclService = calendar.NewAclService(getCalendarService())
	}
	return aclService
}

func getCalendarsService() *calendar.CalendarsService {
	if calendarsService == nil {
		calendarsService = calendar.NewCalendarsService(getCalendarService())
	}
	return calendarsService
}

func getColorsService() *calendar.ColorsService {
	if colorsService == nil {
		colorsService = calendar.NewColorsService(getCalendarService())
	}
	return colorsService
}

func getFreebusyService() *calendar.FreebusyService {
	if freebusyService == nil {
		freebusyService = calendar.NewFreebusyService(getCalendarService())
	}
	return freebusyService
}

func getSettingsService() *calendar.SettingsService {
	if settingsService == nil {
		settingsService = calendar.NewSettingsService(getCalendarService())
	}
	return settingsService
}
