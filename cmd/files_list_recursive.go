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
	"github.com/spf13/viper"
)

// filesListRecursiveCmd represents the recursive command
var filesListRecursiveCmd = &cobra.Command{
	Use:   "recursive",
	Short: "Recursively list files in a folder",
	Long:  "https://developers.google.com/drive/api/v3/reference/files/list",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmdrive.ListFilesRecursive(flags["folderId"].GetString(), flags["fields"].GetString(), viper.GetInt("threads"))
		if err != nil {
			log.Fatalf("Error listing files %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json", compressOutput))
	},
}

func init() {
	gsmhelpers.InitRecursiveCommand(filesListCmd, filesListRecursiveCmd, fileFlags, recursiveFlags)
}