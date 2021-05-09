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
	"log"
	"os"

	"github.com/spf13/cobra"
)

// completionsCmd represents the completion command
var completionsCmd = &cobra.Command{
	Use:   "completions [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

Make sure you have bash-completion installed!
$ source <(./gsm completions bash)

# To load completions for each session, execute once:
Linux:
  $ ./gsm completions bash | sudo tee /etc/bash_completion.d/gsm
MacOS:
  $ ./gsm completions bash | sudo tee /usr/local/etc/bash_completion.d/gsm

PowerShell:

./gsm completions powershell > ./gsm.ps1
> ./gsm.ps1

To load completions for each session, execute once (in powershell):
> mkdir $profile.Substring(0,$profile.LastIndexOf("/"))
> ./gsm completions powershell >> $profile

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ ./gsm completions zsh > "${fpath[1]}/_gsm"

# You will need to start a new shell for this setup to take effect.

Fish:

$ ./gsm completions fish | source

# To load completions for each session, execute once:
$ ./gsm completions fish > ~/.config/fish/completions/gsm.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	DisableAutoGenTag:     true,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			err := cmd.Root().GenBashCompletion(os.Stdout)
			if err != nil {
				log.Fatalln(err)
			}
		case "zsh":
			err := cmd.Root().GenZshCompletion(os.Stdout)
			if err != nil {
				log.Fatalln(err)
			}
		case "fish":
			err := cmd.Root().GenFishCompletion(os.Stdout, true)
			if err != nil {
				log.Fatalln(err)
			}
		case "powershell":
			err := cmd.Root().GenPowerShellCompletion(os.Stdout)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completionsCmd)
}
