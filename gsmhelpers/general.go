/*
Copyright © 2020-2023 Hannes Hayashi

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

// Package gsmhelpers contains helper functions to GSM
package gsmhelpers

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/googleapi"
	"gopkg.in/yaml.v3"
)

// standardRetrier is a retrier object that should be used by every function that calls a Google API
var standardRetrier *backoff.ExponentialBackOff

// RetryOn defines the HTTP error codes that should be retried on.
// Note that GSM will always attempt to retry on a 403 error code with a message that indicates a quota / rate limit error
var RetryOn []int

// GetFileContentAsString returns the content of a file as a string
func GetFileContentAsString(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// GetCSVContent gets the content of a CSV file as [][]string
func GetCSVContent(path string, delimiter rune, skipHeader bool) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	r.Comma = delimiter
	csv, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if skipHeader {
		return csv[1:], nil
	}
	return csv, nil
}

// formatError adds an errKey prefix to an error message
func formatError(err error, errKey string) error {
	return fmt.Errorf("%s: %v", errKey, err)
}

// logError returns a retryable error, indicating that the operation should be reattempted or nil if no error occurred or if the error is not retryable
func logError(err error, d time.Duration) {
	if err != nil {
		log.Printf("%v - Retrying after %s...", err, d)
	}
}

// errorIsRetryable checks if a Google API response returned a retryable error
func errorIsRetryable(err error) bool {
	gerr, ok := err.(*googleapi.Error)
	if !ok {
		return false
	}
	keyWords := []string{
		"quota",
		"limit",
		"rate",
	}
	if gerr.Code == 403 {
		msg := strings.ToLower(gerr.Message)
		for i := range keyWords {
			if strings.Contains(msg, keyWords[i]) {
				return true
			}
		}
	} else if Contains(gerr.Code, RetryOn) {
		return true
	}
	return false
}

// SetStandardRetrier sets the standard retrier
func SetStandardRetrier(standardDelay, maxInterval, maxElapsedTime time.Duration) {
	standardRetrier = backoff.NewExponentialBackOff()
	standardRetrier.InitialInterval = standardDelay
	standardRetrier.MaxInterval = maxInterval
	standardRetrier.MaxElapsedTime = maxElapsedTime
	standardRetrier.Multiplier = 2
}

// Contains checks if a value is inside a slice
func Contains[T comparable](value T, slice []T) bool {
	for i := range slice {
		if value == slice[i] {
			return true
		}
	}
	return false
}

// MaxThreads returns the maximum number of threads (goroutines) that should be spawned
func MaxThreads(fThreads int) int {
	var threads int
	if fThreads != 0 {
		threads = fThreads
	} else {
		cThreads := viper.GetInt("threads")
		if cThreads != 0 {
			threads = cThreads
		} else {
			threads = 4
		}
	}
	maxThreads := 16
	if threads > maxThreads {
		return maxThreads
	}
	return threads
}

// GetJSONEncoder returns a new json encoder
func GetJSONEncoder(indent bool) *json.Encoder {
	enc := json.NewEncoder(os.Stdout)
	if indent {
		enc.SetIndent("", "\t")
	}
	return enc
}

// Output streams output in the specified format to stdout
func Output(i any, format string, compress bool) error {
	if format == "json" {
		enc := GetJSONEncoder(!compress)
		return enc.Encode(i)
	}
	if format == "xml" {
		enc := xml.NewEncoder(os.Stdout)
		if !compress {
			enc.Indent("", "\t")
		}
		return enc.Encode(i)
	}
	if format == "yaml" {
		enc := yaml.NewEncoder(os.Stdout)
		return enc.Encode(i)
	}
	return nil
}

// // CreateDocs creates GSM documentation
// func CreateDocs(cmd *cobra.Command) error {
// 	dir := "../gsm-hosting/gsm.hayashi-ke.online/content"
// 	tmpDir := dir + "/tmp"
// 	err := os.MkdirAll(tmpDir, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}
// 	filePrepender := func(filename string) string {
// 		return filename
// 	}
// 	linkHandler := func(name string) string {
// 		return "/" + strings.ReplaceAll(strings.TrimSuffix(strings.ToLower(name), ".md"), "_", "/")
// 	}
// 	err = doc.GenMarkdownTreeCustom(cmd, tmpDir, filePrepender, linkHandler)
// 	if err != nil {
// 		return err
// 	}
// 	d, err := os.Open(tmpDir)
// 	if err != nil {
// 		return err
// 	}
// 	defer d.Close()
// 	files, err := d.Readdir(-1)
// 	if err != nil {
// 		return err
// 	}
// 	for i := range files {
// 		name := strings.TrimSuffix(files[i].Name(), ".md")
// 		split := strings.Split(name, "_")
// 		url := "/" + strings.Join(split, "/")
// 		oldPath := tmpDir + "/" + files[i].Name()
// 		newPath := dir + url
// 		err = os.MkdirAll(newPath, os.ModePerm)
// 		if err != nil {
// 			return err
// 		}
// 		f, err := os.Open(oldPath)
// 		if err != nil {
// 			return err
// 		}
// 		defer f.Close()
// 		scanner := bufio.NewScanner(f)
// 		lines := []string{
// 			"---",
// 			fmt.Sprintf(`title: "%s"`, split[len(split)-1]),
// 			fmt.Sprintf(`url: "%s"`, url),
// 			`---`,
// 		}
// 		i := 0
// 		for scanner.Scan() {
// 			if i < 2 {
// 				i++
// 				continue
// 			}
// 			lines = append(lines, scanner.Text())
// 		}
// 		if err = scanner.Err(); err != nil {
// 			return err
// 		}
// 		n, err := os.Create(newPath + "/_index.md")
// 		if err != nil {
// 			return err
// 		}
// 		defer n.Close()
// 		w := bufio.NewWriter(n)
// 		for i := range lines {
// 			fmt.Fprintln(w, lines[i])
// 		}
// 		err = w.Flush()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return os.Remove(tmpDir)
// }

// getCSVReader uses a FlagSet to read a CSV file and parse it accordingly
func getCSVReader(flags map[string]*Value) (*csv.Reader, error) {
	path := flags["path"].GetString()
	var delimiter rune
	if flags["delimiter"].Changed {
		delimiter = flags["delimiter"].GetRune()
	} else {
		delimiter = ';'
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	r.Comma = delimiter
	return r, nil
}

// GetBatchMaps returns a channel containing maps to be used for batch requests to the Google API
func GetBatchMaps(cmd *cobra.Command, cmdFlags map[string]*Flag) (<-chan map[string]*Value, error) {
	flags, err := consolidateFlags(cmd, cmdFlags)
	if err != nil {
		return nil, fmt.Errorf("error consolidating flags: %v", err)
	}
	csvReader, err := getCSVReader(flags)
	if err != nil {
		return nil, fmt.Errorf("error with CSV file: %v", err)
	}
	threads := MaxThreads(flags["batchThreads"].GetInt())
	maps := make(chan map[string]*Value, threads)
	line, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	err = checkBatchFlags(flags, cmdFlags, int64(len(line)))
	if err != nil {
		return nil, fmt.Errorf("error with batch flag index: %v", err)
	}
	cmdName := cmd.Parent().Use
	if !flags["skipHeader"].GetBool() {
		maps <- batchFlagsToMap(flags, cmdFlags, line, cmdName)
	}
	i := 0
	go func() {
		defer close(maps)
		for {
			i++
			line, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Printf("Error reading line %d: %v\n", i, err)
				continue
			}
			maps <- batchFlagsToMap(flags, cmdFlags, line, cmdName)
		}
	}()
	return maps, nil
}

// GetObjectRetry performs an action that returns an object, retrying on failure when appropriate
func GetObjectRetry(errKey string, c func() (any, error)) (any, error) {
	result, err := backoff.RetryNotifyWithData(func() (any, error) {
		defer Sleep()
		result, err := c()
		if err != nil {
			ferr := formatError(err, errKey)
			if errorIsRetryable(err) {
				return nil, ferr
			}
			return nil, backoff.Permanent(ferr)
		}
		return result, nil
	}, standardRetrier, logError)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ActionRetry performs an action that does not return an object, retrying on failure when appropriate
func ActionRetry(errKey string, c func() error) (bool, error) {
	err := backoff.RetryNotify(func() error {
		defer Sleep()
		err := c()
		if err != nil {
			ferr := formatError(err, errKey)
			if errorIsRetryable(err) {
				return ferr
			}
			return backoff.Permanent(ferr)
		}
		return nil
	}, standardRetrier, logError)
	if err != nil {
		return false, err
	}
	return true, nil
}

// FormatErrorKey formats an error key.
// Error keys are used on error messages to make it easier to debug where an error occurred
func FormatErrorKey(s ...string) string {
	return strings.Join(s, " - ")
}

// Sleep sleeps for standardDelay ms plus a random jitter between 0 and 50
func Sleep() {
	time.Sleep(standardRetrier.InitialInterval + time.Duration(rand.Intn(50))*time.Millisecond)
}

// IsCommandOrChild returns true if the provided command or one of its children was called
func IsCommandOrChild(command ...*cobra.Command) bool {
	for i := range command {
		if command[i].CalledAs() != "" {
			return true
		} else {
			children := command[i].Commands()
			for j := range command[i].Commands() {
				if children[j].CalledAs() != "" {
					return true
				}
			}
		}
	}
	return false
}

func EnsurePrefix(s, p string) string {
	if !strings.HasPrefix(s, p) {
		return p + s
	}
	return s
}
