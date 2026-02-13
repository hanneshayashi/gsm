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
	"testing"
)

// --- FlagValue constructors ---

func TestStringVal(t *testing.T) {
	fv := StringVal("hello")
	if fv.tag != FlagString {
		t.Errorf("StringVal tag: got %v, want %v", fv.tag, FlagString)
	}
	if fv.str != "hello" {
		t.Errorf("StringVal str: got %q, want %q", fv.str, "hello")
	}
}

func TestBoolVal(t *testing.T) {
	fv := BoolVal(true)
	if fv.tag != FlagBool {
		t.Errorf("BoolVal tag: got %v, want %v", fv.tag, FlagBool)
	}
	if fv.boolean != true {
		t.Errorf("BoolVal boolean: got %v, want %v", fv.boolean, true)
	}
}

func TestInt64Val(t *testing.T) {
	fv := Int64Val(42)
	if fv.tag != FlagInt64 {
		t.Errorf("Int64Val tag: got %v, want %v", fv.tag, FlagInt64)
	}
	if fv.i64 != 42 {
		t.Errorf("Int64Val i64: got %v, want %v", fv.i64, 42)
	}
}

func TestFloat64Val(t *testing.T) {
	fv := Float64Val(3.14)
	if fv.tag != FlagFloat64 {
		t.Errorf("Float64Val tag: got %v, want %v", fv.tag, FlagFloat64)
	}
	if fv.f64 != 3.14 {
		t.Errorf("Float64Val f64: got %v, want %v", fv.f64, 3.14)
	}
}

func TestUint64Val(t *testing.T) {
	fv := Uint64Val(100)
	if fv.tag != FlagUint64 {
		t.Errorf("Uint64Val tag: got %v, want %v", fv.tag, FlagUint64)
	}
	if fv.u64 != 100 {
		t.Errorf("Uint64Val u64: got %v, want %v", fv.u64, 100)
	}
}

// --- Value getters ---

func TestGetString(t *testing.T) {
	v := Value{FlagValue: StringVal("test")}
	if got := v.GetString(); got != "test" {
		t.Errorf("GetString: got %q, want %q", got, "test")
	}
}

func TestGetBool(t *testing.T) {
	v := Value{FlagValue: BoolVal(true)}
	if got := v.GetBool(); got != true {
		t.Errorf("GetBool: got %v, want %v", got, true)
	}
}

func TestGetInt64(t *testing.T) {
	v := Value{FlagValue: Int64Val(99)}
	if got := v.GetInt64(); got != 99 {
		t.Errorf("GetInt64: got %v, want %v", got, 99)
	}
}

func TestGetInt(t *testing.T) {
	v := Value{FlagValue: Int64Val(42)}
	if got := v.GetInt(); got != 42 {
		t.Errorf("GetInt: got %v, want %v", got, 42)
	}
}

func TestGetFloat64(t *testing.T) {
	v := Value{FlagValue: Float64Val(2.718)}
	if got := v.GetFloat64(); got != 2.718 {
		t.Errorf("GetFloat64: got %v, want %v", got, 2.718)
	}
}

func TestGetUint64(t *testing.T) {
	v := Value{FlagValue: Uint64Val(12345)}
	if got := v.GetUint64(); got != 12345 {
		t.Errorf("GetUint64: got %v, want %v", got, 12345)
	}
}

func TestGetStringSlice(t *testing.T) {
	fv := FlagValue{tag: FlagStringSlice, strSlice: []string{"a", "b", "c"}}
	v := Value{FlagValue: fv}
	got := v.GetStringSlice()
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("GetStringSlice: got %v, want [a b c]", got)
	}
}

func TestGetRune(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want rune
	}{
		{"semicolon", ";", ';'},
		{"comma", ",", ','},
		{"tab", "\t", '\t'},
		{"unicode", "ü", 'ü'},
		{"empty", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{FlagValue: StringVal(tt.str)}
			if got := v.GetRune(); got != tt.want {
				t.Errorf("GetRune(%q): got %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}

// --- IsSet ---

func TestIsSet(t *testing.T) {
	var nilVal *Value
	if nilVal.IsSet() {
		t.Error("nil Value.IsSet() should return false")
	}

	unchanged := &Value{Changed: false}
	if unchanged.IsSet() {
		t.Error("unchanged Value.IsSet() should return false")
	}

	changed := &Value{Changed: true}
	if !changed.IsSet() {
		t.Error("changed Value.IsSet() should return true")
	}
}

// --- FlagToMap ---

func TestFlagToMap(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]string
	}{
		{"empty", "", nil},
		{"single", "key=value", map[string]string{"key": "value"}},
		{"multiple", "key1=value1;key2=value2", map[string]string{"key1": "value1", "key2": "value2"}},
		{"equals_in_value", "key=val=ue", map[string]string{"key": "val=ue"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlagToMap(tt.input)
			if tt.want == nil {
				if got != nil {
					t.Errorf("FlagToMap(%q): got %v, want nil", tt.input, got)
				}
				return
			}
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("FlagToMap(%q)[%q]: got %q, want %q", tt.input, k, got[k], v)
				}
			}
		})
	}
}

