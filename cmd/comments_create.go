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
	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// commentsCreateCmd represents the create command
var commentsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new comment on a file.",
	Long:  "https://developers.google.com/drive/api/v3/reference/comments/create",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToComment(flags)
		if err != nil {
			log.Fatalf("Error building comment object: %v", err)
		}
		result, err := gsmdrive.CreateComment(flags["fileId"].GetString(), flags["fields"].GetString(), c)
		if err != nil {
			log.Fatalf("Error creating comment on file %s: %v", flags["fileId"].GetString(), err)
		}
		gsmhelpers.StreamOutput(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(commentsCmd, commentsCreateCmd, commentFlags)
}
