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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// flagType represents the type of a flag value as a typed constant.
type flagType uint8

const (
	// FlagString is the type for string flags.
	FlagString flagType = iota
	// FlagBool is the type for bool flags.
	FlagBool
	// FlagInt64 is the type for int64 flags.
	FlagInt64
	// FlagStringSlice is the type for string slice flags.
	FlagStringSlice
	// FlagStringArray is the type for string array flags.
	FlagStringArray
	// FlagFloat64 is the type for float64 flags.
	FlagFloat64
	// FlagUint64 is the type for uint64 flags.
	FlagUint64
)

// FlagValue holds a single typed flag value. Only the field corresponding to
// the tag is valid. Storing concrete typed fields avoids interface boxing and
// the associated heap allocations.
type FlagValue struct {
	tag      flagType
	str      string
	boolean  bool
	i64      int64
	u64      uint64
	f64      float64
	strSlice []string
}

// StringVal creates a FlagValue holding a string.
func StringVal(s string) FlagValue { return FlagValue{tag: FlagString, str: s} }

// BoolVal creates a FlagValue holding a bool.
func BoolVal(b bool) FlagValue { return FlagValue{tag: FlagBool, boolean: b} }

// Int64Val creates a FlagValue holding an int64.
func Int64Val(i int64) FlagValue { return FlagValue{tag: FlagInt64, i64: i} }

// Float64Val creates a FlagValue holding a float64.
func Float64Val(f float64) FlagValue { return FlagValue{tag: FlagFloat64, f64: f} }

// Uint64Val creates a FlagValue holding a uint64.
func Uint64Val(u uint64) FlagValue { return FlagValue{tag: FlagUint64, u64: u} }

// Flag represents a flag configuration that can be easily reused for multiple commands
type Flag struct {
	Defaults       map[string]FlagValue
	Type           flagType
	Description    string
	Required       []string
	AvailableFor   []string
	Recursive      []string
	ExcludeFromAll bool
}

// Value is the value representation of a flag
type Value struct {
	FlagValue
	Index   int64
	Changed bool
	AllFlag bool
}

// IsSet returns true if the flag has been set by a user
func (v *Value) IsSet() bool {
	return v != nil && v.Changed
}

// GetStringSlice returns the value of the flag as a []string
func (v Value) GetStringSlice() []string {
	return v.strSlice
}

// GetBool returns the value of the flag as a bool
func (v Value) GetBool() bool {
	return v.boolean
}

// GetString returns the value of the flag as a string
func (v Value) GetString() string {
	return v.str
}

// GetInt64 returns the value of the flag as an int64
func (v Value) GetInt64() int64 {
	return v.i64
}

// GetInt returns the value of the flag as an int (convenience wrapper over int64)
func (v Value) GetInt() int {
	return int(v.i64)
}

// GetFloat64 returns the value of the flag as a float64
func (v Value) GetFloat64() float64 {
	return v.f64
}

// GetUint64 returns the value of the flag as a uint64
func (v Value) GetUint64() uint64 {
	return v.u64
}

// GetRune returns the first rune of the string value
func (v Value) GetRune() rune {
	for _, r := range v.str {
		return r
	}
	return 0
}

// batchFlagToStringArray returns a string slice with the column as a single field
func batchFlagToStringArray(line []string, index int64) []string {
	if index != 0 {
		return []string{line[index-1]}
	}
	return nil
}

// batchFlagToStringSlice returns a value from a slice based on an index
func batchFlagToStringSlice(line []string, index int64) []string {
	if index != 0 {
		return strings.Split(line[index-1], ",")
	}
	return nil
}

// batchFlagToString returns a value from a slice based on an index and default value
func batchFlagToString(line []string, index int64, def FlagValue) string {
	if index != 0 {
		return line[index-1]
	}
	return def.str
}

// batchFlagToInt64 returns a value from a slice based on an index and default value
func batchFlagToInt64(line []string, index int64, def FlagValue) (int64, error) {
	if index != 0 {
		v, err := strconv.ParseInt(line[index-1], 10, 64)
		if err != nil {
			return def.i64, err
		}
		return v, nil
	}
	return def.i64, nil
}

