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
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/api/googleapi"
	"gopkg.in/yaml.v3"
)

// StandardRetrier is a retrier object that should be used by every function that calls a Google API
var StandardRetrier = newStandardRetrier()

// StandardDelay is the delay (plus a random jitter between 0 and 20) that will be applied after every command.
// This is can be configured either via the config file or via the --standardDelay flag
var StandardDelay int

// GetFileContentAsString returns the content of a file as a string
func GetFileContentAsString(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(f)
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
	if skipHeader {
		csv = csv[1:]
	}
	if err != nil {
		return nil, err
	}
	return csv, nil
}

// FormatError adds an errKey prefix to an error message
func FormatError(err error, errKey string) error {
	return fmt.Errorf("%s: %v", errKey, err)
}

// retryLog returns a retryable error, indicating that the operation should be reattempted or nil if no error ocurred or if the error is not retryable
func retryLog(err error, errKey string) bool {
	defer sleep(StandardDelay)
	if err != nil {
		if errorIsRetryable(err) {
			log.Println(FormatError(err, errKey), "- Retrying...")
			return true
		}
		return false
	}
	return false
}

// errorIsRetryable checks if a Google API response returned a retryable error
func errorIsRetryable(err error) bool {
	gerr, ok := err.(*googleapi.Error)
	if !ok {
		return false
	}
	keyWords := []string{
		"quota",
		"Quota",
		"limit",
		"Limit",
		"rate",
		"Rate",
	}
	if gerr.Code == 403 {
		for _, kw := range keyWords {
			if strings.Contains(gerr.Message, kw) {
				return true
			}
		}
	}
	return false
}

// newStandardRetrier returns a retrier with default values
func newStandardRetrier() *retrier.Retrier {
	// class := retrier.WhitelistClassifier{
	// 	&googleapi.Error{Code: 403},
	// }
	return retrier.New(retrier.ExponentialBackoff(4, 20*time.Second), nil)
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

// StreamOutput streams output in the specified format to stdout
func StreamOutput(i interface{}, format string, compress bool) error {
	if format == "json" {
		enc := json.NewEncoder(os.Stdout)
		if !compress {
			enc.SetIndent("", "\t")
		}
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

func capFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + string(s[1:])
}

func createCrescendoModuleDefs(cmd *cobra.Command) {
	commands := cmd.Commands()
	type outputHandler struct {
		ParameterSetName string `json:"ParameterSetName,omitempty"`
		Handler          string `json:"Handler,omitempty"`
	}
	type parameter struct {
		Name          string `json:"Name,omitempty"`
		OriginalName  string `json:"OriginalName,omitempty"`
		ParameterType string `json:"ParameterType,omitempty"`
		Description   string `json:"Description,omitempty"`
	}
	type crescendoDef struct {
		Schema                  string          `json:"$Schema,omitempty"`
		Verb                    string          `json:"Verb,omitempty"`
		Noun                    string          `json:"Noun,omitempty"`
		OriginalName            string          `json:"OriginalName,omitempty"`
		OriginalCommandElements []string        `json:"OriginalCommandElements,omitempty"`
		OutputHandlers          []outputHandler `json:"OutputHandlers,omitempty"`
		Description             string          `json:"Description,omitempty"`
		Parameters              []parameter     `json:"Parameters,omitempty"`
	}
	for _, command := range commands {
		cresDef := crescendoDef{
			Schema:                  "./Microsoft.PowerShell.Crescendo.Schema.json",
			Verb:                    capFirstLetter(command.Use),
			Noun:                    capFirstLetter(command.Parent().Use),
			Description:             command.Short,
			OriginalName:            "gsm",
			OriginalCommandElements: append(strings.Split(command.CommandPath(), " ")[1:], "--compressOutput"),
			OutputHandlers: []outputHandler{
				{
					ParameterSetName: "Default",
					Handler:          "$args[0] | ConvertFrom-Json",
				},
			},
		}
		command.Flags().VisitAll(func(f *pflag.Flag) {
			p := parameter{
				OriginalName: "--" + f.Name,
				Name:         strings.ToUpper(string(f.Name[0])) + string(f.Name[1:]),
				Description:  f.Usage,
			}
			switch f.Value.Type() {
			case "bool":
				p.ParameterType = "switch"
			default:
				p.ParameterType = "string"
			}
			cresDef.Parameters = append(cresDef.Parameters, p)
		})
		file, _ := json.MarshalIndent(cresDef, "", "\t")
		fileName := "./crescendo/json/" + strings.Join(cresDef.OriginalCommandElements[:len(cresDef.OriginalCommandElements)-1], "_") + ".json"
		_ = ioutil.WriteFile(fileName, file, 0644)
	}
}

func CreateCrescendoModuleDefs(cmd *cobra.Command) {
	commands := cmd.Commands()
	for _, command := range commands {
		createCrescendoModuleDefs(command)
	}
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
		return nil, fmt.Errorf("Error consolidating flags: %v", err)
	}
	csvReader, err := getCSVReader(flags)
	if err != nil {
		return nil, fmt.Errorf("Error with CSV file: %v", err)
	}
	threads := MaxThreads(flags["batchThreads"].GetInt())
	maps := make(chan map[string]*Value, threads)
	line, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	err = checkBatchFlags(flags, cmdFlags, int64(len(line)))
	if err != nil {
		return nil, fmt.Errorf("Error with batch flag index: %v", err)
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
func GetObjectRetry(errKey string, c func() (interface{}, error)) (interface{}, error) {
	var err error
	var result interface{}
	operation := func() error {
		result, err = c()
		if retryLog(err, errKey) {
			return err
		}
		return nil
	}
	StandardRetrier.Run(operation)
	if err != nil {
		return nil, FormatError(err, errKey)
	}
	return result, nil
}

// ActionRetry performs an action that does not return an object, retrying on failure when appropriate
func ActionRetry(errKey string, c func() error) (bool, error) {
	var err error
	operation := func() error {
		err = c()
		if retryLog(err, errKey) {
			return err
		}
		return nil
	}
	StandardRetrier.Run(operation)
	if err != nil {
		return false, FormatError(err, errKey)
	}
	return true, nil
}

// FormatErrorKey formats an error key.
// Error keys are used on error messages to make it easier to debug where an error ocurred
func FormatErrorKey(s ...string) string {
	return strings.Join(s, " - ")
}

// sleep will sleep for the supplied amount of milliseconds plus a jitter between 0 and 20
func sleep(ms int) {
	ms += rand.Intn(50) + 1
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
