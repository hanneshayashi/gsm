/*
Copyright © 2020-2022 Hannes Hayashi

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

	"github.com/spf13/cobra"
)

// peopleUpdateContactCmd represents the updateContact command
var peopleUpdateContactCmd = &cobra.Command{
	Use:               "updateContact",
	Short:             "Update contact data for an existing contact person.",
	Long:              "Implements the API documented at https://developers.google.com/people/api/rest/v1/people/updateContact",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := gsmpeople.GetContact(flags["resourceName"].GetString(), flags["personFields"].GetString(), flags["sources"].GetString(), "*")
		if err != nil {
			log.Fatalf("Error getting contact: %v", err)
		}
		p, err = mapToPerson(flags, p)
		if err != nil {
			log.Fatalf("Error building person object: %v", err)
		}
		result, err := gsmpeople.UpdateContact(flags["resourceName"].GetString(), flags["updatePersonFields"].GetString(), flags["personFields"].GetString(), flags["sources"].GetString(), flags["fields"].GetString(), p)
		if err != nil {
			log.Fatalf("Error updating contact: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(peopleCmd, peopleUpdateContactCmd, peopleFlags)
}
