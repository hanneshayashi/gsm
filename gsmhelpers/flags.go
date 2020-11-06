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
package gsmhelpers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

type Flag struct {
	Defaults     map[string]interface{}
	Type         string
	Description  string
	Required     []string
	AvailableFor []string
}

type Value struct {
	Index   int64
	Value   interface{}
	Type    string
	Changed bool
}

func (v *Value) IsSet() bool {
	if v != nil && v.Changed {
		return true
	}
	return false
}

func (v Value) GetStringSlice() []string {
	return InterfaceToStringSlice(v.Value)
}

func (v Value) GetBool() bool {
	return InterfaceToBool(v.Value)
}

func (v Value) GetRune() rune {
	return InterfaceToRune(v.Value)
}

func (v Value) GetString() string {
	return InterfaceToString(v.Value)
}

func (v Value) GetUint64() uint64 {
	return InterfaceToUint64(v.Value)
}

func (v Value) GetInt64() int64 {
	return InterfaceToInt64(v.Value)
}

func (v Value) GetFloat64() float64 {
	return InterfaceToFloat64(v.Value)
}

// InterfaceToStringSlice converts an interface to a string slice ([]string] or returns nil if the interface is nil
// Panics if the interface is not a string slice
func InterfaceToStringSlice(i interface{}) []string {
	if i != nil {
		return i.([]string)
	}
	return nil
}

// BatchFlagToStringSlice returns a value from a slice based on an index and default value
func BatchFlagToStringSlice(line []string, index int64) (value []string) {
	if index != 0 {
		value = strings.Split(line[index-1], ",")
	} else {
		value = nil
	}
	return value
}

// InterfaceToRune converts an interface to a rune or returns rune(-1 )if the interface is nil
// Panics if the interface is not a rune
func InterfaceToRune(i interface{}) rune {
	if i != nil {
		s := i.(string)
		if len(s) != 1 {
			log.Fatalf("rune must be exactly one character")
		}
		return []rune(s)[0]
	}
	return rune(-1)
}

// InterfaceToString converts an interface to a string or returns 0 if the interface is nil
// Panics if the interface is not a string
func InterfaceToString(i interface{}) string {
	if i != nil {
		return i.(string)
	}
	return ""
}

// BatchFlagToString returns a value from a slice based on an index and default value
func BatchFlagToString(line []string, index int64, def interface{}) (value string) {
	if index != 0 {
		value = line[index-1]
	} else {
		value = InterfaceToString(def)
	}
	return value
}

// InterfaceToFloat64 converts an interface to an float64 or returns 0 if the interface is nil
// Panics if the interface is not an float64
func InterfaceToFloat64(i interface{}) float64 {
	if i != nil {
		return i.(float64)
	}
	return 0.0
}

// BatchFlagToFloat64 returns a value from a slice based on an index and default value
func BatchFlagToFloat64(line []string, index int64, def interface{}) (value float64, err error) {
	if index != 0 {
		value, err = strconv.ParseFloat(line[index-1], 64)
		if err != nil {
			return InterfaceToFloat64(def), err
		}
	} else {
		value = InterfaceToFloat64(def)
	}
	return value, nil
}

// InterfaceToUint64 converts an interface to a uint64 or returns 0 if the interface is nil
// Panics if the interface is not a uint64
func InterfaceToUint64(i interface{}) uint64 {
	if i != nil {
		return i.(uint64)
	}
	return 0
}

// InterfaceToInt64 converts an interface to an int64 or returns 0 if the interface is nil
// Panics if the interface is not an int64
func InterfaceToInt64(i interface{}) int64 {
	if i != nil {
		return i.(int64)
	}
	return 0
}

// BatchFlagToInt64 returns a value from a slice based on an index and default value
func BatchFlagToInt64(line []string, index int64, def interface{}) (value int64, err error) {
	if index != 0 {
		value, err = strconv.ParseInt(line[index-1], 10, 64)
		if err != nil {
			return InterfaceToInt64(def), err
		}
	} else {
		value = InterfaceToInt64(def)
	}
	return value, nil
}

// InterfaceToBool converts an interface to a bool or returns false if the interface is nil
// Panics if the interface is not a bool
func InterfaceToBool(i interface{}) bool {
	if i != nil {
		return i.(bool)
	}
	return false
}

