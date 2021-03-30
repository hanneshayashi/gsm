/*
Package gsmadmin implements the Admin SDK APIs
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
package gsmadmin

import (
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.BatchCreatePrintersResponse)
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.BatchDeletePrintersResponse)
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(parent), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Printer)
	return r, nil
}

// DeletePrinter deletes a Printer.
func DeletePrinter(name string) (bool, error) {
	srv := getCustomersChromePrintersService()
	c := srv.Delete(name)
	_, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return false, err
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Printer)
	return r, nil
}

func listPrinters(c *admin.CustomersChromePrintersListCall, ch chan *admin.Printer, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*admin.ListPrintersResponse)
	for i := range r.Printers {
		ch <- r.Printers[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listPrinters(c, ch, errKey)
	}
	return err
}

// ListPrinters lists printers configs.
func ListPrinters(parent, filter, fields string, cap int) (<-chan *admin.Printer, <-chan error) {
	srv := getCustomersChromePrintersService()
	c := srv.List(parent).PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c = c.Filter(filter)
	}
	ch := make(chan *admin.Printer, cap)
	err := make(chan error, 1)
	go func() {
		e := listPrinters(c, ch, gsmhelpers.FormatErrorKey(parent, filter))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

func listPrinterModels(c *admin.CustomersChromePrintersListPrinterModelsCall, ch chan *admin.PrinterModel, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*admin.ListPrinterModelsResponse)
	for i := range r.PrinterModels {
		ch <- r.PrinterModels[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listPrinterModels(c, ch, errKey)
	}
	return err
}

// ListPrinterModels lists the supported printer models.
func ListPrinterModels(parent, filter, fields string, cap int) (<-chan *admin.PrinterModel, <-chan error) {
	srv := getCustomersChromePrintersService()
	c := srv.ListPrinterModels(parent).PageSize(1000)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if filter != "" {
		c = c.Filter(filter)
	}
	ch := make(chan *admin.PrinterModel, cap)
	err := make(chan error, 1)
	go func() {
		e := listPrinterModels(c, ch, gsmhelpers.FormatErrorKey(parent, filter))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
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
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Printer)
	return r, nil
}
