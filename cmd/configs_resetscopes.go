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
	"fmt"
	"log"
	"strings"

	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// configsResetScopesCmd represents the resetScopes command
var configsResetScopesCmd = &cobra.Command{
	Use:   "resetScopes",
	Short: "Resets the scopes of a config back to the defaults.",
	Long: fmt.Sprintf(`Sets the the scopes of the specified config or the default config if not specified back to the defaults:
%s`, strings.Join(gsmconfig.GetDefaultScopes(), ",\n")),
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},
	DisableAutoGenTag: true,
	Run: func(_ *cobra.Command, _ []string) {
		result, err := gsmconfig.UpdateConfig(&gsmconfig.GSMConfig{
			Scopes: gsmconfig.GetDefaultScopes(),
		}, cfgFile)
		if err != nil {
			fmt.Printf("Error resetting scopes: %v\n", err)
			return
		}
		err = gsmhelpers.Output(result, "yaml", false)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsResetScopesCmd, configFlags)
}
