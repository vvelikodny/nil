package nil

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Int32 holds data of nullable int32 value.
type Int32 struct {
	valid bool
	value int32
}

// Nil
func (s Int32) Nil() bool {
	return !s.valid
}

// Value
func (s Int32) Value() int32 {
	return s.value
}

// String returns string representation of nillable value.
func (s Int32) String() string {
	if !s.valid {
		return "nil"
	}
	return strconv.FormatInt(int64(s.value), 10)
}

// MarshalJSON
func (s Int32) MarshalJSON() ([]byte, error) {
	if !s.valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.value)
}

// UnmarshalJSON
func (s *Int32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = Int32{}
		return nil
	}

	var res int32

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = Int32{value: res, valid: true}

	return nil
}

// NewNilInt32 creates new nil int32 value.
func NewNilInt32() Int32 {
	return Int32{
		valid: false,
		value: 0,
	}
}

// NewInt32 creates new int32 value.
func NewInt32(value int32) Int32 {
	return Int32{
		valid: true,
		value: value,
	}
}

func FromInt32Ptr(value *int32) Int32 {
	if value == nil {
		return NewNilInt32()
	}

	return NewInt32(*value)
}