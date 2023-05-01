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

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/hanneshayashi/gsm/gsmlog"

	"github.com/spf13/cobra"
)

// logClearCmd represents the clear command
var logClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the current log.",
	Long:  "",
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},
	DisableAutoGenTag: true,
	Run: func(_ *cobra.Command, _ []string) {
		err := gsmlog.Clear(logFile)
		if err != nil {
			log.Fatalf("Error clearing log: %v", err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(logCmd, logClearCmd, logFlags)
}
