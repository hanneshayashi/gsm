/*
Package gsmlog implements the Enterprise License Manager API
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

package gsmlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// PrintLastLines prints the last n lines in the specified file
func PrintLastLines(path string, n int) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	s := string(b)
	lines := strings.Split(s, "\n")
	max := len(lines) - 1
	if n > max || n == 0 {
		n = max
	}
	lastLines := lines[max-n : max]
	for i := range lastLines {
		fmt.Println(lastLines[i])
	}
	return nil
}

// Clear clears the specified file (truncate its content)
func Clear(path string) error {
	f, err := os.OpenFile(path, os.O_TRUNC, os.ModeTemporary)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
