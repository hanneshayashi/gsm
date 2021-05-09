/*
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

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// filesMoveCmd represents the move command
var filesMoveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a file.",
	Long: `You can't move folders to Shared Drives ourside your organization with this command!
Use "files move recursive" instead!`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		f, err := gsmdrive.GetFile(flags["fileId"].GetString(), "id,parents", "")
		if err != nil {
			log.Println(err)
		}
		result, err := gsmdrive.UpdateFile(f.Id, flags["parent"].GetString(), f.Parents[0], "", "", "", nil, nil, false, false)
		if err != nil {
			log.Fatalf("Error during move: %v", err)
		}
		err = gsmhelpers.Output(result, "json", compressOutput)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(filesCmd, filesMoveCmd, fileFlags)
}
