package nil_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vvelikodny/nil"
)

func TestNilString(t *testing.T) {
	str1 := nil.NewNilString()
	require.True(t, str1.Valid())
	require.Equal(t, "nil", fmt.Sprint(str1))

	str2 := nil.NewString("value")
	require.False(t, str2.Valid())
	require.Equal(t, "value", fmt.Sprint(str2))
}

type Json1 struct {
	Field nil.String `json:"field,omitempty"`
}

func TestUnmarshalJsonWithNullString(t *testing.T) {
	var j Json1
	require.NoError(t, json.NewDecoder(bytes.NewBufferString(`{}`)).Decode(&j))

	assert.Equal(t, Json1{
		Field: nil.NewNilString(),
	}, j)

	require.NoError(t, json.NewDecoder(bytes.NewBufferString(`{"field":null}`)).Decode(&j))
	assert.Equal(t, Json1{
		Field: nil.NewNilString(),
	}, j)

	b := true
	value := nil.FromBoolPtr(&b)
	fmt.Println("nil?", value.Nil())
}

func TestJsonWithNullString(t *testing.T) {
	var b bytes.Buffer
	require.NoError(t, json.NewEncoder(&b).Encode(Json1{
		Field: nil.NewNilString(),
	}))
	assert.Equal(t, `{"field":null}
`, b.String())
}
