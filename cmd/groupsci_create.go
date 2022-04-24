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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmci"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// groupsCiCreateCmd represents the create command
var groupsCiCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a Group.",
	Long: `Implements the API documented at https://cloud.google.com/identity/docs/reference/rest/v1/groups/create
Examples:
  - Create a dynamic group:
    gsm groupsCi create --id group@example.org --labels "cloudidentity.googleapis.com/groups.discussion_forum" --queries "resourceType=USER;query=user.organizations.exists(org, org.department=='engineering')"`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		g, err := mapToGroupCi(flags)
		if err != nil {
			log.Fatalf("Error building group object: %v", err)
		}
		if g.Parent == "" {
			var customerID string
			customerID, err = gsmadmin.GetOwnCustomerID()
			if err != nil {
				log.Fatalf("Error determining customer ID: %v", err)
			}
			g.Parent = "customers/" + customerID
		}
		if len(g.Labels) == 0 {
			g.Labels = map[string]string{
				"cloudidentity.googleapis.com/groups.discussion_forum": "",
			}
		}
		result, err := gsmci.CreateGroup(g, flags["initialGroupConfig"].GetString(), flags["fields"].GetString())
		if err != nil {
			log.Fatalf("Error creating group: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(groupsCiCmd, groupsCiCreateCmd, groupCiFlags)
}
