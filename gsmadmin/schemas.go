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

package gsmadmin

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteSchema removes a custom schema.
func DeleteSchema(customerID, schemaKey string) (bool, error) {
	srv := getSchemasService()
	c := srv.Delete(customerID, schemaKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID), func() error {
		return c.Do()
	})
	return result, err
}

// GetSchema retrieves a custom schema.
func GetSchema(customerID, schemaKey, fields string) (*admin.Schema, error) {
	srv := getSchemasService()
	c := srv.Get(customerID, schemaKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Schema)
	return r, nil
}

// InsertSchema creates a custom schema.
func InsertSchema(customerID, fields string, schema *admin.Schema) (*admin.Schema, error) {
	srv := getSchemasService()
	c := srv.Insert(customerID, schema)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Schema)
	return r, nil
}

// ListSchema lists custom schemas.
func ListSchema(customerID, fields string) ([]*admin.Schema, error) {
	srv := getSchemasService()
	c := srv.List(customerID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Schemas)
	return r.Schemas, nil
}

// PatchSchema updates a customs schema.
func PatchSchema(customerID, schemaKey, fields string, schema *admin.Schema) (*admin.Schema, error) {
	srv := getSchemasService()
	c := srv.Patch(customerID, schemaKey, schema)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.Schema)
	return r, nil
}
