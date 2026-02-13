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
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
	"testing"
)

// stdoutMu serializes tests that capture os.Stdout.
// Only one test can swap the global Stdout at a time.
var stdoutMu sync.Mutex

// skipUnlessAcc skips the test if the GSM_ACC environment variable is not set.
// All acceptance tests should call this at the top.
func skipUnlessAcc(t *testing.T) {
	t.Helper()
	if os.Getenv("GSM_ACC") == "" {
		t.Skip("Skipping acceptance test (set GSM_ACC=1 to run)")
	}
}

// runGSM executes a GSM command with the given arguments and captures stdout.
// It returns the raw bytes written to stdout. If the command fails, the test
// is marked as failed.
//
// Usage:
//
//	out := runGSM(t, "files", "get", "--fileId", "root", "--fields", "id,name")
func runGSM(t *testing.T, args ...string) []byte {
	t.Helper()

	stdoutMu.Lock()
	defer stdoutMu.Unlock()

	// Swap stdout with a pipe so we can capture output.
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	os.Stdout = w

	// Run the command.
	rootCmd.SetArgs(args)
	execErr := rootCmd.Execute()

	// Close the write end and restore stdout before reading.
	_ = w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()

	if execErr != nil {
		t.Fatalf("gsm %v failed: %v\nOutput:\n%s", args, execErr, buf.String())
	}
	return buf.Bytes()
}

// runGSMJSON executes a GSM command, captures stdout, and unmarshals the JSON
// output into the target value. The target should be a pointer.
//
// Usage:
//
//	var file drive.File
//	runGSMJSON(t, &file, "files", "get", "--fileId", "root", "--fields", "id,name")
func runGSMJSON(t *testing.T, target any, args ...string) {
	t.Helper()
	out := runGSM(t, args...)
	if err := json.Unmarshal(out, target); err != nil {
		t.Fatalf("failed to unmarshal JSON from gsm %v: %v\nRaw output:\n%s", args, err, string(out))
	}
}
