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
	"fmt"
	"gsm/gsmhelpers"
	"gsm/gsmpeople"
	"log"

	"github.com/spf13/cobra"
)

// peopleCreateContactCmd represents the createContact command
var peopleCreateContactCmd = &cobra.Command{
	Use:   "createContact",
	Short: "Create a new contact and return the person resource for that contact.",
	Long:  "https://developers.google.com/people/api/rest/v1/people/createContact",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		p, err := mapToPerson(flags, nil)
		if err != nil {
			log.Fatalf("Error building person object: %v", err)
		}
		result, err := gsmpeople.CreateContact(p, flags["personFields"].GetString(), flags["sources"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error creating contact %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	gsmhelpers.InitCommand(peopleCmd, peopleCreateContactCmd, peopleFlags)
}
