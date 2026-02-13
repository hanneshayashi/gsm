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

package gsmhelpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"iter"
	"os"
	"strings"
	"testing"
)

// testItem is a simple struct for testing StreamOrCollect.
type testItem struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func TestStreamOrCollect_Stream(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	items := []testItem{
		{Name: "alice", ID: 1},
		{Name: "bob", ID: 2},
		{Name: "charlie", ID: 3},
	}

	seq := func(yield func(testItem, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				return
			}
		}
	}

	StreamOrCollect(seq, true, false)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// NDJSON: each line is a separate JSON object
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines of NDJSON, got %d: %q", len(lines), output)
	}

	for i, line := range lines {
		var got testItem
		if err := json.Unmarshal([]byte(line), &got); err != nil {
			t.Fatalf("line %d: failed to unmarshal: %v", i, err)
		}
		if got != items[i] {
			t.Errorf("line %d: got %+v, want %+v", i, got, items[i])
		}
	}
}

func TestStreamOrCollect_Collect(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	items := []testItem{
		{Name: "alice", ID: 1},
		{Name: "bob", ID: 2},
	}

	seq := func(yield func(testItem, error) bool) {
		for _, item := range items {
			if !yield(item, nil) {
				return
			}
		}
	}

	StreamOrCollect(seq, false, false)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	var got []testItem
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal JSON array: %v\noutput: %s", err, buf.String())
	}
	if len(got) != 2 {
		t.Fatalf("expected 2 items, got %d", len(got))
	}
	for i, item := range got {
		if item != items[i] {
			t.Errorf("item %d: got %+v, want %+v", i, item, items[i])
		}
	}
}

func TestStreamOrCollect_ErrorHandling(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	testErr := errors.New("test error")

	seq := func(yield func(testItem, error) bool) {
		if !yield(testItem{Name: "alice", ID: 1}, nil) {
			return
		}
		if !yield(testItem{}, testErr) { // error item, should be skipped
			return
		}
		yield(testItem{Name: "bob", ID: 2}, nil)
	}

	StreamOrCollect(seq, true, false)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Should have 2 lines (alice and bob), error is logged but not output
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), output)
	}
}

func TestChanToIter(t *testing.T) {
	ch := make(chan testItem, 3)
	ch <- testItem{Name: "alice", ID: 1}
	ch <- testItem{Name: "bob", ID: 2}
	ch <- testItem{Name: "charlie", ID: 3}
	close(ch)

	var got []testItem
	for item, err := range ChanToIter(ch, nil) {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		got = append(got, item)
	}

	if len(got) != 3 {
		t.Fatalf("expected 3 items, got %d", len(got))
	}
	if got[0].Name != "alice" || got[1].Name != "bob" || got[2].Name != "charlie" {
		t.Errorf("unexpected items: %+v", got)
	}
}

func TestChanToIter_WithError(t *testing.T) {
	ch := make(chan testItem, 2)
	errCh := make(chan error, 1)

	ch <- testItem{Name: "alice", ID: 1}
	ch <- testItem{Name: "bob", ID: 2}
	close(ch)

	testErr := errors.New("pagination failed")
	errCh <- testErr
	close(errCh)

	var items []testItem
	var errs []error
	for item, err := range ChanToIter(ch, errCh) {
		if err != nil {
			errs = append(errs, err)
			continue
		}
		items = append(items, item)
	}

	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if len(errs) != 1 {
		t.Fatalf("expected 1 error, got %d", len(errs))
	}
	if !errors.Is(errs[0], testErr) {
		t.Errorf("expected %v, got %v", testErr, errs[0])
	}
}

func TestChanToIter_EarlyTermination(t *testing.T) {
	ch := make(chan int, 100)
	for i := range 100 {
		ch <- i
	}
	close(ch)

	var count int
	for _, err := range ChanToIter(ch, nil) {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		count++
		if count == 5 {
			break
		}
	}

	if count != 5 {
		t.Errorf("expected 5 items, got %d", count)
	}
}

func TestStreamOrCollectJSON(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ch := make(chan testItem, 2)
	ch <- testItem{Name: "alice", ID: 1}
	ch <- testItem{Name: "bob", ID: 2}
	close(ch)

	StreamOrCollectJSON(ch, nil, true, false)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
}

// TestStreamOrCollect_EmptyIterator verifies that an empty iterator produces
// valid output (empty JSON array for collect, nothing for stream).
func TestStreamOrCollect_EmptyIterator(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	empty := func(yield func(testItem, error) bool) {
		// yield nothing
	}

	StreamOrCollect(iter.Seq2[testItem, error](empty), false, false)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Should produce "null\n" since []testItem(nil) encodes as null
	// or "[]" - either is acceptable
	output := strings.TrimSpace(buf.String())
	if output != "null" && output != "[]" {
		t.Errorf("expected null or [] for empty collect, got %q", output)
	}
}
