package nil

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Float64 holds data of nullable float64 value.
type Float64 struct {
	valid bool
	value float64
}

// Nil
func (s Float64) Nil() bool {
	return !s.valid
}

// Value
func (s Float64) Value() float64 {
	return s.value
}

// String returns string representation of nillable value.
func (s Float64) String() string {
	if !s.valid {
		return "nil"
	}
	return strconv.FormatFloat(s.value, 'f', -1, 64)
}

// MarshalJSON
func (s Float64) MarshalJSON() ([]byte, error) {
	if !s.valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.value)
}

// UnmarshalJSON
func (s *Float64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = Float64{}
		return nil
	}

	var res float64

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = Float64{value: res, valid: true}

	return nil
}

// NewNilFloat64 creates new nil float64 value.
func NewNilFloat64() Float64 {
	return Float64{
		valid: false,
		value: 0,
	}
}

// NewFloat64 creates new float64 value.
func NewFloat64(value float64) Float64 {
	return Float64{
		valid: true,
		value: value,
	}
}

func FromFloat64Ptr(value *float64) Float64 {
	if value == nil {
		return NewNilFloat64()
	}

	return NewFloat64(*value)
}
