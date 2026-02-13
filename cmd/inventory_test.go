/*
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
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// TestGenerateInventory walks the Cobra command tree and writes a markdown
// inventory of all commands and their flags to COMMANDS.md.
//
// Run with:
//
//	GSM_GENERATE_INVENTORY=1 go test -run TestGenerateInventory ./cmd/ -count=1
func TestGenerateInventory(t *testing.T) {
	if os.Getenv("GSM_GENERATE_INVENTORY") == "" {
		t.Skip("Skipping: set GSM_GENERATE_INVENTORY=1 to regenerate COMMANDS.md")
	}

	var sb strings.Builder
	sb.WriteString("# GSM Command Inventory\n\n")
	sb.WriteString("Auto-generated — run `GSM_GENERATE_INVENTORY=1 go test -run TestGenerateInventory ./cmd/` to update.\n\n")
	sb.WriteString("## Summary\n\n")

	// Collect stats first.
	var totalCmds, totalLeafs int
	countCommands(rootCmd, &totalCmds, &totalLeafs)
	fmt.Fprintf(&sb, "- **Total commands**: %d\n", totalCmds)
	fmt.Fprintf(&sb, "- **Leaf commands** (executable): %d\n\n", totalLeafs)

	sb.WriteString("## Commands\n\n")
	writeCommandTree(&sb, rootCmd, 0)

	err := os.WriteFile("../COMMANDS.md", []byte(sb.String()), 0644)
	if err != nil {
		t.Fatalf("Failed to write COMMANDS.md: %v", err)
	}
	t.Logf("Wrote COMMANDS.md (%d bytes, %d commands, %d leaf commands)", sb.Len(), totalCmds, totalLeafs)
}

func countCommands(cmd *cobra.Command, total, leafs *int) {
	children := cmd.Commands()
	if len(children) == 0 {
		*leafs++
	}
	for _, child := range children {
		if child.Hidden {
			continue
		}
		*total++
		countCommands(child, total, leafs)
	}
}

func writeCommandTree(sb *strings.Builder, cmd *cobra.Command, depth int) {
	children := cmd.Commands()
	// Sort children alphabetically.
	sort.Slice(children, func(i, j int) bool {
		return children[i].Name() < children[j].Name()
	})

	for _, child := range children {
		if child.Hidden {
			continue
		}

		fullPath := commandPath(child)
		isLeaf := len(child.Commands()) == 0

		// Write heading: depth 0 = ###, depth 1 = ####, etc.
		headingLevel := depth + 3
		if headingLevel > 6 {
			headingLevel = 6
		}
		heading := strings.Repeat("#", headingLevel)
		fmt.Fprintf(sb, "%s `%s`\n\n", heading, fullPath)

		if child.Short != "" {
			sb.WriteString(child.Short + "\n\n")
		}

		// Write flags for leaf commands only.
		if isLeaf {
			writeFlags(sb, child)
		}

		// Recurse into subcommands.
		writeCommandTree(sb, child, depth+1)
	}
}

func commandPath(cmd *cobra.Command) string {
	parts := []string{}
	for c := cmd; c != nil && c.Name() != "gsm"; c = c.Parent() {
		parts = append([]string{c.Name()}, parts...)
	}
	return "gsm " + strings.Join(parts, " ")
}

func writeFlags(sb *strings.Builder, cmd *cobra.Command) {
	var flags []*pflag.Flag
	cmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		if f.Hidden {
			return
		}
		flags = append(flags, f)
	})

	if len(flags) == 0 {
		sb.WriteString("*No command-specific flags.*\n\n")
		return
	}

	sb.WriteString("| Flag | Type | Required | Description |\n")
	sb.WriteString("|------|------|----------|-------------|\n")
	for _, f := range flags {
		required := ""
		if _, ok := f.Annotations[cobra.BashCompOneRequiredFlag]; ok {
			required = "✓"
		}
		// Truncate long descriptions and escape pipes.
		desc := strings.ReplaceAll(f.Usage, "|", "\\|")
		desc = strings.ReplaceAll(desc, "\n", " ")
		if len(desc) > 120 {
			desc = desc[:117] + "..."
		}
		fmt.Fprintf(sb, "| `%s` | %s | %s | %s |\n", f.Name, f.Value.Type(), required, desc)
	}
	sb.WriteString("\n")
}
