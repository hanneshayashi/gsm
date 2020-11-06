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
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// resourcesFeaturesCmd represents the resourcesFeatures command
var resourcesFeaturesCmd = &cobra.Command{
	Use:   "resourcesFeatures",
	Short: "Manage resource features (Part of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/resources/features",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var resourcesFeatureFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customer": {
		AvailableFor: []string{"delete", "get", "insert", "list", "patch", "rename"},
		Type:         "string",
		Description: `The unique ID for the customer's G Suite account.
As an account administrator, you can also use the my_customer alias to represent your account's customer ID.`,
		Defaults: map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer", "patch": "my_customer", "rename": "my_customer"},
	},
	"featureKey": {
		AvailableFor: []string{"delete", "get", "patch"},
		Type:         "string",
		Description:  `The unique ID of the feature.`,
		Required:     []string{"delete", "get", "patch"},
	},
	"name": {
		AvailableFor: []string{"insert"},
		Type:         "string",
		Description:  `The name of the feature.`,
		Required:     []string{"insert"},
	},
	"oldName": {
		AvailableFor: []string{"rename"},
		Type:         "string",
		Description:  `The unique ID of the feature to rename.`,
		Required:     []string{"rename"},
	},
	"newName": {
		AvailableFor: []string{"rename"},
		Type:         "string",
		Description:  `New name of the feature.`,
		Required:     []string{"rename"},
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}

func init() {
	rootCmd.AddCommand(resourcesFeaturesCmd)
}

func mapToFeature(flags map[string]*gsmhelpers.Value) (*admin.Feature, error) {
	feature := &admin.Feature{}
	if flags["name"].IsSet() {
		feature.Name = flags["name"].GetString()
		if feature.Name == "" {
			feature.ForceSendFields = append(feature.ForceSendFields, "Name")
		}
	}
	return feature, nil
}

func mapToFeatureRename(flags map[string]*gsmhelpers.Value) (*admin.FeatureRename, error) {
	featureRename := &admin.FeatureRename{}
	if flags["newName"].IsSet() {
		featureRename.NewName = flags["newName"].GetString()
		if featureRename.NewName == "" {
			featureRename.ForceSendFields = append(featureRename.ForceSendFields, "NewName")
		}
	}
	return featureRename, nil
}
