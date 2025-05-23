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

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// drivesUpdateCmd represents the update command
var drivesUpdateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Updates the metadata for a shared drive.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/drives/update",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		d, err := mapToDrive(flags)
		if err != nil {
			log.Fatalf("Error building drive object: %v\n", err)
		}
		result, err := gsmdrive.UpdateDrive(flags["driveId"].GetString(), flags["fields"].GetString(), flags["useDomainAdminAccess"].GetBool(), d)
		if err != nil {
			log.Fatalf("Error updating drive: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(drivesCmd, drivesUpdateCmd, driveFlags)
}
