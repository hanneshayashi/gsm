/*
Copyright © 2020-2025 Hannes Hayashi

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

// groupsPatchCmd represents the patch command
var groupsPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates a group's properties. This method supports patch semantics",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/groups/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		g, err := mapToGroup(flags)
		if err != nil {
			log.Fatalf("Error building group object: %v", err)
		}
		result, err := gsmadmin.PatchGroup(flags["groupKey"].GetString(), flags["fields"].GetString(), g)
		if err != nil {
			log.Fatalf("Error patching group: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupsCmd, groupsPatchCmd, groupFlags)
}
