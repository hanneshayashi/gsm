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
package cmd

import (
	"log"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// postmasterTrafficStatsCmd represents the postmasterTrafficStats command
var postmasterTrafficStatsCmd = &cobra.Command{
	Use:   "postmasterTrafficStats",
	Short: "Use Gmail Postmaster Tools to view email traffic statistics (Part of Gmail Postmaster API)",
	Long: `You need to set up your domain(s) at https://postmaster.google.com/u/1/managedomains first.
https://developers.google.com/gmail/postmaster/reference/rest/v1/domains.trafficStats`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var postmasterTrafficStatFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `The resource name of the traffic statistics to get.
E.g., domains/mymail.mydomain.com/trafficStats/20160807.`,
		Required: []string{"get"},
	},
	"parent": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description:  `Fully qualified domain name.`,
		Required:     []string{"list"},
	},
	"startDateDay": {
		AvailableFor: []string{"list"},
		Type:         "int64",
		Description: `The day of the earliest date of the metrics to retrieve inclusive.
If you specify one date flag, you must specify ALL (start and end)!`,
	},
	"startDateMonth": {
		AvailableFor: []string{"list"},
		Type:         "int64",
		Description: `The month of the earliest date of the metrics to retrieve inclusive.
If you specify one date flag, you must specify ALL (start and end)!`,
	},
	"startDateYear": {
		AvailableFor: []string{"list"},
		Type:         "int64",
		Description: `The year of the earliest date of the metrics to retrieve inclusive.
If you specify one date flag, you must specify ALL (start and end)!`,
	},
	"endDateDay": {
		AvailableFor: []string{"list"},
		Type:         "int64",
		Description: `The day of the most recent date of the metrics to retrieve inclusive.
If you specify one date flag, you must specify ALL (start and end)!`,
	},
	"endDateMonth": {
		AvailableFor: []string{"list"},
		Type:         "int64",
		Description: `The month of the most recent date of the metrics to retrieve inclusive.
If you specify one date flag, you must specify ALL (start and end)!`,
	},
	"endDateYear": {
		AvailableFor: []string{"list"},
		Type:         "int64",
		Description: `The year of the most recent date of the metrics to retrieve inclusive.
If you specify one date flag, you must specify ALL (start and end)!`,
	},
	"fields": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var postmasterTrafficStatFlagsALL = gsmhelpers.GetAllFlags(postmasterTrafficStatFlags)

func init() {
	rootCmd.AddCommand(postmasterTrafficStatsCmd)
}
