/*
Copyright © 2020-2024 Hannes Hayashi

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
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		result, err := gsmconfig.ListConfigs()
		defaultPresent := false
		for i := 0; i < len(result); i++ {
			if result[i].Default {
				defaultPresent = true
			}
		}
		if err != nil {
			fmt.Printf("Error listing configs: %v\n", err)
			return
		}
		if len(result) == 0 {
			fmt.Println("Can't find any configs. Please create a config with \"gsm configs new\" to get started.")
			return
		}
		if flags["details"].GetBool() {
			err = gsmhelpers.Output(result, "yaml", false)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			for i := 0; i < len(result); i++ {
				if result[i].Default {
					fmt.Println(result[i].Name, "(Default)")
				} else {
					fmt.Println(result[i].Name)
				}
			}
		}
		if !defaultPresent && !flags["config"].IsSet() {
			fmt.Println("No config loaded. Please run \"gsm configs load --name\" to load a config file")
		}
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsListCmd, configFlags)
}
