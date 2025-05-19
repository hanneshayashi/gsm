/*
Copyright © 2020-2025 Hannes Hayashi

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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/hanneshayashi/gsm/gsmconfig"
	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/skratchdot/open-golang/open"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/impersonate"
)

var ctx context.Context

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

// GetClientUser does user-based authentication via OAuth and returns an *http.Client
func GetClientUser(credentials []byte, tokenName string, redirectPort int, scope ...string) (client *http.Client, err error) {
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(credentials, scope...)
	config.RedirectURL = fmt.Sprintf("http://127.0.0.1:%d/oauth/callback", redirectPort)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	tokenPath := fmt.Sprintf("%s/%s", gsmconfig.CfgDir, tokenName)
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok, err := tokenFromFile(tokenPath)
	if err != nil {
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		srv := &http.Server{Addr: fmt.Sprintf(":%d", redirectPort)}
		done := make(chan bool, 1)
		http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
			queryParts, _ := url.ParseQuery(r.URL.RawQuery)
			code := queryParts["code"][0]
			tok, err = config.Exchange(ctx, code)
			if err != nil {
				log.Fatal(err)
			}
			err = saveToken(tokenPath, tok)
			if err != nil {
				log.Fatal(err)
			}
			_, err = fmt.Fprintf(w, "You can close this window now")
			if err != nil {
				log.Fatal(err)
			}
			done <- true
			close(done)
		})
		err = open.Run(authURL)
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			if <-done {
				errShutdown := srv.Shutdown(ctx)
				if errShutdown != nil {
					log.Fatal(errShutdown)
				}
			}
		}()
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}
	return config.Client(ctx, tok), nil
}

// GetClientADC returns a client to be used for API services
func GetClientADC(subject, serviceAccountEmail string, scope ...string) (client *http.Client, err error) {
	if serviceAccountEmail == "" {
		serviceAccountEmail, err = metadata.EmailWithContext(context.Background(), "")
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

// GetClient returns a client to be used for API services
func GetClient(subject string, credentials []byte, scope ...string) (client *http.Client, err error) {
	config, err := google.JWTConfigFromJSON(credentials, scope...)
	if err != nil {
		return nil, fmt.Errorf("error parsing Service Account credential file to config: %v", err)
	}
	config.Subject = subject
	return config.Client(ctx), nil
}

func init() {
	ctx = context.Background()
}
