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

// nilString holds data of nullable string.
type nilString struct {
	Valid       bool
	StringValue string
}

// Nil
func (s nilString) Nil() bool {
	return !s.Valid
}

// Value
func (s nilString) Value() string {
	return s.StringValue
}

// String
func (s nilString) String() string {
	if !s.Valid {
		return "nil"
	}
	return s.StringValue
}

// MarshalJSON
func (s nilString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(s.StringValue)
}

// UnmarshalJSON
func (s *nilString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = nilString{}
		return nil
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*s = nilString{StringValue: res, Valid: true}

	return nil
}

// NewNilString creates new nil string value.
func NewNilString() String {
	return &nilString{
		Valid:       false,
		StringValue: "",
	}
}

// NewNilString creates new string value.
func NewString(value string) String {
	return &nilString{
		Valid:       true,
		StringValue: value,
	}
}
