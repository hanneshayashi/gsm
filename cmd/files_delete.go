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
	"fmt"
	"gsm/gsmdrive"
	"gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// filesDeleteCmd represents the delete command
var filesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently deletes a file owned by the user without moving it to the trash. If the file belongs to a shared drive the user must be an organizer on the parent. If the target is a folder, all descendants owned by the user are also deleted.",
	Long:  "https://developers.google.com/drive/api/v3/reference/files/delete",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrive.DeleteFile(flags["fileId"].GetString())
		if err != nil {
			log.Fatalf("Error deleting file %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	filesCmd.AddCommand(filesDeleteCmd)
	gsmhelpers.AddFlags(fileFlags, filesDeleteCmd.Flags(), filesDeleteCmd.Use)
	markFlagsRequired(filesDeleteCmd, fileFlags, "")
}
