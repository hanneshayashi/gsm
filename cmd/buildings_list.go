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

package cmd

import (
	"log"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	admin "google.golang.org/api/admin/directory/v1"

	"github.com/spf13/cobra"
)

// buildingsListCmd represents the list command
var buildingsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Retrieves a list of buildings for an account.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/admin/directory/reference/rest/v1/resources.buildings/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.ListBuildings(flags["customer"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*admin.Building{}
			for i := range result {
				final = append(final, i)
			}
			err := gsmhelpers.Output(final, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error listing buildings: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(buildingsCmd, buildingsListCmd, buildingFlags)
}
