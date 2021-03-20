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

	"github.com/hanneshayashi/gsm/gsmadmin"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// customersPatchCmd represents the patch command
var customersPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates a customer. This method supports patch semantics.",
	Long:              "https://developers.google.com/admin-sdk/directory/v1/reference/customers/patch",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToCustomer(flags)
		if err != nil {
			log.Fatalf("Error building customer object: %v", err)
		}
		result, err := gsmadmin.PatchCustomer(flags["customerKey"].GetString(), flags["fields"].GetString(), c)
		if err != nil {
			log.Fatalf("Error patching customer: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(customersCmd, customersPatchCmd, customerFlags)
}
