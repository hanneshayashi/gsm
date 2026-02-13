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
	"iter"
	"log"
)

// StreamOrCollect consumes an iter.Seq2[T, error] and either streams each
// item as NDJSON to stdout (stream=true) or collects all items into a slice
// and outputs them as a JSON array (stream=false).
//
// Errors encountered during iteration are logged. If an error occurs in
// non-streaming mode, the items collected so far are still output.
func StreamOrCollect[T any](seq iter.Seq2[T, error], stream, compress bool) {
	if stream {
		enc := GetJSONEncoder(false)
		for item, err := range seq {
			if err != nil {
				log.Println(err)
				continue
			}
			if err := enc.Encode(item); err != nil {
				log.Println(err)
			}
		}
	} else {
		var items []T
		for item, err := range seq {
			if err != nil {
				log.Println(err)
				continue
			}
			items = append(items, item)
		}
		if err := Output(items, "json", compress); err != nil {
			log.Fatalln(err)
		}
	}
}

// ChanToIter bridges a channel-based producer into an iter.Seq2[T, error].
// This is used for concurrent operations (batch commands, recursive listings)
// that need internal goroutines but want a clean iterator-based consumer API.
//
// If errCh is nil, no error channel is consumed. Otherwise, any error sent
// on errCh is yielded after the data channel is drained.
func ChanToIter[T any](ch <-chan T, errCh <-chan error) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for item := range ch {
			var zero error
			if !yield(item, zero) {
				// Consumer stopped early. Drain the channel to unblock
				// the producer goroutine and prevent leaks.
				go func() {
					for range ch {
					}
					// Also drain error channel if present
					if errCh != nil {
						for range errCh {
						}
					}
				}()
				return
			}
		}
		// Channel drained, now check for errors
		if errCh != nil {
			for e := range errCh {
				var zero T
				if !yield(zero, e) {
					return
				}
			}
		}
	}
}

// StreamOrCollectJSON is a convenience wrapper that streams or collects
// results from a channel directly, bridging via ChanToIter.
func StreamOrCollectJSON[T any](ch <-chan T, errCh <-chan error, stream, compress bool) {
	StreamOrCollect(ChanToIter(ch, errCh), stream, compress)
}
