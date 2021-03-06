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

package gsmreports

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"
	reports "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/googleapi"
)

func getCustomerUsageReport(c *reports.CustomerUsageReportsGetCall, ch chan *reports.UsageReport, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*reports.UsageReports)
	for i := range r.UsageReports {
		ch <- r.UsageReports[i]
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = getCustomerUsageReport(c, ch, errKey)
	}
	return err
}

// GetCustomerUsageReport retrieves a report which is a collection of properties and statistics for a specific customer's account.
// For more information, see the Customers Usage Report guide. For more information about the customer report's parameters, see the Customers Usage parameters reference guides.
func GetCustomerUsageReport(date, customerID, parameters, fields string, cap int) (<-chan *reports.UsageReport, <-chan error) {
	srv := getCustomerUsageReportsService()
	c := srv.Get(date)
	if customerID != "" {
		c.CustomerId(customerID)
	}
	if parameters != "" {
		c.Parameters(parameters)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *reports.UsageReport, cap)
	err := make(chan error, 1)
	go func() {
		e := getCustomerUsageReport(c, ch, gsmhelpers.FormatErrorKey(date, customerID, parameters))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}
