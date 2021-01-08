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
	"log"

	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	ci "google.golang.org/api/cloudidentity/v1beta1"

	"github.com/spf13/cobra"
)

// groupsCiSearchCmd represents the search command
var groupsCiSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for Groups matching a specified query.",
	Long: `https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/search
Examples:
  - Search for security groups: gsm groupsCi search --query "parent == 'customers/{customer_id}' && 'cloudidentity.googleapis.com/groups.security' in labels"
  - Search for dynamic groups: gsm groupsCi search --query "parent == 'customers/{customer_id}' && 'cloudidentity.googleapis.com/groups.dynamic' in labels"`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmci.SearchGroups(flags["query"].GetString(), flags["view"].GetString(), flags["fields"].GetString(), gsmhelpers.MaxThreads(0))
		if streamOutput {
			enc := gsmhelpers.GetJSONEncoder(false)
			for i := range result {
				enc.Encode(i)
			}
		} else {
			final := []*ci.Group{}
			for i := range result {
				final = append(final, i)
			}
			gsmhelpers.Output(final, "json", compressOutput)
		}
		e := <-err
		if e != nil {
			log.Fatalf("Error searching for groups: %v", e)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupsCiCmd, groupsCiSearchCmd, groupCiFlags)
}
