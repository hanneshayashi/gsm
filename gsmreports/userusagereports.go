/*
Package gsmreports implements the Reports API of Admin SDK
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
	"github.com/hanneshayashi/gsm/gsmhelpers"
	reports "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/googleapi"
)

func makeUserUsageReportsGetCallAndAppend(c *reports.UserUsageReportGetCall, usageReports []*reports.UsageReport, errKey string) ([]*reports.UsageReport, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*reports.UsageReports)
	usageReports = append(usageReports, r.UsageReports...)
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		usageReports, err = makeUserUsageReportsGetCallAndAppend(c, usageReports, errKey)
	}
	return usageReports, err
}

// GetUserUsageReport Retrieves a report which is a collection of properties and statistics for a set of users with the account.
// For more information, see the User Usage Report guide. For more information about the user report's parameters, see the Users Usage parameters reference guides.
func GetUserUsageReport(userKey, date, customerID, filters, orgUnitID, parameters, groupIDFilter, fields string) ([]*reports.UsageReport, error) {
	srv := getUserUsageReportsService()
	c := srv.Get(userKey, date)
	if customerID != "" {
		c.CustomerId(customerID)
	}
	if filters != "" {
		c.Filters(filters)
	}
	if orgUnitID != "" {
		c.OrgUnitID(orgUnitID)
	}
	if parameters != "" {
		c.Parameters(parameters)
	}
	if groupIDFilter != "" {
		c.GroupIdFilter(groupIDFilter)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var usageReports []*reports.UsageReport
	usageReports, err := makeUserUsageReportsGetCallAndAppend(c, usageReports, gsmhelpers.FormatErrorKey(userKey, date))
	return usageReports, err
}