func BatchFlagToBool(line []string, index int64, def interface{}) (value bool, err error) {
	if index != 0 {
		value, err = strconv.ParseBool(line[index-1])
		if err != nil {
			return InterfaceToBool(def), err
		}
	} else {
		value = InterfaceToBool(def)
	}
	return value, nil
}

func CheckBatchFlags(flag *pflag.Flag) {
	if flag.Changed && flag.Value.String() == "0" {
		log.Fatalf("Columns must be 1-indexed (don't use 0 to reference columns)")
	}
}

//FlagToMap first splits a string by ";" to get the attributes, then each attribute is
func FlagToMap(value string) (m map[string]string) {
	if value != "" {
		m = make(map[string]string)
		split := strings.Split(value, ";")
		for _, att := range split {
			s2 := strings.SplitN(att, "=", 2)
			if len(s2) > 1 {
				m[s2[0]] = s2[1]
			}
		}
	}
	return
}

// FlagsToMap converts all flags to a map
func FlagsToMap(flags *pflag.FlagSet) (m map[string]*Value) {
	m = make(map[string]*Value)
	foo := func(flag *pflag.Flag) {
		m[flag.Name] = &Value{
			Changed: flag.Changed,
			Type:    flag.Value.Type(),
		}
	}
	flags.VisitAll(foo)
	for k := range m {
		// fmt.Printf("%s is %s\n", k, m[k].Type)
		switch m[k].Type {
		case "int64":
			m[k].Value, _ = flags.GetInt64(k)
		case "bool":
			m[k].Value, _ = flags.GetBool(k)
		case "float64":
			m[k].Value, _ = flags.GetFloat64(k)
		case "stringSlice":
			m[k].Value, _ = flags.GetStringSlice(k)
		case "stringArray":
			m[k].Value, _ = flags.GetStringArray(k)
		case "uint64":
			m[k].Value, _ = flags.GetUint64(k)
		default:
			m[k].Value, _ = flags.GetString(k)
		}
	}
	return m
}

func AddFlagsBatch(m map[string]*Flag, flags *pflag.FlagSet, command string) {
	for f := range m {
		if Contains(command, m[f].AvailableFor) {
			flags.Int64(f, 0, m[f].Description)
		}
	}
}

func AddFlags(m map[string]*Flag, flags *pflag.FlagSet, command string) {
	for f := range m {
		if !Contains(command, m[f].AvailableFor) {
			continue
		}
		def := m[f].Defaults[command]
		switch m[f].Type {
		case "int64":
			flags.Int64(f, InterfaceToInt64(def), m[f].Description)
		case "bool":
			flags.Bool(f, InterfaceToBool(def), m[f].Description)
		case "float64":
			flags.Float64(f, InterfaceToFloat64(def), m[f].Description)
		case "stringSlice":
			flags.StringSlice(f, nil, m[f].Description)
		case "stringArray":
			flags.StringArray(f, nil, m[f].Description)
		case "uint64":
			flags.Uint64(f, InterfaceToUint64(def), m[f].Description)
		default:
			flags.String(f, InterfaceToString(def), m[f].Description)
		}
	}
}

// BatchFlagsToMap converts all information for a single csv line to a map to be used as input for the creation of a struct
func BatchFlagsToMap(flags map[string]*Value, defaultFlags map[string]*Flag, line []string, command string) map[string]*Value {
	m := make(map[string]*Value)
	for k := range flags {
		m[k] = &Value{
			Changed: flags[k].Changed,
		}
		if defaultFlags[k] == nil {
			continue
		}
		var err error
		def := defaultFlags[k].Defaults[command]
		switch defaultFlags[k].Type {
		case "int64":
			m[k].Value, err = BatchFlagToInt64(line, flags[k].GetInt64(), def)
		case "bool":
			m[k].Value, err = BatchFlagToBool(line, flags[k].GetInt64(), def)
		case "float64":
			m[k].Value, err = BatchFlagToFloat64(line, flags[k].GetInt64(), def)
		case "stringSlice":
			m[k].Value = BatchFlagToStringSlice(line, flags[k].GetInt64())
		default:
			m[k].Value = BatchFlagToString(line, flags[k].GetInt64(), def)
		}
		if err != nil {
			fmt.Printf("Error paring %s: %v\n", defaultFlags[k].Type, err)
		}
	}
	return m
}
