/*
Package gsmconfig is responsible for all functions pertaining to the configuration of GSM
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
package gsmconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// GSMConfig represents a GSM configuration
type GSMConfig struct {
	Name            string   `yaml:"name,omitempty"`
	CredentialsFile string   `yaml:"credentialsFile,omitempty"`
	Mode            string   `yaml:"mode,omitempty"`
	Subject         string   `yaml:"subject,omitempty"`
	Threads         int      `yaml:"threads,omitempty"`
	StandardDelay   int      `yaml:"standardDelay,omitempty"`
	Scopes          []string `yaml:"scopes,omitempty"`
	Default         bool     `yaml:"default,omitempty"`
}

// CfgDir should be set to the directory containing all GSM configuration files
var CfgDir string

// CreateConfig creates a new config
func CreateConfig(config *GSMConfig) (string, error) {
	b, err := yaml.Marshal(config)
	if err != nil {
		return "", err
	}
	f, err := os.Create(fmt.Sprintf("%s/%s.yaml", CfgDir, config.Name))
	if err != nil {
		return "", err
	}
	_, err = f.Write(b)
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

// GetConfig returns a single GSM config or an error
func GetConfig(name string) (*GSMConfig, error) {
	f, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.yaml", CfgDir, name))
	if err != nil {
		return nil, err
	}
	config := &GSMConfig{}
	err = yaml.Unmarshal(f, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// LoadConfig sets the default config (.gsm.yaml) and renames the old one to <name>.yaml
func LoadConfig(name string) error {
	var oldConfig *GSMConfig
	_, err := os.Stat(GetConfigPath(".gsm"))
	if !os.IsNotExist(err) {
		oldConfig, err = GetConfig(".gsm")
		if err != nil {
			return err
		}
		fmt.Println("Rename .gsm.yaml to", GetConfigPath(oldConfig.Name))
		err = os.Rename(GetConfigPath(".gsm"), GetConfigPath(oldConfig.Name))
		if err != nil {
			return err
		}
	}
	fmt.Println("Rename", GetConfigPath(name), "to .gsm.yaml")
	err = os.Rename(GetConfigPath(name), GetConfigPath(".gsm"))
	return err
}

func sortConfigs(configs []*GSMConfig) []*GSMConfig {
	var i int
	for i = range configs {
		if configs[i].Default {
			break
		}
	}
	configsSorted := []*GSMConfig{}
	configsSorted = append(configsSorted, configs[i])
	for i = range configs {
		if !configs[i].Default {
			configsSorted = append(configsSorted, configs[i])
		}
	}
	return configsSorted
}

// ListConfigs lists the config files in the config dir
func ListConfigs() ([]*GSMConfig, error) {
	configs := make([]*GSMConfig, 0)
	f, err := os.Open(CfgDir)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}
		b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", CfgDir, file.Name()))
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", file.Name(), err)
			continue
		}
		c := &GSMConfig{}
		err = yaml.Unmarshal(b, c)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", file.Name(), err)
			continue
		}
		if file.Name() == ".gsm.yaml" {
			c.Default = true
		}
		configs = append(configs, c)
	}
	return sortConfigs(configs), nil
}

// RemoveConfig removes a config
func RemoveConfig(name string) error {
	err := os.Remove(GetConfigPath(name))
	return err
}

// GetConfigPath returns the file path to a config file
func GetConfigPath(name string) string {
	return fmt.Sprintf("%s/%s.yaml", CfgDir, name)
}

// GetScopes returns the scopes of a config file so they can be easily added in the Admin Console
func GetScopes(name string) (string, error) {
	config, err := GetConfig(name)
	if err != nil {
		return "", err
	}
	return strings.Join(config.Scopes, ","), nil
}
