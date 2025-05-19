/*
Copyright © 2020-2025 Hannes Hayashi

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

	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// configsNewCmd represents the new command
var configsNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new config file.",
	Long: `Config files are saved using the YAML format under
'~/.config/gsm/<name>.yaml'.`,
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		c, err := mapToConfig(flags)
		if err != nil {
			fmt.Printf("Error building config object: %v\n", err)
			return
		}
		result, err := gsmconfig.CreateConfig(c)
		if err != nil {
			fmt.Printf("Error creating config: %v\n", err)
			return
		}
		fmt.Printf("Config file \"%s\" successfully created. Load with \"gsm configs load --name %s\"\n", result, c.Name)
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsNewCmd, configFlags)
}
