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

// Package gsmsheets provides functions to utilize the Google Sheets API
package gsmsheets

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	client              *http.Client
	sheetsService       *sheets.Service
	spreadsheetsService *sheets.SpreadsheetsService
	// spreadsheetsSheetsService *sheets.SpreadsheetsSheetsService
)

// SetClient is used to inject a *http.Client into the package
func SetClient(c *http.Client) {
	client = c
}

func getSheetsService() *sheets.Service {
	if client == nil {
		log.Fatalf("gsmsheets.client is not set. Set with gsmsheets.SetClient(client)")
	}
	if sheetsService == nil {
		var err error
		sheetsService, err = sheets.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Error creating sheets service: %v", err)
		}
	}
	return sheetsService
}

func getSpreadsheetsService() *sheets.SpreadsheetsService {
	if spreadsheetsService == nil {
		spreadsheetsService = sheets.NewSpreadsheetsService(getSheetsService())
	}
	return spreadsheetsService
}

// func getSpreadsheetsSheetsService() (spreadsheetssheetsService *sheets.SpreadsheetsSheetsService) {
// 	if spreadsheetssheetsService == nil {
// 		spreadsheetssheetsService = sheets.NewSpreadsheetsSheetsService(getSheetsService())
// 	}
// 	return
// }
