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
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmpeople"

	"github.com/spf13/cobra"
)

// peopleSearchDirectoryPeopleCmd represents the searchDirectoryPeople command
var peopleSearchDirectoryPeopleCmd = &cobra.Command{
	Use:               "searchDirectoryPeople",
	Short:             "Provides a list of domain profiles and domain contacts in the authenticated user's domain directory that match the search query.",
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/people/searchDirectoryPeople",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		gsmhelpers.StreamOrCollect(gsmpeople.SearchDirectoryPeople(flags["readMask"].GetString(), flags["sources"].GetString(), flags["query"].GetString(), flags["fields"].GetString(), flags["mergeSources"].GetStringSlice()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(peopleCmd, peopleSearchDirectoryPeopleCmd, peopleFlags)
}
