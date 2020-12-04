package nil

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Int holds data of nullable bool value.
type Bool struct {
	valid bool
	value bool
}

// Nil
func (s Bool) Nil() bool {
	return !s.valid
}

// Value
func (s Bool) Value() bool {
	return s.value
}

// String returns string representation of nillable value.
func (s Bool) String() string {
	if !s.valid {
		return "nil"
	}
	return strconv.FormatBool(s.value)
}

// MarshalJSON
func (s Bool) MarshalJSON() ([]byte, error) {
	if !s.valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.value)
}

// UnmarshalJSON
func (s *Bool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = Bool{}
		return nil
	}

	var res bool

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = Bool{value: res, valid: true}

	return nil
}

// NewNilBool creates new nil bool value.
func NewNilBool() Bool {
	return Bool{
		valid: false,
		value: false,
	}
}

// NewBool creates new bool value.
func NewBool(value bool) Bool {
	return Bool{
		valid: true,
		value: value,
	}
}

func FromBoolPtr(value *bool) Bool {
	if value == nil {
		return NewNilBool()
	}

	return NewBool(*value)
}
