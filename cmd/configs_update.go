/*
Copyright © 2020-2022 Hannes Hayashi

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

// configsUpdateCmd represents the update command
var configsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the current config.",
	Long:  `To update a config that is not currently loaded, use the --config flag to load it temporarily.`,
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
		result, err := gsmconfig.UpdateConfig(c, cfgFile)
		if err != nil {
			fmt.Printf("Error creating config: %v\n", err)
			return
		}
		err = gsmhelpers.Output(result, "yaml", false)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsUpdateCmd, configFlags)
}
