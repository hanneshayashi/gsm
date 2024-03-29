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
	"github.com/hanneshayashi/gsm/gsmpeople"
	"google.golang.org/api/people/v1"

	"github.com/spf13/cobra"
)

// peopleListDirectoryPeopleCmd represents the listDirectoryPeople command
var peopleListDirectoryPeopleCmd = &cobra.Command{
	Use:   "listDirectoryPeople",
	Short: "Provides a list of domain profiles and domain contacts in the authenticated user's domain directory.",
	Long: `Implements the API documented at https://developers.google.com/people/api/rest/v1/people/listDirectoryPeople
Example:
 - gsm people listDirectoryPeople --readMask names --sources DIRECTORY_SOURCE_TYPE_DOMAIN_PROFILE`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmpeople.ListDirectoryPeople(flags["readMask"].GetString(), flags["sources"].GetString(), flags["fields"].GetString(), flags["mergeSources"].GetStringSlice(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				err := enc.Encode(i)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			final := []*people.Person{}
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
			log.Fatalf("Error listing people: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(peopleCmd, peopleListDirectoryPeopleCmd, peopleFlags)
}
