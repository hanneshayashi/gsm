/*
Copyright © 2020-2022 Hannes Hayashi

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

package gsmsheets

import (
	"fmt"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/sheets/v4"
)

// BatchUpdateSpreadsheet applies one or more updates to a spreadsheet.
func BatchUpdateSpreadsheet(spreadsheetID, fields string, batchUpdateSpreadsheetRequest *sheets.BatchUpdateSpreadsheetRequest) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	srv := getSpreadsheetsService()
	c := srv.BatchUpdate(spreadsheetID, batchUpdateSpreadsheetRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(spreadsheetID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*sheets.BatchUpdateSpreadsheetResponse)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// CreateSpreadsheet creates a spreadsheet, returning the newly created spreadsheet.
func CreateSpreadsheet(spreadsheet *sheets.Spreadsheet, fields string) (*sheets.Spreadsheet, error) {
	srv := getSpreadsheetsService()
	c := srv.Create(spreadsheet)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(spreadsheet.Properties.Title), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*sheets.Spreadsheet)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// GetSpreadsheet returns the spreadsheet at the given ID.
func GetSpreadsheet(spreadsheetID, fields string, ranges []string, includeGridData bool) (*sheets.Spreadsheet, error) {
	srv := getSpreadsheetsService()
	c := srv.Get(spreadsheetID).IncludeGridData(includeGridData)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if ranges != nil {
		c.Ranges(ranges...)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(spreadsheetID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, ok := result.(*sheets.Spreadsheet)
	if !ok {
		return nil, fmt.Errorf("result unknown")
	}
	return r, nil
}

// GetSpreadsheetByDateFilter returns the spreadsheet at the given ID.
func GetSpreadsheetByDateFilter(spreadsheetID, fields string, getSpreadsheetByDataFilterRequest *sheets.GetSpreadsheetByDataFilterRequest) (*sheets.Spreadsheet, error) {
	srv := getSpreadsheetsService()
	c := srv.GetByDataFilter(spreadsheetID, getSpreadsheetByDataFilterRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}
