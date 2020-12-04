package nil

import (
	"bytes"
	"encoding/json"
)

// String holds data of nullable string.
type String struct {
	valid bool
	value string
}

// Nil
func (s String) Valid() bool {
	return !s.valid
}

// Value
func (s String) Value() string {
	return s.value
}

// String
func (s String) String() string {
	if !s.valid {
		return "nil"
	}
	return s.value
}

// MarshalJSON
func (s String) MarshalJSON() ([]byte, error) {
	if !s.valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.value)
}

// UnmarshalJSON
func (s *String) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = String{}
		return nil
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = String{value: res, valid: true}

	return nil
}

// NewNilString creates new nil string value.
func NewNilString() String {
	return String{
		valid: false,
		value: "",
	}
}

// NewString creates new nil string value.
func NewString(value string) String {
	return String{
		valid: true,
		value: value,
	}
}

func FromStringPtr(value *string) String {
	if value == nil {
		return NewNilString()
	}

	return NewString(*value)
}
