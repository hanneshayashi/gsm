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
package cmd

import (
	"log"
	"strconv"

	"github.com/hanneshayashi/gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

// schemasCmd represents the schemas command
var schemasCmd = &cobra.Command{
	Use:               "schemas",
	Short:             "Manage custom schemas for user accounts (Part of Admin SDK)",
	Long:              "Implements the API documented at https://developers.google.com/admin-sdk/directory/reference/rest/v1/schemas",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, _ []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

var schemaFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"customerId": {
		AvailableFor: []string{"delete", "get", "insert", "list", "patch"},
		Type:         "string",
		Description:  `Immutable ID of the Workspace account.`,
		Defaults:     map[string]interface{}{"delete": "my_customer", "get": "my_customer", "insert": "my_customer", "list": "my_customer", "patch": "my_customer"},
	},
	"schemaKey": {
		AvailableFor:   []string{"delete", "get", "patch"},
		Type:           "string",
		Description:    `Name or immutable ID of the schema.`,
		Required:       []string{"delete", "get", "patch"},
		ExcludeFromAll: true,
	},
	"schemaName": {
		AvailableFor:   []string{"insert", "patch"},
		Type:           "string",
		Description:    `The schema's name.`,
		Required:       []string{"insert"},
		ExcludeFromAll: true,
	},
	"displayName": {
		AvailableFor:   []string{"insert", "patch"},
		Type:           "string",
		Description:    `Display name for the schema.`,
		Required:       []string{"insert"},
		ExcludeFromAll: true,
	},
	"schemaFields": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The fields that should be present in this schema.
Can be used multiple times in the form of: "--schemaFields "fieldName=<Some Name>;fieldType=<Type>;multValued=[true|false]...
The following properties are available:
fieldName       - The name of the field.
fieldType       - The type of the field.
				  Possible values are:
                    - STRING  - "Text"
                    - DATE    - "Date
                    - INT64   - "Whole Number"
                    - BOOL    - "Yes or no"
                    - DOUBLE  - "Decimal Number"
                    - PHONE   - "Phone"
                    - EMAIL   - "Email"
multiValued     - A boolean specifying whether this is a multi-valued field or not. Default: false.
indexed         - Boolean specifying whether the field is indexed or not. Default: true.
displayName     - Display Name of the field.
readAccessType  - Specifies who can view values of this field. See Retrieve users as a non-administrator for more information.
				  Note: It may take up to 24 hours for changes to this field to be reflected.
minValue        - Minimum value of this field.
				  This is meant to be indicative rather than enforced.
				  Values outside this range will still be indexed, but search may not be as performant.
maxValue        - Maximum value of this field.
				  This is meant to be indicative rather than enforced.
				  Values outside this range will still be indexed, but search may not be as performant.`,
		Required:       []string{"insert"},
		ExcludeFromAll: true,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
}
var schemaFlagsALL = gsmhelpers.GetAllFlags(schemaFlags)

func init() {
	rootCmd.AddCommand(schemasCmd)
}

func mapToSchema(flags map[string]*gsmhelpers.Value) (*admin.Schema, error) {
	schema := &admin.Schema{}
	var err error
	var ok bool
	if flags["schemaName"].IsSet() {
		schema.SchemaName = flags["schemaName"].GetString()
		if schema.SchemaName == "" {
			schema.ForceSendFields = append(schema.ForceSendFields, "SchemaName")
		}
	}
	if flags["displayName"].IsSet() {
		schema.DisplayName = flags["displayName"].GetString()
		if schema.DisplayName == "" {
			schema.ForceSendFields = append(schema.ForceSendFields, "DisplayName")
		}
	}
	if flags["schemaFields"].IsSet() {
		fields := flags["schemaFields"].GetStringSlice()
		if len(fields) > 0 {
			schema.Fields = []*admin.SchemaFieldSpec{}
			for i := range fields {
				field := &admin.SchemaFieldSpec{}
				m := gsmhelpers.FlagToMap(fields[i])
				field.FieldName, ok = m["fieldName"]
				if field.FieldName == "" && ok {
					field.ForceSendFields = append(field.ForceSendFields, "FieldName")
				}
				field.FieldType, ok = m["fieldType"]
				if field.FieldType == "" && ok {
					field.ForceSendFields = append(field.ForceSendFields, "FieldType")
				}
				field.DisplayName, ok = m["displayName"]
				if field.DisplayName == "" && ok {
					field.ForceSendFields = append(field.ForceSendFields, "DisplayName")
				}
				field.ReadAccessType, ok = m["readAccessType"]
				if field.ReadAccessType == "" && ok {
					field.ForceSendFields = append(field.ForceSendFields, "ReadAccessType")
				}
				var multiValued bool
				_, ok = m["multiValued"]
				if ok {
					multiValued, err = strconv.ParseBool(m["multiValued"])
					if err != nil {
						log.Printf("Error parsing %v to bool: %v. Setting to false.", m["multiValued"], err)
					}
					field.MultiValued = multiValued
					if !multiValued {
						field.ForceSendFields = append(field.ForceSendFields, "MultiValued")
					}
				}
				var indexed bool
				_, ok = m["indexed"]
				if ok {
					indexed, err = strconv.ParseBool(m["indexed"])
					if err != nil {
						log.Printf("Error parsing %v to bool: %v. Setting to true.", m["indexed"], err)
					} else {
						field.Indexed = &indexed
					}
				}
				_, minOk := m["minValue"]
				_, maxOk := m["maxValue"]
				if minOk || maxOk {
					field.NumericIndexingSpec = &admin.SchemaFieldSpecNumericIndexingSpec{}
					if minOk {
						minValue, err := strconv.ParseFloat(m["minValue"], 64)
						if err != nil {
							return nil, err
						}
						field.NumericIndexingSpec.MinValue = minValue
						if field.NumericIndexingSpec.MinValue == 0.0 {
							field.NumericIndexingSpec.ForceSendFields = append(field.NumericIndexingSpec.ForceSendFields, "MinValue")
						}
					}
					if maxOk {
						maxValue, err := strconv.ParseFloat(m["maxValue"], 64)
						if err != nil {
							return nil, err
						}
						field.NumericIndexingSpec.MaxValue = maxValue
						if field.NumericIndexingSpec.MaxValue == 0.0 {
							field.NumericIndexingSpec.ForceSendFields = append(field.NumericIndexingSpec.ForceSendFields, "MaxValue")
						}
					}
				}
				schema.Fields = append(schema.Fields, field)
			}
		} else {
			schema.ForceSendFields = append(schema.ForceSendFields, "Fields")
		}
	}
	return schema, nil
}
