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
	"testing"
)

// TestAcc_Files_Get verifies that "files get" returns valid JSON metadata.
// It uses the "root" alias which always exists for any authenticated user.
func TestAcc_Files_Get(t *testing.T) {
	skipUnlessAcc(t)

	var result map[string]any
	runGSMJSON(t, &result, "files", "get", "--fileId", "root", "--compressOutput")

	if result["id"] == nil || result["id"] == "" {
		t.Error("expected non-empty 'id' field in files get response")
	}
	if result["name"] == nil || result["name"] == "" {
		t.Error("expected non-empty 'name' field in files get response")
	}
	t.Logf("files get root: id=%v name=%v", result["id"], result["name"])
}

// TestAcc_Files_List verifies that "files list" returns a JSON array.
func TestAcc_Files_List(t *testing.T) {
	skipUnlessAcc(t)

	// Use a query to limit results instead of listing the entire Drive.
	var results []map[string]any
	runGSMJSON(t, &results, "files", "list",
		"--q", "'root' in parents",
		"--fields", "nextPageToken,files(id,name)",
		"--compressOutput")

	if len(results) == 0 {
		t.Skip("files list returned 0 results — user's Drive root may be empty")
	}

	first := results[0]
	if first["id"] == nil || first["id"] == "" {
		t.Error("expected non-empty 'id' in first file")
	}
	t.Logf("files list returned %d files, first: id=%v name=%v", len(results), first["id"], first["name"])
}
