/*
Package cmd contains the commands available to the end user
Copyright © 2020-2021 Hannes Hayashi

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

// threadsDeleteCmd represents the delete command
var threadsDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: `Immediately and permanently deletes the specified thread.
This operation cannot be undone. Prefer threads trash instead.`,
	Long:              "https://developers.google.com/gmail/api/reference/rest/v1/users.threads/delete",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmgmail.DeleteThread(flags["userId"].GetString(), flags["id"].GetString())
		if err != nil {
			log.Fatalf("Error deleting thread %s: %v", flags["id"].GetString(), err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(threadsCmd, threadsDeleteCmd, threadFlags)
}
