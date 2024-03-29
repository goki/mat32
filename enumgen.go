// Code generated by "enumgen"; DO NOT EDIT.

package mat32

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _DimsValues = []Dims{0, 1, 2, 3}

// DimsN is the highest valid value
// for type Dims, plus one.
const DimsN Dims = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _DimsNoOp() {
	var x [1]struct{}
	_ = x[X-(0)]
	_ = x[Y-(1)]
	_ = x[Z-(2)]
	_ = x[W-(3)]
}

var _DimsNameToValueMap = map[string]Dims{
	`X`: 0,
	`x`: 0,
	`Y`: 1,
	`y`: 1,
	`Z`: 2,
	`z`: 2,
	`W`: 3,
	`w`: 3,
}

var _DimsDescMap = map[Dims]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
}

var _DimsMap = map[Dims]string{
	0: `X`,
	1: `Y`,
	2: `Z`,
	3: `W`,
}

// String returns the string representation
// of this Dims value.
func (i Dims) String() string {
	if str, ok := _DimsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Dims value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Dims) SetString(s string) error {
	if val, ok := _DimsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _DimsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Dims")
}

// Int64 returns the Dims value as an int64.
func (i Dims) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Dims value from an int64.
func (i *Dims) SetInt64(in int64) {
	*i = Dims(in)
}

// Desc returns the description of the Dims value.
func (i Dims) Desc() string {
	if str, ok := _DimsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// DimsValues returns all possible values
// for the type Dims.
func DimsValues() []Dims {
	return _DimsValues
}

// Values returns all possible values
// for the type Dims.
func (i Dims) Values() []enums.Enum {
	res := make([]enums.Enum, len(_DimsValues))
	for i, d := range _DimsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Dims.
func (i Dims) IsValid() bool {
	_, ok := _DimsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Dims) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Dims) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
