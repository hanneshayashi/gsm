/*
Package gsmgmailpostmaster implements the Gmail Postmaster APIs
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
package gsmgmailpostmaster

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmailpostmastertools/v1"
	"google.golang.org/api/googleapi"
)

// GetTrafficStats Get traffic statistics for a domain on a specific date.
// Returns PERMISSION_DENIED if user does not have permission to access TrafficStats for the domain.
func GetTrafficStats(name, fields string) (*gmailpostmastertools.TrafficStats, error) {
	srv := getDomainsTrafficStatsService()
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
	r, _ := result.(*gmailpostmastertools.TrafficStats)
	return r, nil
}

func listTrafficStats(c *gmailpostmastertools.DomainsTrafficStatsListCall, ch chan *gmailpostmastertools.TrafficStats, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*gmailpostmastertools.ListTrafficStatsResponse)
	for _, i := range r.TrafficStats {
		ch <- i
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		err = listTrafficStats(c, ch, errKey)
	}
	return err
}

// ListTrafficStats List traffic statistics for all available days.
// Returns PERMISSION_DENIED if user does not have permission to access TrafficStats for the domain.
func ListTrafficStats(parent, fields string, startDateDay, startDateMonth, startDateYear, endDateDay, endDateMonth, endDateYear int64, cap int) (<-chan *gmailpostmastertools.TrafficStats, <-chan error) {
	srv := getDomainsTrafficStatsService()
	c := srv.List(parent)
	if startDateDay != 0 {
		c.StartDateDay(startDateDay)
	}
	if startDateMonth != 0 {
		c.StartDateMonth(startDateMonth)
	}
	if startDateYear != 0 {
		c.StartDateYear(startDateYear)
	}
	if endDateDay != 0 {
		c.EndDateDay(endDateDay)
	}
	if endDateMonth != 0 {
		c.EndDateMonth(endDateMonth)
	}
	if endDateYear != 0 {
		c.EndDateYear(endDateYear)
	}
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *gmailpostmastertools.TrafficStats, cap)
	err := make(chan error, 1)
	go func() {
		e := listTrafficStats(c, ch, gsmhelpers.FormatErrorKey(parent))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}
