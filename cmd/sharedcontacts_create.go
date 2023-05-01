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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// sharedContactsCreateCmd represents the create command
var sharedContactsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Domain Shared Contact",
	Long: `Implements the API documented at https://developers.google.com/admin-sdk/domain-shared-contacts
Example: gsm sharedContacts create --domain "example.org" --givenName "Jane" --familyName "Doe" --email "displayName=Jane Doe;address=jane@doe.net;primary=false" --email "displayName=Jane Doe;address=jane.doe@somedomain.net;primary=true" --phoneNumber "phoneNumber=+49 127 12381;primary=true;label=Work" --phoneNumber "phoneNumber=+49 21891238;primary=false;label=Home" --organization "orgName=Some Company;orgDepartment=Some Department;orgTitle=Some Title"`,
	Annotations: map[string]string{
		"crescendoFlags": "--json",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		s, err := mapToSharedContact(flags, nil)
		if err != nil {
			log.Fatalf("Error building shared contact object: %v\n", err)
		}
		result, err := gsmadmin.CreateSharedContact(flags["domain"].GetString(), s)
		if err != nil {
			log.Fatalf("Error creating shared contact: %v", err)
		}
		if flags["json"].GetBool() {
			err = gsmhelpers.Output(result, "json", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			err = gsmhelpers.Output(result, "xml", compressOutput)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	gsmhelpers.InitCommand(sharedContactsCmd, sharedContactsCreateCmd, sharedContactFlags)
}
