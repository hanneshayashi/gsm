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

package gsmreports

import (
	"context"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	reports "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/googleapi"
)

// GetEntityUsageReport retrieves a report which is a collection of properties and statistics for entities used by users within the account.
// For more information, see the Entities Usage Report guide.
// For more information about the entities report's parameters, see the Entities Usage parameters reference guides.
func GetEntityUsageReport(entityType, entityKey, date, customerID, filters, parameters, fields string, cap int) (<-chan *reports.UsageReport, <-chan error) {
	srv := getEntityUsageReportsService()
	c := srv.Get(entityType, entityKey, date)
	if customerID != "" {
		c.CustomerId(customerID)
	}
	if filters != "" {
		c.Filters(filters)
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
		e := c.Pages(context.Background(), func(response *reports.UsageReports) error {
			for i := range response.UsageReports {
				ch <- response.UsageReports[i]
			}
			return nil
		})
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}
