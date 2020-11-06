/*
Package cmd contains the commands available to the end user
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
package cmd

import (
	"gsm/gsmconfig"
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/calendar/v3"
	ci "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/groupssettings/v1"
	"google.golang.org/api/licensing/v1"
	"google.golang.org/api/people/v1"
)

// configsCmd represents the config command
var configsCmd = &cobra.Command{
	Use:   "configs",
	Short: "Configure GSM",
	Long: `GSM saves configurations in .yaml files inside the user's home directory under
'~/.config/gsm/<config>.yaml'.
The currently in-use config is always '~/.config/gsm/.gsm.yaml'.
When you load a config with 'gsm configs load --name <config>, the current .gsm.yaml will be renamed to <name>.yaml and the loaded config will be renamed to .gsm.yaml.
You can always explicitly specify a config file with the --config flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var defaultScopes = []string{admin.AdminDirectoryUserScope,
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
	"https://www.google.com/m8/feeds/contacts/",
	drive.DriveScope,
	gmail.MailGoogleComScope,
	gmail.GmailSettingsSharingScope,
	gmail.GmailSettingsBasicScope,
	gmail.GmailModifyScope,
	ci.CloudIdentityGroupsScope,
	groupssettings.AppsGroupsSettingsScope,
	calendar.CalendarScope,
	licensing.AppsLicensingScope,
	people.DirectoryReadonlyScope,
	people.ContactsOtherReadonlyScope,
}

var configFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"name": {
		AvailableFor: []string{"get", "getScopes", "new", "load", "remove"},
		Type:         "string",
		Description: `Name of the configuration.
This (plus ".yaml") will be used as the file name.`,
		Required: []string{"new", "load", "remove"},
		Defaults: map[string]interface{}{"get": ".gsm", "getScopes": ".gsm"},
	},
	"credentialsFile": {
		AvailableFor: []string{"new"},
		Type:         "string",
		Description: `Path to the credential file.
Can be relative to the binary or fully qualified.`,
		Required: []string{"new"},
	},
	"mode": {
		AvailableFor: []string{"new"},
		Type:         "string",
		Description: `The mode to operate in. Can be:
[dwd|user]`,
		Required: []string{"new"},
	},
	"subject": {
		AvailableFor: []string{"new"},
		Type:         "string",
		Description:  `The user who should be impersonated with DWD.`,
	},
	"scopes": {
		AvailableFor: []string{"new"},
		Type:         "stringSlice",
		Description:  `OAuth Scopes to use.`,
	},
	"details": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `List detailed information about configs.`,
	},
}

func init() {
	rootCmd.AddCommand(configsCmd)
}

func mapToConfig(flags map[string]*gsmhelpers.Value) (*gsmconfig.GSMConfig, error) {
	config := &gsmconfig.GSMConfig{
		Scopes: defaultScopes,
	}
	if flags["name"].IsSet() {
		config.Name = flags["name"].GetString()
	}
	if flags["credentialsFile"].IsSet() {
		config.CredentialsFile = flags["credentialsFile"].GetString()
	}
	if flags["mode"].IsSet() {
		config.Mode = flags["mode"].GetString()
	}
	if flags["subject"].IsSet() {
		config.Subject = flags["subject"].GetString()
	}
	if flags["scopes"].IsSet() {
		config.Scopes = flags["scopes"].GetStringSlice()
	}
	return config, nil
}
