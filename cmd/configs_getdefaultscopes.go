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
	"strings"

	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
)

// configsGetDefaultScopesCmd represents the getDefaultScopes command
var configsGetDefaultScopesCmd = &cobra.Command{
	Use:   "getDefaultScopes",
	Short: "Returns the default scopes.",
	Long: `The default scopes may change with every GSM version as new APIs are added.
Use this command to see the current defaults.
Use "gsm configs resetScopes" to set the current default.`,
	Annotations: map[string]string{
		"crescendoOutput": "$args[0]",
	},
	DisableAutoGenTag: true,
	Run: func(_ *cobra.Command, _ []string) {
		result := gsmconfig.GetDefaultScopes()
		fmt.Println(strings.Join(result, ","))
	},
}

func init() {
	gsmhelpers.InitCommand(configsCmd, configsGetDefaultScopesCmd, configFlags)
}