// batchFlagToBool returns a value from a slice based on an index and default value
func batchFlagToBool(line []string, index int64, def FlagValue) (bool, error) {
	if index != 0 {
		v, err := strconv.ParseBool(line[index-1])
		if err != nil {
			return def.boolean, err
		}
		return v, nil
	}
	return def.boolean, nil
}

// batchFlagToFloat64 returns a value from a slice based on an index and default value
func batchFlagToFloat64(line []string, index int64, def FlagValue) (float64, error) {
	if index != 0 {
		v, err := strconv.ParseFloat(line[index-1], 64)
		if err != nil {
			return def.f64, err
		}
		return v, nil
	}
	return def.f64, nil
}

// batchFlagToUint64 returns a value from a slice based on an index and default value
func batchFlagToUint64(line []string, index int64, def FlagValue) (uint64, error) {
	if index != 0 {
		v, err := strconv.ParseUint(line[index-1], 10, 64)
		if err != nil {
			return def.u64, err
		}
		return v, nil
	}
	return def.u64, nil
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
func FlagsToMap(flags *pflag.FlagSet) map[string]*Value {
	m := make(map[string]*Value)
	flags.VisitAll(func(flag *pflag.Flag) {
		v := &Value{
			Changed: flag.Changed,
		}
		switch flag.Value.Type() {
		case "int64":
			v.tag = FlagInt64
			v.i64, _ = flags.GetInt64(flag.Name)
		case "bool":
			v.tag = FlagBool
			v.boolean, _ = flags.GetBool(flag.Name)
		case "stringSlice":
			v.tag = FlagStringSlice
			v.strSlice, _ = flags.GetStringSlice(flag.Name)
		case "stringArray":
			v.tag = FlagStringArray
			v.strSlice, _ = flags.GetStringArray(flag.Name)
		case "float64":
			v.tag = FlagFloat64
			v.f64, _ = flags.GetFloat64(flag.Name)
		case "uint64":
			v.tag = FlagUint64
			v.u64, _ = flags.GetUint64(flag.Name)
		default:
			v.tag = FlagString
			v.str, _ = flags.GetString(flag.Name)
		}
		m[flag.Name] = v
	})
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
		case FlagInt64:
			flags.Int64(f, def.i64, m[f].Description)
		case FlagBool:
			flags.Bool(f, def.boolean, m[f].Description)
		case FlagStringSlice:
			flags.StringSlice(f, nil, m[f].Description)
		case FlagStringArray:
			flags.StringArray(f, nil, m[f].Description)
		case FlagFloat64:
			flags.Float64(f, def.f64, m[f].Description)
		case FlagUint64:
			flags.Uint64(f, def.u64, m[f].Description)
		default:
			flags.String(f, def.str, m[f].Description)
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
			m[k].FlagValue = flags[k].FlagValue
			continue
		}
		var err error
		def := defaultFlags[k].Defaults[command]
		switch defaultFlags[k].Type {
		case FlagInt64:
			m[k].tag = FlagInt64
			m[k].i64, err = batchFlagToInt64(line, flags[k].Index, def)
		case FlagBool:
			m[k].tag = FlagBool
			m[k].boolean, err = batchFlagToBool(line, flags[k].Index, def)
		case FlagStringSlice:
			m[k].tag = FlagStringSlice
			m[k].strSlice = batchFlagToStringSlice(line, flags[k].Index)
		case FlagStringArray:
			m[k].tag = FlagStringArray
			m[k].strSlice = batchFlagToStringArray(line, flags[k].Index)
		case FlagFloat64:
			m[k].tag = FlagFloat64
			m[k].f64, err = batchFlagToFloat64(line, flags[k].Index, def)
		case FlagUint64:
			m[k].tag = FlagUint64
			m[k].u64, err = batchFlagToUint64(line, flags[k].Index, def)
		default:
			m[k].tag = FlagString
			m[k].str = batchFlagToString(line, flags[k].Index, def)
		}
		if err != nil {
			log.Fatalf("Error parsing %v: %v\n", defaultFlags[k].Type, err)
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
			flagsNew[k].tag = cmdFlags[k].Type
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
