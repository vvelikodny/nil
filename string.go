package null

import (
	"bytes"
	"encoding/json"
)

// String represents nullable string value.
type String interface {
	Nullable
	Value() string
}

// nullString holds data of nullable string.
type nullString struct {
	Valid       bool
	StringValue string
}

// Nil
func (s nullString) Null() bool {
	return !s.Valid
}

// Value
func (s nullString) Value() string {
	return s.StringValue
}

// String
func (s nullString) String() string {
	if !s.Valid {
		return "nil"
	}
	return s.StringValue
}

// MarshalJSON
func (s nullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.StringValue)
}

// UnmarshalJSON
func (s *nullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = nullString{}
		return nil
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = nullString{StringValue: res, Valid: true}

	return nil
}

// NewNullString creates new nil string value.
func NewNullString() String {
	return &nullString{
		Valid:       false,
		StringValue: "",
	}
}

// NewNullString creates new string value.
func NewString(value string) String {
	return &nullString{
		Valid:       true,
		StringValue: value,
	}
}
