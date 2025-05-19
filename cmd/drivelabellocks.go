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

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// driveLabelLocksCmd represents the driveLabelLocks command
var driveLabelLocksCmd = &cobra.Command{
	Use:               "driveLabelLocks",
	Short:             "Manages Drive Label Locks (Part of Drive Labels API)",
	Long:              "Implements the API documented at https://developers.google.com/drive/labels/reference/rest/v2/labels.locks",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var driveLabelLockFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"parent": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Label on which Locks are applied. Format: labels/{label}.
If you don't specify the "labels/" prefix, GSM will automatically prepend it to the request.`,
		Required: []string{"list"},
	},
	"fields": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var labelLockFlagsALL = gsmhelpers.GetAllFlags(driveLabelLockFlags)

func init() {
	rootCmd.AddCommand(driveLabelLocksCmd)
}
