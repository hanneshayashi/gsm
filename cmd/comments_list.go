/*
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

	"github.com/spf13/cobra"
)

// commentsListCmd represents the list command
var commentsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "Lists a file's comments.",
	Long:              "Implements the API documented at https://developers.google.com/workspace/drive/api/reference/rest/v3/comments/list",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		gsmhelpers.StreamOrCollect(gsmdrive.ListComments(flags["fileId"].GetString(), flags["startModifiedTime"].GetString(), flags["fields"].GetString(), flags["includeDeleted"].GetBool()), streamOutput, compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(commentsCmd, commentsListCmd, commentFlags)
}
