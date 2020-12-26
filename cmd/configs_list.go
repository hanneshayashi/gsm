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
	"log"

	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// configsListCmd represents the list command
var configsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List current configurations",
	Long:  ``,
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmconfig.ListConfigs()
		if err != nil {
			log.Fatalf("Error listing configs: %v", err)
		}
		if flags["details"].GetBool() {
			gsmhelpers.Output(result, "yaml", false)
		} else {
			if len(result) > 0 {
				fmt.Println(result[0].Name, "(Default)")
			}
			for i := 1; i < len(result); i++ {
				fmt.Println(result[i].Name)
			}
		}
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsListCmd, configFlags)
}
