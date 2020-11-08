/*
Package gsmhelpers contains helper functions to GSM
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
	"bufio"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gsm/gsmadmin"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"google.golang.org/api/googleapi"
	"gopkg.in/yaml.v2"
)

const version = "0.1.12"

// GetVersion returns the current version
func GetVersion() string {
	return version
}

// GetCSV uses a FlagSet to read a CSV file and parse it accordingliny
func GetCSV(flags map[string]*Value) ([][]string, error) {
	path := flags["path"].GetString()
	var delimiter rune
	if flags["delimiter"].Changed {
		delimiter = flags["delimiter"].GetRune()
	} else {
		delimiter = ';'
	}
	skipHeader := flags["skipHeader"].GetBool()
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	r.Comma = delimiter
	csv, err := r.ReadAll()
	if skipHeader {
		csv = csv[1:]
	}
	if err != nil {
		return nil, err
	}
	return csv, nil
}

// ErrorIsRetryable checks if a Google API response returned a retryable error
func ErrorIsRetryable(err error) bool {
	gerr := err.(*googleapi.Error)
	if gerr.Code == 403 && (strings.Contains(gerr.Message, "quota") || strings.Contains(gerr.Message, "limit") || strings.Contains(gerr.Message, "rate")) {
		return true
	}
	return false
}

// NewStandardRetrier returns a retrier with default values
func NewStandardRetrier() *retrier.Retrier {
	// class := retrier.WhitelistClassifier{
	// 	&googleapi.Error{Code: 403},
	// }
	return retrier.New(retrier.ExponentialBackoff(4, 20*time.Second), nil)
}

// GetCustomerID returns either your own customer ID or the provided one
func GetCustomerID(customerID string) string {
	if customerID == "" {
		var err error
		customerID, err = gsmadmin.GetOwnCustomerID()
		if err != nil {
			log.Printf("Error determining customer ID: %v\n", err)
		}
	}
	return customerID
}

// Contains checks if a string is inside a slice
func Contains(s string, slice []string) bool {
	for i := range slice {
		if s == slice[i] {
			return true
		}
	}
	return false
}

// MaxThreads returns the maximum number of threads (goroutines that should be spawned)
func MaxThreads(lines int) int {
	numCPU := runtime.NumCPU() * 2
	if lines < numCPU {
		return lines
	}
	return numCPU
}

// PrettyPrint is used to output the result of an API call in the requested format
func PrettyPrint(i interface{}, format string) string {
	var b []byte
	if format == "json" {
		b, _ = json.MarshalIndent(i, "", "\t")
	}
	if format == "xml" {
		b, _ = xml.MarshalIndent(i, "", "\t")
	}
	if format == "yaml" {
		b, _ = yaml.Marshal(i)
	}
	return string(b)
}

// CreateDocs creates GSM documentation
func CreateDocs(cmd *cobra.Command) {
	dir := "../gsm-hosting/gsm.hayashi-ke.online/content"
	tmpDir := dir + "/tmp"
	os.MkdirAll(tmpDir, os.ModePerm)
	filePrepender := func(filename string) string {
		return filename
	}
	linkHandler := func(name string) string {
		return "/" + strings.ReplaceAll(strings.TrimSuffix(strings.ToLower(name), ".md"), "_", "/")
	}
	err := doc.GenMarkdownTreeCustom(cmd, tmpDir, filePrepender, linkHandler)
	if err != nil {
		log.Fatalln(err)
	}
	d, err := os.Open(tmpDir)
	defer d.Close()
	if err != nil {
		log.Fatalln(err)
	}
	files, err := d.Readdir(-1)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), ".md")
		split := strings.Split(name, "_")
		url := "/" + strings.Join(split, "/")
		oldPath := tmpDir + "/" + file.Name()
		newPath := dir + url
		os.MkdirAll(newPath, os.ModePerm)
		f, err := os.Open(oldPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		lines := []string{
			"---",
			fmt.Sprintf(`title: "%s"`, split[len(split)-1]),
			fmt.Sprintf(`url: "%s"`, url),
			`---`,
		}
		i := 0
		for scanner.Scan() {
			if i < 2 {
				i++
				continue
			}
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		n, err := os.Create(newPath + "/_index.md")
		if err != nil {
			log.Fatal(err)
		}
		defer n.Close()
		w := bufio.NewWriter(n)
		for _, line := range lines {
			fmt.Fprintln(w, line)
		}
		w.Flush()
	}
	os.Remove(tmpDir)
}
