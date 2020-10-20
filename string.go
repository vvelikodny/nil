package null

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
func (s String) Null() bool {
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

// NewNullString creates new nil string value.
func NewNullString() String {
	return String{
		valid: false,
		value: "",
	}
}

// NewNullString creates new string value.
func NewString(value string) String {
	return String{
		valid: true,
		value: value,
	}
}
