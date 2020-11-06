/*
Package cmd contains the commands available to the end user
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
package cmd

import (
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// colorsCmd represents the colors command
var colorsCmd = &cobra.Command{
	Use:   "colors",
	Short: "Show Calendar and Event color definitions (Part of Calendar API)",
	Long:  `https://developers.google.com/calendar/v3/reference/colors`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var colorFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"fields": {
		AvailableFor: []string{"get"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(colorsCmd)
}
