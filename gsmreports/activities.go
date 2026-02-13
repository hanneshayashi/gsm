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
	"errors"
	"context"
	"iter"

	reports "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/googleapi"
)

// ListActivities retrieves a list of activities for a specific customer's account and application such as the Admin console application or the Google Drive application.
// For more information, see the guides for administrator and Google Drive activity reports.
// For more information about the activity report's parameters, see the activity parameters reference guides.
func ListActivities(userKey, applicationName, actorIPAddress, customerID, endTime, eventName, filters, groupIDFilter, orgUnitID, startTime, fields string) iter.Seq2[*reports.Activity, error] {
	srv := getActivitiesService()
	c := srv.List(userKey, applicationName).MaxResults(1000)
	if actorIPAddress != "" {
		c.ActorIpAddress(actorIPAddress)
	}
	if customerID != "" {
		c.CustomerId(customerID)
	}
	if endTime != "" {
		c.EndTime(endTime)
	}
	if eventName != "" {
		c.EventName(eventName)
	}
	if filters != "" {
		c.Filters(filters)
	}
	if groupIDFilter != "" {
		c.GroupIdFilter(groupIDFilter)
	}
	if orgUnitID != "" {
		c.OrgUnitID(orgUnitID)
	}
	if startTime != "" {
		c.StartTime(startTime)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	return func(yield func(*reports.Activity, error) bool) {
		e := c.Pages(context.Background(), func(response *reports.Activities) error {
			for i := range response.Items {
				if !yield(response.Items[i], nil) {
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
