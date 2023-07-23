// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type Float struct {
	Val  float64
	Null bool
	Set  bool
}

func NewFloat(v float64) Float {
	return Float{Val: v, Set: true}
}

func NewNullableFloat(v *float64) Float {
	if v != nil {
		return NewFloat(*v)
	}
	return Float{Null: true, Set: true}
}

// Scan implements the Scanner interface.
func (f *Float) Scan(value interface{}) error {
	ns := sql.NullFloat64{Float64: f.Val}
	err := ns.Scan(value)
	f.Val, f.Null = ns.Float64, !ns.Valid
	return err
}

// Value implements the driver Valuer interface.
func (f Float) Value() (driver.Value, error) {
	if f.Null {
		return nil, nil
	}
	return float64(f.Val), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (f Float) MarshalJSON() ([]byte, error) {
	if f.Null {
		return json.Marshal(nil)
	}
	return json.Marshal(f.Val)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (f *Float) UnmarshalJSON(data []byte) error {
	f.Set = true
	if string(data) == "null" {
		f.Null = true
		return nil
	}
	if err := json.Unmarshal(data, &f.Val); err != nil {
		return err
	}
	return nil
}

// String implements the stringer interface.
func (f Float) String() string {
	if f.Null {
		return "null"
	}
	return strconv.FormatFloat(f.Val, 'f', 2, 64)
}
