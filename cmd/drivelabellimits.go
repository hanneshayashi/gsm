/*
Copyright © 2020-2023 Hannes Hayashi

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

// driveLabelLimitsCmd represents the driveLabelLimits command
var driveLabelLimitsCmd = &cobra.Command{
	Use:               "driveLabelLimits",
	Short:             "Manages Drive Label Limits (Part of Drive Labels API)",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/labels/reference/rest/v2/limits",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var driveLabelLimitFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"getLabel"},
		Type:         "string",
		Description: `Label revision resource name.
API docs say this must be: "limits/label".
However, only an empty string seems to work currently, so leave empty, if you get an error.`,
	},
	"fields": {
		AvailableFor: []string{"getLabel"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

// var labelLimitFlagsALL = gsmhelpers.GetAllFlags(driveLabelLimitFlags)

func init() {
	rootCmd.AddCommand(driveLabelLimitsCmd)
}
