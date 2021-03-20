/*
Package cmd contains the commands available to the end user
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
package cmd

import (
	"log"
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	admin "google.golang.org/api/admin/directory/v1"

	"github.com/spf13/cobra"
)

// chromePrintersCmd represents the chromePrinters command
var chromePrintersCmd = &cobra.Command{
	Use:               "chromePrinters",
	Short:             "Managed Chrome Printers (Part of Admin SDK)",
	Long:              "https://developers.google.com/admin-sdk/chrome-printer/reference/rest/v1/admin.directory.v1.customers.chrome.printers",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

var chromePrinterFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"batchCreate", "batchDelete", "create", "list", "listModels"},
		Type:         "string",
		Description:  `The name of the customer. Format: customers/{customer_id}`,
	},
	"printer": {
		AvailableFor: []string{"batchCreate"},
		Type:         "stringSlice",
		Description: `A printer to create.
If you want to place the printer under particular OU then populate orgUnitId filed.
Otherwise the printer will be placed under root OU.
Can be used multiple times in the form of "--printer "name=...;displayName=...;description=...", etc.
You can use the following properties:
name                 The resource name of the Printer object, in the format customers/{customer-id}/printers/{printer-id} (During printer creation leave empty)
description          Description of printer.
makeAndModel         Editable. Make and model of printer. e.g. Lexmark MS610de.
                     Value must be in format as seen in printers.listPrinterModels response.
uri                  Editable. Printer URI.
orgUnitId            Organization Unit that owns this printer (Only can be set during Printer creation)
useDriverlessConfig  Editable. flag to use driverless configuration or not.
                     If it's set to be true, makeAndModel can be ignored`,
		Required: []string{"batchCreate"},
	},
	"printerIds": {
		AvailableFor: []string{"batchDelete"},
		Type:         "stringSlice",
		Description:  `A list of printer ids that should be deleted. Max 100 at a time.`,
		Required:     []string{"batchDelete"},
	},
	"displayName": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  `Editable. Name of printer.`,
		Required:     []string{"create"},
	},
	"description": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  `Editable. Description of printer.`,
	},
	"makeAndModel": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description: `Editable. Make and model of printer. e.g. Lexmark MS610de
Value must be in format as seen in printers.listPrinterModels response.`,
	},
	"uri": {
		AvailableFor: []string{"create", "patch"},
		Type:         "string",
		Description:  `Editable. Printer URI.`,
		Required:     []string{"create"},
	},
	"orgUnitId": {
		AvailableFor: []string{"create"},
		Type:         "string",
		Description:  `Organization Unit`,
		Required:     []string{"create"},
	},
	"useDriverlessConfig": {
		AvailableFor: []string{"create", "patch"},
		Type:         "bool",
		Description:  `Editable. flag to use driverless configuration or not. If it's set to be true, makeAndModel can be ignored`,
	},
	"name": {
		AvailableFor: []string{"delete", "get", "patch"},
		Type:         "string",
		Description:  `The name of the printer to be updated. Format: customers/{customer_id}/chrome/printers/{printer_id}`,
		Required:     []string{"delete", "get", "patch"},
	},
	"filter": {
		AvailableFor: []string{"list", "listModels"},
		Type:         "string",
		Description:  `Search query. Search syntax is shared between this api and Admin Console printers pages.`,
	},
	"updateMask": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The list of fields to be updated. Note, some of the fields are read only and cannot be updated. Values for not specified fields will be patched.

A comma-separated list of fully qualified names of fields. Example: "user.displayName,photo".`,
	},
	"clearMask": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The list of fields to be cleared. Note, some of the fields are read only and cannot be updated. Values for not specified fields will be patched.

A comma-separated list of fully qualified names of fields. Example: "user.displayName,photo".`,
	},
	"fields": {
		AvailableFor: []string{"batchCreate", "create", "get", "list", "listModels", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var chromePrinterFlagsALL = gsmhelpers.GetAllFlags(chromePrinterFlags)

func init() {
	rootCmd.AddCommand(chromePrintersCmd)
}

func mapToBatchCreatePrintersRequest(flags map[string]*gsmhelpers.Value) (*admin.BatchCreatePrintersRequest, error) {
	batchCreatePrintersRequest := &admin.BatchCreatePrintersRequest{}
	if flags["printer"].IsSet() {
		printers := flags["printer"].GetStringSlice()
		if len(printers) > 0 {
			batchCreatePrintersRequest.Requests = []*admin.CreatePrinterRequest{}
			for i := range printers {
				m := gsmhelpers.FlagToMap(printers[i])
				createPrinterRequest := &admin.CreatePrinterRequest{
					Printer: &admin.Printer{
						Description:  m["description"],
						DisplayName:  m["displayName"],
						MakeAndModel: m["makeAndModel"],
						OrgUnitId:    m["orgUnitId"],
						Uri:          m["uri"],
					},
				}
				useDriverlessConfig, err := strconv.ParseBool(m["useDriverlessConfig"])
				if err != nil {
					log.Printf("Error parsing %v to bool: %v. Setting to false.", m["useDriverlessConfig"], err)
				}
				createPrinterRequest.Printer.UseDriverlessConfig = useDriverlessConfig
				batchCreatePrintersRequest.Requests = append(batchCreatePrintersRequest.Requests, createPrinterRequest)
			}
		} else {
			batchCreatePrintersRequest.ForceSendFields = append(batchCreatePrintersRequest.ForceSendFields, "Requests")
		}
	}
	return batchCreatePrintersRequest, nil
}

func mapToBatchDeletePrintersRequest(flags map[string]*gsmhelpers.Value) (*admin.BatchDeletePrintersRequest, error) {
	batchDeletePrintersRequest := &admin.BatchDeletePrintersRequest{}
	if flags["printerIds"].IsSet() {
		batchDeletePrintersRequest.PrinterIds = flags["printerIds"].GetStringSlice()
		if len(batchDeletePrintersRequest.PrinterIds) == 0 {
			batchDeletePrintersRequest.ForceSendFields = append(batchDeletePrintersRequest.ForceSendFields, "PrinterIds")
		}
	}
	return batchDeletePrintersRequest, nil
}

func mapToChromePrinter(flags map[string]*gsmhelpers.Value) (*admin.Printer, error) {
	printer := &admin.Printer{}
	if flags["displayName"].IsSet() {
		printer.DisplayName = flags["displayName"].GetString()
		if printer.DisplayName == "" {
			printer.ForceSendFields = append(printer.ForceSendFields, "DisplayName")
		}
	}
	if flags["description"].IsSet() {
		printer.Description = flags["description"].GetString()
		if printer.Description == "" {
			printer.ForceSendFields = append(printer.ForceSendFields, "Description")
		}
	}
	if flags["makeAndModel"].IsSet() {
		printer.MakeAndModel = flags["makeAndModel"].GetString()
		if printer.MakeAndModel == "" {
			printer.ForceSendFields = append(printer.ForceSendFields, "MakeAndModel")
		}
	}
	if flags["uri"].IsSet() {
		printer.Uri = flags["uri"].GetString()
		if printer.Uri == "" {
			printer.ForceSendFields = append(printer.ForceSendFields, "Uri")
		}
	}
	if flags["orgUnitId"].IsSet() {
		printer.OrgUnitId = flags["orgUnitId"].GetString()
		if printer.OrgUnitId == "" {
			printer.ForceSendFields = append(printer.ForceSendFields, "OrgUnitId")
		}
	}
	if flags["useDriverlessConfig"].IsSet() {
		printer.UseDriverlessConfig = flags["useDriverlessConfig"].GetBool()
		if !printer.UseDriverlessConfig {
			printer.ForceSendFields = append(printer.ForceSendFields, "UseDriverlessConfig")
		}
	}
	return printer, nil
}
