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
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v5"
)

// RetryTransport is an http.RoundTripper that wraps another RoundTripper and
// retries requests when transient errors are encountered. It is designed
// specifically for Google Workspace APIs.
//
// # Retry Strategy
//
// The following status codes trigger retries:
//   - 429 (Too Many Requests): Respects Retry-After header
//   - 403 with "quota", "limit", or "rate" in the body: Google API rate limits
//   - Any status codes configured in RetryOn (e.g., 502, 503, 504)
//
// All other 4xx/5xx errors are treated as permanent failures (no retry).
type RetryTransport struct {
	// Base is the underlying RoundTripper used to make the actual requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper

	// InitialInterval is the initial backoff interval.
	InitialInterval time.Duration

	// MaxInterval is the maximum backoff interval.
	MaxInterval time.Duration

	// MaxElapsedTime is the maximum total elapsed time for retries.
	MaxElapsedTime time.Duration

	// Multiplier is the backoff multiplier.
	Multiplier float64

	// RetryOn defines additional HTTP status codes that should be retried.
	// 429 and 403+rate-limit are always retried regardless of this setting.
	RetryOn []int
}

func (t *RetryTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// isRetryableStatusCode checks whether the HTTP status code and response body
// indicate a transient error that should be retried.
func (t *RetryTransport) isRetryableStatusCode(statusCode int, body []byte) bool {
	if statusCode == http.StatusTooManyRequests {
		return true
	}
	if statusCode == http.StatusForbidden {
		bodyLower := strings.ToLower(string(body))
		for _, kw := range []string{"quota", "limit", "rate"} {
			if strings.Contains(bodyLower, kw) {
				return true
			}
		}
	}
	return slices.Contains(t.RetryOn, statusCode)
}

// RoundTrip implements http.RoundTripper with retry logic.
func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Buffer the request body so we can re-send on retries.
	var bodyBytes []byte
	if req.Body != nil && req.Body != http.NoBody {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %w", err)
		}
		if closeErr := req.Body.Close(); closeErr != nil {
			log.Printf("[WARN] Error closing original request body: %v", closeErr)
		}
	}

	bo := backoff.NewExponentialBackOff()
	if t.InitialInterval > 0 {
		bo.InitialInterval = t.InitialInterval
	}
	if t.MaxInterval > 0 {
		bo.MaxInterval = t.MaxInterval
	}
	if t.Multiplier > 0 {
		bo.Multiplier = t.Multiplier
	}

	operation := func() (*http.Response, error) {
		// Reset body for each retry attempt.
		if bodyBytes != nil {
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			req.ContentLength = int64(len(bodyBytes))
		}

		resp, err := t.base().RoundTrip(req)
		if err != nil {
			// Network-level errors are retryable.
			return nil, err
		}

		// For non-error status codes, return immediately.
		if resp.StatusCode < 400 {
			return resp, nil
		}

		// Read the body to check for retryable error messages.
		respBody, readErr := io.ReadAll(resp.Body)
		closeErr := resp.Body.Close()
		if readErr != nil {
			return nil, backoff.Permanent(fmt.Errorf("HTTP %d, failed to read body: %w", resp.StatusCode, readErr))
		}
		if closeErr != nil {
			log.Printf("[WARN] Error closing response body: %v", closeErr)
		}

		if t.isRetryableStatusCode(resp.StatusCode, respBody) {
			// Handle Retry-After header for 429s.
			if resp.StatusCode == http.StatusTooManyRequests {
				retryAfter := resp.Header.Get("Retry-After")
				if retryAfter != "" {
					if seconds, parseErr := strconv.Atoi(retryAfter); parseErr == nil {
						return nil, backoff.RetryAfter(seconds)
					}
					if t, parseErr := time.Parse(time.RFC1123, retryAfter); parseErr == nil {
						waitDuration := time.Until(t)
						if waitDuration > 0 {
							seconds := int(waitDuration.Seconds()) + 1
							return nil, backoff.RetryAfter(seconds)
						}
					}
				}
			}
			log.Printf("Retryable error: HTTP %d %s", resp.StatusCode, req.URL.String())
			return nil, fmt.Errorf("retryable HTTP %d: %s", resp.StatusCode, string(respBody))
		}

		// Non-retryable error: reconstruct the response with the body we already read.
		resp.Body = io.NopCloser(bytes.NewReader(respBody))
		return resp, nil
	}

	opts := []backoff.RetryOption{
		backoff.WithBackOff(bo),
		backoff.WithNotify(func(err error, d time.Duration) {
			log.Printf("%v - Retrying after %s...", err, d)
		}),
	}
	if t.MaxElapsedTime > 0 {
		opts = append(opts, backoff.WithMaxElapsedTime(t.MaxElapsedTime))
	}

	return backoff.Retry(req.Context(), operation, opts...)
}
