/*
Package gsmauth provides the authentication mechanisms for Google APIs
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

package gsmauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/hanneshayashi/gsm/gsmconfig"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/impersonate"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, tokenName string) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokenName = fmt.Sprintf("%s/%s", gsmconfig.CfgDir, tokenName)
	tok, err := tokenFromFile(tokenName)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenName, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatalf("Unable to save OAuth token: %v", err)
	}
}

//GetClientUser does user-based authentication via OAuth and returns an *http.Client
func GetClientUser(credentials []byte, tokenName string, scope ...string) (client *http.Client) {
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(credentials, scope...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client = getClient(config, tokenName)
	return
}

//GetClientADC returns a client to be used for API services
func GetClientADC(subject, serviceAccountEmail string, scope ...string) (client *http.Client) {
	ctx := context.Background()
	if serviceAccountEmail == "" {
		var err error
		serviceAccountEmail, err = metadata.Email("")
		if err != nil {
			log.Fatalf("Error getting Service Account email: %v", err)
		}
	}
	ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: serviceAccountEmail,
		Scopes:          scope,
		Subject:         subject,
	})
	if err != nil {
		log.Fatalf("Error getting token source: %v", err)
	}
	client = oauth2.NewClient(ctx, ts)
	return
}

//GetClient returns a client to be used for API services
func GetClient(subject string, credentials []byte, scope ...string) (client *http.Client) {
	config, err := google.JWTConfigFromJSON(credentials, scope...)
	if err != nil {
		log.Fatalf("Error parsing Service Account credential file to config: %v", err)
	}
	config.Subject = subject
	return config.Client(context.Background())
}
