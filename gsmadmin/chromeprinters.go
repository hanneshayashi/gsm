/*
Copyright © 2020 Hannes Hayashi

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

package gsmadmin

import (
	"errors"
	"context"
	"fmt"
	"iter"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// PrinterResults represents the result of a batch operation
type PrinterResults struct {
	Failures   []*admin.FailureInfo `json:"failures,omitempty"`
	Printers   []*admin.Printer     `json:"printers,omitempty"`
	PrinterIDs []string             `json:"printerIDs,omitempty"`
}

// BatchCreatePrinters creates printers under given Organization Unit.
func BatchCreatePrinters(parent, fields string, batchCreatePrintersRequest *admin.BatchCreatePrintersRequest) (*PrinterResults, error) {
	srv := getCustomersChromePrintersService()
	c := srv.BatchCreatePrinters(parent, batchCreatePrintersRequest)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent), err)
	}
	response := &PrinterResults{
		Failures: r.Failures,
		Printers: r.Printers,
	}
	return response, nil
}

// BatchDeletePrinters deletes printers in batch.
func BatchDeletePrinters(parent string, batchDeletePrintersRequest *admin.BatchDeletePrintersRequest) (*PrinterResults, error) {
	srv := getCustomersChromePrintersService()
	c := srv.BatchDeletePrinters(parent, batchDeletePrintersRequest)
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent), err)
	}
	response := &PrinterResults{
		Failures:   r.FailedPrinters,
		PrinterIDs: r.PrinterIds,
	}
	return response, nil
}

// CreatePrinter creates a printer under given Organization Unit.
func CreatePrinter(parent, fields string, printer *admin.Printer) (*admin.Printer, error) {
	srv := getCustomersChromePrintersService()
	c := srv.Create(parent, printer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(parent), err)
	}
	return r, nil
}

// DeletePrinter deletes a Printer.
func DeletePrinter(name string) (bool, error) {
	srv := getCustomersChromePrintersService()
	c := srv.Delete(name)
	_, err := c.Do()
	if err != nil {
		return false, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return true, nil
}

// GetPrinter returns a Printer resource (printer's config).
func GetPrinter(name, fields string) (*admin.Printer, error) {
	srv := getCustomersChromePrintersService()
	c := srv.Get(name)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r, nil
}

// ListPrinters lists printers configs.
func ListPrinters(parent, filter, fields string) iter.Seq2[*admin.Printer, error] {
	srv := getCustomersChromePrintersService()
	c := srv.List(parent).PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c = c.Filter(filter)
	}
	return func(yield func(*admin.Printer, error) bool) {
		e := c.Pages(context.Background(), func(response *admin.ListPrintersResponse) error {
			for i := range response.Printers {
				if !yield(response.Printers[i], nil) {
					return errIterStopped
				}
			}
			return nil
		})
		if e != nil && !errors.Is(e, errIterStopped) {
			yield(nil, e)
		}
	}
}

// ListPrinterModels lists the supported printer models.
func ListPrinterModels(parent, filter, fields string) iter.Seq2[*admin.PrinterModel, error] {
	srv := getCustomersChromePrintersService()
	c := srv.ListPrinterModels(parent).PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c = c.Filter(filter)
	}
	return func(yield func(*admin.PrinterModel, error) bool) {
		e := c.Pages(context.Background(), func(response *admin.ListPrinterModelsResponse) error {
			for i := range response.PrinterModels {
				if !yield(response.PrinterModels[i], nil) {
					return errIterStopped
				}
			}
			return nil
		})
		if e != nil && !errors.Is(e, errIterStopped) {
			yield(nil, e)
		}
	}
}

// PatchPrinter updates a Printer resource.
func PatchPrinter(name, updateMask, clearMask, fields string, printer *admin.Printer) (*admin.Printer, error) {
	srv := getCustomersChromePrintersService()
	c := srv.Patch(name, printer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if updateMask != "" {
		c.UpdateMask(updateMask)
	}
	if clearMask != "" {
		c.ClearMask(clearMask)
	}
	r, err := c.Do()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", gsmhelpers.FormatErrorKey(name), err)
	}
	return r, nil
}
