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

// commentsUpdateCmd represents the update command
var commentsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a comment with patch semantics.",
	Long:  "https://developers.google.com/drive/api/v3/reference/comments/update",
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToComment(flags)
		if err != nil {
			log.Fatalf("Error building comment object: %v", err)
		}
		result, err := gsmdrive.UpdateComment(flags["fileId"].GetString(), flags["commentId"].GetString(), flags["fields"].GetString(), c)
		if err != nil {
			log.Fatalf("Error updating comment: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), gsmhelpers.PrettyPrint(result, "json"))
	},
}

func init() {
	commentsCmd.AddCommand(commentsUpdateCmd)
	gsmhelpers.AddFlags(commentFlags, commentsUpdateCmd.Flags(), commentsUpdateCmd.Use)
	markFlagsRequired(commentsUpdateCmd, commentFlags, "")
}
