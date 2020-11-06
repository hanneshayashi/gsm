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
package gsmadmin

import (
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteResourcesFeature deletes a feature.
func DeleteResourcesFeature(customer, featureKey string) (bool, error) {
	srv := getResourcesFeaturesService()
	err := srv.Delete(customer, featureKey).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetResourcesFeature retrieves a feature.
func GetResourcesFeature(customer, featureKey, fields string) (*admin.Feature, error) {
	srv := getResourcesFeaturesService()
	c := srv.Get(customer, featureKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// InsertResourcesFeature inserts a feature.
func InsertResourcesFeature(customer, fields string, feature *admin.Feature) (*admin.Feature, error) {
	srv := getResourcesFeaturesService()
	c := srv.Insert(customer, feature)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

func makeListResourceFeaturesCallAndAppend(c *admin.ResourcesFeaturesListCall, features []*admin.Feature) ([]*admin.Feature, error) {
	r, err := c.Do()
	if err != nil {
		return nil, err
	}
	for _, f := range r.Features {
		features = append(features, f)
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		features, err = makeListResourceFeaturesCallAndAppend(c, features)
	}
	return features, err
}

// ListResourcesFeatures retrieves a list of features for an account.
func ListResourcesFeatures(customer, fields string) ([]*admin.Feature, error) {
	srv := getResourcesFeaturesService()
	c := srv.List(customer)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var features []*admin.Feature
	features, err := makeListResourceFeaturesCallAndAppend(c, features)
	return features, err
}

// PatchResourcesFeature updates a feature. This method supports patch semantics.
func PatchResourcesFeature(customer, featureKey, fields string, feature *admin.Feature) (*admin.Feature, error) {
	srv := getResourcesFeaturesService()
	c := srv.Patch(customer, featureKey, feature)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	r, err := c.Do()
	return r, err
}

// RenameResourcesFeature renames a feature.
func RenameResourcesFeature(customer, oldName string, featureRename *admin.FeatureRename) (bool, error) {
	srv := getResourcesFeaturesService()
	err := srv.Rename(customer, oldName, featureRename).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}
