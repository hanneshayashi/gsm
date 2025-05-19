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

package cmd

import (
	"log"
	"math/rand"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmsheets"

	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

// spreadsheetsCmd represents the spreadsheets command
var spreadsheetsCmd = &cobra.Command{
	Use:               "spreadsheets",
	Short:             "Manage Google Sheets spreadsheets (Part of Sheets API)",
	Long:              `Implements the API documented at https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var spreadsheetFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"spreadsheetId": {
		AvailableFor: []string{"batchUpdate", "get", "getByDateFilter"},
		Type:         "string",
		Description:  "The ID of the spreadsheet",
		Required:     []string{"batchUpdate", "get", "getByDateFilter"},
	},
	"title": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  "The ID of the spreadsheet",
		Required:     []string{"create"},
	},
	"csvFileToUpload": {
		AvailableFor: []string{"batchUpdate", "create"},
		Type:         "stringSlice",
		Description: `A list of CSV files that should be added to the spreadsheet as new sheets.
Can be used multiple times in the form of "--csvFileToUpload "title=Some Title;path=./path/to/file.csv""
Delimiter must be ","`,
	},
	"ranges": {
		AvailableFor: []string{"get"},
		Type:         "stringSlice",
		Description:  `The ranges to retrieve from the spreadsheet.`,
	},
	"includeGridData": {
		AvailableFor: []string{"get"},
		Type:         "bool",
		Description: `True if grid data should be returned.
This parameter is ignored if a field mask was set in the request.`,
	},
	"fields": {
		AvailableFor: []string{"batchUpdate", "create", "get", "getByDateFilter"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(spreadsheetsCmd)
}

func mapToSpreadsheet(flags map[string]*gsmhelpers.Value) (*sheets.Spreadsheet, error) {
	spreadsheet := &sheets.Spreadsheet{}
	if flags["title"].IsSet() {
		spreadsheet.Properties = &sheets.SpreadsheetProperties{}
		spreadsheet.Properties.Title = flags["title"].GetString()
		if spreadsheet.Properties.Title == "" {
			spreadsheet.Properties.ForceSendFields = append(spreadsheet.Properties.ForceSendFields, "Title")
		}
	}
	if flags["csvFileToUpload"].IsSet() {
		csvFilesToUpload := flags["csvFileToUpload"].GetStringSlice()
		if len(csvFilesToUpload) > 0 {
			spreadsheet.Sheets = []*sheets.Sheet{}
			for i := range csvFilesToUpload {
				s := &sheets.Sheet{}
				m := gsmhelpers.FlagToMap(csvFilesToUpload[i])
				s.Properties = &sheets.SheetProperties{
					Title: m["title"],
				}
				s.Data = []*sheets.GridData{}
				gd := &sheets.GridData{
					RowData: []*sheets.RowData{},
				}
				csv, err := gsmhelpers.GetCSVContent(m["path"], ',', false)
				if err != nil {
					return nil, err
				}
				for i := range csv {
					rd := &sheets.RowData{}
					rd.Values = []*sheets.CellData{}
					for j := range csv[i] {
						cd := &sheets.CellData{}
						cd.UserEnteredValue = &sheets.ExtendedValue{
							StringValue: &csv[i][j],
						}
						rd.Values = append(rd.Values, cd)
					}
					gd.RowData = append(gd.RowData, rd)
				}
				s.Data = append(s.Data, gd)
				spreadsheet.Sheets = append(spreadsheet.Sheets, s)
			}
		} else {
			spreadsheet.ForceSendFields = append(spreadsheet.ForceSendFields, "Sheets")
		}
	}
	return spreadsheet, nil
}

func mapToBatchUpdateSpreadsheetRequest(flags map[string]*gsmhelpers.Value) (*sheets.BatchUpdateSpreadsheetRequest, error) {
	batchUpdateSpreadsheetRequest := &sheets.BatchUpdateSpreadsheetRequest{}
	spreadsheet, err := gsmsheets.GetSpreadsheet(flags["spreadsheetId"].GetString(), "sheets(properties(title,sheetId))", nil, false)
	if err != nil {
		return nil, err
	}
	sheetMap := make(map[string]int64)
	for i := range spreadsheet.Sheets {
		sheetMap[spreadsheet.Sheets[i].Properties.Title] = spreadsheet.Sheets[i].Properties.SheetId
	}
	if flags["csvFileToUpload"].IsSet() {
		csvFilesToUpload := flags["csvFileToUpload"].GetStringSlice()
		if len(csvFilesToUpload) > 0 {
			batchUpdateSpreadsheetRequest.Requests = []*sheets.Request{}
			for i := range csvFilesToUpload {
				m := gsmhelpers.FlagToMap(csvFilesToUpload[i])
				data, err := gsmhelpers.GetFileContentAsString(m["path"])
				if err != nil {
					return nil, err
				}
				sheetID, ok := sheetMap[m["title"]]
				if !ok {
					r := &sheets.Request{}
					sheetID = int64(rand.Int31())
					r.AddSheet = &sheets.AddSheetRequest{}
					r.AddSheet.Properties = &sheets.SheetProperties{
						Title:   m["title"],
						SheetId: sheetID,
					}
					batchUpdateSpreadsheetRequest.Requests = append(batchUpdateSpreadsheetRequest.Requests, r)
				}
				r := &sheets.Request{}
				r.PasteData = &sheets.PasteDataRequest{
					Coordinate: &sheets.GridCoordinate{
						SheetId: sheetID,
					},
					Data:      data,
					Delimiter: ",",
					Type:      "PASTE_NORMAL",
				}
				batchUpdateSpreadsheetRequest.Requests = append(batchUpdateSpreadsheetRequest.Requests, r)
			}
		}
	}
	return batchUpdateSpreadsheetRequest, nil
}
