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
	"gsm/gsmadmin"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// usersSignOutCmd represents the signOut command
var usersSignOutCmd = &cobra.Command{
	Use: "signOut",
	Short: `Sign a user out of all web and device sessions and reset their sign-in cookies.
User will have to sign in by authenticating again.`,
	Long: "https://developers.google.com/admin-sdk/directory/v1/reference/users/signOut",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.SignOutUser(flags["userKey"].GetString())
		if err != nil {
			log.Fatalf("Error signOutting user %v", err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(usersCmd, usersSignOutCmd, userFlags)
}
