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
	"fmt"
	"log"

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// sharedContactsDeleteCmd represents the delete command
var sharedContactsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a shared contact by referencing its id url",
	Long:  "",
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmadmin.DeleteSharedContact(flags["url"].GetString())
		if err != nil {
			log.Fatalf("Error deleting shared contact: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), string(result))
	},
}

func init() {
	gsmhelpers.InitCommand(sharedContactsCmd, sharedContactsDeleteCmd, sharedContactFlags)
}
