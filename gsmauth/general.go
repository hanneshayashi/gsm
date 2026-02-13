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

// Package gsmauth provides the authentication mechanisms for Google APIs
package gsmauth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/pkg/browser"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/impersonate"
)

// Retrieves a token from a local file.
func tokenFromFile(tokenPath string) (*oauth2.Token, error) {
	f, err := os.Open(tokenPath)
	if err != nil {
		return nil, err
	}
	defer gsmhelpers.CloseLog(f, "tokenFile")
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache OAuth token: %v", err)
	}
	defer gsmhelpers.CloseLog(f, "tokenFile")
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		return fmt.Errorf("unable to save OAuth token: %v", err)
	}
	return nil
}

// randomState generates a random state token for OAuth CSRF protection.
func randomState() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random state: %v", err)
	}
	return hex.EncodeToString(b), nil
}

// GetClientUser does user-based authentication via OAuth and returns an *http.Client
func GetClientUser(credentials []byte, tokenName string, redirectPort int, scope ...string) (client *http.Client, err error) {
	ctx := context.Background()
	config, err := google.ConfigFromJSON(credentials, scope...)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	config.RedirectURL = fmt.Sprintf("http://127.0.0.1:%d/oauth/callback", redirectPort)
	tokenPath := fmt.Sprintf("%s/%s", gsmconfig.CfgDir, tokenName)
	tok, err := tokenFromFile(tokenPath)
	if err != nil {
		state, err := randomState()
		if err != nil {
			return nil, err
		}
		authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline)
		mux := http.NewServeMux()
		srv := &http.Server{Addr: fmt.Sprintf(":%d", redirectPort), Handler: mux}
		done := make(chan bool, 1)
		var callbackErr error
		mux.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
			queryParts, _ := url.ParseQuery(r.URL.RawQuery)
			if queryParts.Get("state") != state {
				http.Error(w, "Invalid state parameter", http.StatusBadRequest)
				callbackErr = fmt.Errorf("OAuth callback: invalid state parameter (possible CSRF)")
				done <- true
				close(done)
				return
			}
			code := queryParts["code"][0]
			tok, callbackErr = config.Exchange(ctx, code)
			if callbackErr != nil {
				http.Error(w, "Token exchange failed", http.StatusInternalServerError)
				done <- true
				close(done)
				return
			}
			callbackErr = saveToken(tokenPath, tok)
			if callbackErr != nil {
				http.Error(w, "Failed to save token", http.StatusInternalServerError)
				done <- true
				close(done)
				return
			}
			_, _ = fmt.Fprintf(w, "You can close this window now")
			done <- true
			close(done)
		})
		err = browser.OpenURL(authURL)
		if err != nil {
			return nil, fmt.Errorf("unable to open browser for OAuth: %v", err)
		}
		go func() {
			if <-done {
				_ = srv.Shutdown(ctx)
			}
		}()
		if srvErr := srv.ListenAndServe(); srvErr != http.ErrServerClosed {
			return nil, fmt.Errorf("OAuth callback server error: %v", srvErr)
		}
		if callbackErr != nil {
			return nil, callbackErr
		}
	}
	return config.Client(ctx, tok), nil
}

// GetClientADC returns a client using Application Default Credentials with impersonation.
// This is the original "adc" mode - it always impersonates a service account.
func GetClientADC(subject, serviceAccountEmail string, scope ...string) (client *http.Client, err error) {
	ctx := context.Background()
	if serviceAccountEmail == "" {
		serviceAccountEmail, err = metadata.EmailWithContext(ctx, "")
		if err != nil {
			return nil, fmt.Errorf("error getting Service Account email: %v", err)
		}
	}
	ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: serviceAccountEmail,
		Scopes:          scope,
		Subject:         subject,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting token source: %v", err)
	}
	client = oauth2.NewClient(ctx, ts)
	return
}

// GetClient returns a client using a service account key with domain-wide delegation.
// The subject parameter specifies the user to impersonate.
func GetClient(subject string, credentials []byte, scope ...string) (client *http.Client, err error) {
	config, err := google.JWTConfigFromJSON(credentials, scope...)
	if err != nil {
		return nil, fmt.Errorf("error parsing Service Account credential file to config: %v", err)
	}
	config.Subject = subject
	return config.Client(context.Background()), nil
}

// GetClientSA returns a client using a service account key without domain-wide delegation.
// This authenticates as the service account itself, useful for accessing resources
// shared directly with the SA (e.g., shared drives, calendars).
func GetClientSA(credentials []byte, scope ...string) (*http.Client, error) {
	config, err := google.JWTConfigFromJSON(credentials, scope...)
	if err != nil {
		return nil, fmt.Errorf("error parsing Service Account credential file to config: %v", err)
	}
	return config.Client(context.Background()), nil
}

// GetClientADCDirect returns a client using Application Default Credentials directly,
// without impersonation. This supports Workload Identity Federation, Cloud Run,
// GCE, and local development with `gcloud auth application-default login`.
func GetClientADCDirect(scope ...string) (*http.Client, error) {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, scope...)
	if err != nil {
		return nil, fmt.Errorf("error getting default client from ADC: %v", err)
	}
	return client, nil
}