// --- batchFlagTo* functions ---

func TestBatchFlagToString(t *testing.T) {
	line := []string{"alpha", "beta", "gamma"}
	def := StringVal("default")

	// Index 0 returns default
	if got := batchFlagToString(line, 0, def); got != "default" {
		t.Errorf("batchFlagToString(0): got %q, want %q", got, "default")
	}

	// Index 2 returns "beta"
	if got := batchFlagToString(line, 2, def); got != "beta" {
		t.Errorf("batchFlagToString(2): got %q, want %q", got, "beta")
	}
}

func TestBatchFlagToInt64(t *testing.T) {
	line := []string{"10", "abc", "30"}
	def := Int64Val(99)

	// Index 0 returns default
	got, err := batchFlagToInt64(line, 0, def)
	if err != nil || got != 99 {
		t.Errorf("batchFlagToInt64(0): got %v/%v, want 99/nil", got, err)
	}

	// Index 1 returns parsed value
	got, err = batchFlagToInt64(line, 1, def)
	if err != nil || got != 10 {
		t.Errorf("batchFlagToInt64(1): got %v/%v, want 10/nil", got, err)
	}

	// Index 2 returns error on "abc"
	got, err = batchFlagToInt64(line, 2, def)
	if err == nil {
		t.Errorf("batchFlagToInt64(2): expected error for 'abc', got %v", got)
	}
}

func TestBatchFlagToBool(t *testing.T) {
	line := []string{"true", "false", "invalid"}
	def := BoolVal(false)

	got, err := batchFlagToBool(line, 0, def)
	if err != nil || got != false {
		t.Errorf("batchFlagToBool(0): got %v/%v, want false/nil", got, err)
	}

	got, err = batchFlagToBool(line, 1, def)
	if err != nil || got != true {
		t.Errorf("batchFlagToBool(1): got %v/%v, want true/nil", got, err)
	}

	got, err = batchFlagToBool(line, 3, def)
	if err == nil {
		t.Errorf("batchFlagToBool(3): expected error for 'invalid', got %v", got)
	}
}

func TestBatchFlagToFloat64(t *testing.T) {
	line := []string{"3.14", "abc"}
	def := Float64Val(1.0)

	// Index 0 returns default
	got, err := batchFlagToFloat64(line, 0, def)
	if err != nil || got != 1.0 {
		t.Errorf("batchFlagToFloat64(0): got %v/%v, want 1.0/nil", got, err)
	}

	// Index 1 returns parsed value
	got, err = batchFlagToFloat64(line, 1, def)
	if err != nil || got != 3.14 {
		t.Errorf("batchFlagToFloat64(1): got %v/%v, want 3.14/nil", got, err)
	}

	// Index 2 returns error on "abc"
	got, err = batchFlagToFloat64(line, 2, def)
	if err == nil {
		t.Errorf("batchFlagToFloat64(2): expected error for 'abc', got %v", got)
	}
}

func TestBatchFlagToUint64(t *testing.T) {
	line := []string{"42", "abc"}
	def := Uint64Val(0)

	// Index 0 returns default
	got, err := batchFlagToUint64(line, 0, def)
	if err != nil || got != 0 {
		t.Errorf("batchFlagToUint64(0): got %v/%v, want 0/nil", got, err)
	}

	// Index 1 returns parsed value
	got, err = batchFlagToUint64(line, 1, def)
	if err != nil || got != 42 {
		t.Errorf("batchFlagToUint64(1): got %v/%v, want 42/nil", got, err)
	}

	// Index 2 returns error on "abc"
	got, err = batchFlagToUint64(line, 2, def)
	if err == nil {
		t.Errorf("batchFlagToUint64(2): expected error for 'abc', got %v", got)
	}
}

func TestBatchFlagToStringSlice(t *testing.T) {
	line := []string{"a,b,c", "x"}

	// Index 0 returns nil
	if got := batchFlagToStringSlice(line, 0); got != nil {
		t.Errorf("batchFlagToStringSlice(0): got %v, want nil", got)
	}

	// Index 1 returns ["a", "b", "c"]
	got := batchFlagToStringSlice(line, 1)
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("batchFlagToStringSlice(1): got %v, want [a b c]", got)
	}
}

func TestBatchFlagToStringArray(t *testing.T) {
	line := []string{"single", "other"}

	// Index 0 returns nil
	if got := batchFlagToStringArray(line, 0); got != nil {
		t.Errorf("batchFlagToStringArray(0): got %v, want nil", got)
	}

	// Index 1 returns ["single"]
	got := batchFlagToStringArray(line, 1)
	if len(got) != 1 || got[0] != "single" {
		t.Errorf("batchFlagToStringArray(1): got %v, want [single]", got)
	}
}

// --- MaxThreads ---

func TestMaxThreads(t *testing.T) {
	// A value of 0 should use default (4 or config)
	// A value > 16 should be capped at 16
	if got := MaxThreads(20); got != 16 {
		t.Errorf("MaxThreads(20): got %v, want 16", got)
	}

	if got := MaxThreads(8); got != 8 {
		t.Errorf("MaxThreads(8): got %v, want 8", got)
	}
}
