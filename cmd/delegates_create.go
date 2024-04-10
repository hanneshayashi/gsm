/*
Copyright © 2020-2024 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmgmail"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// delegatesCreateCmd represents the create command
var delegatesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Adds a delegate with its verification status set directly to accepted, without sending any verification email.",
	Long: `The delegate user must be a member of the same Workspace organization as the delegator user.

Gmail imposes limitations on the number of delegates and delegators each user in a Workspace organization can have. These limits depend on your organization, but in general each user can have up to 25 delegates and up to 10 delegators.

Note that a delegate user must be referred to by their primary email address, and not an email alias.

Also note that when a new delegate is created, there may be up to a one minute delay before the new delegate is available for use.

Implements the API documented at https://developers.google.com/gmail/api/reference/rest/v1/users.settings.delegates/create`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		d, err := mapToDelegate(flags)
		if err != nil {
			log.Fatalf("Error building delegate object: %v", err)
		}
		result, err := gsmgmail.CreateDelegate(flags["userId"].GetString(), flags["fields"].GetString(), d)
		if err != nil {
			log.Fatalf("Error creating delegate: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(delegatesCmd, delegatesCreateCmd, delegateFlags)
}
