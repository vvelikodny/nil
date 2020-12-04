package nil

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Int64 holds data of nullable int64 value.
type Int64 struct {
	valid bool
	value int64
}

// Nil returns true of nil value, otherwise - false.
func (s Int64) Nil() bool {
	return !s.valid
}

// Value returns actual non-nil value.
func (s Int64) Value() int64 {
	return s.value
}

// String returns string representation of nillable value.
func (s Int64) String() string {
	if !s.valid {
		return "nil"
	}
	return strconv.FormatInt(s.value, 10)
}

// MarshalJSON
func (s Int64) MarshalJSON() ([]byte, error) {
	if !s.valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.value)
}

// UnmarshalJSON
func (s *Int64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = Int64{}
		return nil
	}

	var res int64

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = Int64{value: res, valid: true}

	return nil
}

// NewNilInt64 creates new nil int64 value.
func NewNilInt64() Int64 {
	return Int64{
		valid: false,
		value: 0,
	}
}

// NewInt64 creates new int64 value.
func NewInt64(value int64) Int64 {
	return Int64{
		valid: true,
		value: value,
	}
}

func FromInt64Ptr(value *int64) Int64 {
	if value == nil {
		return NewNilInt64()
	}

	return NewInt64(*value)
}
