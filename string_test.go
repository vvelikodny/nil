package null_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vvelikodny/null"
)

func TestNilString(t *testing.T) {
	str1 := null.NewNullString()
	require.True(t, str1.Null())
	require.Equal(t, "nil", fmt.Sprint(str1))

	str2 := null.NewString("value")
	require.False(t, str2.Null())
	require.Equal(t, "value", fmt.Sprint(str2))
}

type Json1 struct {
	Field null.String `json:"field,omitempty"`
}

func TestUnmarshalJsonWithNullString(t *testing.T) {
	var j Json1
	require.NoError(t, json.NewDecoder(bytes.NewBufferString(`{}`)).Decode(&j))

	assert.Equal(t, Json1{
		Field: null.NewNullString(),
	}, j)

	require.NoError(t, json.NewDecoder(bytes.NewBufferString(`{"field":null}`)).Decode(&j))
	assert.Equal(t, Json1{
		Field: null.NewNullString(),
	}, j)
}

func TestJsonWithNullString(t *testing.T) {
	var b bytes.Buffer
	require.NoError(t, json.NewEncoder(&b).Encode(Json1{
		Field: null.NewNullString(),
	}))
	assert.Equal(t, `{"field":null}
`, b.String())
}
