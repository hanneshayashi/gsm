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
	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"log"

	"github.com/spf13/cobra"
)

// configsRemoveCmd represents the remove command
var configsRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a configuration",
	Long: `This will delete the configuration's .yaml file from your '~/.config/gsm' directory.
Credential files and tokens will not be removed!`,
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},	
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		flags := gsmhelpers.FlagsToMap(cmd.Flags())
		name := flags["name"].GetString()
		err := gsmconfig.RemoveConfig(name)
		if err != nil {
			log.Fatalf("Error removing config: %v", err)
		}
		fmt.Printf("'%s' successfully removed.\n", name)
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsRemoveCmd, configFlags)
}
