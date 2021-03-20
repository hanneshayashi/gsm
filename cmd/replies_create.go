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

	"github.com/hanneshayashi/gsm/gsmdrive"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// repliesCreateCmd represents the create command
var repliesCreateCmd = &cobra.Command{
	Use:               "create",
	Short:             "Creates a new reply to a comment.",
	Long:              "https://developers.google.com/drive/api/v3/reference/replies/create",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		r, err := mapToReply(flags)
		if err != nil {
			log.Fatalf("Error building reply object: %v", err)
		}
		result, err := gsmdrive.CreateReply(flags["fileId"].GetString(), flags["commentId"].GetString(), flags["fields"].GetString(), r)
		if err != nil {
			log.Fatalf("Error creating reply: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(repliesCmd, repliesCreateCmd, replyFlags)
}
