/*
Package gsmadmin implements the Admin SDK APIs
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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteFeature deletes a feature.
func DeleteFeature(customer, featureKey string) (bool, error) {
	srv := getResourcesFeaturesService()
	c := srv.Delete(customer, featureKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customer, featureKey), func() error {
		return c.Do()
	})
	return result, err
}

// GetFeature retrieves a feature.
func GetFeature(customer, featureKey, fields string) (*admin.Feature, error) {
	srv := getResourcesFeaturesService()
	c := srv.Get(customer, featureKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, featureKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Feature)
	return r, nil
}

// InsertFeature inserts a feature.
func InsertFeature(customer, fields string, feature *admin.Feature) (*admin.Feature, error) {
	srv := getResourcesFeaturesService()
	c := srv.Insert(customer, feature)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, feature.Name), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Feature)
	return r, nil
}

func listFeatures(c *admin.ResourcesFeaturesListCall, ch chan *admin.Feature, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*admin.Features)
	for _, i := range r.Features {
		ch <- i
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		err = listFeatures(c, ch, errKey)
	}
	return err
}

// ListFeatures retrieves a list of features for an account.
func ListFeatures(customer, fields string, cap int) (<-chan *admin.Feature, <-chan error) {
	srv := getResourcesFeaturesService()
	c := srv.List(customer).MaxResults(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	ch := make(chan *admin.Feature, cap)
	err := make(chan error, 1)
	go func() {
		e := listFeatures(c, ch, gsmhelpers.FormatErrorKey(customer))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// PatchFeature updates a feature. This method supports patch semantics.
func PatchFeature(customer, featureKey, fields string, feature *admin.Feature) (*admin.Feature, error) {
	srv := getResourcesFeaturesService()
	c := srv.Patch(customer, featureKey, feature)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customer, featureKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Feature)
	return r, nil
}

// RenameFeature renames a feature.
func RenameFeature(customer, oldName string, featureRename *admin.FeatureRename) (bool, error) {
	srv := getResourcesFeaturesService()
	c := srv.Rename(customer, oldName, featureRename)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customer, oldName), func() error {
		return c.Do()
	})
	return result, err
}
