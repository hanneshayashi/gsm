/*
Copyright © 2020-2025 Hannes Hayashi

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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Flag represents a flag configuration that can be easily reused for multiple commands
type Flag struct {
	Defaults       map[string]any
	Type           string
	Description    string
	Required       []string
	AvailableFor   []string
	Recursive      []string
	ExcludeFromAll bool
}

// Value is the value representation of a flag
type Value struct {
	Value   any
	Type    string
	Index   int64
	Changed bool
	AllFlag bool
}

// IsSet returns a true or false, depending on if a flag has been set by a user
func (v *Value) IsSet() bool {
	if v != nil && v.Changed {
		return true
	}
	return false
}

// GetStringSlice returns the value of the flag as a []string
func (v Value) GetStringSlice() []string {
	return interfaceToStringSlice(v.Value)
}

// GetBool returns the value of the flag as a bool
func (v Value) GetBool() bool {
	return interfaceToBool(v.Value)
}

// GetRune returns the value of the flag as a rune
func (v Value) GetRune() rune {
	return interfaceToRune(v.Value)
}

// GetString returns the value of the flag as a string
func (v Value) GetString() string {
	return interfaceToString(v.Value)
}

// GetUint64 returns the value of the flag as a uint64
func (v Value) GetUint64() uint64 {
	return interfaceToUint64(v.Value)
}

// GetInt64 returns the value of the flag as an int64
func (v Value) GetInt64() int64 {
	return interfaceToInt64(v.Value)
}

// GetInt returns the value of the flag as an int
func (v Value) GetInt() int {
	return interfaceToInt(v.Value)
}

// GetFloat64 returns the value of the flag as a float64
func (v Value) GetFloat64() float64 {
	return interfaceToFloat64(v.Value)
}

// interfaceToStringSlice converts an interface to a string slice ([]string] or returns nil if the interface is nil
// Panics if the interface is not a string slice
func interfaceToStringSlice(i any) []string {
	if i != nil {
		return i.([]string)
	}
	return nil
}

// interfaceToRune converts an interface to a rune or returns rune(-1 )if the interface is nil
// Panics if the interface is not a rune
func interfaceToRune(i any) rune {
	if i != nil {
		s := i.(string)
		if len(s) != 1 {
			log.Fatalf("rune must be exactly one character")
		}
		return []rune(s)[0]
	}
	return rune(-1)
}

// interfaceToString converts an interface to a string or returns 0 if the interface is nil
// Panics if the interface is not a string
func interfaceToString(i any) string {
	if i != nil {
		return i.(string)
	}
	return ""
}

// batchFlagToStringArray returns a string slice with the column as a single field
func batchFlagToStringArray(line []string, index int64) (value []string) {
	if index != 0 {
		value = []string{line[index-1]}
	} else {
		value = nil
	}
	return value
}

// batchFlagToStringSlice returns a value from a slice based on an index and default value
func batchFlagToStringSlice(line []string, index int64) (value []string) {
	if index != 0 {
		value = strings.Split(line[index-1], ",")
	} else {
		value = nil
	}
	return value
}

// batchFlagToString returns a value from a slice based on an index and default value
func batchFlagToString(line []string, index int64, def any) (value string) {
	if index != 0 {
		value = line[index-1]
	} else {
		value = interfaceToString(def)
	}
	return value
}

// interfaceToFloat64 converts an interface to an float64 or returns 0 if the interface is nil
// Panics if the interface is not an float64
func interfaceToFloat64(i any) float64 {
	if i != nil {
		return i.(float64)
	}
	return 0.0
}

// batchFlagToFloat64 returns a value from a slice based on an index and default value
func batchFlagToFloat64(line []string, index int64, def any) (value float64, err error) {
	if index != 0 {
		value, err = strconv.ParseFloat(line[index-1], 64)
		if err != nil {
			return interfaceToFloat64(def), err
		}
	} else {
		value = interfaceToFloat64(def)
	}
	return value, nil
}

// interfaceToUint64 converts an interface to a uint64 or returns 0 if the interface is nil
// Panics if the interface is not a uint64
func interfaceToUint64(i any) uint64 {
	if i != nil {
		return i.(uint64)
	}
	return 0
}

// interfaceToInt64 converts an interface to an int64 or returns 0 if the interface is nil
// Panics if the interface is not an int64
func interfaceToInt64(i any) int64 {
	if i != nil {
		return i.(int64)
	}
	return 0
}

// interfaceToInt converts an interface to an int or returns 0 if the interface is nil
// Panics if the interface is not an int
func interfaceToInt(i any) int {
	if i != nil {
		return i.(int)
	}
	return 0
}

// batchFlagToInt64 returns a value from a slice based on an index and default value
func batchFlagToInt64(line []string, index int64, def any) (value int64, err error) {
	if index != 0 {
		value, err = strconv.ParseInt(line[index-1], 10, 64)
		if err != nil {
			return interfaceToInt64(def), err
		}
	} else {
		value = interfaceToInt64(def)
	}
	return value, nil
}

// interfaceToBool converts an interface to a bool or returns false if the interface is nil
// Panics if the interface is not a bool
func interfaceToBool(i any) bool {
	if i != nil {
		return i.(bool)
	}
	return false
}

// batchFlagToBool returns a value from a slice based on an index and default value
func batchFlagToBool(line []string, index int64, def any) (value bool, err error) {
	if index != 0 {
		value, err = strconv.ParseBool(line[index-1])
		if err != nil {
			return interfaceToBool(def), err
		}
	} else {
		value = interfaceToBool(def)
	}
	return value, nil
}

// checkBatchFlags checks if the supplied flag values for a batch command are valid in regards to the supplied CSV file
func checkBatchFlags(flags map[string]*Value, defaultFlags map[string]*Flag, length int64) error {
	for k := range flags {
		if defaultFlags[k] == nil || !flags[k].Changed || flags[k].AllFlag {
			continue
		}
		flags[k].Index = flags[k].GetInt64()
		if flags[k].Index == 0 {
			return fmt.Errorf("columns must be 1-indexed (don't use 0 to reference columns)")
		}
		if flags[k].Index > length {
			return fmt.Errorf("index used for %s is out of range. %d > %d. Did you set the delimiter correctly?", k, flags[k].Index, length)
		}
	}
	return nil
}

// FlagToMap first splits a string by ";" to get the attributes, then each attribute is split by "=" to get the key / value pair
func FlagToMap(value string) (m map[string]string) {
	if value != "" {
		m = make(map[string]string)
		split := strings.Split(value, ";")
		for i := range split {
			s2 := strings.SplitN(split[i], "=", 2)
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
		case "int":
			m[k].Value, _ = flags.GetInt(k)
		default:
			m[k].Value, _ = flags.GetString(k)
		}
	}
	return m
}

// addFlagsBatch adds a Int64 flag for all normal flags of a command to be used to reference the column index in a CSV file
func addFlagsBatch(m map[string]*Flag, flags *pflag.FlagSet, command string) {
	for f := range m {
		if Contains(command, m[f].AvailableFor) {
			flags.Int64(f, 0, m[f].Description)
		}
	}
}

// addFlags adds flags to a command
func addFlags(m map[string]*Flag, flags *pflag.FlagSet, command string, recursive bool) {
	for f := range m {
		if !Contains(command, m[f].AvailableFor) || (recursive && !Contains(command, m[f].Recursive)) {
			continue
		}
		def := m[f].Defaults[command]
		switch m[f].Type {
		case "int64":
			flags.Int64(f, interfaceToInt64(def), m[f].Description)
		case "bool":
			flags.Bool(f, interfaceToBool(def), m[f].Description)
		case "float64":
			flags.Float64(f, interfaceToFloat64(def), m[f].Description)
		case "stringSlice":
			flags.StringSlice(f, nil, m[f].Description)
		case "stringArray":
			flags.StringArray(f, nil, m[f].Description)
		case "uint64":
			flags.Uint64(f, interfaceToUint64(def), m[f].Description)
		case "int":
			flags.Int(f, interfaceToInt(def), m[f].Description)
		default:
			flags.String(f, interfaceToString(def), m[f].Description)
		}
	}
}

// batchFlagsToMap converts all information for a single csv line to a map to be used as input for the creation of a struct
func batchFlagsToMap(flags map[string]*Value, defaultFlags map[string]*Flag, line []string, command string) map[string]*Value {
	m := make(map[string]*Value)
	for k := range flags {
		m[k] = &Value{
			Changed: flags[k].Changed,
		}
		if defaultFlags[k] == nil {
			continue
		}
		if flags[k].AllFlag {
			m[k].Value = flags[k].Value
			continue
		}
		var err error
		def := defaultFlags[k].Defaults[command]
		switch defaultFlags[k].Type {
		case "int64":
			m[k].Value, err = batchFlagToInt64(line, flags[k].Index, def)
		case "bool":
			m[k].Value, err = batchFlagToBool(line, flags[k].Index, def)
		case "float64":
			m[k].Value, err = batchFlagToFloat64(line, flags[k].Index, def)
		case "stringSlice":
			m[k].Value = batchFlagToStringSlice(line, flags[k].Index)
		case "stringArray":
			m[k].Value = batchFlagToStringArray(line, flags[k].Index)
		default:
			m[k].Value = batchFlagToString(line, flags[k].Index, def)
		}
		if err != nil {
			log.Fatalf("Error paring %s: %v\n", defaultFlags[k].Type, err)
		}
	}
	return m
}

func markFlagsRequired(cmd *cobra.Command, flags map[string]*Flag, command string) {
	for k := range flags {
		if Contains(command, flags[k].Required) {
			if cmd.Use == "recursive" && !Contains(command, flags[k].Recursive) {
				continue
			}
			err := cmd.MarkFlagRequired(k)
			if err != nil {
				log.Fatalln(cmd.Parent().Parent().Use, cmd.Parent().Use, cmd.Use, command, k, err)
			}
		}
	}
}

// GetAllFlags creates copies of all normal flags with the _ALL suffix.
// These flags are used for batch commands where normal flags get converted to int64 flags that are used to reference columns in CSV files
func GetAllFlags(flags map[string]*Flag) map[string]*Flag {
	flagsAll := map[string]*Flag{}
	for k := range flags {
		if flags[k].ExcludeFromAll {
			continue
		}
		nk := k + "_ALL"
		flagsAll[nk] = &Flag{
			AvailableFor: flags[k].AvailableFor,
			Description:  fmt.Sprintf("Same as %s but value is applied to all lines in the CSV file", k),
			Type:         flags[k].Type,
		}
	}
	return flagsAll
}

// consolidateFlags consolidates a batch commands "normal" and "all" flags
func consolidateFlags(cmd *cobra.Command, cmdFlags map[string]*Flag) (map[string]*Value, error) {
	flags := FlagsToMap(cmd.Flags())
	flagsNew := map[string]*Value{}
	for k := range flags {
		if strings.HasSuffix(k, "_ALL") {
			continue
		}
		flagsNew[k] = flags[k]
	}
	for k := range flagsNew {
		ak := k + "_ALL"
		if flags[k].IsSet() && flags[ak].IsSet() {
			return nil, fmt.Errorf("you can't set a normal flag and its _ALL equivalent at the same time. %s", k)
		}
		if cmdFlags[k] != nil && Contains(cmd.Parent().Use, cmdFlags[k].Required) && !flags[k].IsSet() && !flags[ak].IsSet() {
			return nil, fmt.Errorf("required flag %s is not set", k)
		}
		if !flags[k].IsSet() && flags[ak].IsSet() {
			flagsNew[k] = flags[ak]
			flagsNew[k].AllFlag = true
			flagsNew[k].Type = cmdFlags[k].Type
		}
	}
	return flagsNew, nil
}

// InitBatchCommand sets flags for a batch command appropriately
func InitBatchCommand(parentCmd, childCmd *cobra.Command, cmdFlags, cmdAllFlags, batchFlags map[string]*Flag) {
	parentCmd.AddCommand(childCmd)
	flags := childCmd.Flags()
	addFlagsBatch(cmdFlags, flags, parentCmd.Use)
	addFlags(batchFlags, flags, childCmd.Use, false)
	markFlagsRequired(childCmd, batchFlags, childCmd.Use)
	addFlags(cmdAllFlags, flags, parentCmd.Use, false)
}

// InitCommand sets flags for a command appropriately
func InitCommand(parentCmd, childCmd *cobra.Command, cmdFlags map[string]*Flag) {
	parentCmd.AddCommand(childCmd)
	addFlags(cmdFlags, childCmd.Flags(), childCmd.Use, false)
	markFlagsRequired(childCmd, cmdFlags, childCmd.Use)
}

// InitRecursiveCommand sets flags for a recursive command appropriately
func InitRecursiveCommand(parentCmd, childCmd *cobra.Command, cmdFlags, recursiveFlags map[string]*Flag) {
	parentCmd.AddCommand(childCmd)
	flags := childCmd.Flags()
	addFlags(cmdFlags, flags, parentCmd.Use, true)
	markFlagsRequired(childCmd, cmdFlags, parentCmd.Use)
	addFlags(recursiveFlags, flags, childCmd.Use, false)
	markFlagsRequired(childCmd, recursiveFlags, childCmd.Use)
}

// StringSliceToMapSlice converts a slice of strings to a slice of maps
func StringSliceToMapSlice(slice []string) []map[string]string {
	mapS := make([]map[string]string, 0)
	for i := range slice {
		m := FlagToMap(slice[i])
		mapS = append(mapS, m)
	}
	return mapS
}
