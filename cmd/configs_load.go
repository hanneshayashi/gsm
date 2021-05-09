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
	"fmt"

	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// configsLoadCmd represents the load command
var configsLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a config file",
	Long: `Specify a config file to load by name (without file extension).
The current default config will be renamed to '<name>.yaml' and the loaded config will be renamed to '.gsm.yaml'`,
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		name := flags["name"].GetString()
		err := gsmconfig.LoadConfig(name)
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}
		fmt.Printf("'%s' successfully loaded\n", name)
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsLoadCmd, configFlags)
}
