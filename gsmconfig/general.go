/*
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

// Package gsmconfig is responsible for all functions pertaining to the configuration of GSM
package gsmconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hanneshayashi/gsm/gsmhelpers"
	"github.com/mitchellh/go-homedir"
	admin "google.golang.org/api/admin/directory/v1"
	reports "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/calendar/v3"
	ci "google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/gmailpostmastertools/v1"
	"google.golang.org/api/groupssettings/v1"
	"google.golang.org/api/licensing/v1"
	"google.golang.org/api/people/v1"
	"google.golang.org/api/sheets/v4"
	"gopkg.in/yaml.v3"
)

// GSMConfig represents a GSM configuration
type GSMConfig struct {
	Name            string   `yaml:"name,omitempty"`
	CredentialsFile string   `yaml:"credentialsFile,omitempty"`
	ServiceAccount  string   `yaml:"serviceAccount,omitempty"`
	Mode            string   `yaml:"mode,omitempty"`
	Subject         string   `yaml:"subject,omitempty"`
	LogFile         string   `yaml:"logFile,omitempty"`
	Scopes          []string `yaml:"scopes,omitempty"`
	Threads         int      `yaml:"threads,omitempty"`
	StandardDelay   int      `yaml:"standardDelay,omitempty"`
	Default         bool     `yaml:"default,omitempty"`
}

// CfgDir should be set to the directory containing all GSM configuration files
var CfgDir string

// UpdateConfig updates a new config
func UpdateConfig(config *GSMConfig, name string) (*GSMConfig, error) {
	configOld, err := GetConfig(name)
	if err != nil {
		return nil, err
	}
	if config.CredentialsFile != "" {
		if configOld.Mode == "adc" {
			return nil, fmt.Errorf("credentialsFile is not used with %s mode", configOld.Mode)
		}
		configOld.CredentialsFile = config.CredentialsFile
	}
	if config.Subject != "" {
		if configOld.Mode == "user" {
			return nil, fmt.Errorf("subject is not used with %s mode", configOld.Mode)
		}
		configOld.Subject = config.Subject
	}
	if config.ServiceAccount != "" {
		if configOld.Mode != "adc" {
			return nil, fmt.Errorf("serviceAccount is not used with %s mode", configOld.Mode)
		}
		configOld.ServiceAccount = config.ServiceAccount
	}
	if config.LogFile != "" {
		configOld.LogFile = config.LogFile
	}
	if config.Name != "" {
		_, err = GetConfig(config.Name)
		if err == nil {
			return nil, fmt.Errorf("%s already exists", config.Name)
		}
		configOld.Name = config.Name
	}
	if config.Scopes != nil {
		configOld.Scopes = config.Scopes
	}
	if config.StandardDelay != 0 {
		configOld.StandardDelay = config.StandardDelay
	}
	if config.Threads != 0 {
		configOld.Threads = gsmhelpers.MaxThreads(config.Threads)
	}
	b, err := yaml.Marshal(configOld)
	if err != nil {
		return nil, err
	}
	configPath := GetConfigPath(name)
	err = ioutil.WriteFile(configPath, b, os.ModeAppend)
	if name != ".gsm" && configOld.Name != name {
		err = os.Rename(configPath, GetConfigPath(configOld.Name))
	}
	if err != nil {
		return nil, err
	}
	return configOld, err
}

// CreateConfig creates a new config
func CreateConfig(config *GSMConfig) (string, error) {
	if !gsmhelpers.Contains(config.Mode, []string{"dwd", "user", "adc"}) {
		return "", fmt.Errorf("%s is not a valid mode", config.Mode)
	}
	if config.Mode == "adc" && config.CredentialsFile != "" {
		return "", fmt.Errorf("credentialsFile is not used with %s mode", config.Mode)
	}
	if config.Mode == "dwd" || config.Mode == "user" {
		if config.CredentialsFile == "" {
			return "", fmt.Errorf("credentialsFile is required with %s mode", config.Mode)
		}
		if config.ServiceAccount != "" {
			return "", fmt.Errorf("serviceAccount is not used with %s mode", config.Mode)
		}
	}
	if (config.Mode == "dwd" || config.Mode == "adc") && config.Subject == "" {
		return "", fmt.Errorf("subject is required with %s mode", config.Mode)
	}
	if config.Mode == "user" && config.Subject != "" {
		return "", fmt.Errorf("subject is not used with %s", config.Mode)
	}
	if config.Threads == 0 {
		config.Threads = gsmhelpers.MaxThreads(0)
	}
	if config.StandardDelay == 0 {
		config.StandardDelay = 500
	}
	if config.LogFile == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		config.LogFile = fmt.Sprintf("%s/gsm.log", home)
	}
	if config.Scopes == nil {
		config.Scopes = []string{admin.AdminDirectoryUserScope,
			admin.AdminDirectoryCustomerScope,
			admin.AdminDirectoryGroupScope,
			admin.AdminDirectoryGroupMemberScope,
			admin.AdminDirectoryOrgunitScope,
			admin.AdminDirectoryRolemanagementScope,
			admin.AdminDirectoryUserSecurityScope,
			admin.AdminDirectoryDomainScope,
			admin.AdminDirectoryDeviceMobileScope,
			admin.AdminDirectoryDeviceChromeosScope,
			admin.AdminDirectoryResourceCalendarScope,
			admin.AdminDirectoryUserschemaScope,
			"https://www.google.com/m8/feeds/contacts/",
			drive.DriveScope,
			gmail.MailGoogleComScope,
			gmail.GmailSettingsSharingScope,
			gmail.GmailSettingsBasicScope,
			gmail.GmailModifyScope,
			ci.CloudIdentityGroupsScope,
			"https://www.googleapis.com/auth/cloud-identity.userinvitations",
			"https://www.googleapis.com/auth/cloud-identity.devices",
			"https://www.googleapis.com/auth/cloud-identity.devices.lookup",
			groupssettings.AppsGroupsSettingsScope,
			calendar.CalendarScope,
			licensing.AppsLicensingScope,
			people.DirectoryReadonlyScope,
			people.ContactsOtherReadonlyScope,
			sheets.SpreadsheetsScope,
			reports.AdminReportsAuditReadonlyScope,
			reports.AdminReportsUsageReadonlyScope,
			gmailpostmastertools.PostmasterReadonlyScope,
			"https://www.googleapis.com/auth/admin.contact.delegation",
			"https://www.googleapis.com/auth/admin.chrome.printers",
		}
	}
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
	_, err := GetConfig(name)
	if err != nil {
		return err
	}
	var oldConfig *GSMConfig
	_, err = os.Stat(GetConfigPath(".gsm"))
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
	if err != nil {
		return nil, err
	}
	defer f.Close()
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for i := range files {
		if !strings.HasSuffix(files[i].Name(), ".yaml") {
			continue
		}
		b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", CfgDir, files[i].Name()))
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", files[i].Name(), err)
			continue
		}
		c := &GSMConfig{}
		err = yaml.Unmarshal(b, c)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", files[i].Name(), err)
			continue
		}
		if files[i].Name() == ".gsm.yaml" {
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
