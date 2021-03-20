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

	"github.com/hanneshayashi/gsm/gsmcalendar"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// calendarACLPatchCmd represents the patch command
var calendarACLPatchCmd = &cobra.Command{
	Use:               "patch",
	Short:             "Updates an access control rule. This method supports patch semantics.",
	Long:              `https://developers.google.com/calendar/v3/reference/acl/patch`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		a, err := mapToCalendarACLRule(flags)
		if err != nil {
			log.Fatalf("Error building acl rule object: %v", err)
		}
		result, err := gsmcalendar.PatchACL(flags["calendarId"].GetString(), flags["ruleId"].GetString(), flags["fields"].GetString(), a, flags["sendNotifications"].GetBool())
		if err != nil {
			log.Fatalf("Error patching calendar acl rule: %v", err)
		}
		gsmhelpers.Output(result, "json", compressOutput)
	},
}

func init() {
	gsmhelpers.InitCommand(calendarACLCmd, calendarACLPatchCmd, calendarACLFlags)
}
