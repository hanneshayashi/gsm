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

// otherContactsCopyOtherContactToMyContactsGroupCmd represents the copyOtherContactToMyContactsGroup command
var otherContactsCopyOtherContactToMyContactsGroupCmd = &cobra.Command{
	Use:   "copyOtherContactToMyContactsGroup",
	Short: `Copies an "Other contact" to a new contact in the user's "myContacts" group.`,
	Long:  "https://developers.google.com/people/api/rest/v1/otherContacts/copyOtherContactToMyContactsGroup",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToCopyOtherContactToMyContactsGroupRequest(flags)
		if err != nil {
			log.Fatalf("Error building CopyOtherContactToMyContactsGroupRequest object: %v", err)
		}
		result, err := gsmpeople.CopyOtherContactToMyContactsGroup(flags["resourceName"].GetString(), flags["fields"].GetString(), flags["sources"].GetStringSlice(), c)
		if err != nil {
			log.Fatalf("Error copying other contact %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitCommand(otherContactsCmd, otherContactsCopyOtherContactToMyContactsGroupCmd, otherContactFlags)
}
