package nil

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Int holds data of nullable int value.
type Int struct {
	valid bool
	value int
}

// Nil
func (s Int) Nil() bool {
	return !s.valid
}

// Value
func (s Int) Value() int {
	return s.value
}

// String returns string representation of nillable value.
func (s Int) String() string {
	if !s.valid {
		return "nil"
	}
	return strconv.Itoa(s.value)
}

// MarshalJSON
func (s Int) MarshalJSON() ([]byte, error) {
	if !s.valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.value)
}

// UnmarshalJSON
func (s *Int) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = Int{}
		return nil
	}

	var res int

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = Int{value: res, valid: true}

	return nil
}

// NewNilInt creates new nil int value.
func NewNilInt() Int {
	return Int{
		valid: false,
		value: 0,
	}
}

// NewInt creates new int value.
func NewInt(value int) Int {
	return Int{
		valid: true,
		value: value,
	}
}

func FromIntPtr(value *int) Int {
	if value == nil {
		return NewNilInt()
	}

	return NewInt(*value)
}
